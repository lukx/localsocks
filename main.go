package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/things-go/go-socks5"
)

// InterfaceDialer is a custom dialer that binds to a specific network interface
// and handles temporary outages.
type InterfaceDialer struct {
	ifaceName   string
	mu          sync.Mutex
	lastAttempt time.Time
	backoff     time.Duration
}

// NewInterfaceDialer creates a new InterfaceDialer.
func NewInterfaceDialer(ifaceName string) *InterfaceDialer {
	return &InterfaceDialer{
		ifaceName: ifaceName,
		backoff:   time.Second, // Initial backoff duration
	}
}

// DialContext attempts to dial a connection through the configured network interface.
// If the interface is down, it will log the error and apply an exponential backoff.
func (d *InterfaceDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.backoff > time.Minute {
		log.Fatalf("Interface %s has been down for too long. Exiting.", d.ifaceName)
	}

	// Simple backoff logic
	if time.Since(d.lastAttempt) < d.backoff {
		return nil, fmt.Errorf("interface %s is down, backing off", d.ifaceName)
	}

	dialer, err := d.getNetDialer()
	if err != nil {
		log.Printf("ERROR: Interface %s unavailable: %v. Backing off for %s.", d.ifaceName, err, d.backoff)
		d.lastAttempt = time.Now()
		d.backoff *= 2 // Exponential backoff
		if d.backoff > time.Minute {
			d.backoff = time.Minute // Cap backoff at 1 minute
		}
		return nil, err
	}

	// If we succeed, reset the backoff
	d.backoff = time.Second
	d.lastAttempt = time.Time{} // Reset last attempt time

	return dialer.DialContext(ctx, network, address)
}

func (d *InterfaceDialer) getNetDialer() (*net.Dialer, error) {
	iface, err := net.InterfaceByName(d.ifaceName)
	if err != nil {
		return nil, fmt.Errorf("could not find interface %s: %w", d.ifaceName, err)
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, fmt.Errorf("could not get addresses for interface %s: %w", d.ifaceName, err)
	}

	var localAddr net.Addr
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localAddr = &net.TCPAddr{IP: ipNet.IP}
				break
			}
		}
	}

	if localAddr == nil {
		return nil, fmt.Errorf("no suitable IPv4 address found for interface %s", d.ifaceName)
	}

	return &net.Dialer{LocalAddr: localAddr}, nil
}

func main() {
	port := 10800
	var listener net.Listener
	var err error

	for {
		addr := "127.0.0.1:" + strconv.Itoa(port)
		listener, err = net.Listen("tcp", addr)
		if err == nil {
			break
		}
		log.Printf("Port %d is busy, trying next port", port)
		port++
	}
	defer listener.Close()

	log.Printf("SOCKS proxy listening on %s", listener.Addr().String())

	var dialer func(ctx context.Context, network, address string) (net.Conn, error)
	ifaceName := os.Getenv("LOCALSOCKS_INTERFACE")
	if ifaceName == "" {
		log.Println("WARNING: LOCALSOCKS_INTERFACE not set. Using default system interface.")
		dialer = (&net.Dialer{}).DialContext
	} else {
		log.Printf("Binding upstream connections to interface %s", ifaceName)
		customDialer := NewInterfaceDialer(ifaceName)
		dialer = customDialer.DialContext
	}

	server := socks5.NewServer(
		socks5.WithDial(dialer),
	)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
