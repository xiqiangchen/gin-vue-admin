package openrtb2

const (
	AUCTION_ID         = "${AUCTION_ID}"         //ID of the bid request; from BidRequest.id attribute.
	AUCTION_BID_ID     = "${AUCTION_BID_ID}"     //ID of the bid; from BidResponse.bidid attribute.
	AUCTION_IMP_ID     = "${AUCTION_IMP_ID}"     //ID of the impression just won; from imp.id attribute.
	AUCTION_SEAT_ID    = "${AUCTION_SEAT_ID}"    //ID of the bidder seat for whom the bid was made.
	AUCTION_AD_ID      = "${AUCTION_AD_ID}"      //ID of the ad markup the bidder wishes to serve; from bid.adid attribute.
	AUCTION_PRICE      = "${AUCTION_PRICE}"      //Clearing price using the same currency and units as the bid.
	AUCTION_CURRENCY   = "${AUCTION_CURRENCY}"   //The currency used in the bid (explicit or implied); for confirmation only.
	AUCTION_MBR        = "${AUCTION_MBR}"        //Market Bid Ratio defined as: clearance price / bid price.
	AUCTION_LOSS       = "${AUCTION_LOSS}"       //Loss reason codes. Refer to List: Loss Reason Codes in OpenRTB 3.0.
	AUCTION_MIN_TO_WIN = "${AUCTION_MIN_TO_WIN}" //Minimum bid to win the exchange's auction, using the same currency andunits as the bid.
)

// 定义设备类型常量
const (
	MobileTabletVersion2        = 1 // Mobile/Tablet Version 2.0
	PersonalComputerVersion2    = 2 // Personal Computer Version 2.0
	ConnectedTVVersion2         = 3 // Connected TV Version 2.0
	PhoneNewVersion22           = 4 // Phone New for Version 2.2
	TabletNewVersion22          = 5 // Tablet New for Version 2.2
	ConnectedDeviceNewVersion22 = 6 // Connected Device New for Version 2.2
	SetTopBoxNewVersion22       = 7 // Set Top Box New for Version 2.2
)

// 定义网络类型常量
const (
	UnknownNetwork            = 0 // Unknown
	EthernetNetwork           = 1 // Ethernet
	WifiNetwork               = 2 // WIFI
	CellularUnknownGeneration = 3 // Cellular Network – Unknown Generation
	Cellular2GNetwork         = 4 // Cellular Network – 2G
	Cellular3GNetwork         = 5 // Cellular Network – 3G
	Cellular4GNetwork         = 6 // Cellular Network – 4G
)

const (
	MarkupTypeBanner = 1 // banner
	MarkupTypeVideo  = 2 // video
	MarkupTypeAudio  = 3 //
	MarkupTypeNative = 4 // native
)
