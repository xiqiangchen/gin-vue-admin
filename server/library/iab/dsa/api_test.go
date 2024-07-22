package dsa

import (
	"encoding/json"
	"testing"
)

func Test_GetDsa(t *testing.T) {
	var err error

	var reqExt map[string]json.RawMessage
	err = json.Unmarshal([]byte(reqData), &reqExt)
	if err != nil {
		t.Fatal(err)
	}
	GetDsa(reqExt)

	var respExt map[string]json.RawMessage
	err = json.Unmarshal([]byte(respData), &respExt)
	if err != nil {
		t.Fatal(err)
	}
	SetDsaDefault(respExt)

	t.Log(string(respExt["dsa"]))

	var respExt2 map[string]interface{}
	err = json.Unmarshal([]byte(respData), &respExt2)
	if err != nil {
		t.Fatal(err)
	}
	respExt2["dsa2"] = DSAResponse{
		Behalf:       "test",
		Paid:         "",
		Transparency: nil,
		Adrender:     0,
	}

	a, _ := json.Marshal(respExt2)
	t.Log(string(a))

}

var respData = `{
   			 "dsa": {
   				 "behalf": "Advertiser",
   				 "paid": "Advertiser",
   				 "transparency": [{
   					 "domain": "dsp1domain.com",
   					 "dsaparams": [1,2]
   				 }],
   				 "adrender": 1
   			 }
   		 }`
var reqData = `{
            "dsa": {
                "dsarequired": 3, 
                "pubrender": 0,
                "datatopub": 2,
                "transparency": [{
                    "domain": "platform1domain.com",
                    "dsaparams": [1]},
                    {"domain": "SSP2domain.com",
                    "dsaparams": [1,2]
                    }]
            }
        }`
