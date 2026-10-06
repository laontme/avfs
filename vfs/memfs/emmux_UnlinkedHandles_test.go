package memfs

import (
	"github.com/laontme/avfs"
	"io"
	"os"
	"testing"
)

func TestEmmuxUnlinkedHandles(t *testing.T) {
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
	a, err := m.OpenFile("/a", os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer a.Close()
	if _, err = a.Write([]byte("abc")); err != nil {
		t.Fatal(err)
	}
	b, err := m.OpenFile("/a", os.O_RDWR, 0600)
	if err != nil {
		t.Fatal(err)
	}
	defer b.Close()
	if err = m.Remove("/a"); err != nil {
		t.Fatal(err)
	}
	if _, err = a.Write([]byte("d")); err != nil {
		t.Fatal(err)
	}
	if err = a.Close(); err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(b)
	if err != nil || string(data) != "abcd" {
		t.Fatalf("%q %v", data, err)
	}
	if _, err = m.Stat("/a"); !avfs.IsNotExist(err) {
		t.Fatalf("path recreated: %v", err)
	}
}
