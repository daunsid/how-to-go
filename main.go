package main

import (
	"fmt"
	"time"

	"github.com/daunsid/how-to-go/pkg"
)

func main() {
	// begining of composability
	payload := []byte("hello high value software engineer")
	fmt.Println(payload)
	pkg.HashAndBroadCast(pkg.NewHashReader(payload))

	// begining of goroutines
	start := time.Now()
	userName := pkg.FetchUser()
	likes := pkg.FetchUserLikes(userName)
	match := pkg.FetchUserMatch(userName)

	fmt.Println("likes: ", likes)
	fmt.Println("match: ", match)
	fmt.Println(time.Since(start))

}
