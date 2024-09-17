package main

import (
	"sync"
	"time"
	"log"
)

type TimeoutHandler struct {
	mu        sync.Mutex
	timer     *time.Timer
	waiting   bool
	timeout   time.Duration
}

func NewTimeoutHandler(timeout time.Duration) *TimeoutHandler {
	return &TimeoutHandler{
		timeout: timeout,
	}
}

func (t *TimeoutHandler) Trigger(action func()) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.waiting {
		log.Println("⏳ Timeout already in progress")
		return
	}

	t.waiting = true

	t.timer = time.AfterFunc(t.timeout, func() {
		t.mu.Lock()
		defer t.mu.Unlock()
    	log.Println("⌛ Processing list")
    	action()
		t.waiting = false
	})
}

func (t *TimeoutHandler) Cancel() {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.timer != nil {
		t.timer.Stop()
		t.waiting = false
	}
}


