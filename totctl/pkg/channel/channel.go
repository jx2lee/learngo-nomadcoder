package channel

import (
	"fmt"
	"time"
)

func ChannelExam() {
	c := make(chan string)
	people := [5]string{"a", "b", "c", "d", "e"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
