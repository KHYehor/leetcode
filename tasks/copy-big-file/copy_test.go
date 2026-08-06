package copy_big_file

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func useTempWorkingDirectory(t *testing.T) string {
	t.Helper()

	original, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	workingDirectory := t.TempDir()
	if err := os.Chdir(workingDirectory); err != nil {
		t.Fatalf("change working directory: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(original); err != nil {
			t.Errorf("restore working directory: %v", err)
		}
	})

	return workingDirectory
}

func TestSaveToTmpAndReadFromTmp(t *testing.T) {
	workingDirectory := useTempWorkingDirectory(t)
	if err := os.Mkdir("tmp", 0o755); err != nil {
		t.Fatalf("create tmp directory: %v", err)
	}

	want := []string{"first line", "second line", "third line"}
	if err := SaveToTmp(want, 7); err != nil {
		t.Fatalf("SaveToTmp() error = %v", err)
	}

	got, err := ReadFromTmp(7)
	if err != nil {
		t.Fatalf("ReadFromTmp() error = %v", err)
	}
	withoutNewlines := []string{"first line", "second line", "third line"}
	if !reflect.DeepEqual(got, withoutNewlines) {
		t.Errorf("ReadFromTmp() = %#v, want %#v", got, withoutNewlines)
	}

	contents, err := os.ReadFile(filepath.Join(workingDirectory, "tmp", "7.txt"))
	if err != nil {
		t.Fatalf("read saved file: %v", err)
	}
	if got, want := string(contents), "first line\nsecond line\nthird line"; got != want {
		t.Errorf("saved contents = %q, want %q", got, want)
	}
}

func TestSaveToTmpWithEmptyBufferDoesNotCreateFile(t *testing.T) {
	workingDirectory := useTempWorkingDirectory(t)

	if err := SaveToTmp(nil, 3); err != nil {
		t.Fatalf("SaveToTmp() error = %v", err)
	}

	_, err := os.Stat(filepath.Join(workingDirectory, "tmp", "3.txt"))
	if !os.IsNotExist(err) {
		t.Errorf("expected no file to be created, stat error = %v", err)
	}
}

func TestSaveToTmpReturnsCreateError(t *testing.T) {
	useTempWorkingDirectory(t)

	if err := SaveToTmp([]string{"data"}, 1); err == nil {
		t.Fatal("SaveToTmp() error = nil, want an error when tmp directory is missing")
	}
}

func TestReadFromTmpReturnsOpenError(t *testing.T) {
	useTempWorkingDirectory(t)

	if _, err := ReadFromTmp(99); err == nil {
		t.Fatal("ReadFromTmp() error = nil, want an error for a missing file")
	}
}

func TestWriteToFinalFile(t *testing.T) {
	file, err := os.CreateTemp(t.TempDir(), "output-*.txt")
	if err != nil {
		t.Fatalf("create output file: %v", err)
	}
	t.Cleanup(func() { _ = file.Close() })

	if err := WriteToFinalFile(file, []string{"alpha", "beta", "gamma"}); err != nil {
		t.Fatalf("WriteToFinalFile() error = %v", err)
	}
	if err := file.Close(); err != nil {
		t.Fatalf("close output file: %v", err)
	}

	contents, err := os.ReadFile(file.Name())
	if err != nil {
		t.Fatalf("read output file: %v", err)
	}
	if got, want := string(contents), "alpha\nbeta\ngamma"; got != want {
		t.Errorf("output contents = %q, want %q", got, want)
	}
}

func TestRemoveDirContents(t *testing.T) {
	directory := t.TempDir()
	nestedDirectory := filepath.Join(directory, "nested")
	if err := os.Mkdir(nestedDirectory, 0o755); err != nil {
		t.Fatalf("create nested directory: %v", err)
	}
	for name, contents := range map[string]string{
		filepath.Join(directory, "chunk.txt"):      "chunk",
		filepath.Join(nestedDirectory, "data.txt"): "nested data",
	} {
		if err := os.WriteFile(name, []byte(contents), 0o600); err != nil {
			t.Fatalf("write test file %q: %v", name, err)
		}
	}

	if err := RemoveDirContents(directory); err != nil {
		t.Fatalf("RemoveDirContents() error = %v", err)
	}

	entries, err := os.ReadDir(directory)
	if err != nil {
		t.Fatalf("read directory after cleanup: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("directory contains %d entries after cleanup, want 0", len(entries))
	}
}

func TestRemoveDirContentsReturnsReadError(t *testing.T) {
	missingDirectory := filepath.Join(t.TempDir(), "missing")

	if err := RemoveDirContents(missingDirectory); err == nil {
		t.Fatal("RemoveDirContents() error = nil, want an error for a missing directory")
	}
}

func TestCopyBigFile(t *testing.T) {
	workingDirectory := useTempWorkingDirectory(t)
	if err := os.Mkdir("tmp", 0o755); err != nil {
		t.Fatalf("create tmp directory: %v", err)
	}
	if err := os.WriteFile(filepath.Join("tmp", "stale.txt"), []byte("stale"), 0o600); err != nil {
		t.Fatalf("write stale temporary file: %v", err)
	}

	source := filepath.Join(workingDirectory, "source.txt")
	destination := filepath.Join(workingDirectory, "destination.txt")
	want := "one line"
	if err := os.WriteFile(source, []byte(want), 0o600); err != nil {
		t.Fatalf("write source file: %v", err)
	}

	if err := CopyBigFile(FileToCopy{Src: source, Dst: destination}); err != nil {
		t.Fatalf("CopyBigFile() error = %v", err)
	}

	contents, err := os.ReadFile(destination)
	if err != nil {
		t.Fatalf("read destination file: %v", err)
	}
	if got := string(contents); got != want {
		t.Errorf("destination contents = %q, want %q", got, want)
	}

	entries, err := os.ReadDir(filepath.Join(workingDirectory, "tmp"))
	if err != nil {
		t.Fatalf("read tmp directory after copy: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("tmp directory contains %d entries after copy, want 0", len(entries))
	}
}

func TestCopyBigFileReversesLines(t *testing.T) {
	workingDirectory := useTempWorkingDirectory(t)
	if err := os.Mkdir("tmp", 0o755); err != nil {
		t.Fatalf("create tmp directory: %v", err)
	}

	source := filepath.Join(workingDirectory, "source.txt")
	destination := filepath.Join(workingDirectory, "destination.txt")
	if err := os.WriteFile(source, []byte("first\nsecond\nthird\n"), 0o600); err != nil {
		t.Fatalf("write source file: %v", err)
	}

	if err := CopyBigFile(FileToCopy{Src: source, Dst: destination}); err != nil {
		t.Fatalf("CopyBigFile() error = %v", err)
	}

	contents, err := os.ReadFile(destination)
	if err != nil {
		t.Fatalf("read destination file: %v", err)
	}
	if got, want := string(contents), "third\nsecond\nfirst"; got != want {
		t.Errorf("destination contents = %q, want %q", got, want)
	}
}

func TestCopyBigFileReturnsSourceOpenError(t *testing.T) {
	workingDirectory := useTempWorkingDirectory(t)

	err := CopyBigFile(FileToCopy{
		Src: filepath.Join(workingDirectory, "missing.txt"),
		Dst: filepath.Join(workingDirectory, "destination.txt"),
	})
	if err == nil {
		t.Fatal("CopyBigFile() error = nil, want an error for a missing source")
	}
}

func TestTest(t *testing.T) {
	CopyBigFile(FileToCopy{
		Src: "input.txt",
		Dst: "output.txt",
	})
}
