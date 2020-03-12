package pitaya

import (
	"net"
	"testing"
)

// dsajflk
func GetFreePort(t testing.TB) int {
	t.Helper()
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}
