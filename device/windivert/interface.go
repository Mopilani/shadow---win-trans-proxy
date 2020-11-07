package windivert

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/imgk/divert-go"
)

func DialIPv4(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp4", "8.8.8.8:53", time.Second)
	if err != nil {
		return
	}

	conn.Close()
}

func DialIPv6(wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp6", "[2001:4860:4860::8888]:53", time.Second)
	if err != nil {
		return
	}

	conn.Close()
}

func GetInterfaceIndex() (uint32, uint32, error) {
	const filter = "not loopback and outbound and (ip.DstAddr = 8.8.8.8 or ipv6.DstAddr = 2001:4860:4860::8888) and tcp.DstPort = 53"
	hd, err := divert.Open(filter, divert.LayerNetwork, divert.PriorityDefault, divert.FlagSniff)
	if err != nil {
		return 0, 0, fmt.Errorf("open interface handle error: %w", err)
	}
	defer hd.Close()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go DialIPv4(wg)

	wg.Add(1)
	go DialIPv6(wg)

	a := new(divert.Address)
	b := make([]byte, 1500)

	if _, err := hd.Recv(b, a); err != nil {
		return 0, 0, err
	}

	if err := hd.Shutdown(divert.ShutdownBoth); err != nil {
		return 0, 0, fmt.Errorf("shutdown interface handle error: %w", err)
	}

	if err := hd.Close(); err != nil {
		return 0, 0, fmt.Errorf("close interface handle error: %w", err)
	}

	wg.Wait()

	nw := a.Network()
	return nw.InterfaceIndex, nw.SubInterfaceIndex, nil
}
