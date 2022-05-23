package Until

import (
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 5053, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 5053, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
