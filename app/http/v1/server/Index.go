package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/cache"
	"pro/app/common/response"
)

func Index(c *gin.Context) {
	redis := cache.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
	}
	response.Success(c, "ok", r)
}
//Input:
func Search(c *gin.Context) {
	type  searchInput struct {

	}


}
func Test(c *gin.Context) {
	type testParams struct {
		Id   int    `json:"id"   binding:"required"`
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBind(&testParams{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}
	fmt.Println(testParams{})
	response.Success(c,"ok","")
	//dec := json.NewDecoder(strings.NewReader(jsonstring))
	//fmt.Println(dec)
}
