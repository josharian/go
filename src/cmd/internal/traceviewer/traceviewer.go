// Package traceviewer contains data types and helpers for the catapult trace viewer.
// https://github.com/catapult-project/catapult
package traceviewer

import "time"

// Data is the top level trace viewer object.
type Data struct {
	Events   []*Event         `json:"traceEvents"`
	Frames   map[string]Frame `json:"stackFrames"`
	TimeUnit string           `json:"displayTimeUnit"`

	// This is where mandatory part of the trace starts (e.g. thread names)
	Footer int
}

// Event is a single trace viewer event.
type Event struct {
	Name     string      `json:"name,omitempty"`
	Category string      `json:"cat,omitempty"`
	Phase    string      `json:"ph"`
	Scope    string      `json:"s,omitempty"`
	Time     float64     `json:"ts"`
	Dur      float64     `json:"dur,omitempty"`
	Pid      uint64      `json:"pid"`
	Tid      uint64      `json:"tid"`
	ID       uint64      `json:"id,omitempty"`
	Stack    int         `json:"sf,omitempty"`
	EndStack int         `json:"esf,omitempty"`
	Arg      interface{} `json:"args,omitempty"`
}

// Frame is a trace viewer frame.
type Frame struct {
	Name   string `json:"name"`
	Parent int    `json:"parent,omitempty"`
}

type NameArg struct {
	Name string `json:"name"`
}

type SortIndexArg struct {
	Index int `json:"sort_index"`
}

func Elapsed(start, stop time.Time) float64 {
	return Microseconds(stop.Sub(start))
}

func Microseconds(dur time.Duration) float64 {
	return float64(dur) / float64(time.Microsecond)
}
