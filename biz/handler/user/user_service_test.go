package user

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
)

func TestUserMethod(t *testing.T) {
	h := server.Default()
	h.GET("/user/me", UserMe)
	w := ut.PerformRequest(h.Engine, "GET", "/user/me", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestSendCode(t *testing.T) {
	h := server.Default()
	h.GET("/user/code", SendCode)
	w := ut.PerformRequest(h.Engine, "POST", "/user/code", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUserLogin(t *testing.T) {
	h := server.Default()
	h.GET("/user/login", UserLogin)
	w := ut.PerformRequest(h.Engine, "POST", "/user/login", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}

func TestUserInfo(t *testing.T) {
	h := server.Default()
	h.GET("/user/:id", UserInfo)
	w := ut.PerformRequest(h.Engine, "GET", "/user/:id", &ut.Body{Body: bytes.NewBufferString(""), Len: 1},
		ut.Header{})
	resp := w.Result()
	assert.DeepEqual(t, 201, resp.StatusCode())
	assert.DeepEqual(t, "", string(resp.Body()))
	// todo edit your unit test.
}
