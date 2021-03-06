package main

import (
	"fmt"
	"github.com/reechou/holmes"
	"github.com/leesper/tao"
	"github.com/leesper/tao/examples/chat"
	"runtime"
)

type ChatServer struct {
	tao.Server
}

func NewChatServer(addr string) *ChatServer {
	return &ChatServer{
		tao.NewTCPServer(addr),
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer holmes.Start().Stop()
	tao.MonitorOn(12345)

	tao.Register(chat.CHAT_MESSAGE, chat.DeserializeChatMessage, chat.ProcessChatMessage)

	chatServer := NewChatServer(fmt.Sprintf("%s:%d", "0.0.0.0", 18341))

	chatServer.SetOnConnectCallback(func(conn tao.Connection) bool {
		holmes.Info("%s", "On connect")
		return true
	})

	chatServer.SetOnErrorCallback(func() {
		holmes.Info("%s", "On error")
	})

	chatServer.SetOnCloseCallback(func(conn tao.Connection) {
		holmes.Info("%s", "Closing chat client")
	})

	chatServer.Start()
}
