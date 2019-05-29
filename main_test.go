// +build unit



package main

import (
	"testing"
	
)

//jgjghjghjg
func TestCryptoSha256(t *testing.T) {
	t.Log("Testing CryptoSha256 funct... (expected : r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c=)")
	expected := "r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c="
	
	if result := CryptoSha256("johnkhbeis@gmail.com", "secret"); result != expected {
		t.Errorf("Expected result is r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c=, but it was %s instead.", result)
	}
}

func TestExists(t *testing.T) {
	t.Log("Testing Exists func... (expected : r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c=)")
	expected := "r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c="
	
	if result := CryptoSha256("johnkhbeis@gmail.com", "secret"); result != expected {
		t.Errorf("Expected result is r53pPpbhHSwk0L70Mlx23IJKseX7zOhg69MhxApx69c=, but it was %s instead.", result)
	}
}


