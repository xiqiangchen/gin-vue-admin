package base

type VASTInfo struct {
	ImpTrack   []string
	ClickTrack []string
	// todo 其他先不管，先处理曝光和点击的部分
	ClickThrough []string

	EventTrack map[string][]string
	// todo 还有播放事件的监测。
	/*
		creativeView,
		start, midpoint,
		firstQuartile,
		thirdQuartile,
		complete, mute,
		unmute, pause,
		rewind, resume,
		fullscreen,
		expand, collapse,
		acceptInvitation,
		close
	*/

}

const (
	VAST10         = 1
	VAST20         = 2
	VAST30         = 3
	VAST10Wrapper  = 4
	VAST20Wrapper  = 5
	VAST30Wrapper  = 6
	VAST40         = 7
	VAST40Wrapper  = 8
	DAAST10        = 9
	DAAST10Wrapper = 10
	VAST41         = 11
	VAST41Wrapper  = 12
	VAST42         = 13
	VAST42Wrapper  = 14
)
