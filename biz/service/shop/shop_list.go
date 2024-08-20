package shop

import (
	"context"
	"encoding/json"
	"log"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"

	"xzdp/biz/dal/mysql"
	shop "xzdp/biz/model/shop"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type ShopListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewShopListService(Context context.Context, RequestContext *app.RequestContext) *ShopListService {
	return &ShopListService{RequestContext: RequestContext, Context: Context}
}

func (h *ShopListService) Run(req *shop.Empty) (resp *[]*shop.ShopType, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	// todo edit your code

	var shopTypeList []*shop.ShopType

	// 从 Redis 获取数据
	shopTypeJsonList, err := redis.RedisClient.LRange(h.Context, constants.CACHE_SHOP_TYPE_LIST_KEY, 0, -1).Result()
	if err == nil && len(shopTypeJsonList) > 0 {
		for _, shopTypeJson := range shopTypeJsonList {
			var shopType shop.ShopType
			err := json.Unmarshal([]byte(shopTypeJson), &shopType)
			if err != nil {
				hlog.CtxErrorf(h.Context, "Error unmarshalling shop type: %v", err)
				continue
			}
			shopTypeList = append(shopTypeList, &shopType)
		}
		return &shopTypeList, nil
	}

	hlog.CtxInfof(h.Context, "hello asdfas yangqi")

	// 如果 Redis 没有数据，从数据库获取数据
	shopTypeList, err = mysql.QueryShopType(h.Context)
	if err != nil {
		return nil, err
	}

	// 将数据存储到 Redis
	for _, shopType := range shopTypeList {
		shopTypeJson, err := json.Marshal(shopType)
		if err != nil {
			log.Printf("Error marshalling shop type: %v", err)
			continue
		}
		redis.RedisClient.RPush(h.Context, constants.CACHE_SHOP_TYPE_LIST_KEY, shopTypeJson)
	}

	return &shopTypeList, nil
}
