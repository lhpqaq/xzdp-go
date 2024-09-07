package mysql

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"xzdp/biz/dal/redis"
	"xzdp/biz/model/shop"
	"xzdp/biz/pkg/constants"

	redis2 "github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func QueryShopType(ctx context.Context) ([]*shop.ShopType, error) {
	var shopTypeList []*shop.ShopType

	err = DB.WithContext(ctx).Order("sort asc").Find(&shopTypeList).Error
	if err != nil {
		return nil, err
	}

	if len(shopTypeList) == 0 {
		return nil, fmt.Errorf("no exist shop type")
	}

	return shopTypeList, nil
}

func QueryByID(ctx context.Context, id int64) (*shop.Shop, error) {
	var shop shop.Shop
	err = DB.WithContext(ctx).First(&shop, id).Error

	if err != nil {
		return nil, err
	}
	if shop.ID == 0 {
		return nil, fmt.Errorf("shop isn't exist")
	}

	return &shop, nil
}

func QueryAllShop(ctx context.Context, shops *[]*shop.Shop) (*[]*shop.Shop, error) {
	for i := range *shops {
		dist := (*shops)[i].Distance
		shop, err := QueryByID(ctx, (*shops)[i].ID)
		if err != nil {
			return nil, err
		}
		shop.Distance = dist
		(*shops)[i] = shop

	}
	return shops, nil
}

func queryByID1(ctx context.Context, id int64) (*shop.Shop, error) {
	key := fmt.Sprintf("%s%d", constants.CACHE_SHOP_KEY, id)

	// Query cache from Redis
	shopJson, err := redis.RedisClient.Get(ctx, key).Result()
	if err == nil && shopJson != "" {
		var shop shop.Shop
		if err := json.Unmarshal([]byte(shopJson), &shop); err != nil {
			return nil, err
		}
		return &shop, nil
	}

	if err != redis2.Nil {
		return nil, err
	}

	// Query database
	var shop shop.Shop
	if err := DB.First(&shop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			redis.RedisClient.Set(ctx, key, "", constants.CACHE_NULL_TTL).Err()
		}
		return nil, err
	}

	// Cache the result
	shopJsonByte, err := json.Marshal(shop)
	shopJson = string(shopJsonByte)
	if err != nil {
		return nil, err
	}
	redis.RedisClient.Set(ctx, key, shopJson, constants.CACHE_SHOP_TTL).Err()

	return &shop, nil
}
func queryByID2(ctx context.Context, id int64) (*shop.Shop, error) {
	key := fmt.Sprintf("%s%d", constants.CACHE_SHOP_KEY, id)
	lockKey := fmt.Sprintf("%s%d", constants.LOCK_SHOP_KEY, id)

	// 1. 从 Redis 获取数据
	result, err := redis.GetShopFromCache(ctx, key)
	if err == nil {
		return result, nil
	}
	lock := redis.NewLock(ctx, lockKey, key, 10)
	// 2. 缓存未命中，尝试获取锁
	isLocked := lock.TryLock()
	defer lock.UnLock(key)
	if !isLocked {
		// 锁获取失败，等待后重试
		time.Sleep(50 * time.Millisecond)
		return queryByID2(ctx, id)
	}

	// 2.2 获取锁成功，再次检查缓存
	result, err = redis.GetShopFromCache(ctx, key)
	if err == nil {
		return result, nil
	}

	// 3. 从数据库查询数据
	var shop shop.Shop
	if err := DB.First(&shop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			redis.RedisClient.Set(ctx, key, "", constants.CACHE_NULL_TTL).Err()
			return nil, err
		}
		return nil, err
	}

	// 4. 数据库中存在，缓存数据
	shopJson, err := json.Marshal(shop)
	if err != nil {
		return nil, err
	}
	redis.RedisClient.Set(ctx, key, string(shopJson), constants.CACHE_SHOP_TTL).Err()
	return &shop, nil
}

func LoadShopListToCache(ctx context.Context) error {
	shopMap := make(map[int64][]shop.Shop)
	var shops []*shop.Shop
	err := DB.Find(&shops).Error
	if err != nil {
		return err
	}
	for _, v := range shops {
		shopMap[v.TypeId] = append(shopMap[v.TypeId], *v)
	}
	for typeId, shops := range shopMap {
		key := constants.SHOP_GEO_KEY + fmt.Sprint(typeId)

		var locations []*redis2.GeoLocation
		for _, shop := range shops {
			location := &redis2.GeoLocation{
				Name:      fmt.Sprint(shop.ID),
				Longitude: shop.X,
				Latitude:  shop.Y,
			}
			locations = append(locations, location)
		}

		_, err := redis.RedisClient.GeoAdd(ctx, key, locations...).Result()
		if err != nil {
			return err
		}
	}
	return nil
}
