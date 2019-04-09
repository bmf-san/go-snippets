// References:
// https://cloud.google.com/pubsub/docs/overview?hl=ja
// https://www.linkedin.com/pulse/implementing-pub-sub-golang-sakshyam-ghimire/
// https://medium.com/globant/pub-sub-in-golang-an-introduction-8be4c65eafd4
// https://blog.logrocket.com/building-pub-sub-service-go/
// https://ably.com/blog/pubsub-golang
// https://github.com/saksham-ghimire/software-architectures
// https://github.com/glober-vaibhav/go-pub-sub/tree/main
// https://github.com/krazygaurav/pubsub-go

package main

import (
	"fmt"
	"sync"
)

type Manager struct {
	mu   sync.Mutex
	subs map[string][]chan string
	quit chan struct{}
}

func NewManager() *Manager {
	return &Manager{
		subs: make(map[string][]chan string),
		quit: make(chan struct{}),
	}
}

func (b *Manager) Publish(topic string, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, ch := range b.subs[topic] {
		ch <- msg
	}
}

func (b *Manager) Subscribe(topic string) <-chan string {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan string)
	b.subs[topic] = append(b.subs[topic], ch)
	return ch
}

func (b *Manager) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	close(b.quit)

	for _, ch := range b.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}

func main() {
	mg := NewManager()
	subFoo := mg.Subscribe("foo")
	subBar := mg.Subscribe("bar")
	go func() {
		defer mg.Close()
		mg.Publish("foo", "foo-1")
		mg.Publish("foo", "foo-2")
		mg.Publish("bar", "bar-1")
		mg.Publish("bar", "bar-2")
	}()
	fmt.Println(<-subFoo)
	fmt.Println(<-subFoo)
	fmt.Println(<-subBar)
	fmt.Println(<-subBar)
}
