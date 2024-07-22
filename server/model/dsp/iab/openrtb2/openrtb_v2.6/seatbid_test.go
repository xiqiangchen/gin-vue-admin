package openrtb_v2_6

import (
	"errors"
	"testing"
)

func TestSeatBid_Validate(t *testing.T) {
	subject := &SeatBid{}
	if exp, got := ErrInvalidSeatBidBid, subject.Validate(); !errors.Is(exp, got) {
		t.Fatalf("expected %v, got %v", exp, got)
	}
}
