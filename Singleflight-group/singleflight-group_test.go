package singleflightgroup_test

import (
	"math/rand"
	"reflect"
	"singleflightgroup"
	"testing"
	"time"
)

type (
	sharingStatus      map[int]bool
	resultChangeStatus map[int]bool
)

func TestSingleFlight(t *testing.T) {
	cases := []struct {
		name               string
		sharingResult      []bool
		resultChangeResult []bool
		numFlights         int
		flightSleepTimems  int
	}{
		{
			name:               "1 flight",
			numFlights:         5,
			flightSleepTimems:  10,
			sharingResult:      []bool{true, true, true, true, true, true},
			resultChangeResult: []bool{false, false, false, false, false, false},
		},
	}
	//        Sample case of 2 flights
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var g singleflightgroup.Group

			numGoroutines := c.numFlights
			var flights []*SpyFlight
			for i := 0; i < numGoroutines; i++ {
				sf := NewSpyFlight(i, "key-test-1-flight")
				flights = append(flights, &sf)
				go g.AddFlight(&sf, fetchData)
				time.Sleep(time.Duration(c.flightSleepTimems) * time.Millisecond)
			}

			expectSharingStatus := func(numflights int, sharingResult []bool) sharingStatus {
				expect := sharingStatus{}
				for i := 0; i < numflights; i++ {
					expect[i] = sharingResult[i]
				}
				return expect
			}(c.numFlights, c.sharingResult)
			assertSharing(t, flights, expectSharingStatus)

			// expectResultChangeStatus := func(numflights int, resultChangeResult []bool) []resultChangeStatus {
			// 	expect := []resultChangeStatus{}
			// 	for i := 0; i < numflights; i++ {
			// 		expect = append(expect, resultChangeStatus{i, resultChangeResult[i]})
			// 	}
			// 	return expect
			// }(c.numFlights, c.resultChangeResult)
			// assertResultChange(t, flights, expectResultChangeStatus)
		})
	}
}

func assertSharing(t testing.TB, flights []*SpyFlight, expectSharingStatus sharingStatus) {
	t.Helper()
	gotSharingStatus := sharingStatus{}
	for _, flight := range flights {
		gotSharingStatus[flight.flightId] = flight.shared
	}
	if !reflect.DeepEqual(gotSharingStatus, expectSharingStatus) {
		t.Errorf("Expected %v , Got %v", gotSharingStatus, expectSharingStatus)
	}
}

// func assertResultChange(t testing.TB, flights []*SpyFlight, expectResultChangeStatus []resultChangeStatus) {
// 	t.Helper()
// 	gotResultChangeStatus := []resultChangeStatus{}
// 	for i, flight := range flights {
// 		resultChange := resultChangeStatus{flight.flightId}
// 	}
// }

// This is spy Flight to record results
type SpyFlight struct {
	// This result will be used to compare expected and call results
	result       interface{}
	flightTicket string
	flightId     int
	shared       bool
}

func (sf *SpyFlight) GetTicket() string {
	return sf.flightTicket
}

func (sf *SpyFlight) GettingOff(res singleflightgroup.Result) {
	if res.Err != nil {
		sf.result = res.Val
	} else {
		sf.result = nil
	}
	sf.shared = res.Shared
}

func NewSpyFlight(id int, ticket string) SpyFlight {
	return SpyFlight{
		flightId:     id,
		flightTicket: ticket,
		result:       nil,
		shared:       false,
	}
}

func fetchData() (interface{}, error) {
	time.Sleep(100 * time.Millisecond)
	return rand.Intn(100), nil
}
