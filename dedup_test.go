package dedup

import (
	"testing"
)

func TestNilKnown(t *testing.T) {
	var s Seen
	if s.Known("nothing") != false {
		t.Error("got true, want false")
	}
}

func TestNilVisited(t *testing.T) {
	var s Seen
	if res := s.VisitedKeys(); res != nil {
		t.Errorf("got %+v, want nil", res)
	} else {
		t.Logf("got %+v", res)
	}
}
