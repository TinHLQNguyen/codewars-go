package singleflightgroup

import "sync"

// NOTE: reproduce behavior of g *singleflight.Group.Do(key string, f func)(result, error, shared)
// HINTS: https://victoriametrics.com/blog/go-singleflight/index.html
//
// Practice w/ goroutine, channel

type Group struct {
	m  map[string]*call // map flight key-call
	mu sync.Mutex       // lock protect Group's mapping
}

type call struct {
	result interface{}
	err    error
	dups   int
}

type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

func (g *Group) AddFlight(flight Flight, f func() (interface{}, error)) error {
	result, err, shared := g.Do(flight.GetTicket(), f)
	flight.GettingOff(Result{result, err, shared})
	return err
}

func (g *Group) Do(key string, f func() (interface{}, error)) (interface{}, error, bool) {
	g.mu.Lock()
	defer g.mu.Unlock()

	return nil, nil, false
}

type Flight interface {
	GetTicket() string
	GettingOff(Result)
}
