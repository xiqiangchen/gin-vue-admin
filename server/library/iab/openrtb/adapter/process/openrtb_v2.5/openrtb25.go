package openrtb_v2_5

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/openrtb/adapter/base"
	openrtb_v2_52 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.5"
	openrtb_v2_62 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6"
	//openrtb_v2_6_1 "github.com/flipped-aurora/gin-vue-admin/server/model/dsp/iab/openrtb2/openrtb_v2.6.1"
)

//todo  深拷贝对别下 1.序列化,2.反射,3.硬编码 哪种更快

type OpenRTB25 struct {
	base.Common
}

func (o *OpenRTB25) SerializationType() int {
	return base.OPEN_RTB_SERIALIZATION_JSON
}
func (o *OpenRTB25) ProtocalVersion() string {
	return base.OPEN_RTB_VERSION_2_5
}

// 2.5 []byte <-> 2.5 obj
func (o *OpenRTB25) UnmarshalRequest(data []byte) (interface{}, error) {
	req := new(openrtb_v2_52.BidRequest)
	err := json.Unmarshal(data, req)
	return req, err
}
func (o *OpenRTB25) MarshalRequest(ori interface{}) ([]byte, error) {
	req, ok := ori.(*openrtb_v2_52.BidRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request obj")
	}
	return json.Marshal(req)
}
func (o *OpenRTB25) UnmarshalResponse(data []byte) (interface{}, error) {
	resp := new(openrtb_v2_52.BidResponse)
	err := json.Unmarshal(data, resp)
	return resp, err
}
func (o *OpenRTB25) MarshalResponse(ori interface{}) ([]byte, error) {
	resp, ok := ori.(*openrtb_v2_52.BidResponse)
	if !ok {
		return nil, fmt.Errorf("invalid response obj")
	}
	return json.Marshal(resp)
}

// 2.5 obj <-> 2.6 obj
func (o *OpenRTB25) TransformReqTo(ori interface{}) (req *openrtb_v2_62.BidRequest) {
	// todo 临时
	data, err := json.Marshal(ori)
	if err != nil {
		return nil
	}
	req = new(openrtb_v2_62.BidRequest)
	err = json.Unmarshal(data, req)
	if err != nil {
		return nil
	}
	return req

	oreq := ori.(*openrtb_v2_52.BidRequest)
	req = &openrtb_v2_62.BidRequest{
		ID:                oreq.ID,
		Impressions:       getImpressions(oreq.Imp),
		Site:              getSite(oreq.Site),
		App:               getApp(oreq.App),
		Device:            getDevice(oreq.Device),
		User:              getUser(oreq.User),
		Test:              int(oreq.Test),
		AuctionType:       int(oreq.AT),
		TimeMax:           int(oreq.TMax),
		Seats:             base.DeepCopyStringSlice(oreq.WSeat),
		BlockedSeats:      base.DeepCopyStringSlice(oreq.BSeat),
		Languages:         base.DeepCopyStringSlice(oreq.WLang),
		LanguagesB:        nil, // 2.5 没有这个字段，todo 后续在ext中检查 oreq.Ext
		AllImpressions:    int(oreq.AllImps),
		Currencies:        base.DeepCopyStringSlice(oreq.Cur),
		BlockedCategories: deepCopyCategories(oreq.BCat),
		BlockedAdvDomains: base.DeepCopyStringSlice(oreq.BAdv),
		BlockedApps:       base.DeepCopyStringSlice(oreq.BApp),
		Source:            deepCopySource(oreq.Source),
		Regulations:       deepCopyRegulations(oreq.Regs),
		Ext:               base.DeepCopyByte(oreq.Ext),
	}

	// todo
	return req
}
func (o *OpenRTB25) TransformReqFrom(ori *openrtb_v2_62.BidRequest) (result interface{}) {
	data, err := json.Marshal(ori)
	if err != nil {
		return nil
	}
	req := new(openrtb_v2_52.BidRequest)
	err = json.Unmarshal(data, req)
	if err != nil {
		return nil
	}

	if ori.Source != nil && ori.Source.SupplyChain != nil && req.Source != nil {
		supplyChain26to25(ori.Source, req.Source)
	}

	return req
}
func (o *OpenRTB25) TransformRespTo(ori interface{}) (resp *openrtb_v2_62.BidResponse) {
	// todo 临时
	data, err := json.Marshal(ori)
	if err != nil {
		return nil
	}
	res := new(openrtb_v2_62.BidResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil
	}
	return res
}
func (o *OpenRTB25) TransformRespFrom(ori *openrtb_v2_62.BidResponse) (resp interface{}) {
	// todo 临时
	data, err := json.Marshal(ori)
	if err != nil {
		return nil
	}
	res := new(openrtb_v2_52.BidResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil
	}
	return res
}

