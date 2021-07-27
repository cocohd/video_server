package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	vid string
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Add", testGetUser)
	t.Run("Add", testDeleteUser)
	t.Run("Add", testRegetUser)

}

func testAddUser(t *testing.T) {
	err := AddUserCredential("handy", "!@#")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	_, err := GetUserCredential("handy")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("handy")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("handy")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Adduser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DeleteVideo", testDeleteVideo)
}

func testAddVideoInfo(t *testing.T) {
	v, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideo: %v", err)
	}
	vid = v.Id
	fmt.Println("test videoID: ", v.Id)
	fmt.Println("test vid: ", vid)

}

func testGetVideoInfo(t *testing.T) {
	v, err := GetVideoInfo(vid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
	fmt.Println("video info: ", v)
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(vid)
	if err != nil {
		t.Errorf("Error of DeleteVideo: %v", err)
	}
}

func TestCommentsWorkFlow(t *testing.T) {
	clearTables()
	t.Run("Adduser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	err := AddComment("video123", 1, "转载投自制？")
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	// 测试时from采用两周前的一个值，秒数
	from := 1514764800
	// UnixNano返回纳秒； FormatInt返回i的base进制的字符串表示
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	res, err := ListComments("video123", from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, v := range res {
		fmt.Printf("comment: %d, %v \n", i, v)
	}
}
