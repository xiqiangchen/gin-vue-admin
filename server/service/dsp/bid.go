package dsp

import (
	"errors"
	"fmt"
	dbid "github.com/flipped-aurora/gin-vue-admin/server/dsp/bid"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ad"
	bid_adapter "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/bid"
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

func (bidService *BidService) Bid(req *bid_adapter.BidRequest, c *gin.Context) (*bid_adapter.BidResponse, bool) {

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
func offerByCampaign(req *bid_adapter.BidRequest, campaign *ad.Campaign) (resp *bid_adapter.BidResponse, err error) {

	bids := make([]*bid_adapter.BidResponse_SeatBid_Bid, 0, len(req.GetImp()))
	for _, imp := range req.GetImp() {
		// 底价超过12人民币
		if imp.GetBidfloor() > 13 {
			continue
		}
		_bid, err := getRespBid(req.GetAdxid(), req.GetId(), imp, campaign)
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

	resp = &bid_adapter.BidResponse{
		Id: req.Id,
		Seatbid: []*bid_adapter.BidResponse_SeatBid{
			{
				Bid: bids,
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
func filters(req *bid_adapter.BidRequest) (campaigns []*ad.Campaign) {
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

func filterByFrequencies(req *bid_adapter.BidRequest, cs []*ad.Campaign) (campaigns []*ad.Campaign) {
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

func filterByFrequency(req *bid_adapter.BidRequest, frequencyKey, frequency int) bool {
	if v, exists := dbid.AdFrequency[frequencyKey]; !exists {
		return false
	} else if dev := req.GetDevice(); dev != nil {
		switch strings.ToLower(dev.GetOs()) {
		case "ios":
			if len(dev.GetCaid()) > 0 {
				if times, e := v.Get(dev.GetCaid()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		case "android":
			if len(dev.GetOaidmd5()) > 0 {
				if times, e := v.Get(dev.GetOaidmd5()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			} else if len(dev.GetDidmd5()) > 0 {
				if times, e := v.Get(dev.GetDidmd5()); e {
					if times.(int) >= frequency {
						return true
					}
				}
			}
		}
	}
	return false
}

func getRespBid(adxId int32, id string, imp *bid_adapter.BidRequest_Imp, campaign *ad.Campaign) (bid *bid_adapter.BidResponse_SeatBid_Bid, err error) {

	var cid uint
	var v *ad.Creative
	var exist bool

	// TODO
	if len(campaign.Creatives) == 0 {
		return nil, fmt.Errorf("活动%d不存在创意", campaign.ID)
	}
	//v = campaign.Creatives[0]

	//spotId, randC := kehudsp.GetYorkUCreative()
	if imp.Banner != nil {
		var creativeUrl string
		//v, creativeUrl, exist = GetNearlyCreative(imp.Banner.GetW(), imp.Banner.GetH())
		if !exist {
			//return nil, fmt.Errorf("创意尺寸不存在，广告位: %d, 尺寸:%dx%d", spotId, imp.Banner.GetW(), imp.Banner.GetH())
		}

		if v, exist = campaign.SelectCreative(1, int(imp.Banner.GetW()), int(imp.Banner.GetH())); v != nil && exist && v.Material != nil {
			creativeUrl = v.Material.GetAbsoluteUrl()
		}

		bid = &bid_adapter.BidResponse_SeatBid_Bid{
			Id:    proto.String(id),
			Impid: proto.String(imp.GetId()),
			Price: proto.Float64(campaign.GetBidPrice()),
			//Nurl:  proto.String(""),
		}

		bid.W = proto.Int32(imp.Banner.GetW())
		bid.H = proto.Int32(imp.Banner.GetH())
		//bid.CreativeUrl = proto.String(getImg(imp.Banner.GetW(), imp.Banner.GetH()))
		bid.CreativeUrl = proto.String(creativeUrl)
	}
	if imp.Native != nil {

		bid = &bid_adapter.BidResponse_SeatBid_Bid{
			Id:    proto.String(id),
			Impid: proto.String(imp.GetId()),
			Price: proto.Float64(campaign.GetBidPrice()),
			//Nurl:  proto.String(""),
		}

		tpls := imp.Native.GetTemplates()
		if len(tpls) > 0 {
			tpl := tpls[0]
			var native *bid_adapter.NativeResponse
			v, native = getNative(adxId, tpl, campaign)
			bid.AdmOneof = &bid_adapter.BidResponse_SeatBid_Bid_AdmNative{AdmNative: native}
		}

	}
	if video := imp.Video; video != nil {
		v, exist = campaign.SelectCreative(2, int(imp.Video.GetW()), int(imp.Video.GetH()))
		bid = &bid_adapter.BidResponse_SeatBid_Bid{
			Id:    proto.String(id),
			Impid: proto.String(imp.GetId()),
			Price: proto.Float64(35),
			//Nurl:  proto.String(""),
		}

		bid.W = proto.Int32(imp.Video.GetW())
		bid.H = proto.Int32(imp.Video.GetH())
		bid.Crtype = proto.Int32(int32(3))
		//v, _, exist = GetNearlyCreative(spotId, imp.Video.GetW(), imp.Video.GetH())
		if !exist {
			//return nil, fmt.Errorf("创意尺寸不存在，广告位: %d, 尺寸:%dx%d", spotId, imp.Banner.GetW(), imp.Banner.GetH())
		}
		bid.CreativeUrl = proto.String(v.Material.GetAbsoluteUrl())
	}
	if v != nil {
		cid = v.ID
	}
	bid.Crid = proto.String(fmt.Sprintf("%v_%v_%v", campaign.PlanId, campaign.ID, cid))

	// 监测
	//bid.Nurl = proto.String(kehudsp.GetGlobalImpUrl(randC.ImpUrl))
	bid.BidExt = &bid_adapter.BidResponse_SeatBid_Bid_BidExt{
		LdpType:    bid_adapter.BidResponse_SeatBid_Bid_BidExt_Web.Enum(),
		Ldp:        proto.String(campaign.H5),
		PvtrackUrl: []string{campaign.ImpTrackUrl},
		CtrackUrl:  []string{campaign.ClickTrackUrl},
	}
	if len(campaign.Deeplink)+len(campaign.UniversalLink) > 0 {
		bid.BidExt.LdpType = bid_adapter.BidResponse_SeatBid_Bid_BidExt_DeepLink.Enum()
	}

	//bid.DspSpotId = proto.Int64(int64(spotId))

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
