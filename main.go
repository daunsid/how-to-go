package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"

// 	"github.com/daunsid/how-to-go/pkg"
// )

// func main() {
// 	// begining of composability
// 	// payload := []byte("hello high value software engineer")
// 	// fmt.Println(payload)
// 	// pkg.HashAndBroadCast(pkg.NewHashReader(payload))

// 	// begining of goroutines
// 	start := time.Now()
// 	userName := pkg.FetchUser()

// 	respch := make(chan any, 10)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go pkg.FetchUserLikes(userName, respch, wg)
// 	go pkg.FetchUserMatch(userName, respch, wg)
// 	wg.Wait()
// 	close(respch)

// 	for resp := range respch {
// 		fmt.Println("likes: ", resp)
// 	}
// 	fmt.Println(time.Since(start))

// }
