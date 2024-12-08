package dsp

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/constant"
	dbid "github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/dsp/bid/pricer"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	protocol "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6/native/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6/native/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/vast"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/ip2geo"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BidService struct{}

func (bidService *BidService) SendMsg(msg []byte) {
	utils.SendMsg(global.GVA_KAFKA_PRODUCER, global.GVA_CONFIG.Dsp.Bid.Topic, msg)
}

// func (bidService *BidService) Bid(req *bid_adapter.BidRequest, c *gin.Context) (*bid_adapter.BidResponse, bool) {
func (bidService *BidService) Bid(req *protocol.BidRequest, c *gin.Context) (*protocol.BidResponse, bool) {

	adxStr, _ := c.GetQuery("adx")
	adx, _ := strconv.Atoi(adxStr)
	var campaigns []*ad.Campaign
	if campaigns = filters(req, adx); len(campaigns) == 0 {
		return nil, false
	}
	if winCampaign, weight := selectCampaign(campaigns); winCampaign != nil {
		if resp, offer, err := offerByCampaign(adx, req, winCampaign, weight); err != nil {
			global.GVA_LOG.Error("出价活动转换bidResp协议失败", zap.Error(err))
		} else if offer {
			return resp, true
		} else {
			return nil, false
		}
	}
	return nil, false

}

func (bidService *BidService) FillParams(adx string, req *protocol.BidRequest, c *gin.Context) error {
	if req.Device != nil && req.Device.Geo != nil && len(req.Device.Geo.Country) == 0 {

		var ip string
		switch {
		case len(req.Device.IP) > 0:
			ip = req.Device.IP
		case len(req.Device.IPv6) > 0:
			ip = req.Device.IPv6
		}
		if len(ip) > 0 {
			if trans, err := ip2geo.Parse(ip); err == nil && len(trans.Country) > 0 {
				if cnt, ok := constant.CountryMap[trans.CountryCode]; ok {
					req.Device.Geo.Country = cnt
					if len(req.Device.Geo.City) == 0 {
						req.Device.Geo.City = trans.CountryEn
					}
					if len(req.Device.Geo.Region) == 0 {
						req.Device.Geo.Region = trans.ProvinceEn
					}
				}
			}

		}
	}
	return nil
}

