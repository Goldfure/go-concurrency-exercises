//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

var Tweetfeet = make(chan *Tweet)

func producer(stream Stream) {

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(Tweetfeet)
			return
		}

		Tweetfeet <- tweet
	}
}

func consumer() {
	for t := range Tweetfeet {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// Producer
	go producer(stream)

	// Consumer
	consumer()

	fmt.Printf("Process took %s\n", time.Since(start))
}
