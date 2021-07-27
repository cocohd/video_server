package session

import (
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

// 并发读写的map
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB() {
	sess, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	// 遍历并存入sessionMap。试试看能不能使用普通遍历
	sess.Range(func(key, value interface{}) bool {
		// 先将接口类型转换为SimpleSession
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

// GenerateNewSessionId 生成sessionId
func GenerateNewSessionId(un string) string {
	uuid, _ := utils.NewUUID()

	// ，毫秒
	ct := time.Now().UnixNano() / 1000000
	// session 30分钟后过期
	ttl := ct + 30*60*1000

	// 存入读写map中，便于后续比较
	sess := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(uuid, sess)

	dbops.InsertSession(uuid, ttl, un)

	return uuid
}

// DeleteExpiredSession 删除map和数据库中过期的session
func DeleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

// IsSessionExpired session是否过期
// @return 返回uname，是否过期
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		curTime := time.Now().UnixNano() / 1000000
		if ss.(*defs.SimpleSession).TTL < curTime {
			//如果过期，则删除sessionMap中的session
			DeleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
