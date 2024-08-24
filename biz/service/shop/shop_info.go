package shop

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"

	shop "xzdp/biz/model/shop"

	"xzdp/biz/dal/mysql"

	"github.com/cloudwego/hertz/pkg/app"
)

type ShopInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopInfoService(Context context.Context, RequestContext *app.RequestContext) *ShopInfoService {
	return &ShopInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopInfoService) Run(id int64) (resp *shop.Shop, err error) {
	key := fmt.Sprintf("%s%d", constants.CACHE_SHOP_KEY, id)

	result, err := redis.GetStringLogical(h.Context, key, constants.CACHE_SHOP_TTL, WrappedQueryByID, h.Context, id)
	if err != nil || result == "" {
		return nil, err
	}

	err = json.Unmarshal([]byte(result), &resp)
	if err != nil {
		return nil, err
	}

	//hlog.Debugf("test! : %s", resp)

	return resp, nil
}

func WrappedQueryByID(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("incorrect number of arguments")
	}

	ctx, ok := args[0].(context.Context)
	if !ok {
		return nil, fmt.Errorf("first argument is not context.Context")
	}

	id, ok := args[1].(int64)
	if !ok {
		return nil, fmt.Errorf("second argument is not int")
	}

	result, err := mysql.QueryByID(ctx, id)
	return result, err
}
