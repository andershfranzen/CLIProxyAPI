package managementasset

import (
	"path/filepath"
	"testing"
)

func TestFilePathUsesConfigDirectory(t *testing.T) {
	t.Setenv("MANAGEMENT_STATIC_PATH", "")
	t.Setenv("WRITABLE_PATH", "")
	t.Setenv("writable_path", "")

	dir := t.TempDir()
	got := FilePath(filepath.Join(dir, "config.yaml"))
	want := filepath.Join(dir, "static", ManagementFileName)
	if got != want {
		t.Fatalf("FilePath() = %q, want %q", got, want)
	}
}

func TestFilePathHonoursStaticPathOverride(t *testing.T) {
	dir := t.TempDir()

	t.Setenv("MANAGEMENT_STATIC_PATH", dir)
	if got, want := FilePath("/ignored/config.yaml"), filepath.Join(dir, ManagementFileName); got != want {
		t.Fatalf("FilePath() with directory override = %q, want %q", got, want)
	}

	file := filepath.Join(dir, ManagementFileName)
	t.Setenv("MANAGEMENT_STATIC_PATH", file)
	if got := FilePath("/ignored/config.yaml"); got != file {
		t.Fatalf("FilePath() with file override = %q, want %q", got, file)
	}
	if got := StaticDir("/ignored/config.yaml"); got != dir {
		t.Fatalf("StaticDir() with file override = %q, want %q", got, dir)
	}
}
