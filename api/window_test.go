package api

import (
	"testing"
	"time"
)

func TestWindow_Blocked(t *testing.T) {


	window := &Window{
		Expiration: time.Now().Add(time.Second * 10),
	}

	for i := 0; i < 11; i++ {

		now := time.Now()
		if blocked := window.Blocked(now); blocked && i != 10 {
			t.Errorf("not expected window blocking at %d", i)
		} else if blocked == false && i == 10 {
			t.Error("expected window blocking at 10")
		}
		window.Requests++
	}

}
