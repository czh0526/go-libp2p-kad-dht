//go:build !windows && !wasm

package dht

import (
	"log"
	"os"
	"testing"

	"syscall"
)

func TestMain(m *testing.M) {
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{
		Cur: 4096,
		Max: 4096,
	})
	if err != nil {
		log.Println("failed to increase open file descriptor limit, can't run tests")
		os.Exit(1)
	}
	os.Exit(m.Run())
}
