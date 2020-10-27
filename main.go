package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
)

func s(w http.ResponseWriter, r *http.Request) {
	sum := 0.0
	e := 10000000 * (rand.Intn(10) + 1)
	for i := 1; i < e; i++ {
		sum += math.Sqrt(float64(i))
	}
	fmt.Fprintln(w, sum)
}

func h(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", s)
	http.HandleFunc("/health", h)
	http.ListenAndServe(":8080", nil)
}
