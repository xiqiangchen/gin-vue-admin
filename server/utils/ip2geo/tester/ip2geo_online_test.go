package tester

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo/csv"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var token = "938501824e4611ef807700163e167ffb"

type QueryResponse struct {
	Code int          `json:"code"`
	Data ResponseData `json:"data"`
	Msg  string       `json:"msg"`
}

type ResponseData struct {
	Location QueryData `json:"location"`
}

type QueryData struct {
	AreaCode       string        `json:"area_code"`
	City           string        `json:"city"`
	CityCode       string        `json:"city_code"`
	Continent      string        `json:"continent"`
	Country        string        `json:"country"`
	CountryCode    string        `json:"country_code"`
	District       string        `json:"district"`
	Elevation      string        `json:"elevation"`
	IP             string        `json:"ip"`
	Isp            string        `json:"isp"`
	Latitude       string        `json:"latitude"`
	Longitude      string        `json:"longitude"`
	MultiStreet    []MultiStreet `json:"multi_street"`
	Province       string        `json:"province"`
	Street         string        `json:"street"`
	TimeZone       string        `json:"time_zone"`
	WeatherStation string        `json:"weather_station"`
	ZipCode        string        `json:"zip_code"`
}

type MultiStreet struct {
	Lng          string `json:"lng"`
	Lat          string `json:"lat"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Street       string `json:"street"`
	StreetNumber string `json:"street_number"`
	Radius       string `json:"radius"`
	ZipCode      string `json:"zip_code"`
}

func TestIPDataAPILongConnection(t *testing.T) {
	transport := &http.Transport{
		MaxIdleConns:        10,               // 最大空闲连接数
		IdleConnTimeout:     50 * time.Second, // 空闲连接的超时时间
		MaxIdleConnsPerHost: 5,                // 每个主机的最大空闲连接数
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   50 * time.Second, // 请求超时时间
	}

	list, err := csv.GetAdxIpList(csv.Ipv4, 1)
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, record := range list {
		req, err := http.NewRequest("GET", "https://api.ipdatacloud.com/v2/query?ip="+record.IP+"&key="+token, nil)
		if err != nil {
			t.Fatalf("failed to create request: %v", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("failed to send request: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("expected status 200, got %v", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read response body: %v", err)
		}
		var data QueryResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			t.Fatalf("failed to unmarshal response body: %v", err)
		}
		if data.Code != http.StatusOK {
			t.Errorf("unexpected code: got %v, want %v", data.Code, "200")
		}
		t.Logf("%#v \n", data.Data.Location)
		resp.Body.Close()
	}

}
