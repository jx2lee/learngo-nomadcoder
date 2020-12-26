package goroutine

import (
	"fmt"
	"time"
)

func GoRoutineExam() {
	go sexyCount("jlee")
	sexyCount("jyoung")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
