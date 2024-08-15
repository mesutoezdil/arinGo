package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calcpi(5000))
}

// calcpi startet n Goroutinen, um eine
// Näherung von Pi zu berechnen.
func calcpi(n int) float64 {
	ch := make(chan float64, n)
	for k := 0; k < n; k++ {
		// alle n Werte nichtsequentiell berechnen
		go calcsubterm(ch, float64(k))
	}
	// das Ergebnis mit Null initialisieren
	f := float64(0.0)
	for k := 0; k < n; k++ {
		// alle n Werte addieren
		f += <-ch
	}
	return f
}

func calcsubterm(ch chan<- float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}
