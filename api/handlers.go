package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
)

func SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "注册用户")
}

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	// 将json转成ubody
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 添加用户到数据库
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	// 正确创建用户后，建立session
	id := session.GenerateNewSessionId(ubody.Username)
	// 把session信息序列化
	su := &defs.SignedUp{
		Success:   true,
		SuccessId: id,
	}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalError)
		return
	} else {
		sendNormalResponse(w, string(resp), 200)
	}
}

func LogIn(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}
