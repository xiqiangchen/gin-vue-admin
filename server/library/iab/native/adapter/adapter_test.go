package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalRequest(t *testing.T) {

	t.Run("测试反过序列化适配", func(t *testing.T) {
		req1, err := UnmarshalRequest(data1, "")
		if err != nil {
			t.Fatal(err)
		}
		req2, err := UnmarshalRequest(data2, "")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, req2.Assets[0].ID, req1.Assets[0].ID)
	})

	t.Run("测试序列化适配", func(t *testing.T) {
		req1, err := UnmarshalRequest(data1, "")
		if err != nil {
			t.Fatal(err)
		}
		req2, err := UnmarshalRequest(data2, "")
		if err != nil {
			t.Fatal(err)
		}
		adm1, err := MarshalRequest(req1, NATIVE_1_0)
		if err != nil {
			t.Fatal(err)
		}
		adm2, err := MarshalRequest(req2, NATIVE_1_2)
		if err != nil {
			t.Fatal(err)
		}
		assert.Contains(t, adm1[:30], "native")
		assert.NotContains(t, adm2[:30], "native")
	})

}

func TestUnmarshalRequest(t *testing.T) {

	t.Run("测试解析适配", func(t *testing.T) {
		req1, err := UnmarshalRequest(data1, "")
		if err != nil {
			t.Fatal(err)
		}
		req2, err := UnmarshalRequest(data2, "")
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, req2.Assets[0].ID, req1.Assets[0].ID)
	})

}

func TestUnmarshalResponse(t *testing.T) {
	//var result *request.Request

	t.Run("测试反序列化适配", func(t *testing.T) {
		resp1, err := UnmarshalResponse(respData)
		if err != nil {
			t.Fatal(err)
		}
		adm, err := MarshalResponse(resp1)
		if err != nil {
			t.Fatal(err)
		}
		assert.Contains(t, adm[:30], "native")
	})

}

var data1 = `
{
	"native":{
  "ver": "1.1",
  "context": 2,
  "contextsubtype": 20,
  "plcmttype": 11,
  "plcmtcnt": 1,
  "seq": 2,
  "assets": [
    {
      "id": 123,
      "required": 1,
      "title": {
        "len": 140
      }
    },
    {
      "id": 128,
      "required": 0,
      "img": {
        "w": 1000,
        "h": 800,
        "mimes": ["image/jpg"],
        "wmin": 836,
        "hmin": 627,
        "type": 3
      }
    },
    {
      "id": 126,
      "required": 1,
      "data": {
        "type": 1,
        "len": 25
      }
    },
    {
      "id": 127,
      "required": 1,
      "data": {
        "type": 2,
        "len": 140
      }
    },
    {
      "id": 4,
      "video": {
        "linearity": 1,
        "minduration": 15,
        "maxduration": 30,
        "protocols": [2, 3],
        "mimes": ["video/mp4"]
      }
    }
  ]
}}
`
var data2 = `{
  "ver": "1.1",
  "context": 2,
  "contextsubtype": 20,
  "plcmttype": 11,
  "plcmtcnt": 1,
  "seq": 2,
  "assets": [
    {
      "id": 123,
      "required": 1,
      "title": {
        "len": 140
      }
    },
    {
      "id": 128,
      "required": 0,
      "img": {
        "w": 1000,
        "h": 800,
        "mimes": ["image/jpg"],
        "wmin": 836,
        "hmin": 627,
        "type": 3
      }
    },
    {
      "id": 126,
      "required": 1,
      "data": {
        "type": 1,
        "len": 25
      }
    },
    {
      "id": 127,
      "required": 1,
      "data": {
        "type": 2,
        "len": 140
      }
    },
    {
      "id": 4,
      "video": {
        "linearity": 1,
        "minduration": 15,
        "maxduration": 30,
        "protocols": [2, 3],
        "mimes": ["video/mp4"]
      }
    }
  ]
}
`

