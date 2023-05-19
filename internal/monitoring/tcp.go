package monitoring

import (
	"fmt"
	"net"
	"time"
)

// MeasureRT measures the Round Trip Time (RT) of a TCP connection to a given address.
func MeasureRT(address string) (time.Duration, error) {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return 0, fmt.Errorf("failed to establish TCP connection: %w", err)
	}
	defer conn.Close()

	startTime := time.Now()

	// You can optionally send a sample data packet
	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		return 0, fmt.Errorf("failed to send data: %w", err)
	}

	// You can optionally receive a response
	response := make([]byte, 1024)
	_, err = conn.Read(response)
	if err != nil {
		return 0, fmt.Errorf("failed to receive response: %w", err)
	}

	rt := time.Since(startTime)
	return rt, nil
}
