package constant

const (
	X_ORTB_VERSION = "x-openrtb-version"
)

const (
	SEAT_BID_TYPE_EXCLUSIVE = 0
	SEAT_BID_TYPE_OPTIONAL  = 1
	SEAT_BID_TYPE_MIXED     = 2
)

// nbr
const (
	UnknownError             = 0 // 未知错误
	TechnicalError           = 1 // 技术错误
	InvalidRequest           = 2 // 非法请求
	KnownWebSpider           = 3 // 已知的网络爬虫
	SuspectedNonHumanTraffic = 4 // 怀疑非人为请求
	CloudDataCenterOrProxyIP = 5 // 来自云，数据中心或者代理的IP
	UnsupportedDevice        = 6 // 不支持的设备类型
	BlockedPublisherOrSite   = 7 // 来自受限展示者或站点
	UnmatchedUser            = 8 // 用户不匹配
)

// todo Group
const ()

///////// DigitalServicesAct /////////

const (
	DSADomain = `oskey.cn` // 公司的域名

	VendorId = 333 // todo domark 公司申请的vendorid
)

var (
	// 声明会如何使用用户数据
	DSAParams = []int{1, 2, 3}

	TCFPurpose = []int{2, 3, 4, 6, 7, 8, 9}
)

var DefaultCat = []string{"IAB3", "IAB3-1", "IAB22", "IAB22-1"}
