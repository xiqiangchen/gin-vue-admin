package openrtb_v2_5

const (
	AUCTION_PRICE = "${AUCTION_PRICE}"
)

/*
${AUCTION_ID} ID of the bid request; from BidRequest.id attribute.
${AUCTION_BID_ID} ID of the bid; from BidResponse.bidid attribute.
${AUCTION_IMP_ID} ID of the impression just won; from imp.id attribute.
${AUCTION_SEAT_ID} ID of the bidder seat for whom the bid was made.
${AUCTION_AD_ID} ID of the ad markup the bidder wishes to serve; from bid.adid attribute.
${AUCTION_PRICE} Clearing price using the same currency and units as the bid.
${AUCTION_CURRENCY} The currency used in the bid (explicit or implied); for confirmation only.
${AUCTION_MBR} Market Bid Ratio defined as: clearance price / bid price.
${AUCTION_LOSS} Loss reason codes. Refer to List 5.25.
*/
