package pkg

import "time"

func FetchUser() string {

	time.Sleep(time.Millisecond * 100)
	return "BOB"
}

func FetchUserLikes(userName string) int {
	time.Sleep(time.Millisecond * 100)
	return 11
}

func FetchUserMatch(userName string) string {
	time.Sleep(time.Millisecond * 100)

	return "ANNA"
}
