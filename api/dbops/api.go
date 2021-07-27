package dbops

import (
	"database/sql"
	"fmt"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

// AddUserCredential 添加用户到数据表
func AddUserCredential(loginName, pwd string) error {
	// 不能这样写，每次进行操作都会开启一个数据库连接，需要额外实现
	//db := OpenConn()

	stmt, err := dbConn.Prepare("INSERT INTO users (username, passwd) VALUES (?, ?)")
	if err != nil {
		fmt.Println("AddUserCredential Prepare error")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(loginName, pwd)
	if err != nil {
		return err
	}

	return nil
}

// GetUserCredential 查询用户
func GetUserCredential(loginName string) (string, error) {
	stmt, err := dbConn.Prepare("select passwd from users where username=?")
	if err != nil {
		fmt.Println("GetUserCredential Prepare error")
		return "", err
	}
	defer stmt.Close()

	var passwd string

	// error处理：分两种。sql.ErrNoRows指queryRow未查询到结果，将空行作为结果返回，然后错误有scan带出来
	err = stmt.QueryRow(loginName).Scan(&passwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return passwd, nil
}

// DeleteUser 删除用户
func DeleteUser(loginName string) error {
	stmt, err := dbConn.Prepare("delete from users where username=?")
	if err != nil {
		fmt.Println("DeleteUser Prepare error")
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(loginName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

// AddNewVideo 添加新视频
// creatime -> db ->
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//  创建uuid，Universally Unique Identifier
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmt, err := dbConn.Prepare(`insert into video_info (id, author_id, name, display_ctime) values (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}

	return res, err
}

// GetVideoInfo 根据id获取video信息
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmt, err := dbConn.Prepare("select author_id, name, display_ctime from video_info where id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var videoInfo defs.VideoInfo
	err = stmt.QueryRow(vid).Scan(&videoInfo.AuthorId, &videoInfo.Name, &videoInfo.DisplayCtime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	// 返回空时
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return &videoInfo, nil
}

// DeleteVideo 删除video
func DeleteVideo(vid string) error {
	stmt, err := dbConn.Prepare("delete from video_info where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(vid)
	if err != nil {
		return err
	}
	return nil
}

// AddComment 添加评论
func AddComment(vid string, aid int, content string) error {
	uuid, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmt, err := dbConn.Prepare("insert into comments (id, video_id, author_id, content) values (?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, vid, aid, content)
	if err != nil {
		return err
	}
	return nil
}

// ListComments 获取特定视频下评论列表
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	//
	stmt, err := dbConn.Prepare("select comments.id, users.username, comments.content " +
		"from comments INNER JOIN users " +
		"ON comments.author_id = users.id " +
		"WHERE comments.video_id=? AND comments.time > FROM_UNIXTIME(?)" +
		"AND comments.time <= FROM_UNIXTIME(?)")
	if err != nil {
		return nil, err
	}
	defer dbConn.Close()

	var cList []*defs.Comment

	rows, err := stmt.Query(vid, from, to)
	if err != nil {
		return cList, err
	}

	var cmt *defs.Comment
	for rows.Next() {
		if err := rows.Scan(&cmt.Id, &cmt.Author, &cmt.Content); err != nil {
			return cList, err
		}
		cmt.VideoId = vid

		cList = append(cList, cmt)
	}

	return cList, nil
}
