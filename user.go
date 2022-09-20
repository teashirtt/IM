package main

import (
	"net"
)

type User struct {
	Name string
	Addr string
	Chan chan string
	conn net.Conn
}

// 创建用户API
func NewUser(conn net.Conn) *User {
	UserAddr := conn.RemoteAddr().String()
	user := &User{
		Name: UserAddr,
		Addr: UserAddr,
		Chan: make(chan string),
		conn: conn,
	}

	go user.ListenMessage()

	return user
}

// 监听User channel
func (this *User) ListenMessage() {
	for {
		//FIXME：强踢功能需要关闭管道，这里暂不处理会造成死循环，待解决
		msg := <-this.Chan
		this.conn.Write([]byte(msg + "\n"))
	}
}
