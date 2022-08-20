package goTest

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
	"testing"
	"time"
)

type (
	Def interface {
		GetUser() (string, error)
	}

	Default struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) Def {
	return &Default{
		cli: cli,
	}
}

func (m *Default) GetUser() (string, error) {
	return "123", nil
}

func TestRpc(t *testing.T) {
	var rpc zrpc.RpcClientConf
	fmt.Println(rpc)
	var etcd discov.EtcdConf
	var s []string
	s = append(s, "192.168.0.100:23791")
	etcd.Hosts = s
	etcd.Key = "user.rpc"
	rpc.Etcd = etcd

	var t1 = time.Now().UnixMilli()
	client, e := zrpc.NewClient(rpc)
	var a1 = time.Now().UnixMilli()
	fmt.Println(a1 - t1)
	if e != nil {
		fmt.Println("连接rpc失败无法执行")
	}
	fmt.Println(client)

	Def := NewUser(client)

	str, e := Def.GetUser()
	fmt.Println(str)
}
