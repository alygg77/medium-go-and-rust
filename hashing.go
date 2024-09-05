package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"time"
)

func main() {
	// Generate a new public/private keypair
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Message to sign
	message := []byte("Hello, World!")

	// Measure time to sign
	totalIterations := 100000
	var totalDuration time.Duration

	for i := 0; i < totalIterations; i++ {
		start := time.Now()
		_ = ed25519.Sign(privateKey, message)
		duration := time.Since(start)
		totalDuration += duration
	}

	avgDuration := totalDuration / time.Duration(totalIterations)

	fmt.Printf("Average hashing/signing time: %v\n", avgDuration)
	fmt.Printf("Total time for %d iterations: %v\n", totalIterations, totalDuration)
}
