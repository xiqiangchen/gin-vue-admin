package trans

// IpTrans 洲|国家|国家英文|洲/省|洲/省英文（中国的是拼音）|城市|城市英文（中国的是拼音）|经度|纬度|国家缩写|邮政编码|互联网服务提供商|行政区码|时区|省份ID
type IpTrans struct {
	Continent     string  `json:"continent"`       //洲
	Country       string  `json:"country"`         //国家
	CountryEn     string  `json:"country_en"`      //国家英文
	Province      string  `json:"province"`        //省份
	ProvinceEn    string  `json:"province_en"`     //省份
	City          string  `json:"city"`            //城市
	CityEn        string  `json:"city_en"`         //城市英文（中国的是拼音）
	Longitude     float64 `json:"longitude"`       //经度
	Latitude      float64 `json:"latitude"`        //纬度
	CountryCode   string  `json:"country_code"`    //国家缩写
	ZipCode       string  `json:"zip_code"`        //邮编
	Carrier       string  `json:"carrier"`         //互联网服务提供商
	LbsCityId     string  `json:"lbs_city_id"`     //地址表的ID - 市级
	TimeZone      string  `json:"time_zone"`       //时区
	LbsProvinceId string  `json:"lbs_province_id"` //地址表的ID - 省级
}
