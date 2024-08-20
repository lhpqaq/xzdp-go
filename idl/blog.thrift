// idl/blog.thrift
namespace go blog
include "./user.thrift"
struct Empty {}


struct Blog {
    1: i64 id (go.tag='gorm:"id"');
    2: i64 shopId (go.tag='gorm:"shop_id"')
    3: i64 userId (go.tag='gorm:"user_id"')
    4: string title (go.tag='gorm:"title"')
    5: string images (go.tag='gorm:"images"')
    6: string content (go.tag='gorm:"content"')
    7: i64 liked (go.tag='gorm:"liked"')
    8: i64 comments (go.tag='gorm:"comments"')
    9: string createTime (go.tag='gorm:"create_time"');
    10: string updateTime (go.tag='gorm:"update_time"');
    11: string icon (go.tag='gorm:"-"');
    12: string nickName (go.tag='gorm:"-"');
    13: bool isLiked (go.tag='gorm:"-"');
}

struct BlogReq {
    1: i64 current (api.query="current");
}

struct LikeResp {
    1: bool isLiked;
}
struct FollowBlogReq {
    1: string lastId (api.query="lastId");
    2: i64 offset (api.query="offset");
}
struct FollowBlogRresp {
    1: list<Blog> list;
    2: string minTime (api.query="minTime");
    3: i64 offset (api.query="offset");
}
service BlogService {
    list<Blog> GetHotBlog(1: BlogReq request) (api.get="/blog/hot");
    list<Blog> GetUserBlog(1: BlogReq request) (api.get="/blog/user/:id");
    list<Blog> BlogOfMe(1: BlogReq request) (api.get="/blog/of/me");
    // 发布博客
    Blog PostBlog(1: Blog request) (api.post="/blog");
    // 查看博客
    Blog GetBlog(1: string request) (api.get="/blog/:id");
    // 删除博客
    Empty DeleteBlog(1: string request) (api.delete="/blog/:id");
    // 点赞博客
    LikeResp LikeBlog(1: string request) (api.put="/blog/like/:id");
    // 点赞用户排行
    list<user.UserDTO> GetLikes(1: string request) (api.get="/blog/likes/:id");
    // 关注的博客
    FollowBlogRresp GetFollowBlog(1: FollowBlogReq request) (api.get="/blog/of/follow");
}