// 2.5 []byte <-> 2.6 obj
func (o *OpenRTB25) NormalizeRequest(data []byte) (*openrtb_v2_62.BidRequest, error) {
	// todo domark 先暂时 序列化 2.6兼容2.5 (后续还要看下)
	reqtmp := new(openrtb_v2_62.BidRequest)
	err := json.Unmarshal(data, reqtmp)
	if err != nil {
		return nil, err
	}

	// ext
	if reqtmp.Source != nil && reqtmp.Source.SupplyChain == nil {
		supplyChainAdapte(reqtmp.Source) // ext 放到 supplychain
	}
	return reqtmp, nil
}
func (o *OpenRTB25) DenormalizeRequest(ori *openrtb_v2_62.BidRequest) (data []byte, err error) {
	// todo 这里应该是做深度拷贝返回2.5 的[]byte的

	// ext
	if ori.Source != nil && ori.Source.SupplyChain != nil {
		supplyChainAdapte2(ori.Source) // supplychain 放到ext
	}

	//ori.Impressions[0].Native.Version = ""
	//ori.Impressions[0].Native.Request = []byte{}
	//ori.Impressions[0].Native.APIs = []openrtb_v2_6.APIFramework{}
	//ori.Impressions[0].Native.BlockedAttrs = []openrtb_v2_6.CreativeAttribute{}
	//ori.Impressions[0].Native.Ext = json.RawMessage([]byte{})
	//ori.Impressions[0].Native = nil
	data, err = json.Marshal(ori)
	if err != nil {
		return nil, err
	}

	return
}
func (o *OpenRTB25) NormalizeResponse(data []byte) (*openrtb_v2_62.BidResponse, error) {
	// todo domark 先暂时 序列化 2.6兼容2.5 (后续还要看下)
	resptmp := new(openrtb_v2_62.BidResponse)
	err := json.Unmarshal(data, resptmp)
	return resptmp, err
}
func (o *OpenRTB25) DenormalizeResponse(resp *openrtb_v2_62.BidResponse) ([]byte, error) {

	// todo domark 先暂时 序列化 2.6兼容2.5 (后续还要看下)
	return json.Marshal(resp)
}

