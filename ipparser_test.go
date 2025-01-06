package ipparser

import (
	"fmt"
	"testing"
)

func TestIsValidIPv4 (t *testing.T) {
	ip := "192.168.2.0"
	if !IsValidIPv4(ip) {
		t.Fatalf("%v is not IPv4", ip)
	}
}

func TestIPv4toDecimal(t *testing.T) {
	ip := "192.178.0.1"
	decimal, err := IPv4toDecimal(ip)
	if (err != nil) {
		t.Fatalf("%s", err.Error())
	}
	fmt.Printf("decimal: %v\n", decimal)
}

func TestIPv4toDecimal2(t *testing.T) {
	ip := "192.178.0.1"
	decimal, err := IPv4toDecimal2(ip)
	if (err != nil) {
		t.Fatalf("%s", err.Error())
	}
	fmt.Printf("decimal: %v\n", decimal)
	
}