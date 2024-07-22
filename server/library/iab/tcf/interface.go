package tcf

type TCFProxy interface {
	ParseConsent(consent string) error // 解析 tcf 的同意字符串

	IsAllowed(purpose []int, vendor int) bool // 检查 用户是否允许我们使用他的数据

	IsCMPValid(cmpid int) bool // 检查cmp是否合法，虽然流量来源不关我们的事，但是避免被牵连，不合法的流量最好不要

	GetCMPInfo() (cmpid int, cmpversion int)
}
