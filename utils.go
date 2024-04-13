import (
	"crypto/rand"
	"fmt"
)

// generateClientID generates a unique client ID.
func generateClientID() uint32 {
	// Simple random generation, replace with a more robust method as needed
	var n uint32
	rand.Read(((*[4]byte)(nil))[:])
	return n
}

// generateGUID generates a globally unique identifier.
func generateGUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
package yrs

// Utility functions for JSON serialization, encoding, and decoding.
