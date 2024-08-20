package message

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	redis2 "github.com/go-redis/redis/v8"
	"strconv"
	"xzdp/biz/dal/redis"
	"xzdp/biz/pkg/cache"
	"xzdp/biz/pkg/constants"
	"xzdp/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
	message "xzdp/biz/model/message"
)

type SseService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSseService(Context context.Context, RequestContext *app.RequestContext) *SseService {
	return &SseService{RequestContext: RequestContext, Context: Context}
}

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

func (h *SseService) consumeMq(id string) ([]redis2.XStream, error) {
	u := utils.GetUser(h.Context).GetID()
	idStr := strconv.FormatInt(u, 10)
	key := constants.MESSAGE_STREAM_KEY + idStr
	consumer := constants.STREAM_CONSUMER + idStr
	redis.CreateConsumerGroup(h.Context, key)
	// 读redis消息队列
	xSet, err := redis.ConsumeMq(h.Context, key, consumer, 20, 1, id)
	return xSet, err
}

func (h *SseService) Ack(id string) error {
	u := utils.GetUser(h.Context).GetID()
	idStr := strconv.FormatInt(u, 10)
	key := constants.MESSAGE_STREAM_KEY + idStr
	err := redis.AckMq(h.Context, key, id)
	if err != nil {
		return err
	}
	return nil
}
func (h *SseService) handleMessage(xSet []redis2.XStream) (*message.MessageResp, error) {
	if len(xSet) == 0 {
		return nil, nil
	}
	m := xSet[0].Messages
	if len(m) == 0 {
		return nil, nil
	}
	hlog.CtxDebugf(h.Context, "xSet = %+v", xSet)
	var msg message.Message
	jsonStr := m[0].Values["message"].(string)
	if err := utils.UnSerializeStruct(jsonStr, &msg); err != nil {
		hlog.CtxErrorf(h.Context, "MapToStructByJson failed with data: %v, error: %v", m[0].Values, err)
		return nil, err
	}
	hlog.CtxDebugf(h.Context, "msg = %+v", msg)
	to := msg.To
	userDto, err := cache.GetUserDtoFromCacheOrDB(h.Context, to)
	if err != nil {
		hlog.CtxErrorf(h.Context, "GetUserDtoFromCacheOrDB failed with data: %v, error: %v", to, err)
		return nil, err
	}
	msgResp := &message.MessageResp{
		Message: &msg,
		User:    userDto,
	}
	err = h.Ack(m[0].ID)
	if err != nil {
		return nil, err
	}
	return msgResp, nil
}
