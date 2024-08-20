package dsp

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	dbid "github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	bid_adapter "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid/protocol/bsw"
	protocol "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/shopspring/decimal"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BidService struct {
}

func (bidService *BidService) SendMsg(msg []byte) {
	utils.SendMsg(global.GVA_KAFKA_PRODUCER, global.GVA_CONFIG.Dsp.Bid.Topic, msg)
}

// func (bidService *BidService) Bid(req *bid_adapter.BidRequest, c *gin.Context) (*bid_adapter.BidResponse, bool) {
func (bidService *BidService) Bid(req *protocol.BidRequest, c *gin.Context) (*protocol.BidResponse, bool) {

	var campaigns []*ad.Campaign
	if campaigns = filters(req); len(campaigns) == 0 {
		return nil, false
	}
	if winCampaign, weight := selectCampaign(campaigns); winCampaign != nil {
		if resp, offer, err := offerByCampaign(req, winCampaign, weight); err != nil {
			global.GVA_LOG.Error("出价活动转换bidResp协议失败", zap.Error(err))
		} else if offer {
			return resp, true
		} else {
			return nil, false
		}
	}
	return nil, false

}

// 根据活动填充出价响应
// func offerByCampaign(req *bid_adapter.BidRequest, campaign *ad.Campaign) (resp *bid_adapter.BidResponse, err error) {
func offerByCampaign(req *protocol.BidRequest, campaign *ad.Campaign, weight float64) (resp *protocol.BidResponse, offer bool, err error) {

	bids := make([]protocol.Bid, 0, len(req.Impressions))
	// 暂时只支持对第一个曝光进行处理
	for _, imp := range req.Impressions {
		if imp.BidFloor > campaign.GetBidPrice() || (rand.Float64() > (campaign.GetBidRate() / weight / 100)) {
			continue
		}
		_bid, err := getRespBid(req.ID, req, imp, campaign)
		if err != nil {
			continue
		}

		bids = append(bids, _bid)

		break
	}

	if len(bids) == 0 {
		offer = false
		//err = errors.New("offerByCampaign异常，没找到匹配的素材")
		return
	}

	resp = &protocol.BidResponse{
		ID:       req.ID,
		BidID:    req.ID,
		Currency: "USD",
		SeatBids: []protocol.SeatBid{
			{
				Seat: fmt.Sprintf("seat-%d", campaign.CreatedBy),
				Bids: bids,
			},
		},
	}
	offer = true
	return
}

// 选择其中一个活动作为出价
func selectCampaign(campaigns []*ad.Campaign) (*ad.Campaign, float64) {
	var totalRate float64
	weights := make([]float64, 0, len(campaigns))
	for _, c := range campaigns {
		totalRate = totalRate + c.GetBidRate()
		weights = append(weights, c.GetBidRate())
	}
	i, rate := utils.WeightedRandomIndex(weights)
	return campaigns[i], rate
}

// 筛选符合条件的账户、计划、活动
// func filters(req *bid_adapter.BidRequest) (campaigns []*ad.Campaign) {
func filters(req *protocol.BidRequest) (campaigns []*ad.Campaign) {
	for _, campaign := range dbid.ActiveCampaigns {

		// 基础过滤
		// 计划状态过滤
		// 计划投放周期过滤
		// 计划预算过滤
		// 计划曝光频控过滤
		// 计划点击频控过滤
		// 活动状态过滤
		// 虚拟活动过滤
		// 活动投放周期过滤
		// 活动投放时段过滤
		// 活动预算过滤
		// 活动价格初步过滤
		// 活动报告频控
		// 活动点击频控
		// 活动黑白名单过滤
		// 活动定向包过滤
		// 模板创意素材过滤
		// 价格过滤
		// 计算出价
		// 填充创意
		// 填充曝光
		// 填充点击
		// 响应

		// 状态过滤、投放周期过滤、投放时间段、预算过滤都在dbid.ActiveCampaigns

		// 预算过滤、曝光数过滤
		if filterByBudget(campaign) {
			continue
		}

		// 定向过滤
		if filterByTarget(req, campaign) {
			continue
		}

		if filterByWhiteBlackList(req, campaign) {
			continue
		}

		// 曝光频次过滤
		if filterByFrequencies(req, campaign) {
			continue
		}

		campaigns = append(campaigns, campaign)
	}
	return
}

