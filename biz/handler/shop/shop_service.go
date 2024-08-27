package shop

import (
	"context"
	"strconv"

	shop "xzdp/biz/model/shop"
	service "xzdp/biz/service/shop"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// ShopList .
// @router /shop-type/list [GET]
func ShopList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req shop.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewShopListService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ShopOfType .
// @router /shop/of/type [GET]
func ShopOfType(ctx context.Context, c *app.RequestContext) {
	var err error
	var req shop.ShopOfTypeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewShopOfTypeService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ShopInfo .
// @router /shop/:id [GET]
func ShopInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req shop.Empty
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

	resp, err := service.NewShopInfoService(ctx, c).Run(id)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ShopOfTypeGeo .
// @router /shop/of/type/geo [GET]
func ShopOfTypeGeo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req shop.ShopOfTypeGeoReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewShopOfTypeGeoService(ctx, c).Run(&req)

	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
