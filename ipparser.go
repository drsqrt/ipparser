package ipparser

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

//check the validity of IPv4
func IsValidIPv4(ipStr string) bool {
	if strings.TrimSpace(ipStr) == "" {
		log.Printf("IPv4 is empty string")
	}

	ip := net.ParseIP(ipStr)
	if ip == nil || strings.Count(ipStr, ".") != 3 {
		log.Printf("invalid ip address")
		return false
	}

	log.Printf("ip is : %v ", ip)
	ip = ip.To4()
	if ip == nil {
		log.Printf("%v not an IPv4 address", ip)
		return false
	}

	for _, ele := range ip {
		log.Println(ele)
	}

	log.Printf("%v IPv4 address", ip)
	return true
}

// unsigned int 32 is the internet way of handling ip addresses
func IPv4toDecimal(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IPv4 address")
	}
	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("not an IPv4 address")
	}

	var decimal uint32
	for i, b := range ip {
		decimal += uint32(b) << (8 * (3 - i))
	}
	return decimal, nil
}

//another way of converting IPv4 string to decimal uint32
func IPv4toDecimal2(ipStr string) (uint32, error) {
	parts := strings.Split(ipStr, ".")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid IPv4 address")
	}

	var result uint32
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 0 || num > 255 {
			return 0, fmt.Errorf("invalid octet in IPv4 address: %s", part)
		}
		// Shift the result left by 8 bits and OR with the current octet
		result = (result << 8) | uint32(num)
	}
	return result, nil
}
