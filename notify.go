package gokvo

import (
	"fmt"
	"reflect"
	"sync"
)

// EventName ..
type EventName string

// NotifiAction ..
type NotifiAction func(interface{})

// NotificationCenter ..
type NotificationCenter struct {
	mu       sync.Mutex
	oberPool map[EventName][]*Subscriber
}

// Subscriber ..
type Subscriber struct {
	Name   EventName
	Ober   Observer
	Action NotifiAction
}

func (n *NotificationCenter) addObserver(ober Observer, name EventName, action NotifiAction) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.oberPool == nil {
		n.oberPool = make(map[EventName][]*Subscriber)
	}

	if _, ok := n.oberPool[name]; !ok {
		n.oberPool[name] = make([]*Subscriber, 0)
	}

	s := &Subscriber{Name: name, Ober: ober, Action: action}

	n.oberPool[name] = append(n.oberPool[name], s)
}

func (n *NotificationCenter) removeObserver(ober Observer, name EventName) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.oberPool == nil {
		return
	}

	if _, ok := n.oberPool[name]; !ok {
		return
	}

	for i, v := range n.oberPool[name] {
		if reflect.DeepEqual(v.Ober, ober) {
			n.oberPool[name] = append(n.oberPool[name][:i], n.oberPool[name][i+1:]...)
			break
		}
	}
}

func (n *NotificationCenter) post(name EventName, info interface{}) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.oberPool == nil {
		return fmt.Errorf("post info fail：%s", ErrInvalidPool)
	}

	if _, ok := n.oberPool[name]; !ok {
		return fmt.Errorf("post info fail：%s", ErrInvalidEvent)
	}

	subs := n.oberPool[name]

	if len(subs) == 0 {
		return fmt.Errorf("post info fail：%s", ErrInvalidSubscriber)
	}

	for _, s := range subs {
		s.Action(info)
	}
	return nil
}
