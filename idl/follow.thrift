// idl/blog.thrift
namespace go follow
include "./user.thrift"


struct Follow {
    1: i64 id(go.tag='gorm:"id"')
    2: i64 userId(go.tag='gorm:"user_id"')
    3: i64 followUserId(go.tag='gorm:"follow_user_id"')
    4: string createTime (go.tag='gorm:"create_time default:CURRENT_TIMESTAMP"')
}

struct FollowReq {
    1: bool isFollow(api.query="isFollow")
    2: i64 targetUser(api.query="targetUser")
}

struct FollowResp {
    1: Follow RespBody;
}

struct isFollowedResp{
    1: bool isFollowed;
}
struct commonFollowReq{
    1: i64 userId(api.query="userId");
}
struct commonFollowResp{
    1: list<user.UserDTO> commonFollows;
}
service FollowService {
    FollowResp Follow(1: FollowReq request) (api.put="/follow");
    isFollowedResp isFollowed(1: string request) (api.get="/follow/isFollowed/:id");
    commonFollowResp commonFollow(1: string request) (api.get="/follow/common/:id");
}
