package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// *. 中间件验证
type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// session验证
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", LogIn)

	return router
}

func main() {
	// 此处的r返回的即为http.Handler类型
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

//main -> middleware（校验等处理） -> defs(message, err) -> handlers -> dbops -> response
