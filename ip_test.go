package ipv4

import "testing"

func TestLocalIP(t *testing.T) {
	ip := LocalIP()
	if ip == "" {
		t.Fatalf("No IP was obtained!")
	}
}

func TestLocalIPv4s(t *testing.T) {
	_, err := LocalIPv4s()
	if err != nil {
		t.Fatalf("No IP was obtained!")
	}
}
