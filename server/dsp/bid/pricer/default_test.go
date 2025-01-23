package pricer

import "testing"

func TestPricer(t *testing.T) {
	if p, err := New("dac78876796e5479fc66ca252790cb1c3cb2687a04f205acb54ccbdb8504b775",
		"a6b120682f56d2b7df82087b79bee9cd2435bbfa6f2dbf35638aeafc039350ee"); err != nil {
		t.Fatal(err)
	} else {

		if ee, err := p.Encode(1.32); err == nil {
			t.Log(ee)
			if pr, err2 := p.Decode(ee); err2 == nil {
				t.Log(pr)
			} else {
				t.Fatal(err2)

			}
		} else {
			t.Fatal(err)

		}
	}
}
