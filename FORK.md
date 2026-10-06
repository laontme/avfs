# laontme/avfs fork

Upstream: https://github.com/avfs/avfs
Base: 1e793cb300c31dfa8194f52c53c4f74790e0ca2f (2026-10-05).
License: Apache-2.0, unchanged.

master includes all changes. Independent branches rooted at the upstream base:

- codex/fix-unlinked-open-files: preserve data through open handles after unlink.
- codex/fix-sparse-writes: zero-fill gaps; empty writes do not extend files.
- codex/fix-append-position: choose current EOF while holding the shared node lock.
- codex/fix-hardlink-rename: Linux same-inode rename preserves both names.
- codex/fix-symlink-target: preserve exact target strings and own their memory.
- codex/fix-symlink-size: symlink metadata uses target byte length.
- codex/virtual-umask: virtual-only Unix umask; no process umask syscall.
- codex/portable-paths-default: portable OS path emulation enabled by default.

The six fix branches include scoped regression tests and retain the upstream
module path for future upstream PRs. The two integration-policy branches are
separate from filesystem correctness fixes. codex/fork-module-path changes the
combined fork to github.com/laontme/avfs (and reconciles the overlapping append
and sparse-write changes). No upstream PRs have been opened.

The virtual umask policy intentionally changes OSFS behavior. Emmux uses only
MemFS. Build tag avfs_hostpaths opts into host paths, unless avfs_setostype is
also selected. MemFS still uses host timestamps; this fork is not a complete
sandbox or POSIX-conformance implementation.

Verified in Emmux's Nix devshell on Darwin:

    CGO_ENABLED=0 go test -race ./vfs/memfs -skip 'TestMemFS/.*/.*Perm$'
    CGO_ENABLED=0 go test ./... -run '^TestEmmux'

Native permission goldens are excluded; the second command compiles all packages
but runs only the added regression tests. A full upstream suite is not claimed.
