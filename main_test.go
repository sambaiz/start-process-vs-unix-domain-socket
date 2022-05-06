package main

import (
	"net"
	"os/exec"
	"syscall"
	"testing"
	"time"
)

const value = "abcdef"

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := exec.Command("./main", "--value", value).Output(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSocketDialEveryTime(b *testing.B) {
	cmd := exec.Command("./main", "--server", "--socket", "test.sock")
	if err := cmd.Start(); err != nil {
		b.Fatal(err)
	}
	defer cmd.Process.Signal(syscall.SIGTERM)

	time.Sleep(1e9)

	buf := make([]byte, 512)

	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("unix", "test.sock")
		if err != nil {
			b.Fatal(err)
		}

		conn.Write([]byte(value))
		if _, err := conn.Read(buf); err != nil {
			b.Fatal(err)
		}
		_ = conn.Close()
	}
}

func BenchmarkSocketDialOnce(b *testing.B) {
	cmd := exec.Command("./main", "--server", "--socket", "test.sock")
	if err := cmd.Start(); err != nil {
		b.Fatal(err)
	}
	defer cmd.Process.Signal(syscall.SIGTERM)

	time.Sleep(1e9)

	conn, err := net.Dial("unix", "test.sock")
	if err != nil {
		b.Fatal(err)
	}

	buf := make([]byte, 512)
	for i := 0; i < b.N; i++ {
		conn.Write([]byte(value))
		if _, err := conn.Read(buf); err != nil {
			b.Fatal(err)
		}
	}
	_ = conn.Close()
}

func BenchmarkSocketAlreadyProcessStarted(b *testing.B) {
	cmd := exec.Command("./main", "--server", "--socket", "test.sock")
	if err := cmd.Start(); err != nil {
		b.Fatal(err)
	}
	defer cmd.Process.Signal(syscall.SIGTERM)

	time.Sleep(1e9)

	b.ResetTimer()

	conn, err := net.Dial("unix", "test.sock")
	if err != nil {
		b.Fatal(err)
	}

	buf := make([]byte, 512)
	for i := 0; i < b.N; i++ {
		conn.Write([]byte(value))
		if _, err := conn.Read(buf); err != nil {
			b.Fatal(err)
		}
	}
	_ = conn.Close()
}