// 根据活动填充出价响应
// func offerByCampaign(req *bid_adapter.BidRequest, campaign *ad.Campaign) (resp *bid_adapter.BidResponse, err error) {
func offerByCampaign(adx int, req *protocol.BidRequest, campaign *ad.Campaign, weight float64) (resp *protocol.BidResponse, offer bool, err error) {

	bids := make([]protocol.Bid, 0, len(req.Impressions))
	// 暂时只支持对第一个曝光进行处理
	for _, imp := range req.Impressions {
		if imp.BidFloor > campaign.GetBidPrice() || (rand.Float64() > (campaign.GetBidRate() / weight / 100)) {
			continue
		}
		_bid, err := getRespBid(adx, req.ID, req, imp, campaign)
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
func filters(req *protocol.BidRequest, adx int) (campaigns []*ad.Campaign) {
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
		if filterByTarget(req, campaign, adx) {
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
func filterByTarget(req *protocol.BidRequest, c *ad.Campaign, adx int) bool {
	if !c.InAdx(adx) {
		return true
	}
	// 先增加地区定向
	var geo *protocol.Geo
	if req.Device != nil && req.Device.Geo != nil {
		geo = req.Device.Geo
	} else if req.User != nil && req.User.Geo != nil {
		geo = req.User.Geo
	}
	if geo != nil && !c.InRegion(geo.Country) {
		switch len(geo.Country) {
		case 2:
			if country, ok := constant.CountryMap[geo.Country]; ok && c.InRegion(country) {
				return false
			}
		case 3:
			return !c.InRegion(geo.Country)
		}
		return true
	}

	// 设备过滤
	var os string
	if req.Device != nil && len(req.Device.OS) > 0 {
		os = req.Device.OS
	}
	if !c.InOs(os) {
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
func getRespBid(adx int, id string, req *protocol.BidRequest, imp protocol.Impression, campaign *ad.Campaign) (bid protocol.Bid, err error) {

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

	// 公共部分
	switch campaign.GetBidMode() {
	case constant.BidModeFixed:
		bid.Price = campaign.GetBidPrice()
	case constant.BidModeAvg:
		// 	bid.Price = utils.Ceil((campaign.GetBidPrice()-imp.BidFloor)*rand.Float64()+imp.BidFloor, 2)
		bid.Price = utils.Ceil(((campaign.GetBidPrice()-imp.BidFloor)*0.3+imp.BidFloor)+0.45*rand.Float64()*(campaign.GetBidPrice()-imp.BidFloor), 2)
	default:
		return bid, errors.New("不支持的出价模式")
	}
	switch campaign.GetBidMethod() {
	case constant.BidMethodCpm:
	default:
		return bid, errors.New("不支持的出价方法")
	}
	paramsMap := buildTrackParamsMap(req, imp)
	dp := campaign.VoteDeeplink()
	campaign.FillTrackParams(paramsMap)
	params := BuildTrackParams(paramsMap)
	pr, _ := pricer.DefaultPricer.Encode(bid.Price)
	impTrack := BuildImpTrack(params, pr)
	clkTrack := BuildClkTrack(params)
	var imps, clks []string

	impTrack = strings.ReplaceAll(impTrack, constant.DspCampaignId, strconv.Itoa(int(campaign.ID)))
	clkTrack = strings.ReplaceAll(clkTrack, constant.DspCampaignId, strconv.Itoa(int(campaign.ID)))
	if req.Device != nil && len(req.Device.OS) > 0 {
		impTrack = strings.ReplaceAll(impTrack, constant.DspOs, strings.ToLower(req.Device.OS))
		clkTrack = strings.ReplaceAll(clkTrack, constant.DspOs, strings.ToLower(req.Device.OS))
	}
	imps = append(imps, impTrack)
	clks = append(clks, clkTrack)

	if len(campaign.ImpTrackUrl) > 0 {
		var impTemp = strings.ReplaceAll(campaign.ImpTrackUrl, constant.DspCampaignId, strconv.Itoa(int(campaign.ID)))
		impTemp = strings.ReplaceAll(impTemp, constant.DspOs, strings.ToLower(req.Device.OS))
		imps = append(imps, impTemp)
	}
	if len(campaign.ClickTrackUrl) > 0 {
		var clkTemp = strings.ReplaceAll(campaign.ClickTrackUrl, constant.DspCampaignId, strconv.Itoa(int(campaign.ID)))
		clkTemp = strings.ReplaceAll(clkTemp, constant.DspOs, strings.ToLower(req.Device.OS))
		clks = append(clks, clkTemp)
	}

	bid.CampaignID = protocol.StringOrNumber(strconv.Itoa(int(campaign.ID)))
	if len(campaign.Brand) > 0 {
		bid.AdvDomains = []string{campaign.Brand}
		bid.AdID = campaign.Brand
	}
	tracks := protocol.ExtTracks{
		//ImpressionTracks: []string{impTrack},
		//ClickTracks:      []string{clkTrack},
		Deeplink:      dp,
		LandingUrl:    campaign.H5,
		UniversalLink: campaign.UniversalLink,
		BillingId:     162000188148,
	}
	if adx == 10 && imp.Ext != nil {
		var billingId struct {
			BillingId []string `json:"billing_id,omitempty"`
		}
		if err1 := json.Unmarshal(imp.Ext, &billingId); err1 == nil && len(billingId.BillingId) > 0 {
			tracks.BillingId, _ = strconv.ParseInt(billingId.BillingId[0], 10, 64)
		}
	}

	if ext, e := json.Marshal(tracks); e == nil {
		bid.Ext = ext
	}

	// 处理扩展功能
	var (
		finalLink string
		os        = "android"
	)
	if req.Device != nil && len(req.Device.OS) > 0 {
		os = strings.ToLower(req.Device.OS)
	}
	if campaign.IsLinkSystem() && dbid.LinkSystemClient != nil {
		if r, err := dbid.LinkSystemClient.GetClickLog(strconv.Itoa(adx), os, "0"); err == nil && r.Success && len(r.Data) > 0 {
			finalLink = r.Data
		}
	}

	if !campaign.IsLinkSystem() || len(finalLink) == 0 {
		if req.Device != nil {
			if len(dp) > 0 && strings.ToLower(req.Device.OS) == "android" {
				finalLink = dp
				campaign.DPIncr()
			} else if len(campaign.UniversalLink) > 0 && strings.ToLower(req.Device.OS) == "ios" {
				finalLink = campaign.UniversalLink
			} else if len(campaign.H5) > 0 {
				finalLink = campaign.H5
			}
		}
	}

	//spotId, randC := kehudsp.GetYorkUCreative()
	if imp.Banner != nil {
		var creativeUrl string

		if v, exist = campaign.SelectCreative(1, imp.Banner.Width, imp.Banner.Height); v != nil && exist && v.Material != nil {
			creativeUrl = v.Material.GetAbsoluteUrl()
		}

		bid.Width = imp.Banner.Width
		bid.Height = imp.Banner.Height
		//bid.CreativeUrl = proto.String(getImg(imp.Banner.GetW(), imp.Banner.GetH()))

		/*if len(campaign.ImpTrackUrl) > 0 {
			tracks.ImpressionTracks = append(tracks.ImpressionTracks, campaign.ImpTrackUrl)
		}*/

		/*if len(campaign.ClickTrackUrl) > 0 {
			tracks.ClickTracks = append(tracks.ClickTracks, campaign.ClickTrackUrl)
		}*/

		if len(campaign.Images) > 0 {
			if imgs, exists := campaign.Images[imp.Banner.Width*10000+imp.Banner.Height]; exists {
				if len(imgs) > 0 {
					if m := imgs[rand.Intn(len(imgs))].Material; m != nil && len(m.Url) > 0 {
						bid.ImageURL = m.GetAbsoluteUrl()
						bid.Width = imp.Banner.Width
						bid.Height = imp.Banner.Height
						bid.CreativeID = utils.MD5(m.GetAbsoluteUrl())
					}
				}
			} else if c := getNearlyCreative(campaign.Images, imp.Banner.Width, imp.Banner.Height); c != nil {
				if m := c.Material; m != nil && len(m.Url) > 0 {
					bid.ImageURL = m.GetAbsoluteUrl()
					bid.Width = imp.Banner.Width
					bid.Height = imp.Banner.Height
					bid.CreativeID = utils.MD5(m.GetAbsoluteUrl())
				}
			}
		}

		if adm := campaign.GetAdm(); len(adm) > 0 {
			bid.AdMarkup = campaign.BuildAdmForCode(impTrack, clkTrack)
			bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspLandingPage, finalLink)
			bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspCampaignId, strconv.Itoa(int(campaign.ID)))
			if req.Device != nil && len(req.Device.OS) > 0 {
				bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspOs, strings.ToLower(req.Device.OS))
			}

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
			if len(bid.ImageURL) > 0 {
				bid.AdMarkup = strings.ReplaceAll(bid.AdMarkup, constant.DspCreativeUrl, bid.ImageURL)
				bid.CreativeID = utils.MD5(bid.ImageURL)
			} else {
				bid.CreativeID = utils.MD5(adm)
			}
			//bid.CreativeID = strconv.Itoa(len(adm))
		} else {
			//global.GVA_LOG.Error("暂时只支持adm代码投放")
			return bid, errors.New("暂时只支持adm代码投放")
		}

		_ = creativeUrl
	} else if imp.Native != nil {
		var natResp *response.Response
		var err2 error
		switch imp.Native.Version {
		case "1.0":
			var tmpNative struct {
				Req *request.Request `json:"request"`
			}

			var raw string
			if err1 := json.Unmarshal([]byte(imp.Native.Request), &raw); err1 != nil {
				return bid, err1
			}

			if err1 := json.Unmarshal([]byte(raw), &tmpNative); err1 == nil {
				natResp, err2 = processNativeRequest(tmpNative.Req, campaign, imps, clks, finalLink)
			}
		default:
			var raw string
			if err1 := json.Unmarshal([]byte(imp.Native.Request), &raw); err1 != nil {
				return bid, err1
			}

			var native request.Request
			if err1 := json.Unmarshal([]byte(raw), &native); err1 == nil {
				natResp, err2 = processNativeRequest(&native, campaign, imps, clks, finalLink)
			} else {
				return bid, err1
			}
		}
		if err2 == nil && natResp != nil {
			if req, err3 := json.Marshal(natResp); err3 == nil {
				bid.AdMarkup = string(req)
				bid.CreativeID = utils.MD5(bid.AdMarkup)
			}
		}
	} else if video := imp.Video; video != nil {
		// TODO
	}
	return
}

// 根据Data Type生成相应的内容
func generateDataValue(dataType request.DataTypeID, campaign *ad.Campaign) string {
	switch dataType {
	case request.DataTypeSponsored:
		return "Sponsored by Example Brand"
	case request.DataTypeDesc:
		return "This is a detailed description of the product or service being advertised."
	case request.DataTypeRating:
		return "4.5"
	case request.DataTypeLikes:
		return "10.5K"
	case request.DataTypeDownloads:
		return "1M+"
	case request.DataTypePrice:
		return "$99.99"
	case request.DataTypeSalePrice:
		return "$79.99"
	case request.DataTypePhone:
		return "+1-800-555-0123"
	case request.DataTypeAddress:
		return "123 Example Street, City, Country"
	case request.DataTypeDescAdditional:
		return "Additional details about the product or service"
	case request.DataTypeDisplayURL:
		return "www.example.com/product"
	case request.DataTypeCTADesc:
		return "Shop Now"
	default:
		return "Unknown data type"
	}
}

// 处理native请求并生成响应的主要函数
func processNativeRequest(nativeReq *request.Request, campaign *ad.Campaign, impTracks, clkTracks []string, landing string) (*response.Response, error) {
	nativeResp := &response.Response{
		Version:     "1.2",
		Assets:      make([]response.Asset, 0, len(nativeReq.Assets)),
		ImpTrackers: impTracks,
		Link:        response.Link{URL: landing, ClickTrackers: clkTracks},
	}

	var title, desc string
	if len(campaign.Creatives) > 0 {
		title = campaign.Creatives[0].Title
		desc = campaign.Creatives[0].Desc
	}

	// 遍历处理每个资产
	for _, asset := range nativeReq.Assets {
		assetResp := response.Asset{
			ID: asset.ID,
		}

		// 处理标题类型
		if asset.Title != nil {
			assetResp.Title = &response.Title{
				Text: title,
			}
		}

		// 处理图片类型
		if asset.Image != nil {
			assetResp.Image = &response.Image{
				//URL:    "https://example.com/ad-image.jpg",
				Width:  asset.Image.Width,
				Height: asset.Image.Height,
			}

			if len(campaign.Images) > 0 {
				var w, h = asset.Image.Width, asset.Image.Height
				if w == 0 {
					w = asset.Image.WidthMin
				}
				if h == 0 {
					h = asset.Image.HeightMin
				}

				if imgs, exists := campaign.Images[w*10000+h]; exists {
					if len(imgs) > 0 {
						if cr := imgs[rand.Intn(len(imgs))]; cr != nil {
							title = cr.Title
							desc = cr.Desc
							if m := cr.Material; m != nil && len(m.Url) > 0 {
								assetResp.Image.URL = m.GetAbsoluteUrl()
							}
						}
					}
				} else if c := getNearlyCreative(campaign.Images, w, h); c != nil {
					title = c.Title
					desc = c.Desc
					if m := c.Material; m != nil && len(m.Url) > 0 {
						assetResp.Image.URL = m.GetAbsoluteUrl()
					}
				}
			}

		}

		// 处理视频类型
		if asset.Video != nil {
			var videoUrl string

			if len(campaign.Videos) > 0 {
				if m, e := campaign.SelectCreative(2, 1024, 1024); e && m.Material != nil && len(m.Material.Url) > 0 {
					videoUrl = m.Material.Url

				}
			}

			if vas, err1 := GenerateVAST(videoUrl, landing, title, desc, campaign.Brand, impTracks, clkTracks, 1024, 1024); err1 == nil {
				if byt, err2 := xml.Marshal(vas); err2 == nil {
					assetResp.Video = &response.Video{
						VASTTag: string(byt),
					}
				}
			}
		}

		// 处理数据类型
		if asset.Data != nil {
			assetResp.Data = &response.Data{
				Value: desc, //generateDataValue(asset.Data.TypeID, campaign),
			}
		}

		nativeResp.Assets = append(nativeResp.Assets, assetResp)
	}

	for i := range nativeResp.Assets {
		if nativeResp.Assets[i].Data != nil {
			nativeResp.Assets[i].Data.Value = desc
		} else if nativeResp.Assets[i].Title != nil {
			nativeResp.Assets[i].Title.Text = title
		}
	}

	return nativeResp, nil
}

func BuildImpTrack(params string, price string) string {
	return fmt.Sprintf("%s/track%s?pr=%v&%s", global.GVA_CONFIG.Dsp.Domain, global.GVA_CONFIG.Dsp.Track.Impression.Uri, price, params)
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
	//fmt.Sprintf("rid={BidRequest.id}&sp={BidRequest.imp[0].id}&ap={BidRequest.app.id}&site={BidRequest.site.id}&puer={BidRequest.app.publisher.id or BidRequest.site.publisher.id}&os={BidRequest.device.os}&ifa5={MD5(BidRequest.device.ifa)}&ip={BidRequest.device.ip}&ua={BidRequest.device.ua}&oid={BidRequest.device.oaid}&cid={BidRequest.device.caid}&cidv={BidRequest.device.caid_version}")
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
	if req.Source != nil && req.Source.SupplyChain != nil && len(req.Source.SupplyChain.Nodes) > 0 {
		nodes := make([]string, 0, len(req.Source.SupplyChain.Nodes))
		for _, node := range req.Source.SupplyChain.Nodes {
			nodes = append(nodes, node.Domain)
		}
		params["schain"] = strings.Join(nodes, ",")
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
	params["ts"] = strconv.FormatInt(time.Now().Unix(), 10)
	return params
}

func getNearlyCreative(cmap map[int][]*ad.Creative, w, h int) *ad.Creative {
	var min *ad.Creative
	var minRate = math.MaxFloat64
	rate := float64(w) / float64(h)
	for k, v := range cmap {
		if len(v) > 0 {
			if w == 0 || h == 0 {
				return v[rand.Intn(len(v))]
			}
			r := math.Abs(rate - float64(k/10000)/float64(k%10000))
			if r < minRate || r == math.NaN() {
				minRate = r
				min = v[rand.Intn(len(v))]
			}
		}

	}
	return min
}

// 生成VAST XML的函数
func GenerateVAST(videoURL, clickThrough, title, desc, adv string, imps, clks []string, width, height int) (*vast.VAST, error) {
	id := utils.MD5(videoURL)

	var impss []vast.Impression
	for i, im := range imps {
		impss = append(impss, vast.Impression{URI: im, ID: fmt.Sprintf("imp-%d", i)})
	}
	var clkss []vast.VideoClick
	for i, clk := range clks {
		clkss = append(clkss, vast.VideoClick{URI: clk, ID: fmt.Sprintf("clk-%d", i)})
	}
	vast := &vast.VAST{
		Version: "4.0",
		Ads: []vast.Ad{
			{
				ID: "ad-" + time.Now().Format("20060102150405"),
				InLine: &vast.InLine{
					AdSystem: &vast.AdSystem{
						Version: "1.0",
						Name:    "Ad Server",
					},
					AdTitle:     title,
					Impressions: impss,
					Description: desc,
					Advertiser:  adv,
					//Pricing:     "25.00",
					Creatives: []vast.Creative{
						{
							ID:       id,
							Sequence: 1,
							AdID:     id,
							Linear: &vast.Linear{
								Duration: vast.Duration(30),
								MediaFiles: []vast.MediaFile{
									{
										ID:                  id,
										Delivery:            "progressive",
										Type:                "video/mp4",
										Width:               width,
										Height:              height,
										Codec:               "H.264",
										Bitrate:             2000,
										MinBitrate:          1500,
										MaxBitrate:          2500,
										Scalable:            true,
										MaintainAspectRatio: true,
										URI:                 videoURL,
									},
								},
								VideoClicks: &vast.VideoClicks{
									ClickThroughs: []vast.VideoClick{
										{
											ID:  "ct-1",
											URI: clickThrough,
										},
									},
									ClickTrackings: clkss,
								},
								/*TrackingEvents: []vast.Tracking{
									{Event: "start", URI: impURL},
									//{Event: "firstQuartile", URI: "https://example.com/track/firstQuartile"},
									//{Event: "midpoint", URI: "https://example.com/track/midpoint"},
									//{Event: "thirdQuartile", URI: "https://example.com/track/thirdQuartile"},
									//{Event: "complete", URI: "https://example.com/track/complete"},
								},*/
							},
						},
					},
				},
			},
		},
	}

	return vast, nil
}
