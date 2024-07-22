package openrtb_v2_6_proto

import (
	"encoding/json"
	"github.com/gogo/protobuf/proto"
	"testing"
)

func Test_OpenRtb(t *testing.T) {
	ortbReq := new(BidRequest)
	err := json.Unmarshal([]byte(reqNativeDemo), ortbReq)
	if err != nil {
		t.Error(err)
	}
	// req 有问题。
	//1。js字段
	// 2。clickbrowser
	// 3。secure 字段
	t.Log(ortbReq.GetId())

	//////////////////////////
	ortbResp := new(BidResponse)
	err = json.Unmarshal([]byte(respNativeDemo), ortbResp)
	if err != nil {
		t.Error(err)
	}
	t.Log(ortbReq.GetImp()[0].GetSecure())

	var a interface{}
	a = ortbReq
	data, err := json.Marshal(a)
	t.Log(string(data))

	pdata, err := proto.Marshal(ortbResp)
	if err != nil {
		t.Error(err)
	}
	t.Log(pdata)

	newortbResp := new(BidResponse)
	err = proto.Unmarshal(pdata, newortbResp)
	if err != nil {
		t.Error(err)
	}
	t.Log(newortbResp)

	switch a.(type) {
	case *BidRequest:
		t.Log("bidreq")
	default:
		t.Log("default")
	}

}

