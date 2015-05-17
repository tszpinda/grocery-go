package main

import (
	"time"
)

// Fruid is main model
type Fruit struct {
	ID      int
	Name    string
	Price   float64
	Stock   int
	Updated time.Time
}
