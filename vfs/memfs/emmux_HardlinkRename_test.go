package memfs

import (
	"github.com/laontme/avfs"
	"io"
	"os"
	"testing"
)

func TestEmmuxHardlinkRename(t *testing.T) {
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
	if err = m.WriteFile("/a", []byte("abc"), 0600); err != nil {
		t.Fatal(err)
	}
	if err = m.Link("/a", "/b"); err != nil {
		t.Fatal(err)
	}
	if err = m.Rename("/a", "/b"); err != nil {
		t.Fatal(err)
	}
	a, err := m.Stat("/a")
	if err != nil {
		t.Fatal(err)
	}
	b, err := m.Stat("/b")
	if err != nil {
		t.Fatal(err)
	}
	if !m.SameFile(a, b) || a.Size() != 3 {
		t.Fatal("rename changed inode")
	}
}
