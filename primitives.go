package main

import (
	"fmt"
	"strconv"
	"time"
)

func someFunc(num string, channel chan string) {
	sec, _ := strconv.Atoi(num)
	time.Sleep(time.Duration(sec) * time.Second)
	channel <- num
}

func main() {
	myChannel1 := make(chan string)
	myChannel2 := make(chan string)
	myChannel3 := make(chan string)
	go someFunc("1", myChannel1)
	go someFunc("2", myChannel2)
	go someFunc("3", myChannel3)
	//time.Sleep(time.Second * 2)
	msg1 := <-myChannel1
	msg2 := <-myChannel2
	msg3 := <-myChannel3
	fmt.Println(msg1, msg2, msg3)
}
