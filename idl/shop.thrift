// idl/shop.thrift
namespace go shop

// import "github.com/hertz-contrib/sessions"

struct Empty {}

struct ShopType {
    1: i64 id
    2: string name
    3: string icon
    4: i32 sort
}

struct ShopOfTypeReq {
    1: i32 typeId (api.query="typeId");
    2: i64 current (api.query="current");
}

struct ShopOfTypeGeoReq {
    1: i32 typeId (api.query="typeId");
    2: i64 current (api.query="current");
    3: double longitude (api.query="x");
    4: double latitude (api.query="y");
    5: double distance (api.query="dist");
}

struct Shop {
    1: i64 id (go.tag='gorm:"id",redis:"id"');
    2: string name (go.tag='gorm:"name",redis:"name"');
    3: i64 typeId (go.tag='gorm:"type_id",redis:"typeId"');
    4: string images (go.tag='gorm:"images",redis:"images"');
    5: string area (go.tag='gorm:"area",redis:"area"');
    6: string address (go.tag='gorm:"address",redis:"address"');
    7: double x (go.tag='gorm:"longitude",redis:"longitude"');
    8: double y (go.tag='gorm:"latitude",redis:"latitude"');
    9: i64 avgPrice (go.tag='gorm:"avg_price",redis:"avgPrice"');
    10: i32 sold (go.tag='gorm:"sold",redis:"sold"');
    11: i32 comments (go.tag='gorm:"comments",redis:"comments"');
    12: i32 score (go.tag='gorm:"score",redis:"score"');
    13: string openHours (go.tag='gorm:"open_hours",redis:"openHours"');
    14: string createTime (go.tag='gorm:"create_time",redis:"createTime"');
    15: string updateTime (go.tag='gorm:"update_time",redis:"updateTime"');
    16: double distance (go.tag='gorm:"-",redis:"-"');
}
service ShopService {
    list<ShopType> ShopList(1: Empty request) (api.get="/shop-type/list");
    list<Shop> ShopOfType(1: ShopOfTypeReq request) (api.get="/shop/of/type");
    list<Shop> ShopOfTypeGeo(1: ShopOfTypeGeoReq request) (api.get="/shop/of/type/geo");
    Shop ShopInfo(1: Empty request) (api.get="/shop/:id");
}
