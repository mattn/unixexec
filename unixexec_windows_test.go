package unixexec

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSimple(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(filepath.Join(dir, "testdir"))
	if err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(dir)

	p, err := LookPath("git")
	if filepath.HasPrefix(p, dir) {
		t.Fatalf("LookPath should not return commands in current directory: %v", p)
	}
}
