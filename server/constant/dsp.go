package constant

const (
	StatusOn  = 1
	StatusOff = 0

	Virtually = 1
	Real      = 0
	Filter    = 1
	Pass      = 0

	BidModeFixed = 1
	BidModeAvg   = 2

	BidMethodCpm = 1
	BidMethodCpc = 2
	BidMethodCpa = 3

	MarkupTypeBanner = 1
)

const (
	DefaultAdxId      = 1001
	DefaultProtocol   = 0
	DefaultPriceMacro = "%%BSW_PRICE%%"
)

const (
	DeviceIdOaid      = "oaid"
	DeviceIdIdfa      = "idfa"
	DeviceIdGaid      = "gaid"
	DeviceIdImei      = "imei"
	DeviceIdAndroidId = "andid"
	DeviceIdMd5Imei   = "mimei"
	DeviceIdMd5Idfa   = "midfa"
	DeviceIdMd5Oaid   = "moaid"
)

const (
	DspImpTrack     = "${DSP_IMP_TRACK}"
	DspClkTrack     = "${DSP_CLK_TRACK}"
	DspLdTrack      = "${DSP_LD_TRACK}"
	DspBundle       = "${DSP_BUNDLE}"
	DspPublisher    = "${DSP_PUBLISHER}"
	DspOfferDayHour = "${DSP_OFFER_DAY_HOUR}"
)
