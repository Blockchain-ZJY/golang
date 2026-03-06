package main

import (
	"log"

	"github.com/go-mysql-org/go-mysql/canal"
)

/*
监听 binlog（CDC）	解析 MySQL binlog 行事件
*/
type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	log.Printf("Table: %s Action: %s Rows: %+v\n",
		e.Table.Name, e.Action, e.Rows)
	return nil
}

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:3306"
	cfg.User = "root"
	cfg.Password = "root"
	cfg.Flavor = "mysql"
	cfg.Dump.ExecutionPath = ""
	cfg.Dump.TableDB = ""
	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	//当 binlog 有事件时，调用我Handler的OnRow 方法。
	c.SetEventHandler(&MyEventHandler{})

	// 从最新位点开始监听
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
