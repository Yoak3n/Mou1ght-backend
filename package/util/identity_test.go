package util

import "testing"

func TestGenerateIdentity(t *testing.T) {
	// Test with different lengths
	identity := GenerateIdentity()
	t.Log(identity)
}
