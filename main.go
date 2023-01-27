package main

import (
	"sync"
)

var wg = sync.WaitGroup{}
var mtx = sync.Mutex{}

func main() {
	revTriangle(9)

}
