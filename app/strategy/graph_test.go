package strategy

import (
	"fmt"
	"pro/app/model"
	"pro/config"
	"testing"
)

func Test_Initialize(t *testing.T) {
	//main()
	config.Run()
	if err := model.Run(); err != nil {
		fmt.Println("数据库链接失败:", err)
		return
	}
	Initialization()
}


