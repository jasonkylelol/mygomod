package main

import (
	"fmt"
	"time"
)

var subscribeList []*subscriber

type subscriber struct {
	// cli                        mqtt.Client
	topic                      string
	subExitChan, subExitedChan chan struct{}
}

func subMsg() {
	topics := []string{
		"aaa",
		"bbb",
		"ccc",
	}
	for _, topic := range topics {
		sub := &subscriber{topic: topic}
		go sub.start()
		subscribeList = append(subscribeList, sub)
	}
}

func stopSub() {
	for _, sub := range subscribeList {
		sub.stop()
	}
}

func (s *subscriber) start() {
	s.subExitChan = make(chan struct{})
	s.subExitedChan = make(chan struct{})

	s.loop()
}

func (s *subscriber) loop() {
	fmt.Printf("[mqtt subscribe] [%v] subscribe started\n", s.topic)
	loop := true
	for loop {
		select {
		case <-s.subExitChan:
			fmt.Println("break:", s.topic)
			loop = false
			// break
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Printf("[mqtt subscribe] [%v] subscribe exited\n", s.topic)
	s.subExitedChan <- struct{}{}
}

func (s *subscriber) stop() {
	s.subExitChan <- struct{}{}
	<-s.subExitedChan
}
