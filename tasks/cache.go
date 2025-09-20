package tasks

import (
	"context"
	"runtime"
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value any, ttl time.Duration)
	Get(key string) (any, bool)
	Delete(key string)
	Close()
}

type Value struct {
	v   any
	ttl time.Time
	mu  sync.RWMutex
}

type SafeCache struct {
	mu      sync.RWMutex
	storage map[string]*Value

	delCh  chan string
	cancel context.CancelFunc
}

func NewCache() Cache {
	ctx, cancel := context.WithCancel(context.Background())

	sc := &SafeCache{
		mu:      sync.RWMutex{},
		storage: make(map[string]*Value),
		delCh:   make(chan string, 1024),
		cancel:  cancel,
	}

	for i := 0; i < runtime.NumCPU()-1; i++ {
		go sc.poolCleaner()
	}

	go sc.ttlChecker(ctx)

	return sc
}

func (s *SafeCache) poolCleaner() {
	for key := range s.delCh {
		s.deleteIfExpired(key)
	}
}

func (s *SafeCache) ttlChecker(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()

			var expiredKeys []string
			s.mu.RLock()
			for k, vp := range s.storage {
				vp.mu.RLock()
				expired := now.After(vp.ttl)
				vp.mu.RUnlock()
				if expired {
					expiredKeys = append(expiredKeys, k)
				}
			}
			s.mu.RUnlock()

			for _, key := range expiredKeys {
				select {
				case s.delCh <- key:
				default:
					s.deleteIfExpired(key)
				}
			}
		}
	}
}

func (s *SafeCache) Set(key string, value any, ttl time.Duration) {
	exp := time.Now().Add(ttl)

	s.mu.Lock()
	vp, ok := s.storage[key]

	// Empty struct - can init
	if !ok {
		s.storage[key] = &Value{
			v:   value,
			ttl: exp,
			mu:  sync.RWMutex{},
		}
		s.mu.Unlock()

		return
	}

	// Value exists - capture the mutex before releasing the main one
	vp.mu.Lock()

	// Release main mutex
	s.mu.Unlock()

	// Overwrite
	vp.v = value
	vp.ttl = exp

	vp.mu.Unlock()

	return
}

func (s *SafeCache) Get(key string) (any, bool) {
	s.mu.RLock()
	vp, ok := s.storage[key]
	s.mu.RUnlock()

	// No value - return nil
	if !ok {
		return nil, false
	}

	vp.mu.RLock()
	expired := time.Now().After(vp.ttl)
	// Not expired -> return the result
	if !expired {
		res := vp.v
		vp.mu.RUnlock()
		return res, true
	}
	vp.mu.RUnlock()

	s.mu.Lock()
	// Forced to get the value again, because it could be changed
	vp2, ok := s.storage[key]
	// Check if the value still exists
	if ok {
		vp2.mu.RLock()
		stillExpired := time.Now().After(vp2.ttl)
		// The value is still expired - can be deleted
		if stillExpired {
			delete(s.storage, key)
			vp2.mu.RUnlock()
			s.mu.Unlock()
			return nil, false
		}
		// The value was overwritten
		res := vp2.v
		vp2.mu.RUnlock()
		s.mu.Unlock()
		return res, true
	}
	// The value no longer exists
	s.mu.Unlock()
	return nil, false
}

func (s *SafeCache) deleteIfExpired(key string) {
	s.mu.Lock()
	vp, ok := s.storage[key]
	if ok {
		vp.mu.RLock()
		if time.Now().After(vp.ttl) {
			delete(s.storage, key)
		}
		vp.mu.RUnlock()
	}
	s.mu.Unlock()
}

func (s *SafeCache) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.storage, key)
}

func (s *SafeCache) Close() {
	s.cancel()
	close(s.delCh)
}
