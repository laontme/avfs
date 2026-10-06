package memfs

import (
	"github.com/laontme/avfs"
	"io"
	"os"
	"testing"
)

func TestEmmuxAppendHandles(t *testing.T) {
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
	a, err := m.OpenFile("/a", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
	b, err := m.OpenFile("/a", os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	for i, f := range []avfs.File{a, b, a} {
		if _, err = f.Seek(0, io.SeekStart); err != nil {
			t.Fatal(err)
		}
		if _, err = f.Write([]byte{byte('a' + i)}); err != nil {
			t.Fatal(err)
		}
	}
	data, err := m.ReadFile("/a")
	if err != nil || string(data) != "abc" {
		t.Fatalf("%q %v", data, err)
	}
}