var respData = `{
    "native": {
        "ver": "1.2",
        "assets": null,
        "assetsurl": "",
        "dcourl": "",
        "link": {
            "url": "lazada://ph/web/www/marketing/gateway/rta?null&dsource=sml&exlaz=e_EOkf+u3yftnGip8qo24MCZFmNCmP6gU/DAlXOlSwx0HCXvTKdNAtBQ/9joiMivH8N0V5K9iWbnm/vxXbJRG+Gpt3vboX/T/sD3WnASIas1Y=&rta_token=&rta_event_id=chb2i14a1lcvn8qsdgig_861002&os=android&gps_adid=89859846-203d-4d73-97fd-81300e458970&android_id=__android_id__&idfa=&idfv=__idfv__&bundle_id=com.bstar.intl&device_model=TECNO+LE7&device_make=TECNO&sub_id1=&sub_id2=67700080004877&sub_id3=chb2i14a1lcvn8qsdgig_861002&sub_id4=dsp&sub_id5=7fbe5ba6-4013-4d52-ba78-2d6ed2880de7&lzdcid=__lzdcid__&adtype=__adtype__",
            "clicktrackers": [
                "https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9edqznZEuLcXtGboH4VjFW_ak3-voWyLPvq97vMzYieaLmr-pIWeXAs2iuw4tYleFr_JXuRLG3dH3XxhaeOFbCmdByF391fRiq8FPIe64Fee"
            ],
            "fallback": "https://c.lazada.com.ph/t/c.YJPqgU?rta_token=&rta_event_id=chb2i14a1lcvn8qsdgig_861002&os=android&gps_adid=89859846-203d-4d73-97fd-81300e458970&imei=__imei__&android_id=__android_id__&idfa=&idfv=__idfv__&bundle_id=com.bstar.intl&device_model=TECNO+LE7&device_make=TECNO&sub_id1=appicplay&sub_id2=67700080004877&sub_id3=chb2i14a1lcvn8qsdgig_861002&sub_id4=dsp&sub_id5=7fbe5ba6-4013-4d52-ba78-2d6ed2880de7&lzdcid=__lzdcid__&adtype=__adtype__&lazada_randomcid={lazada_randomcid}"
        },
        "imptrackers": [
            "https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9f4ZyPvv_bdeIV304Fo7DEGyVsh0fyAV3rdcfUIEQYKw7NZ1ZmFc90zcFL26iKU6XSRe3OsAkBi6hNbNGIWbi6wIZFnOV0BXRIuaYg5tx4MG?price=${AUCTION_PRICE}"
        ],
        "jstracker": "",
        "eventtrackers": [
            {
                "event": 1,
                "method": 1,
                "url": "https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9f4ZyPvv_bdeIV304Fo7DEGyVsh0fyAV3rdcfUIEQYKw7NZ1ZmFc90zcFL26iKU6XSRe3OsAkBi6hNbNGIWbi6wIZFnOV0BXRIuaYg5tx4MG?price=${AUCTION_PRICE}"
            }
        ],
        "privacy": ""
    }
}`

var test1 = "{\"native\":{\"ver\":\"1.1\",\"plcmtcnt\":1,\"assets\":[{\"id\":100,\"required\":1,\"title\":{\"len\":90}},{\"id\":201,\"required\":1,\"img\":{\"type\":1,\"w\":150,\"h\":150}},{\"id\":203,\"required\":1,\"img\":{\"type\":3,\"w\":1200,\"wmin\":150,\"h\":627,\"hmin\":150}},{\"id\":401,\"required\":0,\"data\":{\"type\":1}},{\"id\":402,\"required\":1,\"data\":{\"type\":2}},{\"id\":403,\"required\":0,\"data\":{\"type\":3}},{\"id\":404,\"required\":0,\"data\":{\"type\":4}},{\"id\":412,\"required\":1,\"data\":{\"type\":12}}]}}"
