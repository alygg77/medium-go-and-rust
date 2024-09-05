package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("/Users/vladimir/Downloads/book-war-and-peace.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Start timing the process

	// Initialize the counter for "Pierre"
	count := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, "Pierre")
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Calculate the elapsed time
	duration := time.Since(start)

	// Output the result
	fmt.Printf("Number of occurrences of 'Pierre': %d\n", count)
	fmt.Printf("Time taken: %v\n", duration)
}
