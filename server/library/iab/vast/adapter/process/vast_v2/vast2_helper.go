package vast_v2

import (
	"encoding/xml"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/vast/adapter/base"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/vast_v2"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"strings"
)

type VAST2Helper struct {
	base.Common
}

func (h *VAST2Helper) GetTrackInfo(adm string) (vi *base.VASTInfo, err error) {
	vi = new(base.VASTInfo)
	var vast = new(vast_v2.VAST)
	err = xml.Unmarshal([]byte(adm), vast)
	if err != nil {
		// todo domark
		return nil, err
	}
	getTrack(vi, vast)
	for i := range vi.ClickTrack {
		vi.ClickTrack[i] = strings.TrimSpace(vi.ClickTrack[i])
	}
	for i := range vi.ImpTrack {
		vi.ImpTrack[i] = strings.TrimSpace(vi.ImpTrack[i])
	}
	return vi, nil
}
func (h *VAST2Helper) SetTrackInfo(adm string, vi *base.VASTInfo) (result string, err error) {
	var vast = new(vast_v2.VAST)
	err = xml.Unmarshal([]byte(adm), vast)
	if err != nil {
		// todo domark
		return adm, err
	}
	setTrack(vi, vast)

	//data, err := xml.Marshal(vi)
	//return string(data), nil

	return vast.ToString()
}
func (h *VAST2Helper) GetProtcol() []int {
	return []int{base.VAST20, base.VAST20Wrapper}
}

