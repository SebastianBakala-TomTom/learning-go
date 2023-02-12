package main

import (
	"sync/atomic"
	"syscall"
)

// Sample code to show how its importat to use value or pointer semantics
// in a consistent way. Choose the semantic that is reasonable and practical
// for the given type and be consistent. One exception is an unmarshal
// operation since that always requires the address of the value.

type IP []byte
type IPMask []byte

// Mask is using a value reciever and returning a value of type IP.
// This method is using value semantics for type IP.
func (ip IP) Mask(mask IPMask) IP {
	if len(mask) == IPv6len && len(ip) == IPv4len && allFF(mask[:12]) {
		mask = mask[12:]
	}
	if len(mask) == IPv4len && len(ip) == IPv6len && bytesEqual(ip[:12], v4InV6Prefix) {
		ip = ip[12:]
	}
	n := len(ip)
	if n != len(mask) {
		return nil
	}
	out := make(IP, n)
	for i := 0; i < n; i++ {
		out[i] = ip[i] & mask[i]
	}
	return out
}

// ipEmptyString accepts a value of type IP and returns a value of type string.
// The function is using value semantics for type IP.
func ipEmptyString(ip IP) string {
	if len(ip) == 0 {
		return ""
	}
	return ip.String()
}

// Should time use value or pointer semantics? If you need to modify a time
// value should you mutate the value or create a new one?
type Time struct {
	sec  int64
	nsec int32
	loc  *Location
}

// Factory functions dictate the semantics that will be used. The Now function
// returns a value of type Time. This means we should be using value
// semantics and copy Time values.
func Now() Time {
	sec, nsec := now()
	return Time{sec + unixToInternal, nsec, Local}
}

// Add is using a value receiver and returning a value of type Time. This
// method is using value semantics for Time.
func (t Time) Add(d Duration) Time {
	t.sec += int64(d / 1e9)
	nsec := int32(t.nsec) + int32(d%1e9)
	if nsec >= 1e9 {
		t.sec++
		nsec -= 1e9
	} else if nsec < 0 {
		t.sec--
		nsec += 1e9
	}
	t.nsec = nsec
	return t
}

// div accepts a value of type Time and returns values of built-in types.
// The function is using value semantics for type Time.
func div(t Time, d Duration) (qmod2 int, r Duration) {
	// Code here
}

// The only use pointer semantics for the `Time` api are these
// unmarshal related functions.

func (t *Time) UnmarshalBinary(data []byte) error {}
func (t *Time) GobDecode(data []byte) error {}
func (t *Time) UnmarshalJSON(data []byte) error {}
func (t *Time) UnmarshalText(data []byte) error {}