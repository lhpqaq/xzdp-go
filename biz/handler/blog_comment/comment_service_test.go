package blog_comment

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetHotComment(t *testing.T) {
	h := server.Default()
	h.GET("/comment/hot", GetHotComment)
	w := ut.PerformRequest(h.Engine, "GET", "/comment/hot", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetComment(t *testing.T) {
	h := server.Default()
	h.GET("/comment/:blogID", GetComment)
	w := ut.PerformRequest(h.Engine, "GET", "/comment/:blogID", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestPostComment(t *testing.T) {
	h := server.Default()
	h.GET("/comment/post", PostComment)
	w := ut.PerformRequest(h.Engine, "POST", "/comment/post", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestLikeComment(t *testing.T) {
	h := server.Default()
	h.GET("/comment/like/:id", LikeComment)
	w := ut.PerformRequest(h.Engine, "PUT", "/comment/like/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteComment(t *testing.T) {
	h := server.Default()
	h.GET("/comment/:id", DeleteComment)
	w := ut.PerformRequest(h.Engine, "DELETE", "/comment/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
