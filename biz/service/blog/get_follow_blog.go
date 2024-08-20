package blog

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"xzdp/biz/dal/mysql"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	blog "xzdp/biz/model/blog"
)

type GetFollowBlogService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetFollowBlogService(Context context.Context, RequestContext *app.RequestContext) *GetFollowBlogService {
	return &GetFollowBlogService{RequestContext: RequestContext, Context: Context}
}

func (h *GetFollowBlogService) Run(req *blog.FollowBlogReq) (resp *blog.FollowBlogRresp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	u := utils.GetUser(h.Context).GetID()
	key := constants.FEED_KEY + strconv.FormatInt(u, 10)
	zSet, err := redis.GetBlogsByKey(h.Context, key, req.LastId, req.Offset)
	var bids []string
	for _, z := range zSet {
		bids = append(bids, z.Member.(string))
	}
	var blogs []*blog.Blog
	err = mysql.DB.Where("id in ?", bids).Find(&blogs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("没有更多数据")
	}
	if err != nil {
		return nil, err
	}
	//fmt.Printf("blogs: %v\n", blogs)
	var res blog.FollowBlogRresp
	res.List = blogs
	res.MinTime = "0"
	if len(zSet) > 0 {
		res.MinTime = strconv.FormatInt(int64(zSet[len(zSet)-1].Score), 10)
	}
	// 取最小分数的记录数
	var offset int64 = 0
	minScore := zSet[len(zSet)-1].Score
	for _, element := range zSet {
		if element.Score == minScore {
			offset++
		}
	}
	res.Offset = offset
	return &res, nil
}
