package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"xzdp/biz/model/shop"
	"xzdp/biz/pkg/constants"

	"github.com/go-redis/redis/v8"
)

func GetShopFromCache(ctx context.Context, key string) (*shop.Shop, error) {
	shopJson, err := RedisClient.Get(ctx, key).Result()
	if err == nil && shopJson != "" {
		var shop shop.Shop
		if err := json.Unmarshal([]byte(shopJson), &shop); err != nil {
			return nil, err
		}
		return &shop, nil
	}

	if err != redis.Nil {
		return nil, err
	}

	if shopJson == "" {
		return nil, errors.New("shop not found in cache")
	}

	return nil, errors.New("unknown error")
}

func QueryShopWithDistance(ctx context.Context, req *shop.ShopOfTypeGeoReq) (resp *[]*shop.Shop, err error) {
	current := req.Current
	from := (current - 1) * constants.DEFAULT_PAGE_SIZE
	end := current * constants.DEFAULT_PAGE_SIZE

	key := constants.SHOP_GEO_KEY + strconv.Itoa(int(req.TypeId))
	geoRadiusQuery := redis.GeoRadiusQuery{
		Radius:   req.Distance,
		WithDist: true,
		Sort:     "ASC",
		Count:    int(end),
	}
	locations, err := RedisClient.GeoRadius(ctx, key, req.Longitude, req.Latitude, &geoRadiusQuery).Result()

	if err != nil {
		fmt.Printf("Error querying Redis: %v\n", err)
		return nil, err
	}
	if len(locations) == 0 {
		return &[]*shop.Shop{}, nil
	}
	if len(locations) <= int(from) {
		return &[]*shop.Shop{}, nil
	}

	shops := make([]*shop.Shop, 0, len(locations))
	for _, loc := range locations[from:end] {
		id, _ := strconv.ParseInt(loc.Name, 10, 64)
		shops = append(shops, &shop.Shop{
			ID:       id,
			X:        loc.Longitude,
			Y:        loc.Latitude,
			Distance: loc.Dist,
			TypeId:   int64(req.TypeId),
		})
	}
	return &shops, nil
}