// 参与竞价，选出竞得者
func bids() {

}

// 补充必要信息
func fill() {

}

func filterByBudget(c *ad.Campaign) bool {
	return dbid.BudgetControl.CheckOver(c.GetBudgetKey())
}

// func filterByFrequencies(req *bid_adapter.BidRequest, cs []*ad.Campaign) (campaigns []*ad.Campaign) {
func filterByFrequencies(req *protocol.BidRequest, c *ad.Campaign) bool {
	if filterByFrequency(req, c.GetImpFrequencyKey(), c.GetImpFrequencyMinute()) {
		return true
	}
	if filterByFrequency(req, c.GetClkFrequencyKey(), c.GetClkFrequencyMinute()) {
		return true
	}
	return false
}

// 定向过滤
func filterByTarget(req *protocol.BidRequest, c *ad.Campaign) bool {
	// 先增加地区定向
	var geo *protocol.Geo
	if req.Device != nil && req.Device.Geo != nil {
		geo = req.Device.Geo
	} else if req.User != nil && req.User.Geo != nil {
		geo = req.User.Geo
	}
	if !c.InRegion(geo.Country) {
		return true
	}
	return false
}

// 黑名名单过来
func filterByWhiteBlackList(req *protocol.BidRequest, c *ad.Campaign) bool {
	if bl := c.BlackWhiteList; bl != nil && req != nil {
		if dev := req.Device; bl.HasDeviceWhileBlackList() && dev != nil {
			if len(dev.IFA) > 0 {
				if bl.IsDeviceWhileList(constant.DeviceIdGaid, dev.IFA) {
					return false
				} else if bl.IsDeviceBlackList(constant.DeviceIdGaid, dev.IFA) {
					return true
				} else if bl.IsDeviceWhileList(constant.DeviceIdOaid, dev.IFA) {
					return false
				} else if bl.IsDeviceBlackList(constant.DeviceIdOaid, dev.IFA) {
					return true
				}
			}
			if len(dev.IDMD5) > 0 {
				if bl.IsDeviceWhileList(constant.DeviceIdMd5Idfa, dev.IDMD5) {
					return false
				} else if bl.IsDeviceBlackList(constant.DeviceIdMd5Idfa, dev.IDMD5) {
					return true
				} else if bl.IsDeviceWhileList(constant.DeviceIdMd5Imei, dev.IDMD5) {
					return false
				} else if bl.IsDeviceBlackList(constant.DeviceIdMd5Imei, dev.IDMD5) {
					return true
				}
			}
		}
	}
	return false
}

// func filterByFrequency(req *bid_adapter.BidRequest, frequencyKey, frequency int) bool {
func filterByFrequency(req *protocol.BidRequest, frequencyKey, frequency int) bool {
	if v, exists := dbid.AdFrequency.Load(frequencyKey); !exists {
		return false
	} else {
		if _cache, ok := v.(local_cache.Cache); ok {
			if dev := req.Device; dev != nil {
				switch strings.ToLower(dev.OS) {
				case "ios":
					if len(dev.IDMD5) > 0 {
						if times, e := _cache.Get(dev.IDMD5); e {
							if times.(int) >= frequency {
								return true
							}
						}
					}
				case "android":
					if len(dev.IDMD5) > 0 {
						if times, e := _cache.Get(dev.IDMD5); e {
							if times.(int) >= frequency {
								return true
							}
						}
					}
				}
			}
		}

	}
	return false
}

