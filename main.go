package main

import (
	"fmt"
	"math/rand"
	"time"
)

func hasSameBirthday(people int) bool {
	// Create a map to track birthdays
	birthdays := make(map[int]bool)

	// Generate a random birthday for each person
	for i := 0; i < people; i++ {
		// Randomly generate a birthday between 1 and 365
		birthday := rand.Intn(365) + 1

		// Check if this birthday is already in the map
		if birthdays[birthday] {
			// If it is, we have found two people with the same birthday
			return true
		}

		// Otherwise, add this birthday to the map
		birthdays[birthday] = true
	}

	// If we finish the loop without finding a duplicate, return false
	return false
}

func calculateProbability(people int, trials int) float64 {
	sameBirthdayCount := 0

	// Perform the simulation for the given number of trials
	for i := 0; i < trials; i++ {
		if hasSameBirthday(people) {
			sameBirthdayCount++
		}
	}

	// Calculate the probability as the fraction of trials where two people shared a birthday
	return float64(sameBirthdayCount) / float64(trials)
}

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Number of trials for the simulation
	trials := 10000

	// Print the table header
	fmt.Printf("%-25s%-25s
", "People in the room", "Probability of 2 (or more) having the same birthday")

	// Simulate for groups of 10, 20, ..., 100 people
	for people := 10; people <= 100; people += 10 {
		// Calculate the probability
		probability := calculateProbability(people, trials)

		// Print the results
		fmt.Printf("%-25d%-25.10f
", people, probability)
	}
}

			   import (
    "crypto/rand"
    // "encoding/base64"
    // "encoding/hex"
    "fmt"
)

func Key() string {
    buf := make([]byte, 16)
    _, err := rand.Read(buf)
    if err != nil {
        panic(err)  // out of randomness, should never happen
    }
    return fmt.Sprintf("%x", buf)
    // or hex.EncodeToString(buf)
    // or base64.StdEncoding.EncodeToString(buf)
}
