package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var mtx = sync.Mutex{}

func main() {
	mariaScore := []int32{12, 24, 10, 24, 9}
	fmt.Println(breakingRecords(mariaScore))
}
