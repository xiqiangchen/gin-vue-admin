package native_tpl

import "encoding/json"

const (
	Request_Native_Asset_Image_Main = 3
)

type Request struct {
	Native Request_Native `json:"native,omitempty"`
}
type Request_Native struct {
	Ver string `json:"ver,omitempty"`
	//Layout int                    `json:"layout,omitempty"`
	//Adunit int                    `json:"adunit,omitempty"`
	Assets []Request_Native_Asset `json:"assets,omitempty"`
}

type Request_Native_Asset struct {
	Id       int                         `json:"id"`
	Required int                         `json:"required,omitempty"`
	Image    *Request_Native_Asset_Image `json:"img,omitempty"`
	Title    *Request_Native_Asset_Title `json:"title,omitempty"`
	Data     *Request_Native_Asset_Data  `json:"data,omitempty"`
	Video    *Request_Native_Asset_Video `json:"video,omitempty"`
	Ext      json.RawMessage             `json:"ext,omitempty"`
}

type Request_Native_Asset_Image struct {
	Type int `json:"type,omitempty"`
	W    int `json:"w,omitempty"`
	H    int `json:"h,omitempty"`
	Wmin int `json:"wmin,omitempty"`
	Hmin int `json:"hmin,omitempty"`
	//WRatio int      `json:"wratio,omitempty"`
	//HRatio int      `json:"hratio,omitempty"`
	Mimes []string `json:"mimes,omitempty"`
}

type Request_Native_Asset_Video struct {
	Minduration int      `json:"minduration"`
	Maxduration int      `json:"maxduration"`
	Protocols   []int    `json:"protocols"`
	Mimes       []string `json:"mimes"`
}

type Request_Native_Asset_Title struct {
	Len int32 `json:"len"`
}

type Request_Native_Asset_Data struct {
	Type int `json:"type"`
	Len  int `json:"len,omitempty"`
}

// response
type Response struct {
	Native Response_Native `json:"native,omitempty"`
}
type Response_Native struct {
	Ver         string                  `json:"ver,omitempty"`
	Assets      []Response_Native_Asset `json:"assets"`
	Link        Response_Native_Link    `json:"link"`
	Imptrackers []string                `json:"imptrackers,omitempty"`
	Jstracker   string                  `json:"jstracker,omitempty"`
}

type Response_Native_Asset struct {
	Id       int                          `json:"id"`
	Required int                          `json:"required,omitempty"`
	Title    *Response_Native_Asset_Title `json:"title,omitempty"`
	Image    *Response_Native_Asset_Image `json:"img,omitempty"`
	Data     *Response_Native_Asset_Data  `json:"data,omitempty"`
	Video    *Response_Native_Asset_Video `json:"video,omitempty"`
	Link     *Response_Native_Link        `json:"link,omitempty"`
}
type Response_Native_Asset_Title struct {
	Text string `json:"text"`
}
type Response_Native_Asset_Image struct {
	Url string `json:"url,omitempty"`
	W   int    `json:"w,omitempty"`
	H   int    `json:"h,omitempty"`
}
type Response_Native_Asset_Data struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value"`
}
type Response_Native_Asset_Video struct {
	Vasttag string `json:"vasttag"`
}
type Response_Native_Link struct {
	Url           string   `json:"url"`
	Clicktrackers []string `json:"clicktrackers,omitempty"`
	Fallback      string   `json:"fallback,omitempty"`
}
