package base

// 常用的是2.0和4.0 。后续有需要再加

// 有必要转换吗？直接将dsp的vast 加上 监测链接后，返回给adx 就行了

type Common struct{}

func (c *Common) GetTrackInfo(adm string) (vi *VASTInfo, err error) {
	return
}
func (c *Common) SetTrackInfo(adm string, vi *VASTInfo) (result string, err error) {
	return
}

func (c *Common) GetProtcol() []int {
	return []int{}
}