func getImpressions(os []openrtb_v2_52.Imp) (rs []openrtb_v2_62.Impression) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Impression, 0, len(os))
	for _, imp := range os {
		rimp := openrtb_v2_62.Impression{
			ID:                    imp.ID,
			Banner:                deepCopyBanner(imp.Banner),
			Video:                 deepCopyVideo(imp.Video),
			Audio:                 deepCopyAudio(imp.Audio),
			Native:                deepCopyNative(imp.Native),
			PMP:                   deepCopyPMP(imp.PMP),
			DisplayManager:        imp.DisplayManager,
			DisplayManagerVersion: imp.DisplayManagerVer,
			Interstitial:          int(imp.Instl),
			TagID:                 imp.TagID,
			BidFloor:              imp.BidFloor,
			BidFloorCurrency:      imp.BidFloorCur,
			Quantity:              nil, // 2.5 没有
			Exp:                   int(imp.Exp),
			IFrameBusters:         base.DeepCopyStringSlice(imp.IframeBuster),
			Ext:                   base.DeepCopyByte(imp.Ext),
		}
		if imp.Secure != nil {
			rimp.Secure = openrtb_v2_62.NumberOrString(*imp.Secure)
		}
		rs = append(rs, rimp)
	}
	return rs
}
func deepCopyVideo(o *openrtb_v2_52.Video) (r *openrtb_v2_62.Video) {
	if o == nil {
		return nil
	}
	boxAllowed := int(o.BoxingAllowed)
	r = &openrtb_v2_62.Video{
		MIMEs:       base.DeepCopyStringSlice(o.MIMEs),
		MinDuration: int(o.MinDuration),
		MaxDuration: int(o.MaxDuration),
		Protocols:   deepCopyProtocols(o.Protocols),
		Protocol:    openrtb_v2_62.Protocol(o.Protocol),
		Width:       int(o.W),
		Height:      int(o.H),
		//StartDelay:      openrtb_v2_6.StartDelay(*o.StartDelay),
		Linearity: openrtb_v2_62.VideoLinearity(o.Linearity),
		//Skip:            int(*o.Skip),
		SkipMin:         int(o.SkipMin),
		SkipAfter:       int(o.SkipMin),
		Sequence:        int(o.Sequence),
		BlockedAttrs:    dcCreativeAttribute(o.BAttr),
		MaxExtended:     int(o.MaxExtended),
		MinBitrate:      int(o.MinBitRate),
		MaxBitrate:      int(o.MaxBitRate),
		BoxingAllowed:   &boxAllowed,
		PodID:           "",  // 2.5 没有
		PodDuration:     0,   // 2.5 没有
		PodSequence:     0,   // 2.5 没有
		SlotInPod:       0,   // 2.5 没有
		RqdDurs:         nil, // 2.5 没有
		MinCPMPerSecond: 0,   // 2.5 没有
		PlaybackMethods: deepCopyVideoPlayback(o.PlaybackMethod),
		Delivery:        deepCopyContentDelivery(o.Delivery),
		//Position:        openrtb_v2_6.AdPosition(*o.Pos),
		CompanionAds:   dcCompanionAds(o.CompanionAd),
		APIs:           dcAPIFramework(o.API),
		CompanionTypes: dcCompanionType(o.CompanionType),
		Placement:      openrtb_v2_62.VideoPlacement(o.Placement),
		Plcmt:          0, //  2.5 没有
		Ext:            base.DeepCopyByte(o.Ext),
	}
	if o.Pos != nil {
		r.Position = openrtb_v2_62.AdPosition(*o.Pos)
	}
	if o.StartDelay != nil {
		r.StartDelay = openrtb_v2_62.StartDelay(*o.StartDelay)
	}
	if o.Skip != nil {
		r.Skip = int(*o.Skip)
	}
	return r
}
func deepCopyAudio(o *openrtb_v2_52.Audio) (r *openrtb_v2_62.Audio) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Audio{
		MIMEs:       base.DeepCopyStringSlice(o.MIMEs),
		MinDuration: int(o.MinDuration),
		MaxDuration: int(o.MaxDuration),
		Protocols:   deepCopyProtocols(o.Protocols),
		//StartDelay:     openrtb_v2_6.StartDelay(*o.StartDelay),
		Sequence:       int(o.Sequence),
		BlockedAttrs:   dcCreativeAttribute(o.BAttr),
		MaxExtended:    int(o.MaxExtended),
		MinBitrate:     int(o.MinBitrate),
		MaxBitrate:     int(o.MaxBitrate),
		Delivery:       deepCopyContentDelivery(o.Delivery),
		CompanionAds:   dcCompanionAds(o.CompanionAd),
		APIs:           dcAPIFramework(o.API),
		CompanionTypes: dcCompanionType(o.CompanionType),
		MaxSequence:    int(o.MaxSeq),
		Feed:           openrtb_v2_62.FeedType(o.Feed),
		Stitched:       int(o.Stitched),
		//VolumeNorm:     openrtb_v2_6.VolumeNorm(*o.NVol),
		Ext: base.DeepCopyByte(o.Ext),
	}
	if o.NVol != nil {
		r.VolumeNorm = openrtb_v2_62.VolumeNorm(*o.NVol)
	}
	if o.StartDelay != nil {
		r.StartDelay = openrtb_v2_62.StartDelay(*o.StartDelay)
	}
	return r
}
func deepCopyVideoPlayback(os []openrtb_v2_52.PlaybackMethod) (rs []openrtb_v2_62.VideoPlayback) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.VideoPlayback, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.VideoPlayback(o))
	}
	return rs
}
func deepCopyProtocols(os []openrtb_v2_52.Protocol) (rs []openrtb_v2_62.Protocol) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Protocol, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.Protocol(o))
	}
	return rs
}
func deepCopyNative(o *openrtb_v2_52.Native) (r *openrtb_v2_62.Native) {
	if o == nil {
		return nil
	}
	return &openrtb_v2_62.Native{
		Request:      json.RawMessage(o.Request),
		Version:      o.Ver,
		APIs:         dcAPIFramework(o.API),
		BlockedAttrs: dcCreativeAttribute(o.BAttr),
		Ext:          base.DeepCopyByte(o.Ext),
	}
}
func deepCopyContentDelivery(os []openrtb_v2_52.ContentDeliveryMethod) (rs []openrtb_v2_62.ContentDelivery) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.ContentDelivery, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.ContentDelivery(o))
	}
	return rs
}
func dcCompanionAds(os []openrtb_v2_52.Banner) (rs []openrtb_v2_62.Banner) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Banner, 0, len(os))
	for _, o := range os {
		banner := deepCopyBanner(&o)
		rs = append(rs, *banner)
	}
	return rs
}
func deepCopyBanner(o *openrtb_v2_52.Banner) (r *openrtb_v2_62.Banner) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Banner{
		//Width:        int(*o.W),
		//Height:       int(*o.H),
		Formats:      deepCopyFormats(o.Format),
		WidthMax:     int(o.WMax),
		HeightMax:    int(o.HMax),
		WidthMin:     int(o.WMin),
		HeightMin:    int(o.HMin),
		ID:           o.ID,
		BlockedTypes: deepCopyBlockedTypes(o.BType),
		BlockedAttrs: dcCreativeAttribute(o.BAttr),
		//Position:     openrtb_v2_6.AdPosition(*o.Pos),
		MIMEs:    base.DeepCopyStringSlice(o.MIMEs),
		TopFrame: int(o.TopFrame),
		ExpDirs:  deepCopyExpDirs(o.ExpDir),
		APIs:     dcAPIFramework(o.API),
		VCM:      int(o.VCm),
		Ext:      base.DeepCopyByte(o.Ext),
	}
	if o.Pos != nil {
		r.Position = openrtb_v2_62.AdPosition(*o.Pos)
	}
	if o.W != nil {
		r.Width = int(*o.W)
	}
	if o.H != nil {
		r.Height = int(*o.H)
	}
	return r
}
func deepCopyFormats(os []openrtb_v2_52.Format) (rs []openrtb_v2_62.Format) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Format, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.Format{
			Width:       int(o.W),
			Height:      int(o.H),
			WidthRatio:  int(o.WRatio),
			HeightRatio: int(o.HRatio),
			WidthMin:    int(o.WMin),
			Ext:         base.DeepCopyByte(o.Ext),
		})
	}
	return rs
}
func deepCopyExpDirs(os []openrtb_v2_52.ExpandableDirection) (rs []openrtb_v2_62.ExpDir) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.ExpDir, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.ExpDir(o))
	}
	return rs
}
func deepCopyBlockedTypes(os []openrtb_v2_52.BannerAdType) (rs []openrtb_v2_62.BannerType) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.BannerType, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.BannerType(o))
	}
	return rs
}
func dcCompanionType(os []openrtb_v2_52.CompanionType) (rs []openrtb_v2_62.CompanionType) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.CompanionType, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.CompanionType(o))
	}
	return rs
}
func dcAPIFramework(os []openrtb_v2_52.APIFramework) (rs []openrtb_v2_62.APIFramework) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.APIFramework, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.APIFramework(o))
	}
	return rs
}
func dcCreativeAttribute(os []openrtb_v2_52.CreativeAttribute) (rs []openrtb_v2_62.CreativeAttribute) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.CreativeAttribute, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.CreativeAttribute(o))
	}
	return rs
}
func deepCopyPMP(o *openrtb_v2_52.PMP) (r *openrtb_v2_62.PMP) {
	if o == nil {
		return nil
	}
	return &openrtb_v2_62.PMP{
		Private: int(o.PrivateAuction),
		Deals:   deepCopyDeals(o.Deals),
		Ext:     base.DeepCopyByte(o.Ext),
	}
}
func deepCopyDeals(os []openrtb_v2_52.Deal) (rs []openrtb_v2_62.Deal) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Deal, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.Deal{
			ID:               o.ID,
			BidFloor:         o.BidFloor,
			BidFloorCurrency: o.BidFloorCur,
			Seats:            base.DeepCopyStringSlice(o.WSeat),
			AdvDomains:       base.DeepCopyStringSlice(o.WADomain),
			AuctionType:      int(o.AT),
			Ext:              base.DeepCopyByte(o.Ext),
		})
	}
	return
}
func deepCopyPublisher(o *openrtb_v2_52.Publisher) (r *openrtb_v2_62.Publisher) {
	if o == nil {
		return nil
	}
	p := openrtb_v2_62.Publisher(openrtb_v2_62.ThirdParty{
		ID:         o.ID,
		Name:       o.Name,
		Categories: deepCopyCategories(o.Cat),
		Domain:     o.Domain,
		Ext:        base.DeepCopyByte(o.Ext),
	})
	return &p
}
func deepCopyProducer(o *openrtb_v2_52.Producer) (r *openrtb_v2_62.Producer) {
	if o == nil {
		return nil
	}
	p := openrtb_v2_62.Producer(openrtb_v2_62.ThirdParty{
		ID:         o.ID,
		Name:       o.Name,
		Categories: deepCopyCategories(o.Cat),
		Domain:     o.Domain,
		Ext:        base.DeepCopyByte(o.Ext),
	})
	return &p
}
func deepCopyContent(o *openrtb_v2_52.Content) (r *openrtb_v2_62.Content) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Content{
		ID:               o.ID,
		Episode:          int(o.Episode),
		Title:            o.Title,
		Series:           o.Series,
		Season:           o.Season,
		Artist:           o.Artist,
		Genre:            o.Genre,
		Album:            o.Album,
		ISRC:             o.ISRC,
		Producer:         deepCopyProducer(o.Producer),
		URL:              o.URL,
		CategoryTaxonomy: 0, // 2.5 没有这个字段，todo 在ext 中解析下
		Categories:       deepCopyCategories(o.Cat),
		//ProductionQuality:  openrtb_v2_6.ProductionQuality(*o.ProdQ),
		VideoQuality:       0,
		Context:            openrtb_v2_62.ContentContext(o.Context),
		ContentRating:      o.ContentRating,
		UserRating:         o.UserRating,
		MediaRating:        openrtb_v2_62.IQGRating(o.QAGMediaRating),
		Keywords:           o.Keywords,
		LiveStream:         int(o.LiveStream),
		SourceRelationship: int(o.SourceRelationship),
		Length:             int(o.Len),
		Language:           o.Language,
		LanguageB:          "", // todo 2.5 没有
		Embeddable:         int(o.Embeddable),
		Data:               deepcopyData(o.Data),
		Network:            nil, // 2.5 没有
		Channel:            nil, // 2.5 没有
		KwArray:            nil, // 2.5 没有
		Ext:                base.DeepCopyByte(o.Ext),
	}
	if o.ProdQ != nil {
		r.ProductionQuality = openrtb_v2_62.ProductionQuality(*o.ProdQ)
	}
	return r
}