var respNativeDemo = `{
	"id": "a258c668-ff5c-b84f-ac57-ebac9ead5199",
	"seatbid": [
		{
			"bid": [
				{
					"id": "chb2i14a1lcvn8qsdgig",
					"impid": "1",
					"price": 0.76,
					"nurl": "https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9RheSvKztJytK1nov2hmb3YbJ0OYVz4uFJ4glkvNIKtUEo0nz6vkNt1d4KkSTBn1X_CpIUQVw2vyd5yJKbEbsD9dLOlXOIpoXA3EqEE_o1hb?price=${AUCTION_PRICE}",
					"adm": "{\"native\":{\"ver\":\"1.2\",\"assets\":null,\"assetsurl\":\"\",\"dcourl\":\"\",\"link\":{\"url\":\"lazada://ph/web/www/marketing/gateway/rta?null\u0026dsource=sml\u0026exlaz=e_EOkf%252Bu3yftnGip8qo24MCZFmNCmP6gU%252FDAlXOlSwx0HCXvTKdNAtBQ%252F9joiMivH8N0V5K9iWbnm%252FvxXbJRG%252BGpt3vboX%252FT%252FsD3WnASIas1Y%253D\u0026rta_token=\u0026rta_event_id=chb2i14a1lcvn8qsdgig_861002\u0026os=android\u0026gps_adid=89859846-203d-4d73-97fd-81300e458970\u0026android_id=__android_id__\u0026idfa=\u0026idfv=__idfv__\u0026bundle_id=com.bstar.intl\u0026device_model=TECNO+LE7\u0026device_make=TECNO\u0026sub_id1=\u0026sub_id2=67700080004877\u0026sub_id3=chb2i14a1lcvn8qsdgig_861002\u0026sub_id4=dsp\u0026sub_id5=7fbe5ba6-4013-4d52-ba78-2d6ed2880de7\u0026lzdcid=__lzdcid__\u0026adtype=__adtype__\",\"clicktrackers\":[\"https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9edqznZEuLcXtGboH4VjFW_ak3-voWyLPvq97vMzYieaLmr-pIWeXAs2iuw4tYleFr_JXuRLG3dH3XxhaeOFbCmdByF391fRiq8FPIe64Fee\"],\"fallback\":\"https://c.lazada.com.ph/t/c.YJPqgU?rta_token=\u0026rta_event_id=chb2i14a1lcvn8qsdgig_861002\u0026os=android\u0026gps_adid=89859846-203d-4d73-97fd-81300e458970\u0026imei=__imei__\u0026android_id=__android_id__\u0026idfa=\u0026idfv=__idfv__\u0026bundle_id=com.bstar.intl\u0026device_model=TECNO+LE7\u0026device_make=TECNO\u0026sub_id1=appicplay\u0026sub_id2=67700080004877\u0026sub_id3=chb2i14a1lcvn8qsdgig_861002\u0026sub_id4=dsp\u0026sub_id5=7fbe5ba6-4013-4d52-ba78-2d6ed2880de7\u0026lzdcid=__lzdcid__\u0026adtype=__adtype__\u0026lazada_randomcid={lazada_randomcid}\"},\"imptrackers\":[\"https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9f4ZyPvv_bdeIV304Fo7DEGyVsh0fyAV3rdcfUIEQYKw7NZ1ZmFc90zcFL26iKU6XSRe3OsAkBi6hNbNGIWbi6wIZFnOV0BXRIuaYg5tx4MG?price=${AUCTION_PRICE}\"],\"jstracker\":\"\",\"eventtrackers\":[{\"event\":1,\"method\":1,\"url\":\"https://dsp.atomhike-en.com/tracking/v2/duMy41bm6VNxGRSRNaHn9f4ZyPvv_bdeIV304Fo7DEGyVsh0fyAV3rdcfUIEQYKw7NZ1ZmFc90zcFL26iKU6XSRe3OsAkBi6hNbNGIWbi6wIZFnOV0BXRIuaYg5tx4MG?price=${AUCTION_PRICE}\"}],\"privacy\":\"\"}}",
					"adid": "chb2i14a1lcvn8qsdgig",
					"adomain": [
						""
					],
					"bundle": "com.lazada.android",
					"crid": "0be6da83-cd2f-4039-a4fa-6bdfb1bcdff6",
					"language": "en",
					"w": 1200,
					"h": 627
				}
			],
			"seat": "bukas"
		}
	],
	"bidid": "chb2i14a1lcvn8qsdgig",
	"cur": "USD"
}`
var reqNativeDemo = `{
  "id": "a258c668-ff5c-b84f-ac57-ebac9ead5199",
  "imp": [
    {
      "id": "1",
      "native": {
        "request": "{\"native\":{\"ver\":\"1.1\",\"plcmtcnt\":1,\"assets\":[{\"id\":100,\"required\":1,\"title\":{\"len\":90}},{\"id\":201,\"required\":1,\"img\":{\"type\":1,\"w\":150,\"h\":150}},{\"id\":203,\"required\":1,\"img\":{\"type\":3,\"w\":1200,\"wmin\":150,\"h\":627,\"hmin\":150}},{\"id\":401,\"required\":0,\"data\":{\"type\":1}},{\"id\":402,\"required\":1,\"data\":{\"type\":2}},{\"id\":403,\"required\":0,\"data\":{\"type\":3}},{\"id\":404,\"required\":0,\"data\":{\"type\":4}},{\"id\":412,\"required\":1,\"data\":{\"type\":12}}]}}",
        "ver": "1.1",
        "api": [
          3,
          5,
          6
        ],
        "battr": [
          10
        ]
      },
      "displaymanager": "TradPlus",
      "displaymanagerver": "1.0",
      "tagid": "1142904122",
      "bidfloor": 0.02,
      "bidfloorcur": "USD",
      "clickbrowser": true,
      "secure": false,
      "exp": 10800,
      "ext": {
        "deeplink": 1
      }
    }
  ],
  "app": {
    "id": "C26FF28457E410EF0A206A020EA66591",
    "name": "Bilibili",
    "bundle": "com.bstar.intl",
    "storeurl": "https://play.google.com/store/apps/details?id=com.bstar.intl",
    "ver": "2.34.0",
    "ext": {
      "orientation": 1
    }
  },
  "device": {
    "ua": "Dalvik/2.1.0 (Linux; U; Android 11; TECNO LE7 Build/RP1A.200720.011)",
    "geo": {
      "lat": 9.0227,
      "lon": 125.1791,
      "type": 2,
      "ipservice": 1,
      "country": "PHL"
    },
    "ip": "103.236.176.90",
    "devicetype": 4,
    "make": "TECNO",
    "model": "TECNO LE7",
    "os": "Android",
    "osv": "11",
    "hwv": "mt6769",

    "h": 2208,
    "w": 1080,
    "ppi": 6,
    "pxratio": 3,
    "js": false,
    "language": "en",
    "connectiontype": 2,
    "ifa": "89859846-203d-4d73-97fd-81300e458970",
    "ext": {
      "timezone": "Asia/Manila",
      "gaid": "89859846-203d-4d73-97fd-81300e458970"
    }
  },
  "user": {
    "id": "UID-fcbaf2e7-3974-4a78-b225-8ee8214adbad"
  },
  "at": 1,
  "tmax": 1500,
  "cur": [
    "USD"
  ],
  "bcat": [
    "IAB26-4",
    "IAB26-3",
    "IAB26-2",
    "IAB26-1",
    "IAB25-3",
    "IAB25-2",
    "IAB25-1",
    "IAB25-4",
    "IAB25-5",
    "IAB25-7",
    "IAB25-6",
    "IAB23-1",
    "IAB23-2",
    "IAB23-6",
    "IAB23-8",
    "IAB23-9",
    "IAB23-10",
    "IAB14-8",
    "IAB18-2",
    "IAB24"
  ],
  "badv": [
    ""
  ],
  "bapp": [
    "com.iqiyi.i18n",
    "com.iqiyi.i18n.tv",
    "com.viu.phone",
    "com.viu.tv",
    "com.viu.pad",
    "com.tencent.qqlivei18n",
    "com.youku.international.phone",
    "com.hunantv.imgo.activity.inter",
    "com.mgtv.tv.intl",
    "com.netflix.mediaclient",
    "com.netflix.ninja",
    "com.disney.disneyplus",
    "com.google.android.youtube",
    "com.google.android.youtube.tv",
    "com.ss.android.ugc.trill"
  ],
  "source": {
    "ext": {
      "schain": {
        "complete": 1,
        "ver": "1.0",
        "nodes": [
          {
            "asi": "tradplusad.com",
            "sid": "93882-10001",
            "rid": "a258c668-ff5c-b84f-ac57-ebac9ead5199",
            "hp": 1
          }
        ]
      }
    }
  },
  "regs": {
    "ext": {
      "gdpr": 0,
      "ccpa": 0
    }
  },
  "ext": {}
}
`
