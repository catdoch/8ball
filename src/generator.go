package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	things := []string {
		"Buy puppies",
		"Buy kittens",
		"Buy a Golden Retriever farm",
	}

	n := rand.Intn(len(things))

	//message := fmt.Sprint( "We should...", things[rand.Intn(len(things))] )
	fmt.Println("We should...", things[n])
}