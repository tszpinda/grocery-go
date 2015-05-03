package main

import (
	"time"
)

type Fruit struct {
	Id      int
	Name    string
	Price   float64
	Stock   int
	Updated time.Time
}
