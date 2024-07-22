package openrtb_v2_6

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSource(t *testing.T) {
	var subject *Source
	if err := fixture("source", &subject); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	exp := &Source{
		FinalSaleDecision: 1,
		TransactionID:     "transaction-id",
		PaymentChain:      "payment-chain",
		Ext:               json.RawMessage("{}"),
	}
	if got := subject; !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %+v, got %+v", exp, got)
	}
}
