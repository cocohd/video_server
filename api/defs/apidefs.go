package defs

// UserCredential 用户模型
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// SignedUp response
type SignedUp struct {
	Success   bool   `json:"success"`
	SuccessId string `json:"session_id"`
}

// VideoInfo 视频数据结构
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
}

// Comment 评论数据结构
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
}

// SimpleSession session
type SimpleSession struct {
	Username string
	TTL      int64
}

func CreateUser(loginName, pwd string) {

}

func GetUser() {

}