// func deepcopyChannelEntity(o json.RawMessage) (r *openrtb_v2_6.ChannelEntity) {}
func getSite(o *openrtb_v2_52.Site) (r *openrtb_v2_62.Site) {
	if o == nil {
		return nil
	}
	privacyPolicy := int(o.PrivacyPolicy)
	return &openrtb_v2_62.Site{
		Inventory: openrtb_v2_62.Inventory{
			ID:                o.ID,
			Name:              o.Name,
			Domain:            o.Domain,
			Categories:        deepCopyCategories(o.Cat),
			SectionCategories: deepCopyCategories(o.SectionCat),
			PageCategories:    deepCopyCategories(o.PageCat),
			PrivacyPolicy:     &privacyPolicy,
			Publisher:         deepCopyPublisher(o.Publisher),
			Content:           deepCopyContent(o.Content),
			Keywords:          o.Keywords,
			Ext:               base.DeepCopyByte(o.Ext),
		},
		Page:     o.Page,
		Referrer: o.Ref,
		Search:   o.Search,
		Mobile:   int(o.Mobile),
	}
}
func getApp(o *openrtb_v2_52.App) (r *openrtb_v2_62.App) {
	if o == nil {
		return nil
	}
	privacyPolicy := int(o.PrivacyPolicy)
	return &openrtb_v2_62.App{
		Inventory: openrtb_v2_62.Inventory{
			ID:                o.ID,
			Name:              o.Name,
			Domain:            o.Domain,
			Categories:        deepCopyCategories(o.Cat),
			SectionCategories: deepCopyCategories(o.SectionCat),
			PageCategories:    deepCopyCategories(o.PageCat),
			PrivacyPolicy:     &privacyPolicy,
			Publisher:         deepCopyPublisher(o.Publisher),
			Content:           deepCopyContent(o.Content),
			Keywords:          o.Keywords,
			Ext:               base.DeepCopyByte(o.Ext),
		},
		Bundle:   o.Bundle,
		StoreURL: o.StoreURL,
		Version:  o.Ver,
		Paid:     int(o.Paid),
	}
}
func getDevice(o *openrtb_v2_52.Device) (r *openrtb_v2_62.Device) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Device{
		UA:           o.UA,
		Sua:          nil,
		Geo:          nil,
		IP:           o.IP,
		IPv6:         o.IPv6,
		DeviceType:   openrtb_v2_62.DeviceType(o.DeviceType),
		Make:         o.Make,
		Model:        o.Model,
		OS:           o.OS,
		OSVersion:    o.OSV,
		HWVersion:    o.HWV,
		Height:       int(o.H),
		Width:        int(o.W),
		PPI:          int(o.PPI),
		PixelRatio:   o.PxRatio,
		JS:           int(o.JS),
		GeoFetch:     int(o.GeoFetch),
		FlashVersion: o.FlashVer,
		Language:     o.Language,
		LanguageB:    "", // todo 2.5 没有这个字段
		Carrier:      o.Carrier,
		MCCMNC:       o.MCCMNC,
		//ConnType:     openrtb_v2_6.ConnType(*o.ConnectionType),
		IFA:     o.IFA,
		IDSHA1:  o.DIDSHA1,
		IDMD5:   o.DIDMD5,
		PIDSHA1: o.DPIDSHA1,
		PIDMD5:  o.DPIDMD5,
		MacSHA1: o.MACSHA1,
		MacMD5:  o.MACMD5,
		Ext:     base.DeepCopyByte(o.Ext),
	}
	if o.DNT != nil {
		r.DNT = int(*o.DNT)
	}
	if o.Lmt != nil {
		r.LMT = int(*o.Lmt)
	}
	if o.ConnectionType != nil {
		r.ConnType = openrtb_v2_62.ConnType(*o.ConnectionType)
	}
	return r
}
func getUser(o *openrtb_v2_52.User) (r *openrtb_v2_62.User) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.User{
		ID:          o.ID,
		BuyerID:     o.BuyerUID,
		BuyerUID:    o.BuyerUID,
		YearOfBirth: int(o.Yob),
		Gender:      o.Gender,
		Keywords:    o.Keywords,
		CustomData:  o.CustomData,
		Geo:         dcGeo(o.Geo),
		Data:        deepcopyData(o.Data),
		Ext:         base.DeepCopyByte(o.Ext),
	}
	return r
}
func dcGeo(o *openrtb_v2_52.Geo) (r *openrtb_v2_62.Geo) {
	if o == nil {
		return nil
	}
	return &openrtb_v2_62.Geo{
		Latitude:      o.Lat,
		Longitude:     o.Lon,
		Type:          openrtb_v2_62.LocationType(o.Type),
		Accuracy:      int(o.Accuracy),
		LastFix:       int(o.LastFix),
		IPService:     openrtb_v2_62.IPLocation(o.IPService),
		Country:       o.Country,
		Region:        o.Region,
		RegionFIPS104: o.RegionFIPS104,
		Metro:         o.Metro,
		City:          o.City,
		ZIP:           o.ZIP,
		UTCOffset:     int(o.UTCOffset),
		Ext:           base.DeepCopyByte(o.Ext),
	}
}
func deepcopyData(os []openrtb_v2_52.Data) (rs []openrtb_v2_62.Data) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Data, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.Data{
			ID:      o.ID,
			Name:    o.Name,
			Segment: dcSegment(o.Segment),
			Ext:     base.DeepCopyByte(o.Ext),
		})
	}
	return rs
}
func dcSegment(os []openrtb_v2_52.Segment) (rs []openrtb_v2_62.Segment) {
	if os == nil {
		return nil
	}
	rs = make([]openrtb_v2_62.Segment, 0, len(os))
	for _, o := range os {
		rs = append(rs, openrtb_v2_62.Segment{
			ID:    o.ID,
			Name:  o.Name,
			Value: o.Value,
			Ext:   base.DeepCopyByte(o.Ext),
		})
	}
	return rs
}
func deepCopyCategories(original []string) (copiedSlice []openrtb_v2_62.ContentCategory) {
	if original == nil {
		return nil
	}
	copiedSlice = make([]openrtb_v2_62.ContentCategory, len(original))
	for i, v := range original {
		copiedSlice[i] = openrtb_v2_62.ContentCategory(v)
	}
	return copiedSlice
}
func deepCopySource(o *openrtb_v2_52.Source) (r *openrtb_v2_62.Source) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Source{
		FinalSaleDecision: int(o.FD),
		TransactionID:     o.TID,
		PaymentChain:      o.PChain,
		SupplyChain:       nil,
		Ext:               base.DeepCopyByte(o.Ext),
	}
	// todo parse ext

	var ext map[string]json.RawMessage
	json.Unmarshal(o.Ext, &ext)
	if d, ok := ext["schain"]; ok {
		r.SupplyChain = new(openrtb_v2_62.SupplyChain)
		json.Unmarshal(d, r.SupplyChain)
	}
	return r
}
func deepCopyRegulations(o *openrtb_v2_52.Regs) (r *openrtb_v2_62.Regulations) {
	if o == nil {
		return nil
	}
	r = &openrtb_v2_62.Regulations{
		COPPA:     int(o.COPPA),
		GDPR:      0, // 2.5 没有 todo 后续在ext 中找
		USPrivacy: "",
		Ext:       base.DeepCopyByte(o.Ext),
	}
	// todo parse ext
	return r
}

