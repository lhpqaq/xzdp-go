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

struct Shop {
    1: i64 id (go.tag='gorm:"id"');
    2: string name (go.tag='gorm:"name"');
    3: i64 typeId (go.tag='gorm:"type_id"');
    4: string images (go.tag='gorm:"images"');
    5: string area (go.tag='gorm:"area"');
    6: string address (go.tag='gorm:"address"');
    7: double x (go.tag='gorm:"longitude"');
    8: double y (go.tag='gorm:"latitude"');
    9: i64 avgPrice (go.tag='gorm:"avg_price"');
    10: i32 sold (go.tag='gorm:"sold"');
    11: i32 comments (go.tag='gorm:"comments"');
    12: i32 score (go.tag='gorm:"score"');
    13: string openHours (go.tag='gorm:"open_hours"');
    14: string createTime (go.tag='gorm:"create_time"');
    15: string updateTime (go.tag='gorm:"update_time"');
    16: double distance (go.tag='gorm:"-"');
}
service ShopService {
    list<ShopType> ShopList(1: Empty request) (api.get="/shop-type/list");
    list<Shop> ShopOfType(1: ShopOfTypeReq request) (api.get="/shop/of/type");
}
