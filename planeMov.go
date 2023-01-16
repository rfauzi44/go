package main

import "fmt"

func planeMov(planeDur int) {
	movieDur := [...]int{2, 5, 3, 4, 8, 6}
	for i := 0; i < (len(movieDur)); i++ {
		for j := i + 1; j < len(movieDur); j++ {
			var result = movieDur[i] + movieDur[j]
			if result == planeDur {
				fmt.Println(movieDur[i], movieDur[j])
			}
		}

	}
}