func deepCopySource2(o *openrtb_v2_62.Source) (r *openrtb_v2_52.Source) {
	if o == nil {
		return nil
	}

	r = &openrtb_v2_52.Source{
		FD:     int8(o.FinalSaleDecision),
		TID:    o.TransactionID,
		PChain: o.PaymentChain,
		Ext:    nil,
	}

	if o.SupplyChain != nil {
		sc, _ := json.Marshal(o.SupplyChain)
		var ext = make(map[string]json.RawMessage, 2)
		json.Unmarshal(o.Ext, &ext)
		ext["schain"] = sc
		extRaw, _ := json.Marshal(ext)
		r.Ext = base.DeepCopyByte(extRaw)
	}
	return r
}

// 2.6 的 supplychain 放到 2.5 的 ext
func supplyChain26to25(o *openrtb_v2_62.Source, r *openrtb_v2_52.Source) (err error) {
	sc, _ := json.Marshal(o.SupplyChain)
	var ext = make(map[string]json.RawMessage, 2)
	json.Unmarshal(o.Ext, &ext)
	ext["schain"] = sc
	extRaw, err := json.Marshal(ext)
	if err != nil {
		return nil
	}
	r.Ext = extRaw
	return nil
}

// 2.5 的 ext 放到 2.6 的 supplychain
func supplyChain25to26(o *openrtb_v2_52.Source, r *openrtb_v2_62.Source) (err error) {
	//sc, _ := json.Marshal(o.SupplyChain)
	var ext = make(map[string]json.RawMessage, 2)
	json.Unmarshal(o.Ext, &ext)
	sc := ext["schain"]
	json.Unmarshal(sc, &r.SupplyChain)
	delete(ext, "schain")
	extRaw, err := json.Marshal(ext)
	if err != nil {
		return nil
	}
	r.Ext = extRaw
	return nil
}

// 2.6 的 ext 放到 supplyChain
func supplyChainAdapte(o *openrtb_v2_62.Source) (err error) {
	//sc, _ := json.Marshal(o.SupplyChain)
	var ext = make(map[string]json.RawMessage, 2)
	err = json.Unmarshal(o.Ext, &ext)
	if err != nil {
		return err
	}
	sc := ext["schain"]
	json.Unmarshal(sc, &o.SupplyChain)
	delete(ext, "schain")
	extRaw, err := json.Marshal(ext)
	if err != nil {
		return nil
	}
	o.Ext = extRaw
	return nil
}

// 2.6 的 supplyChain 放到 ext
func supplyChainAdapte2(o *openrtb_v2_62.Source) (err error) {
	sc, _ := json.Marshal(o.SupplyChain)
	var ext = make(map[string]json.RawMessage, 2)
	err = json.Unmarshal(o.Ext, &ext)
	if err != nil {
		return err
	}
	ext["schain"] = sc
	o.SupplyChain = nil
	extRaw, err := json.Marshal(ext)
	if err != nil {
		return nil
	}
	o.Ext = extRaw
	return nil
}
