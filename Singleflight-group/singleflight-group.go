package singleflightgroup

import (
	"sync"
)

// NOTE: reproduce behavior of g *singleflight.Group.Do(key string, f func)(result, error, shared)
// HINTS: https://victoriametrics.com/blog/go-singleflight/index.html
//
// Practice w/ goroutine, channel

type Group struct {
	m  map[string]*call // map flight key-call
	mu sync.Mutex       // lock protect Group's mapping
}

// 1 call managed by 1 key
type call struct {
	result interface{}    // result of function call
	err    error          // error from function call
	dups   int            // number of routines in waiting line
	wg     sync.WaitGroup // waiting for the 1st call to finish
}

type Result struct {
	Val    interface{}
	Err    error
	Shared bool
}

func (g *Group) AddFlight(flight Flight, f func() (interface{}, error), wg *sync.WaitGroup) error {
	// Need to check only GettingOff finish
	defer wg.Done()
	result, err, shared := g.Do(flight.GetTicket(), f)
	flight.GettingOff(Result{result, err, shared})
	return err
}

func (g *Group) Do(key string, f func() (interface{}, error)) (interface{}, error, bool) {
	g.mu.Lock()

	// NOTE: The idea should be:
	// - Lock and only unlock after group modification (map edit) is finished
	// - check map existence, make if not
	// - each call will have 1 key
	// - if existed, add increment dups, waits for previous to finish
	// - if not already exist, make new key & call to g.m

	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		g.mu.Unlock() // can unlock here, this one already finished checking
		c.wg.Wait()   // need to wait til the 1st caller finish

		return c.result, c.err, true
	} else {
		c := new(call)
		c.wg.Add(1) // new call being invoked, everyone waiting
		g.m[key] = c
		g.mu.Unlock() // can unlock here, already added the new key

		// get the result, notify it's done, and remove the key
		c.result, c.err = f()
		g.mu.Lock()
		c.wg.Done()
		delete(g.m, key)
		g.mu.Unlock()

		return c.result, c.err, c.dups > 0
	}
}

type Flight interface {
	GetTicket() string
	GettingOff(Result)
}