// func getRespBid(adxId int32, id string, imp *bid_adapter.BidRequest_Imp, campaign *ad.Campaign) (bid *bid_adapter.BidResponse_SeatBid_Bid, err error) {
func getRespBid(id string, req *protocol.BidRequest, imp protocol.Impression, campaign *ad.Campaign) (bid protocol.Bid, err error) {

	var v *ad.Creative
	var exist bool
	bid = protocol.Bid{
		ID:         id,
		ImpID:      imp.ID,
		MarkupType: constant.MarkupTypeBanner,
		// 暂时默认写死
		Categories: []protocol.ContentCategory{protocol.ContentCategoryContestsFreebies},
		AdID:       "shopee",
		AdvDomains: []string{"shopee.com"},
	}

	if len(campaign.Creatives) == 0 && len(campaign.Adm) == 0 {
		return bid, fmt.Errorf("活动%d不存在创意/动态代码", campaign.ID)
	}
	//v = campaign.Creatives[0]

	//spotId, randC := kehudsp.GetYorkUCreative()
	if imp.Banner != nil {
		var creativeUrl string
		//v, creativeUrl, exist = GetNearlyCreative(imp.Banner.GetW(), imp.Banner.GetH())
		if !exist {
			//return nil, fmt.Errorf("创意尺寸不存在，广告位: %d, 尺寸:%dx%d", spotId, imp.Banner.GetW(), imp.Banner.GetH())
		}

		if v, exist = campaign.SelectCreative(1, imp.Banner.Width, imp.Banner.Height); v != nil && exist && v.Material != nil {
			creativeUrl = v.Material.GetAbsoluteUrl()
		}
		switch campaign.GetBidMode() {
		case constant.BidModeFixed:
			bid.Price = campaign.GetBidPrice()
		case constant.BidModeAvg:
			// 	bid.Price = utils.Ceil((campaign.GetBidPrice()-imp.BidFloor)*rand.Float64()+imp.BidFloor, 2)
			bid.Price = utils.Ceil(((campaign.GetBidPrice()-imp.BidFloor)*0.3+imp.BidFloor)+0.35*rand.Float64()*(campaign.GetBidPrice()-imp.BidFloor), 2)
		default:
			return bid, errors.New("不支持的出价模式")
		}
		switch campaign.GetBidMethod() {
		case constant.BidMethodCpm:
		default:
			return bid, errors.New("不支持的出价方法")
		}
		bid.Width = imp.Banner.Width
		bid.Height = imp.Banner.Height
		//bid.CreativeUrl = proto.String(getImg(imp.Banner.GetW(), imp.Banner.GetH()))

		if adm := campaign.GetAdm(); len(adm) > 0 {
			paramsMap := buildTrackParamsMap(req, imp)
			campaign.FillTrackParams(paramsMap)
			params := BuildTrackParams(paramsMap)
			impTrack := BuildImpTrack(params)
			clkTrack := BuildClkTrack(params)
			bid.AdMarkup = campaign.BuildAdmForCode(impTrack, clkTrack)
			// 替换dsp宏
			if req.App != nil {
				if len(req.App.Bundle) > 0 {
					bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspBundle, req.App.Bundle)
				} else if len(req.App.ID) > 0 {
					bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspBundle, req.App.ID)
				}
				if req.App.Publisher != nil {
					if len(req.App.Publisher.ID) > 0 {
						bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspPublisher, req.App.Publisher.ID)
					} else if len(req.App.Publisher.Domain) > 0 {
						bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspPublisher, req.App.Publisher.Domain)
					} else if len(req.App.Publisher.Name) > 0 {
						bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspPublisher, req.App.Publisher.Name)
					}
				}
			}
			// ${DSP_OFFER_DAY_HOUR}
			now := time.Now()
			dayHour := strconv.Itoa(now.Day()*100 + now.Hour())
			bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspOfferDayHour, dayHour)
			//bid.CreativeID = strconv.Itoa(len(adm))
			bid.CreativeID = utils.MD5(adm)
			bid.CampaignID = protocol.StringOrNumber(strconv.Itoa(int(campaign.ID)))
		} else {
			global.GVA_LOG.Error("暂时只支持adm代码投放")
			return bid, errors.New("暂时只支持adm代码投放")
		}
		_ = creativeUrl
	}
	if imp.Native != nil {
		// TODO
	}
	if video := imp.Video; video != nil {
		// TODO
	}
	return
}

