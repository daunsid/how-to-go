package main

import (
	"fmt"
	"sync"
	"time"
	//"google.golang.org/appengine/channel"
)

func FetchUser() string {

	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func FetchUserLikes(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	respch <- 11
	wg.Done()
}

func FetchUserMatch(userName string, respch chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	respch <- "ANNA"
	wg.Done()
}

func main() {
	// begining of composability
	// payload := []byte("hello high value software engineer")
	// fmt.Println(payload)
	// pkg.HashAndBroadCast(pkg.NewHashReader(payload))

	// begining of goroutines
	start := time.Now()
	userName := FetchUser()

	respch := make(chan any, 10)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go FetchUserLikes(userName, respch, wg)
	go FetchUserMatch(userName, respch, wg)
	wg.Wait()
	close(respch)

	for resp := range respch {
		fmt.Println("likes: ", resp)
	}
	fmt.Println(time.Since(start))

}
