package observer

import (
	"fmt"
	"time"
)

type Observer struct {
	name         string
	notification chan string
}

type Subject interface {
	Register(observer Observer)
	NotifyAll(msg string)
}

type NewsAgency struct {
	observers []*Observer
}

func (a *NewsAgency) Register(observer Observer) {
	a.observers = append(a.observers, &observer)
}

func (a *NewsAgency) NotifyAll(msg string) {
	for _, o := range a.observers {
		m := fmt.Sprintf("message %v is sent to the observer named as: %v ",msg, o.name)
		o.notification <- m
	}
}

func ObserverPattern() {
	sub := &NewsAgency{
		observers: []*Observer{},
	}

	o1 := Observer{
		name: "Times of India",
		notification: make(chan string),
	}

	o2 := Observer{
		name: "The Himalayan Post",
		notification: make(chan string),
	}

	sub.Register(o1)
	sub.Register(o2)

	go sub.NotifyAll("breaking news!!!")

	for {
		select {
		case msg := <-o1.notification:
			fmt.Println("received message on observer1:", msg)
		case msg := <-o2.notification:
			fmt.Println("received message on observer2:", msg)
		case <-time.After(2 * time.Second):
			fmt.Println("no message received on any observer in 2 seconds")
		}
	}


}
