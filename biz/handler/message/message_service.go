package message

import (
	"context"
	"time"

	service "xzdp/biz/service/message"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/sse"
)

// Sse .
// @router /sse [POST]
func Sse(ctx context.Context, c *app.RequestContext) {
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
}

// PublishWithRetry 尝试发布事件，如果失败则重试
func PublishWithRetry(s *sse.Stream, event *sse.Event) error {
	maxRetries := 3           // 最大重试次数
	retryDelay := time.Second // 重试间隔时间

	for i := 0; i <= maxRetries; i++ {
		err := s.Publish(event)
		if err == nil {
			return nil
		}

		if i < maxRetries {
			hlog.Errorf("Publish failed, retrying... (attempt %d/%d)", i+1, maxRetries)
			time.Sleep(retryDelay)
		} else {
			hlog.Errorf("Publish failed after %d attempts, giving up.", maxRetries)
			return err
		}
	}

	return nil
}
