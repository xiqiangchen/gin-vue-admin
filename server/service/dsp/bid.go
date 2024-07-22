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
	"go.uber.org/zap"
	"math/rand"
	"strings"
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
	if winCampaign := selectCampaign(campaigns); winCampaign != nil {
		if resp, err := offerByCampaign(req, winCampaign); err != nil {
			global.GVA_LOG.Error("出价活动转换bidResp协议失败", zap.Error(err))
		} else {
			return resp, true
		}
	}
	return nil, false

}

// 根据活动填充出价响应
// func offerByCampaign(req *bid_adapter.BidRequest, campaign *ad.Campaign) (resp *bid_adapter.BidResponse, err error) {
func offerByCampaign(req *protocol.BidRequest, campaign *ad.Campaign) (resp *protocol.BidResponse, err error) {

	bids := make([]protocol.Bid, 0, len(req.Impressions))
	for _, imp := range req.Impressions {
		if imp.BidFloor > campaign.GetBidPrice() || rand.Float64() > campaign.GetBidRate()/100 {
			continue
		}
		_bid, err := getRespBid(req.ID, imp, campaign)
		if err != nil {
			continue
		}

		bids = append(bids, _bid)

		break
	}

	if len(bids) == 0 {
		err = errors.New("offerByCampaign异常，没找到匹配的素材")
		return
	}

	resp = &protocol.BidResponse{
		ID:       req.ID,
		BidID:    req.ID,
		Currency: "USD",
		SeatBids: []protocol.SeatBid{
			{
				Bids: bids,
			},
		},
	}
	return
}

// 选择其中一个活动作为出价
func selectCampaign(campaigns []*ad.Campaign) *ad.Campaign {
	return campaigns[rand.Intn(len(campaigns))]
}

// 筛选符合条件的账户、计划、活动
// func filters(req *bid_adapter.BidRequest) (campaigns []*ad.Campaign) {
func filters(req *protocol.BidRequest) (campaigns []*ad.Campaign) {
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
	// 曝光频次过滤
	campaigns = filterByFrequencies(req, dbid.ActiveCampaigns)

	return
}

// 参与竞价，选出竞得者
func bids() {

}

// 补充必要信息
func fill() {

}

// func filterByFrequencies(req *bid_adapter.BidRequest, cs []*ad.Campaign) (campaigns []*ad.Campaign) {
func filterByFrequencies(req *protocol.BidRequest, cs []*ad.Campaign) (campaigns []*ad.Campaign) {
	for _, c := range cs {
		if filterByFrequency(req, c.GetImpFrequencyKey(), c.GetImpFrequencyMinute()) {
			continue
		}
		if filterByFrequency(req, c.GetClkFrequencyKey(), c.GetClkFrequencyMinute()) {
			continue
		}
		campaigns = append(campaigns, c)
	}
	return
}

// func filterByFrequency(req *bid_adapter.BidRequest, frequencyKey, frequency int) bool {
func filterByFrequency(req *protocol.BidRequest, frequencyKey, frequency int) bool {
	if v, exists := dbid.AdFrequency[frequencyKey]; !exists {
		return false
	} else if dev := req.Device; dev != nil {
		switch strings.ToLower(dev.OS) {
		case "ios":
			if len(dev.IDMD5) > 0 {
				if times, e := v.Get(dev.IDMD5); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		case "android":
			if len(dev.IDMD5) > 0 {
				if times, e := v.Get(dev.IDMD5); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		}
	}
	return false
}

// func getRespBid(adxId int32, id string, imp *bid_adapter.BidRequest_Imp, campaign *ad.Campaign) (bid *bid_adapter.BidResponse_SeatBid_Bid, err error) {
func getRespBid(id string, imp protocol.Impression, campaign *ad.Campaign) (bid protocol.Bid, err error) {

	var v *ad.Creative
	var exist bool
	bid = protocol.Bid{
		ID:         id,
		ImpID:      imp.ID,
		MarkupType: constant.MarkupTypeBanner,
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
			bid.Price = utils.Ceil((campaign.GetBidPrice()-imp.BidFloor)*rand.Float64()+imp.BidFloor, 2)
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
			bid.AdMarkup = campaign.GetAdm()
			//bid.CreativeID = strconv.Itoa(len(adm))
			bid.CreativeID = utils.MD5(adm)
		} else {
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
