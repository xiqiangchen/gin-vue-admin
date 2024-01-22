package dsp

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

type BidApi struct {
}

var bidService = service.ServiceGroupApp.DspGroup.BidService

func (bidApi *BidApi) Req(c *gin.Context) {
	adx, ok := c.GetQuery("adx")
	if !ok {
		response.IllegalWithMessage("adx必填", c)
		return
	}
	adxId, _ := strconv.Atoi(adx)
	var resp []byte
	var offer bool
	if bodyBytes, err := io.ReadAll(c.Request.Body); err != nil {
		response.NoContent(c)
		global.GVA_LOG.Error("渠道解释失败", zap.Error(err))
		return
	} else if resp, offer = bidService.Bid(adxId, bodyBytes); !offer {
		response.NoContent(c)
		return
	} else {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write(resp)
		return
	}
}
