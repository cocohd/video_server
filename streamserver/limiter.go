package main

import "log"

type ConnLimiter struct {
	concurrentConn int // z最大链接数
	bucket         chan int
}

// NewConnLimiter 实现ConLimiter类型
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		// 使用go的channel，实现流控制
		bucket: make(chan int, cc),
	}
}

// GetConn 获得一个链接资源
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		log.Printf("Reached the rate limitation!")
		return false
	}

	cl.bucket <- 1
	return true
}

// ReleaseConn 释放一个链接资源
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("Connection released:%d!", c)
}
