package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/cache"
	"pro/app/common/response"
	"pro/app/strategy"
)

//Visitor Section
func Index(c *gin.Context) {
	redis := cache.RedisInter.Get()
	r, err := redis.Do("Set", "test", 111)
	if err != nil {
		fmt.Println(err)
	}
	response.Success(c, "ok", r)
}
//fixme:是id
func Search(c *gin.Context) {
	type  searchInput struct {
		Source string
		Destination string
	}
	if err := c.ShouldBind(&searchInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}
	//Strategy Dijkastra: ID only
	sourceId:=c.Query("source")
	destinationId:=c.Query("destination")
	source:= strategy.CyberPortMap.NodeMap[sourceId]
	destination:= strategy.CyberPortMap.NodeMap[destinationId]

	shortestDistance,nextStep:=strategy.CyberPortMap.Dijkstra(source.NodeID(),destination.NodeID())
	nextNode:= strategy.CyberPortMap.NodeMap[nextStep]
	direction,angle:=strategy.GetAngle(source.Longitude,source.Latitude,nextNode.Longitude,nextNode.Latitude)
	fmt.Printf("Source:%s to Destination%s with next step %s with a total weight %f,with a direction of %s, with an angle of %f",source,destination,nextStep,shortestDistance,direction,angle)
	responseData := gin.H{
		"direction":     direction,
		"angle": 		 angle,
	}
	response.Success(c,"ok",responseData)
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

//Admin Section
func Add(c *gin.Context) {
	type addInput struct {
		Id int
		NameEnglish string
		NameChinese string
		NameTraditionalChinese string
		Latitude float64
		Longitude float64
	}
	if err := c.ShouldBind(&addInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}

	response.Success(c,"ok","")
}
