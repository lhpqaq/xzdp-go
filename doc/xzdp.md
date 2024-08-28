# 一、项目详情
## 1.1项目介绍
接口文档：[接口文档](https://doc.apipost.net/docs/30594359f464000)
本项目是黑马程序员的Redis实战项目，使用Go语言重构的版本。目前项目还在开发中，本文会持续更新。
![模块划分](https://i-blog.csdnimg.cn/direct/cbb8b728c35443b2b40c4b4268ca6995.png)
## 1.2使用技术栈

- 框架：字节跳动开源框架-[Hertz](https://www.cloudwego.io/zh/docs/hertz/)
- IDL:[thrift](https://www.cloudwego.io/zh/docs/hertz/tutorials/toolkit/)
- Redis：set、zset、stream
- SSE（Server-Send Events）：服务端推送

# [二、达人探店](https://blog.csdn.net/homonym/article/details/141157945)
达人探店，类似于博客、笔记等功能。涉及到基于基于SortSet的点赞列表和点赞排行。
![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/6c18d0163cfd424e9a3803076709e294.png)
## 2.1发布博客
![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/a6df356962a64e01a544a2c6fb6cc438.png)

图片不直接存在数据库中，而是存储图片的地址，故发布博客的实现分为两个步骤 流程如上图：

1. 上传图片，获得图片地址
2. 发布博客
   前端提交后向数据库插入一条记录

```go
	u := utils.GetUser(h.Context).GetID()
	req.UserId = u
	if !errors.Is(mysql.DB.Create(&req).Error, nil) {
		return nil, errors.New("创建失败")
	}
	req.Icon = utils.GetUser(h.Context).GetIcon()
	req.NickName = utils.GetUser(h.Context).GetNickName()
	req.IsLiked = false
```
## 2.2 查看博客
查看博客时，除了显示博客内容以外还要显示用户头像、是否关注等信息
![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/df9db74d21e14d6ab4711ad8f1b4db11.png)
用户信息使用userDto结构体定义，需要通过user对象转换而来，避免频繁转换，直接将字段定义才Blog结构体中，并使用gorm:-表示不属于blog表的字段，需要忽略
```go
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
```
查询逻辑分为3步，1、查询博客数据，2、查询用户信息，3、查询点赞状态

```go
	if !errors.Is(mysql.DB.First(&resp, "id = ?", req).Error, nil) {
		return nil, errors.New("未找到该博客")
	}
	userId := resp.UserId
	user, err := mysql.GetById(h.Context, userId)
	if err != nil {
		return nil, err
	}
	resp.Icon = user.Icon
	resp.NickName = user.NickName
	resp.IsLiked = false
	// 获取点赞状态
	u := utils.GetUser(h.Context).GetID()
	key := constants.BLOG_LIKED_KEY + *req
	isLike, err := redis.IsLiked(h.Context, key, strconv.FormatInt(u, 10))
	if err != nil {
		return nil, err
	}
	resp.IsLiked = isLike
	return resp, nil
```
## 2.3 删除博客
删除博客前要注意同时删除点赞信息和评论信息，最后再删除数据库记录

```go
	key := constants.BLOG_LIKED_KEY + *req
	// 从redis删除点赞数据
	ok, err := redis.HasLikes(h.Context, key)
	if err != nil {
		return nil, err
	}
	if ok {
		if !errors.Is(redis.DeleteLikes(h.Context, key), nil) {
			return nil, err
		}
	}
	// 删除评论信息
	bid, err := strconv.ParseInt(*req, 10, 64)
	if err != nil {
		return nil, err
	}
	err = mysql.DeleteBlogComment(h.Context, bid)
	if !errors.Is(err, nil) {
		return nil, err
	}
	// 从数据库删除博客
	err = mysql.DB.Where("id = ?", req).Delete(&blog.Blog{}).Error
	if !errors.Is(err, nil) {
		return nil, err
	}
```

## 2.4点赞博客
点赞功能实现时要判断该博客是否已经点赞过，为了方便实现时将点赞状态写到blog的islike字段中。点赞前判断是否已经点赞，如果已经点赞则取消点赞

```go
	var interBlog blog.Blog
	err = mysql.DB.Where("id=?", req).First(&interBlog).Error
	if !errors.Is(err, nil) {
		return nil, errors.New("博客不存在")
	}
	// 判断是否已经点赞
	u := utils.GetUser(h.Context).GetID()
	idStr := strconv.FormatInt(u, 10)
	key := constants.BLOG_LIKED_KEY + *req
	isLike, err := redis.IsLiked(h.Context, key, idStr)
	if err != nil {
		hlog.Debugf("like redis error: %+v", err)
		return nil, err
	}
	fmt.Printf("isLike = %+v", isLike)
	// 如果已经点赞则取消点赞
	if isLike {
		if !errors.Is(redis.RedisClient.ZRem(h.Context, key, idStr).Err(), nil) {
			return nil, errors.New("取消点赞失败")
		}
		// 同步减少点赞数
		mysql.DB.Model(&blog.Blog{}).Where("id = ?", req).UpdateColumn("liked", gorm.Expr("liked - ?", 1))
		return &blog.LikeResp{IsLiked: false}, nil
	}
	// 否则点赞
	if !errors.Is(redis.RedisClient.ZAdd(h.Context, key, &redis2.Z{
		Score:  float64(time.Now().Unix()),
		Member: idStr,
	}).Err(), nil) {
		return nil, errors.New("点赞失败")
	}
	// 同步增加点赞数
	mysql.DB.Model(&blog.Blog{}).Where("id = ?", req).UpdateColumn("liked", gorm.Expr("liked + ?", 1))
```
## 2.5 点赞排行
通过点赞时候将用户id存入到sortedset中这里。在获取点赞排行时可以将点赞人按照score排行，取前五条最早的进行返回

```go
func (h *GetLikesService) Run(req *string) (resp *[]*user.UserDTO, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	key := constants.BLOG_LIKED_KEY + *req
	ids, err := redis.RedisClient.ZRange(h.Context, key, 0, 4).Result()
	if err != nil {
		return nil, err
	}
	var users []*user.User
	if !errors.Is(mysql.DB.Where("id in ?", ids).Find(&users).Error, nil) {
		return nil, errors.New("获取失败")
	}
	var userDtos []*user.UserDTO
	for _, u := range users {
		d := &user.UserDTO{
			ID:       u.ID,
			NickName: u.NickName,
			Icon:     u.Icon,
		}
		userDtos = append(userDtos, d)
	}
	if len(userDtos) == 0 {
		userDtos = make([]*user.UserDTO, 0)
	}
	return &userDtos, nil
}
```
## 2.6 获取指定用户博客
用访问主页时获取博客

```go
func (h *GetUserBlogService) Run(req *blog.BlogReq, uerID int64) (resp *[]*blog.Blog, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	u, err := mysql.GetById(h.Context, uerID)
	if err != nil {
		return nil, err
	}
	d := &user.UserDTO{
		ID:       u.ID,
		NickName: u.NickName,
		Icon:     u.Icon,
	}
	blogList, err := mysql.QueryBlogByUserID(h.Context, int(req.Current), d)
	if err != nil {
		return nil, err
	}
	return &blogList, nil
}
```
## 2.7 blog模块的工具代码
mysql
```go
//mysql.go

func QueryBlogByUserID(ctx context.Context, current int, user *user.UserDTO) (resp []*blog.Blog, err error) {
	var blogs []*blog.Blog
	pageSize := constants.MAX_PAGE_SIZE
	err = DB.Where("user_id = ?", user.ID).Order("liked desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&blogs).Error
	if err != nil {
		hlog.CtxErrorf(ctx, "query error: %v", err)
		return nil, err
	}
	for i := range blogs {
		blogs[i].NickName = user.NickName
		blogs[i].Icon = user.Icon
	}

	return blogs, nil
}

func QueryHotBlog(ctx context.Context, current int) (resp []*blog.Blog, err error) {
	var blogs []*blog.Blog
	pageSize := constants.MAX_PAGE_SIZE

	if err := DB.Order("liked desc").Limit(pageSize).Offset((current - 1) * pageSize).Find(&blogs).Error; err != nil {
		hlog.CtxErrorf(ctx, "err = %s", err.Error())
		return nil, err
	}

	for i := range blogs {
		u, err := GetById(ctx, blogs[i].UserId)
		if err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		if err := DB.First(&u, blogs[i].UserId).Error; err != nil {
			hlog.CtxErrorf(ctx, "err = %s", err.Error())

			return nil, err
		}
		blogs[i].NickName = u.NickName
		blogs[i].Icon = u.Icon
	}

	return blogs, nil
}

```
redis

```go
func IsLiked(ctx context.Context, key string, member string) (ok bool, err error) {
	_, err = RedisClient.ZScore(ctx, key, member).Result()
	if err != nil {
		if errors.Is(err, redis2.Nil) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func HasLikes(ctx context.Context, key string) (bool, error) {
	exists, err := RedisClient.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

func DeleteLikes(ctx context.Context, key string) error {
	err := RedisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetBlogsByKey(ctx context.Context, key string, max string, offset int64) ([]redis2.Z, error) {
	size := constants.MAX_PAGE_SIZE
	result, err := RedisClient.ZRevRangeByScoreWithScores(
		ctx,
		key,
		&redis2.ZRangeBy{Max: max, Min: "0", Offset: offset, Count: int64(size)},
	).Result()
	if err != nil {
		return nil, err
	}
	return result, nil
}

```

# [三、好友关注](https://blog.csdn.net/homonym/article/details/141161128)
好友关注涉及到取Set的增加、删除、取交集，Feed流推送
## 3.1 关注和取关
查看博客或则个人主页时候都会有关注/取消关注按钮，取决于用户是否关注了该博主。因此查看博客时除了请求加载博客内容，还需要发送请求获取关注状态。
| 已关注 | 未关注 |
|--|--|
| ![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/5d93ed16b3f34093a4ba3586f8386774.png) |  ![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/dfffbcdd1f0b4cc2862afc4b8da2da6c.png)|
获取关注状态，这里可以直接从redis的set中获取

```go
	user := utils.GetUser(h.Context).GetID()
	//查找是否关注
	if !errors.Is(redis.RedisClient.SIsMember(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(user, 10), targetUserID).Err(), nil) {
		return &follow.IsFollowedResp{IsFollowed: false}, nil
	}
	return &follow.IsFollowedResp{IsFollowed: true}, nil
```

未关注时点击则关注，否则取消关注

```go
myID := utils.GetUser(h.Context).GetID()
	isFollow := req.GetIsFollow()
	targetUserId := req.GetTargetUser()
	f := follow.Follow{
		UserId:       myID,
		FollowUserId: targetUserId,
	}
	// 如果是true,则添加关注，将用户id和被关注用户的id存入数据库
	if isFollow {
		// 判断是否已经关注
		if !errors.Is(redis.RedisClient.SIsMember(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
			return &follow.FollowResp{RespBody: &f}, nil
		}
		// 将关注的用户存入redis的set中
		if !errors.Is(redis.RedisClient.SAdd(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
			hlog.CtxErrorf(h.Context, "err = %s", err.Error())
			return nil, err
		}
		if !errors.Is(mysql.DB.Create(&f).Error, nil) {
			return nil, errors.New("关注失败")
		}
		return &follow.FollowResp{RespBody: &f}, nil
	}
	// 如果是false,则取消关注
	if !errors.Is(mysql.DB.Where("user_id = ? and follow_user_id = ?", myID, targetUserId).Delete(&f).Error, nil) {
		return nil, errors.New("取消关注失败")
	}
	// 将取消关注的用户从redis的set中删除
	if !errors.Is(redis.RedisClient.SRem(h.Context, constants.FOLLOW_USER_KEY+strconv.FormatInt(myID, 10), targetUserId).Err(), nil) {
		hlog.CtxErrorf(h.Context, "err = %s", err.Error())
		return nil, err
	}
	return &follow.FollowResp{RespBody: &f}, nil
```
## 3.2 共同关注
这里是再用户主页查看你们共同关注的用户，用到set求交集来完成


![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/7974a09302eb45acbe0c46f2ee9bf042.png)

- 步骤一：再关注时候除了将关注的用户存到数据库，还需要存入redis的set中（实现方法查看上一届节）
- 获取时先从redis用双方的关注列表找出共同的用户（set取交集）
- 从数据库查询交集用户的信息并转换未uderDTO



```go
	user := utils.GetUser(h.Context).GetID()
	key1 := constants.FOLLOW_USER_KEY + strconv.FormatInt(user, 10)
	key2 := constants.FOLLOW_USER_KEY + targetUserID
	arr, err := redis.RedisClient.SInter(h.Context, key1, key2).Result()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	if !errors.Is(mysql.DB.Where("id in ?", arr).Find(&users).Error, nil) {
		return nil, errors.New("查询失败")
	}
	var userDto []*model.UserDTO
	// 遍历arr，转换为userDTO
	for _, u := range users {
		d := utils.UserToUserDTO(u)
		userDto = append(userDto, d)
	}
	return &follow.CommonFollowResp{
		CommonFollows: userDto,
	}, nil
```

## 3.3 Feed流推送

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/a614c8d23db44ce28c8ad1c68d4fe42f.png)
### Feed流
* Feed 流：是提供给用户的内容流，为用户持续的提供 “沉浸式” 的体验，通过无限下拉刷新获取新的信息。比如微博的关注页，抖音的关注页视频都叫Feed流
* Feed：Feed流中的一条信息，比如朋友发布的一条朋友圈
  本项目中Feed流用再个人主页中，查看关注的用户发布的博客


关于Feed流的这里不做赘述，可以查看[redis实现Feed流推送](https://blog.csdn.net/homonym/article/details/141142877)


在本文中采取推模式作为案例。redis中实现feed流需要使用zset，当博主发布一条动态时往粉丝的收件箱（redis的zset）写一条数据。数据格式为：

```json
{
score：一般为时间戳，
member：消息内容
}
```
具体实现代码
将消息写入粉丝收信箱

```go
	fans, err := mysql.GetFansByID(h.Context, u)
	if err != nil {
		return nil, err
	}
	for _, fan := range fans {
		key := constants.FEED_KEY + strconv.FormatInt(fan.ID, 10)
		err = redis.RedisClient.ZAdd(h.Context, key, &redis2.Z{
			Score:  float64(time.Now().Unix()),
			Member: req.ID,
		}).Err()
	}
```
粉丝读取收信箱，需要注意的是，由于Feed流中的数据是随时间变化不断更新的，传统的分页方式为根据每页几条pageSize和当前第几页页Page来计算查询范围，这对于Feed流中的动态列表而言会有重复读的问题，应当采用滚动分页模式。

```go
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
```

## 3.4 好友关注工具类
mysql

```go

func GetFansByID(ctx context.Context, id int64) (resp []*user.UserDTO, err error) {
	var fs []follow.Follow
	err = DB.Where("follow_user_id = ?", id).Find(&fs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return make([]*user.UserDTO, 0), nil
	}
	if err != nil {
		return nil, errors.New("获取失败")
	}
	for _, f := range fs {
		interUser, e := GetById(ctx, f.UserId)
		if e != nil {
			return nil, e
		}
		u := &user.UserDTO{
			ID:       interUser.ID,
			NickName: interUser.NickName,
			Icon:     interUser.Icon,
		}
		resp = append(resp, u)
	}
	return resp, nil
}

```


# [四、消息通知(新增功能)](https://blog.csdn.net/homonym/article/details/141162740)

在原版的黑马点评中没有消息通知的功能，在这里为了学习，采用了SSE功能实现一个点赞博客通知事件。
## 4.1服务端实时推送技术之SSE（Server-Send Events）
服务端推送是一种允许应用服务器主动将信息发送到客户端的能力，为客户端提供了实时的信息更新和通知。
服务端推送的背景与需求主要基于以下几个诉求：

1. 实时通知：在如点赞，评论、回复等情况下需要试试通知用户。
2. 节省资源：如果没有服务端推送，客户端需要通过轮询的方式来获取新信息，会造成客户端、服务端的资源损耗。
3. 增强用户体验/营销活动：针对特定用户或用户群发送有针对性的内容，如优惠活动、个性化推荐等。

常见的实时消息处理方案：

- 轮询：在没有服务端推送时，要想试获得实时数据智能依赖客户端发起轮询。对实时性要求越高轮询越频繁，服务端和客户端的开销和压力就越大。并且存在数据长时间没有更新的情况会浪费很多轮询。
- Websocket：基于TCP的全双工协议，能够实现客户端和服务端双向通信，非常适合实时性极强的通信场景。
- ==SSE==：基于HTTP协议的推送技术，是 HTML5 的一部分，通过设置content-type为text/stream来告诉客户端，内容不是一次性返回的，而是返回流。允许服务端主动向客户端发送消息，但是不允许客户端通过sse向服务端实时发送数据，即只允许单向数据交互。与websocket相比，更简单、更轻量。
- 第三方推送平台：各家操作系统厂商一般都会提供推送渠道。同时，也有一些跨平台的推送服务，如个推、极光推送、友盟推送等，帮助开发者在不同平台上实现统一的推送功能。


## 4.2案例实现
redis中有三种方式可以实现消息队列，分别是list，pub/sub，stream，他们的区别如下

在本案例中结合redis的stream消息队列来做sse推送。如果对于消息队列有较高的要求，请考虑其他专业的消息队列。
![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/ec0e6df50e7343939a5dbbdcc08244fc.png)


我们先修改点赞博客时的逻辑，当有人点赞时往redis的消息队列中添加一条消息。


```go
	// 推送消息
	streamKey := constants.MESSAGE_STREAM_KEY + strconv.FormatInt(interBlog.UserId, 10)
	msg := &message.Message{
		From:    u,
		To:      interBlog.UserId,
		Content: "点赞了你的博客",
		Type:    "like",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	err = redis.ProduceMq(h.Context, streamKey, msg)
	if err != nil {
		return nil, err
	}
	return &blog.LikeResp{IsLiked: true}, nil
```
这样博主的消息队列中就存在一条待消费的消息，然后编写sse逻辑。当客户端连接到sse后，通过一个无限循环监听消息队列，如果有新消息则返回。同时通过go的上下文监听退出状态，及时释放资源。

```go
s := sse.NewStream(c)
	//c.Status(consts.StatusOK)
	subCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-subCtx.Done():
			hlog.Debugf("SSE stream closed")
			return
		default:
			req := ">"
			serv := service.NewSseService(subCtx, c)
			resp, err := serv.Run(req)
			if err != nil {
				hlog.Errorf("Error running SSE service: %v", err)
				continue
			}
			if resp == nil {
				continue
			}
			event := &sse.Event{
				Event: "message",
				Data:  []byte(*resp),
			}
			hlog.Debugf("SSE event: %v", event)
			err = s.Publish(event)
			if err := PublishWithRetry(s, event); err != nil {
				hlog.Errorf("Error publishing SSE event: %v", err)
			}
		}
	}
```
通过redis查询消息队列时，设置阻塞式查询，可以减少长时间得到空信息的问题。需要注意的是，redis的stream使用消费者组读取方式，读取前需要创建消费者组，读取后需要确认消息。
```go

func (h *SseService) Run(req string) (resp *string, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	xSet, err := h.consumeMq(">")
	// 有更新的话，发送给前端
	if err != nil {
		if errors.Is(err, redis2.Nil) {
			hlog.CtxDebugf(h.Context, "No messages found in Redis")
			return nil, nil
		}
		hlog.CtxErrorf(h.Context, "redis.ConsumeMq err = %+v", err)
		return nil, err
	}
	msgResp, err := h.handleMessage(xSet)
	if err != nil {
		hlog.CtxErrorf(h.Context, "handleMessage failed with data: %v, error: %v", xSet, err)
		return nil, err
	}
	bytes, err := utils.SerializeStruct(msgResp)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

```

## 4.3使用的工具类
redis

```go

// ProduceMq 写stream消息队列
func ProduceMq(ctx context.Context, key string, message interface{}) error {
	messageJSON, err := utils.SerializeStruct(message)
	if err != nil {
		return err
	}
	// xadd写stream
	err = RedisClient.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		ID:     "*",
		Values: []interface{}{"message", messageJSON},
	}).Err()
	// 错误处理
	if err != nil {
		return err
	}
	return nil
}

// ConsumeMq 读取stream消息队列
func ConsumeMq(ctx context.Context, key string, consumer string, block time.Duration, count int64, id string) ([]redis.XStream, error) {
	if id == "" {
		id = ">"
	}
	xSet, err := RedisClient.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    constants.STREAM_READ_GROUP,
		Consumer: consumer,
		Streams:  []string{key, id},
		Count:    count,
		Block:    block,
		NoAck:    false,
	}).Result()
	if err != nil {
		return nil, err
	}
	return xSet, nil
}
// 创建消费者组
func CreateConsumerGroup(ctx context.Context, key string) {
	RedisClient.XGroupCreateMkStream(ctx, key, constants.STREAM_READ_GROUP, "0")
}

//确认消息
func AckMq(ctx context.Context, key string, id string) error {
	return RedisClient.XAck(ctx, key, constants.STREAM_READ_GROUP, id).Err()
}

```
