package copy_big_file

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

const Limit200MB = 200 * 1024 * 1024

type FileToCopy struct {
	Src string
	Dst string
}

func SaveToTmp(buffer []string, num int) error {
	if len(buffer) == 0 {
		return nil
	}
	tmp, err := os.Create("./tmp/" + strconv.Itoa(num) + ".txt")
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(tmp)
	for i, line := range buffer {
		if i > 0 {
			if err := writer.WriteByte('\n'); err != nil {
				_ = tmp.Close()
				return err
			}
		}
		if _, err := writer.WriteString(line); err != nil {
			_ = tmp.Close()
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		_ = tmp.Close()
		return err
	}
	return tmp.Close()
}

func ReadFromTmp(num int) ([]string, error) {
	tmp, err := os.Open("./tmp/" + strconv.Itoa(num) + ".txt")
	if err != nil {
		return nil, err
	}
	defer tmp.Close()

	var lines []string
	scanner := bufio.NewScanner(tmp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteToFinalFile(file *os.File, lines []string) error {
	writer := bufio.NewWriter(file)
	for i, line := range lines {
		if i > 0 {
			if err := writer.WriteByte('\n'); err != nil {
				return err
			}
		}
		if _, err := writer.WriteString(line); err != nil {
			return err
		}
	}
	return writer.Flush()
}

func RemoveDirContents(dirPath string) error {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())
		if err = os.RemoveAll(entryPath); err != nil {
			return err
		}
	}

	return nil
}

func CopyBigFile(input FileToCopy) error {
	src, err := os.Open(input.Src)
	if err != nil {
		return err
	}
	defer src.Close()

	var buffer []string
	var currentSize int64
	tmpNumber := 0

	scanner := bufio.NewScanner(src)
	scanner.Buffer(make([]byte, 64*1024), Limit200MB)
	for scanner.Scan() {
		line := scanner.Text()
		buffer = append(buffer, line)
		currentSize += int64(len(line))
		if currentSize >= Limit200MB {
			slices.Reverse(buffer)
			if err = SaveToTmp(buffer, tmpNumber); err != nil {
				return err
			}
			tmpNumber++
			clear(buffer)
			buffer = buffer[:0]
			currentSize = 0
		}
	}
	if err = scanner.Err(); err != nil {
		return err
	}

	if len(buffer) > 0 {
		slices.Reverse(buffer)
		if err = SaveToTmp(buffer, tmpNumber); err != nil {
			return err
		}
		tmpNumber++
		clear(buffer)
		buffer = buffer[:0]
	}

	dst, err := os.Create(input.Dst)
	if err != nil {
		return err
	}

	for i := tmpNumber - 1; i >= 0; i-- {
		if i < tmpNumber-1 {
			if _, err = dst.WriteString("\n"); err != nil {
				_ = dst.Close()
				return err
			}
		}

		var tmp *os.File
		tmp, err = os.Open(filepath.Join("tmp", strconv.Itoa(i)+".txt"))
		if err != nil {
			_ = dst.Close()
			return err
		}
		_, copyErr := io.Copy(dst, tmp)
		closeErr := tmp.Close()
		if err := errors.Join(copyErr, closeErr); err != nil {
			_ = dst.Close()
			return err
		}
	}
	if err := dst.Close(); err != nil {
		return err
	}

	if err := RemoveDirContents("./tmp"); err != nil {
		return err
	}

	return nil
}
