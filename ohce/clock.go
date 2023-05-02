package main

import "time"

type Clock struct {
}

type ClockI interface {
	GetTime() int
}

func (c *Clock) GetTime() int {
	return time.Now().Hour()
}
