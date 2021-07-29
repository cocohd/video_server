package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// streamHandler:简单的视频点播
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	//打开视频文件
	video, err := os.Open(vl)
	if err != nil {
		fmt.Println(err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	defer video.Close()

	// 在header设置文件格式，这样浏览器会按照mp4格式解析
	w.Header().Set("Content-Type", "video/mp4")
	//将文件以二进制流形式传递给client端
	http.ServeContent(w, r, "", time.Now(), video)
}

// uploadHandler：流数据上传
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 限制请求body大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	// 整个请求正文被解析，其文件部分的总共 maxMemory 字节存储在内存中，其余部分存储在磁盘上的临时文件中
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Internal Error!")
		return
	}

	// 解析表单<form name=“file”>
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error!")
		return
	}

	// 读取数据
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error!")
		return
	}

	// 写入文件
	index := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+index, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return
	}

	// 文件保存成功后，给前端返回一个正确码
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded Successfully")
}

// testUploadHandler
func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w, "testUpload!!!")
	// 渲染模板并返回给客户端
	//t, _ := template.ParseFiles(VIDEO_DIR + "upload .00000000000000000000000000000inthtml")
	t, _ := template.ParseFiles("./upload.html")
	t.Execute(w, nil)

	fmt.Println(r.Body, p.MatchedRoutePath())
}
