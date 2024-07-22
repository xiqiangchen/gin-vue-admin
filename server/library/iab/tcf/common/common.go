package common

import (
	"fmt"
	"github.com/LiveRamp/iabconsent"
	"github.com/flipped-aurora/gin-vue-admin/server/library/iab/tcf/cml"
	"strconv"
)

type CommonTcfConsentStringProxy struct {
	V2 *iabconsent.V2ParsedConsent
}

func (c *CommonTcfConsentStringProxy) ParseConsent(consent string) (err error) {
	version := iabconsent.TCFVersionFromTCString(consent)
	switch version {
	case iabconsent.V1:
		//var v1, err = iabconsent.ParseV1(consent)
		// Use v1 consent.
		return fmt.Errorf("invalid consent version v%s", strconv.Itoa(int(version)))
	case iabconsent.V2:
		v2p, err := iabconsent.ParseV2(consent)
		if err != nil {
			return err
		}
		c.V2 = v2p
		return nil
	default:
		return fmt.Errorf("invalid consent version v%s", strconv.Itoa(int(version)))
	}
	return nil
}
func (c *CommonTcfConsentStringProxy) IsAllowed(purposes []int, vendor int) bool {
	fmt.Println(c.V2.VendorAllowed(vendor))
	fmt.Println(c.V2.EveryPurposeAllowed(purposes))
	fmt.Println(c.V2.PublisherRestricted(purposes, vendor))
	return c.V2.SuitableToProcess(purposes, vendor)
}

// 单独抽出去是否比较好？
func (c *CommonTcfConsentStringProxy) IsCMPValid(cmpid int) bool {
	if cml.CML == nil {
		return false
	}
	_, exist := cml.CML.Cmps[strconv.Itoa(cmpid)]
	return exist
}

func (c *CommonTcfConsentStringProxy) GetCMPInfo() (int, int) {
	return c.V2.CMPID, c.V2.CMPVersion
}
