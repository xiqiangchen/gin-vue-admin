package cml

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ConsentManagementList struct {
	LastUpdated string
	Cmps        map[string]ConsentManagement

	BidexLastSyn int64 `json:"-"`
}
type ConsentManagement struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	IsCommercial bool     `json:"isCommercial"`
	Environments []string `json:"environments"`
}

// 每周四 5:00 PM 发布一次，官方建议只缓存一周
var CML *ConsentManagementList
var CMLUrl string = `https://cmplist.consensu.org/v2/cmp-list.json`

// todo 是否要设置一个本地的保底？
func InitCML() (err error) {

	// todo domark 第一次初始化，这里启动依赖网络，考虑海外地域网络不一定良好，后续是否做成默认读取一个本地的配置文件比较好
	newCml := getLastCML()
	if newCml != nil {
		newCml.BidexLastSyn = time.Now().Unix()
		CML = newCml
	}

	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		for {
			<-ticker.C
			update()
		}
	}()
	return nil
}
func update() (err error) {
	location, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		log.Println("无法加载时区：", err)
		return err
	}
	eut := time.Now().In(location)

	if eut.Weekday() == time.Thursday && eut.Hour() == 17 {
		newCml := getLastCML()
		if newCml != nil {
			newCml.BidexLastSyn = time.Now().Unix()
			CML = newCml
		}
	}
	return nil
}
func getLastCML() *ConsentManagementList {
	resp, err := http.DefaultClient.Get(CMLUrl)
	if err != nil {
		log.Println(err)
		return nil
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	newCML := new(ConsentManagementList)
	err = json.Unmarshal(respBody, newCML)
	if err != nil {
		log.Println(err)
		return nil
	}
	return newCML
}
