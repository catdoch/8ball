package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	things := []string {
		"buy puppies",
		"buy kittens",
		"buy a Golden Retriever farm",
		"go for dinner",
		"find something more exciting than this...",
		"Go for a run",
	}

	n := rand.Intn(len(things))

	fmt.Println("We should...", things[n])
}