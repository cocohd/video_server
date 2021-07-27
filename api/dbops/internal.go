package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

// 不和handler直接交互，如session获取

// InsertSession 插入session到数据表中
func InsertSession(sid string, ttl int64, uname string) error {
	// ttl变成string格式
	ttlstr := strconv.FormatInt(ttl, 10)

	stmt, err := dbConn.Prepare("insert into sessions (session_id, TTL, username) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	return nil
}

// RetrieveSession 验证session是否过期
func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	sess := &defs.SimpleSession{}

	stmt, err := dbConn.Prepare("select TTL, username from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// ttl在数据库中是string类型的，需要经过转换成int64
	var ttl string
	err = stmt.QueryRow(sid).Scan(&ttl, &sess.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	//返回字符串表示的整数值，接受正负号。
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	//返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		sess.TTL = res
	}

	return sess, nil
}

// RetrieveAllSessions 返回全部的session，放入map中
func RetrieveAllSessions() (*sync.Map, error) {
	p := sync.Map{}

	stmt, err := dbConn.Prepare("select * from sessions")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		// 单独处理TTL
		sess := &defs.SimpleSession{}
		var ttl, id string
		err := rows.Scan(&id, &sess.Username, &ttl)
		if err != nil {
			return nil, err
		}

		if ttlInt, err := strconv.ParseInt(ttl, 10, 64); err == nil {
			sess.TTL = ttlInt
			p.Store(id, sess)
			log.Printf("session id: %s, ttl: %d", id, sess.TTL)
		}
	}
	return &p, nil
}

// DeleteSession 删除session
func DeleteSession(sid string) error {
	stmt, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sid)
	if err != nil {
		return nil
	}
	return nil
}
