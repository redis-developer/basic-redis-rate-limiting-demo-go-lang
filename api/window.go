package api

import "time"

type Window struct {
	Expiration time.Time
	Requests   int
}

func (w *Window) Blocked(t time.Time) bool {
	if w.Expiration.Before(t) {
		w.Expiration = t.Add(time.Second * 10)
		w.Requests = 0
	}

	if w.Expiration.After(t) && w.Requests >= 10 {
		return true
	}
	return false
}