func getNative(adxId int32, tpl *bid_adapter.NativeRequest, campaign *ad.Campaign) (*ad.Creative, *bid_adapter.NativeResponse) {
	var imageCnt int
	var assets []*bid_adapter.NativeResponse_Asset
	//var v *ad.Creative
	// TODO
	/*
		for _, asset := range tpl.Assets {
			if img := asset.GetImg(); img != nil {
				if v == nil && img.GetType() != bid_adapter.NativeRequest_Asset_Image_MAIN {
					//	v, _, _ = kehudsp.GetNearlyCreative(dspSpotId, img.GetW(), img.GetH())
				} else if img.GetType() == bid_adapter.NativeRequest_Asset_Image_MAIN {
					//	v, _, _ = kehudsp.GetNearlyCreative(dspSpotId, img.GetW(), img.GetH())
				}
			}
		}*/

	var titleAd *ad.Creative
	for _, asset := range tpl.Assets {
		if img := asset.GetImg(); img != nil {
			switch adxId {
			case 30190:
				if imageCnt > 0 && (tpl.GetTemplateId() == "201900005" || tpl.GetTemplateId() == "201900006") {
					continue
				}
			}
			if v, exist := campaign.SelectCreative(1, int(img.GetW()), int(img.GetH())); exist && v.Material != nil {
				titleAd = v
				assets = append(assets, &bid_adapter.NativeResponse_Asset{
					Id:       asset.Id,
					Required: asset.Required,
					AssetOneof: &bid_adapter.NativeResponse_Asset_Img{Img: &bid_adapter.NativeResponse_Asset_Image{
						Url:  proto.String(v.Material.GetAbsoluteUrl()),
						W:    img.W,
						H:    img.H,
						Type: bid_adapter.NativeResponse_Asset_Image_ImageAssetType(img.GetType()).Enum(),
					}},
				})
				imageCnt++
			}
			//_, creativeUrl, _ := kehudsp.GetNearlyCreative(dspSpotId, img.GetW(), img.GetH())

		}

		if video := asset.GetVideo(); video != nil {
			if v, exist := campaign.SelectCreative(2, int(video.GetW()), int(video.GetH())); exist && v.Material != nil {
				assets = append(assets, &bid_adapter.NativeResponse_Asset{
					Id:       asset.Id,
					Required: asset.Required,
					AssetOneof: &bid_adapter.NativeResponse_Asset_Video_{Video: &bid_adapter.NativeResponse_Asset_Video{
						Url:      proto.String(v.Material.GetAbsoluteUrl()),
						W:        proto.Int32(video.GetW()),
						H:        proto.Int32(video.GetH()),
						Duration: proto.Int32(15),
					}},
				})
			}
		}

		if title := asset.GetTitle(); title != nil {
			//var val kehudsp.CommonTitle
			var val []rune
			if titleAd != nil && len(titleAd.Title) > 0 {
				val = []rune(titleAd.Title)
				switch adxId {
				case 30120:
				case 30190:

				default:
					switch {
					case title.GetLen() == 0:
						if len(val) > 8 {
							val = val[:8]
						}
					default:
						if len(val) > int(title.GetLen()) {
							val = val[:int(title.GetLen())]
						}
					}
				}
			}

			assets = append(assets, &bid_adapter.NativeResponse_Asset{
				Id:       asset.Id,
				Required: asset.Required,
				AssetOneof: &bid_adapter.NativeResponse_Asset_Title_{
					Title: &bid_adapter.NativeResponse_Asset_Title{
						Text: proto.String(string(val)),
					},
				},
			})
		}

		if data := asset.GetData(); data != nil {
			var val = "查看详情"
			switch {
			case data.GetLen() < 3:
				val = "详情"
			case data.GetLen() < 5:
				val = "查看详情"
			default:
				var desc []rune
				//desc := kehudsp.CommonTitle
				if titleAd != nil && len(titleAd.Desc) > 0 {
					desc = []rune(titleAd.Desc)
				}
				//if len(kehudsp.CommonTitle) >= int(data.GetLen()) {
				//	desc = desc[:int(data.GetLen())]
				//}
				val = string(desc)
			}
			assets = append(assets, &bid_adapter.NativeResponse_Asset{
				Id:       asset.Id,
				Required: asset.Required,
				AssetOneof: &bid_adapter.NativeResponse_Asset_Data_{
					Data: &bid_adapter.NativeResponse_Asset_Data{
						Type:  bid_adapter.NativeResponse_Asset_Data_DataAssetType(data.GetType()).Enum(),
						Value: proto.String(val),
					},
				},
			})
		}

	}
	return titleAd, &bid_adapter.NativeResponse{
		Assets:     assets,
		TemplateId: proto.String(tpl.GetTemplateId()),
	}
}

