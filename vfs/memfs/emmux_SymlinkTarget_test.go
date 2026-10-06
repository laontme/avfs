package memfs

import "testing"

func TestEmmuxSymlinkTarget(t *testing.T) {
	m, err := New()
	if err != nil {
		t.Fatal(err)
	}
	target := "./missing/../target"
	if err = m.Symlink(target, "/s"); err != nil {
		t.Fatal(err)
	}
	got, err := m.Readlink("/s")
	if err != nil || got != target {
		t.Fatalf("%q %v", got, err)
	}
}
