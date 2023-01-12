package src

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenPass(password string, level string) {
	var char int
	if level == "low" {
		char = 6
	} else if level == "med" {
		char = 12
	} else if level == "strong" {
		char = 18
	}
	rPassword := []rune(password)
	var s string
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(rPassword); i++ {
		c := string(rPassword[i])
		if rand.Float64() < 0.5 {
			l := strings.ToLower(c)
			s = s + l
		} else {
			u := strings.ToUpper(c)
			s = s + u
		}
	}

	mixPassword := []rune(s)

	leftPassword := char - len(mixPassword)
	for i := 0; i < leftPassword; i++ {
		min := 33
		max := 126
		g := rand.Intn(max-min+1) + min
		mixPassword = append(mixPassword, rune(g))
	}
	mixPasswordFinal := string(mixPassword)
	fmt.Println(mixPasswordFinal)
}
