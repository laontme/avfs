package memfs

import (
	"github.com/laontme/avfs"
	"io"
	"os"
	"testing"
)

func TestEmmuxSymlinkSize(t *testing.T) {
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
	target := "missing-target"
	if err = m.Symlink(target, "/s"); err != nil {
		t.Fatal(err)
	}
	info, err := m.Lstat("/s")
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode()&os.ModeSymlink == 0 || info.Size() != int64(len(target)) {
		t.Fatalf("mode %v size %d", info.Mode(), info.Size())
	}
}
