package blog

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestGetHotBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/hot", GetHotBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/hot", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetUserBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/user/:id", GetUserBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/user/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestBlogOfMe(t *testing.T) {
	h := server.Default()
	h.GET("/blog/of/me", BlogOfMe)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/of/me", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestPostBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog", PostBlog)
	w := ut.PerformRequest(h.Engine, "POST", "/blog", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/:id", GetBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestDeleteBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/:id", DeleteBlog)
	w := ut.PerformRequest(h.Engine, "DELETE", "/blog/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestLikeBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/like/:id", LikeBlog)
	w := ut.PerformRequest(h.Engine, "PUT", "/blog/like/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetLikes(t *testing.T) {
	h := server.Default()
	h.GET("/blog/likes/:id", GetLikes)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/likes/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestGetFollowBlog(t *testing.T) {
	h := server.Default()
	h.GET("/blog/of/follow", GetFollowBlog)
	w := ut.PerformRequest(h.Engine, "GET", "/blog/of/follow", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
