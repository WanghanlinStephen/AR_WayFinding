package model

import (
	"fmt"
	"pro/config"
	"testing"
)

func Test_Connection(t *testing.T) {
	config.Run()
	if err := Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	GetConnectionsList()
}
