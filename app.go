package main

import (
	"context"
	"net"

	udpx "github.com/go75/udpx/pack"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
	name string
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}
var serverAddr = &net.UDPAddr{IP: net.IPv4(127,0,0,1), Port: 10101}
var peerAddr *net.UDPAddr
var conn *net.UDPConn

// c r u d

func (a *App) C(name string) string {
	
	var err error
	conn, err = net.ListenUDP("udp", nil)
	if err != nil {
		return err.Error()
	}

	data := udpx.ObjtoBytes(udpx.Obj{
		ID: 0,
		Payload: []byte(name),
	})

	_, err = conn.WriteToUDP(data, serverAddr)
	if err != nil {
		return err.Error()
	}
	
	buf := make([]byte, 8)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return err.Error()
		}
		println("recv :" + string(buf[:n]))
		println(addr.String(), serverAddr.String())
		if addr.String() != serverAddr.String() {
			continue
		}
		if string(buf[:n]) != "ok" {
			return "加入失败"
		}
		break
	}
	println("ok")

	a.name = name

	go func() {
		buf := make([]byte, 1024)
		for {
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				continue
			}
			if addr.String() == serverAddr.String() {
				var err error
				peerAddr, err = net.ResolveUDPAddr("udp", string(buf[:n]))
				if err != nil {
					runtime.EventsEmit(a.ctx, "addr", "err")
				}
				continue
			}
			runtime.EventsEmit(a.ctx, "msg", string(buf[:n]))
		}
	}()

	return ""
}

func (a *App) Send(msg string) string {
	if peerAddr != nil {
		_, err := conn.WriteToUDP([]byte(msg), peerAddr)
		if err != nil {
			return err.Error()
		}
		return ""
	}
	return "未连接远程节点"
}

func (a *App) Connect(name string) string {
	data := udpx.ObjtoBytes(udpx.Obj{
		ID: 1,
		Payload: []byte(name),
	})
	_, err := conn.WriteToUDP(data, serverAddr)
	if err != nil {
		return err.Error()
	}
	buf := make([]byte, 16)
	n, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		return err.Error()
	}
	if addr.String() != serverAddr.String() {
		return "error"
	}
	peerAddr, err = net.ResolveUDPAddr("udp", string(buf[:n]))
	if err != nil {
		return err.Error()
	}

	println(peerAddr.String())
	return  ""
}