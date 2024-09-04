package voucher

import (
	"context"
	"errors"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	voucher "xzdp/biz/model/voucher"
	service "xzdp/biz/service/voucher"
	"xzdp/biz/utils"
)

// VoucherList .
// @router /voucher/list/:id [GET]
func VoucherList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req voucher.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewVoucherListService(ctx, c).Run(id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// SeckillVoucher .
// @router /voucher-order/seckill/:id [POST]
func SeckillVoucher(ctx context.Context, c *app.RequestContext) {
	var err error
	voucherString := c.Param("id")
	if voucherString == "" {
		utils.SendErrResponse(ctx, c, consts.StatusOK, errors.New("参数错误"))
		return
	}
	voucherID, err := strconv.ParseInt(voucherString, 10, 64)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	resp, err := service.NewSeckillVoucherService(ctx, c).Run(&voucherID)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
