package memfs

import (
	"github.com/avfs/avfs"
	"io"
	"os"
	"testing"
)

func TestEmmuxSparseWrite(t *testing.T) {
	m, err := NewWithOptions(&Options{OSType: avfs.OsLinux})
	if err != nil {
		t.Fatal(err)
	}
	// These regressions specify Linux semantics, independent of the host OS.
	if m.OSType() != avfs.OsLinux {
		t.Skip("requires avfs_setostype")
	}
	_ = io.SeekStart
	_ = os.O_RDWR
	f, err := m.OpenFile("/a", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err = f.Write([]byte("abc")); err != nil {
		t.Fatal(err)
	}
	if _, err = f.Seek(5, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	if _, err = f.Write(nil); err != nil {
		t.Fatal(err)
	}
	info, err := f.Stat()
	if err != nil || info.Size() != 3 {
		t.Fatalf("empty write extended file: %v", err)
	}
	if _, err = f.Write([]byte("x")); err != nil {
		t.Fatal(err)
	}
	if _, err = f.Seek(0, io.SeekStart); err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil || string(data) != "abc\x00\x00x" {
		t.Fatalf("%q %v", data, err)
	}
}
