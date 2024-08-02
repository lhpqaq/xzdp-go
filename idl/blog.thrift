// idl/blog.thrift
namespace go blog
struct Empty {}


struct Blog {
    1: i64 id (go.tag='gorm:"id"');
    2: i64 shopId (go.tag='gorm:"shop_id"')
    3: i64 userId (go.tag='gorm:"user_id"')
    4: string icon (go.tag='gorm:"-"')
    5: string name (go.tag='gorm:"-"')
    6: string title (go.tag='gorm:"title"')
    7: string images (go.tag='gorm:"images"')
    8: string content (go.tag='gorm:"content"')
    9: i64 liked (go.tag='gorm:"liked"')
    10: i64 comments (go.tag='gorm:"comments"')
    11: string createTime (go.tag='gorm:"create_time"');
    12: string updateTime (go.tag='gorm:"update_time"');
}

struct BlogReq {
    1: i64 current (api.query="current");
}
service BlogSerivice {
    list<Blog> GetHotBlog(1: BlogReq request) (api.get="/blog/hot");
    list<Blog> GetBlogOfMe(1: BlogReq request) (api.get="/blog/of/me");
}