package main

import (
	"fmt"
	"net/http"
)

func checkWeb() {
	website := []string{
		"https://stackoverflow.com/",
		"https://github.com/",
		"https://www.linkedin.com/",
		"http://medium.com/",
		"https://golang.org/",
		"https://www.udemy.com/",
		"https://www.coursera.org/",
		"https://wesionary.team/",
	}

	for _, url := range website {
		wg.Add(1)
		mtx.Lock()

		go cekWebsite(url)

	}

	wg.Wait()
}

func cekWebsite(url string) {

	defer wg.Done()
	defer mtx.Unlock()

	if res, err := http.Get(url); err != nil {

		fmt.Println(url, "is done")

	} else {

		fmt.Printf("[%d] %s is up \n", res.StatusCode, url)

	}

}