func hasVideo(tpl *bid_adapter.NativeRequest) bool {
	for _, asset := range tpl.Assets {
		if video := asset.GetVideo(); video != nil {
			return true
		}
	}
	return false
}
func BuildImpTrack(params string) string {
	return fmt.Sprintf("%s/track%s?pr=${AUCTION_PRICE}&%s", global.GVA_CONFIG.Dsp.Domain, global.GVA_CONFIG.Dsp.Track.Impression.Uri, params)
	//return fmt.Sprintf("%s:%d/track%s?pr=${AUCTION_PRICE}&%s", global.GVA_CONFIG.Dsp.Domain, global.GVA_CONFIG.Dsp.Track.Port, global.GVA_CONFIG.Dsp.Track.Impression.Uri, params)
}

func BuildClkTrack(params string) string {
	return fmt.Sprintf("%s/track%s?%s", global.GVA_CONFIG.Dsp.Domain, global.GVA_CONFIG.Dsp.Track.Click.Uri, params)
	//return fmt.Sprintf("%s:%d/track%s?%s", global.GVA_CONFIG.Dsp.Domain, global.GVA_CONFIG.Dsp.Track.Port, global.GVA_CONFIG.Dsp.Track.Click.Uri, params)
}

func BuildTrackParams(params map[string]string) string {
	var keys = make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	values := url.Values{}
	for _, key := range keys {
		values.Add(key, params[key])
	}

	return values.Encode()
}

func buildTrackParamsMap(req *protocol.BidRequest, imp protocol.Impression) map[string]string {
	fmt.Sprintf("rid={BidRequest.id}&sp={BidRequest.imp[0].id}&ap={BidRequest.app.id}&site={BidRequest.site.id}&puer={BidRequest.app.publisher.id or BidRequest.site.publisher.id}&os={BidRequest.device.os}&ifa5={MD5(BidRequest.device.ifa)}&ip={BidRequest.device.ip}&ua={BidRequest.device.ua}&oid={BidRequest.device.oaid}&cid={BidRequest.device.caid}&cidv={BidRequest.device.caid_version}")
	params := make(map[string]string)
	params["rid"] = req.ID
	params["at"] = strconv.Itoa(req.AuctionType)
	if req.App != nil {
		if len(req.App.Bundle) > 0 {
			params["ap"] = req.App.Bundle
		}
		if len(req.App.ID) > 0 {
			params["apid"] = req.App.ID
		}
		if puer := req.App.Publisher; puer != nil {
			if len(puer.ID) > 0 {
				params["puerid"] = puer.ID
			}
			if len(puer.Domain) > 0 {
				params["puer"] = puer.Domain
			} else if len(puer.Name) > 0 {
				params["puer"] = puer.Name
			}
		}
	} else if req.Site != nil {
		if len(req.Site.ID) > 0 {
			params["siteid"] = req.Site.ID
		}
		if len(req.Site.Domain) > 0 {
			params["site"] = req.Site.Domain
		}
		if puer := req.Site.Publisher; puer != nil {
			if len(puer.ID) > 0 {
				params["puerid"] = puer.ID
			}
			if len(puer.Domain) > 0 {
				params["puer"] = puer.Domain
			}
		}
	}
	if dev := req.Device; dev != nil {
		params["os"] = dev.OS
		if len(dev.IFA) > 0 {
			params["ifa5"] = utils.MD5(dev.IFA)
		}
		params["ip"] = dev.IP
		params["ip6"] = dev.IPv6
		params["did5"] = dev.IDMD5
		params["did"] = dev.IFA

		if geo := dev.Geo; geo != nil {
			params["cny"] = geo.Country
		}
	}
	if user := req.User; user != nil {
		params["uid"] = user.ID
		params["gender"] = user.Gender
		params["birth"] = strconv.Itoa(user.YearOfBirth)
		if geo := user.Geo; geo != nil {
			params["cny"] = geo.Country
		}
	}
	params["imid"] = imp.ID
	switch {
	case imp.Banner != nil:
		params["adt"] = "0"
	case imp.Native != nil:
		params["adt"] = "1"
	case imp.Video != nil:
		params["adt"] = "2"
	case imp.Audio != nil:
		params["adt"] = "3"
	default:
		params["adt"] = "0"
	}
	params["bf"] = decimal.NewFromFloat(imp.BidFloor).String()
	params["sp"] = imp.TagID
	return params
}
