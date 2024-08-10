package follow

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestFollow(t *testing.T) {
	h := server.Default()
	h.GET("/follow", Follow)
	w := ut.PerformRequest(h.Engine, "GET", "/follow", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestIsFollowed(t *testing.T) {
	h := server.Default()
	h.GET("/follow/isFollowed", IsFollowed)
	w := ut.PerformRequest(h.Engine, "GET", "/follow/isFollowed", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