func getTrack(vi *base.VASTInfo, v *vast_v2.VAST) {
	if len(v.Ads) == 0 || v.Ads[0].InLine == nil {
		return
	}

	if v.Ads[0].InLine != nil {
		inline := v.Ads[0].InLine
		// 填充曝光监测
		for _, imp := range inline.Impressions {
			vi.ImpTrack = append(vi.ImpTrack, imp.URI)
		}

		// 填充点击监测
		for _, crtv := range inline.Creatives {
			if crtv.Linear != nil {
				if crtv.Linear.VideoClicks != nil {
					for _, ctrack := range crtv.Linear.VideoClicks.ClickTrackings {
						vi.ClickTrack = append(vi.ClickTrack, ctrack.URI)
					}
					for _, click := range crtv.Linear.VideoClicks.ClickThroughs {
						vi.ClickThrough = append(vi.ClickThrough, click.URI)
					}
				}
			}
			if crtv.NonLinearAds != nil {
				for _, nonLinear := range crtv.NonLinearAds.NonLinears {
					for _, ctrack := range nonLinear.NonLinearClickTrackings {
						vi.ClickTrack = append(vi.ClickTrack, ctrack.URI)
					}
					if nonLinear.NonLinearClickThrough != nil {
						vi.ClickThrough = []string{nonLinear.NonLinearClickThrough.CDATA}
					}
				}
			}
		}
	}

	if v.Ads[0].Wrapper != nil {
		wrapper := v.Ads[0].Wrapper
		// 填充曝光监测
		for _, imp := range wrapper.Impressions {
			vi.ImpTrack = append(vi.ImpTrack, imp.URI)
		}
		for _, crtv := range wrapper.Creatives {
			if crtv.Linear != nil {
				if crtv.Linear.VideoClicks != nil {
					for _, ctrack := range crtv.Linear.VideoClicks.ClickTrackings {
						vi.ClickTrack = append(vi.ClickTrack, ctrack.URI)
					}
					for _, click := range crtv.Linear.VideoClicks.ClickThroughs {
						vi.ClickThrough = append(vi.ClickThrough, click.URI)
					}
				}
			}
			if crtv.NonLinearAds != nil {
				for _, nonLinear := range crtv.NonLinearAds.NonLinears {
					for _, ctrack := range nonLinear.NonLinearClickTracking {
						vi.ClickThrough = append(vi.ClickThrough, ctrack.CDATA)
					}
				}
			}
		}
	}
}
func setTrack(vi *base.VASTInfo, v *vast_v2.VAST) {
	if len(v.Ads) == 0 || v.Ads[0].InLine == nil {
		return
	}

	if v.Ads[0].InLine != nil {
		inline := v.Ads[0].InLine
		// 填充曝光监测
		for _, imp := range vi.ImpTrack {
			inline.Impressions = append(inline.Impressions, vast_v2.Impression{
				ID:  utils.MD5(imp),
				URI: imp,
			})
		}

		// 填充点击监测和事件监测
		for _, crtv := range inline.Creatives {
			if crtv.Linear != nil {
				// 填充点击监测
				if crtv.Linear.VideoClicks != nil {
					for _, ctrack := range vi.ClickTrack {
						crtv.Linear.VideoClicks.ClickTrackings = append(crtv.Linear.VideoClicks.ClickTrackings, vast_v2.VideoClick{
							ID:  utils.MD5(ctrack),
							URI: ctrack,
						})
					}
					/// todo 是否处理302点击?
				}
				// 填充事件监测
				for eventName, urls := range vi.EventTrack {
					for _, url := range urls {
						crtv.Linear.TrackingEvents = append(crtv.Linear.TrackingEvents, vast_v2.Tracking{
							Event: eventName,
							//Offset: nil,
							URI: url,
						})
					}
				}
			}
			if crtv.NonLinearAds != nil {
				// 填充点击监测
				for _, nonLinear := range crtv.NonLinearAds.NonLinears {
					for _, click := range vi.ClickTrack {
						nonLinear.NonLinearClickTrackings = append(nonLinear.NonLinearClickTrackings, vast_v2.NonLinearClickTracking{
							ID:  utils.MD5(click),
							URI: click,
						})
					}
					/// todo 是否处理302点击?
				}
				// 填充事件监测
				for eventName, urls := range vi.EventTrack {
					for _, url := range urls {
						crtv.NonLinearAds.TrackingEvents = append(crtv.NonLinearAds.TrackingEvents, vast_v2.Tracking{
							Event: eventName,
							//Offset: nil,
							URI: url,
						})
					}
				}
			}
		}
	}

	if v.Ads[0].Wrapper != nil {
		wrapper := v.Ads[0].Wrapper
		// 填充曝光监测
		for _, imp := range vi.ImpTrack {
			wrapper.Impressions = append(wrapper.Impressions, vast_v2.Impression{
				ID:  utils.MD5(imp),
				URI: imp,
			})
		}
		for _, crtv := range wrapper.Creatives {
			if crtv.Linear != nil {
				if crtv.Linear.VideoClicks != nil {
					for _, ctrack := range vi.ClickTrack {
						crtv.Linear.VideoClicks.ClickTrackings = append(crtv.Linear.VideoClicks.ClickTrackings, vast_v2.VideoClick{
							ID:  utils.MD5(ctrack),
							URI: ctrack,
						})
					}
				}
				// 填充事件监测
				for eventName, urls := range vi.EventTrack {
					for _, url := range urls {
						crtv.Linear.TrackingEvents = append(crtv.Linear.TrackingEvents, vast_v2.Tracking{
							Event: eventName,
							//Offset: nil,
							URI: url,
						})
					}
				}
			}
			if crtv.NonLinearAds != nil {
				for _, nonLinear := range crtv.NonLinearAds.NonLinears {
					for _, ctrack := range vi.ClickTrack {
						nonLinear.NonLinearClickTracking = append(nonLinear.NonLinearClickTracking, vast_v2.CDATAString{
							CDATA: ctrack,
						})
					}
				}
				// 填充事件监测
				for eventName, urls := range vi.EventTrack {
					for _, url := range urls {
						crtv.NonLinearAds.TrackingEvents = append(crtv.NonLinearAds.TrackingEvents, vast_v2.Tracking{
							Event: eventName,
							//Offset: nil,
							URI: url,
						})
					}
				}
			}
		}
	}
}
