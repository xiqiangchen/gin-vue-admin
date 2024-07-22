package dsa

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
)

func GetDsa(ext map[string]json.RawMessage) {

}

func SetDsaDefault(ext map[string]json.RawMessage) {
	//, behalf, paid, domain string, dsaparams []int

	if ext == nil {
		ext = make(map[string]json.RawMessage, 1)
	}

	var dsa *DSAResponse

	d, ok := ext["dsa"]
	if !ok {
		dsa = new(DSAResponse)
	} else {
		err := json.Unmarshal(d, &dsa)
		if err != nil {
			return
		}
	}

	dsa.Transparency = append(dsa.Transparency, Transparency{
		Domain:    constant.DSADomain,
		Dsaparams: constant.TCFPurpose,
	})

	// todo 判断空，填充
	if len(dsa.Behalf) == 0 {
	}
	if len(dsa.Paid) == 0 {
		dsa.Paid = dsa.Behalf
	}

	newdsa, _ := json.Marshal(dsa)
	ext["dsa"] = newdsa
}
