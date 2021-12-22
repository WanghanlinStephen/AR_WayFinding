package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pro/app/cache"
	"pro/app/common/response"
	"pro/app/model"
	"pro/app/models"
	"pro/app/strategy"
	"strconv"
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
	if err := c.ShouldBind(&models.SearchInput{}); err != nil {
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
	angle:=strategy.GetAngle(source.Longitude,source.Latitude,source.IntersectionalAngle,nextNode.Longitude,nextNode.Latitude)
	fmt.Printf("Source:%s to Destination%s with next step %s with a total weight %f, with an angle of %f",source,destination,nextStep,shortestDistance,angle)
	responseData := &models.SearchOutput{
		Angle:     angle,
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
func AddNode(c *gin.Context) {

	if err := c.ShouldBind(&models.AddNodeInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	id,_:=strconv.Atoi(c.Query("id"));
	latitude,_:=strconv.ParseFloat(c.Query("latitude"), 64);
	longitude,_:=strconv.ParseFloat(c.Query("longitude"), 64);
	newNode:=models.AddNodeInput{
		Id:                     id,
		NameEnglish:            c.Query("nameEnglish"),
		NameChinese:            c.Query("nameChinese"),
		NameTraditionalChinese: c.Query("nameTraditionalChinese"),
		Latitude:               latitude,
		Longitude:              longitude,
	}
	err:=model.AddNode(c,newNode)
	if err!=nil{
		response.Error(c,"AddNode 失败")
		return
	}
	response.Success(c,"ok","")
}


func AddConnection(c *gin.Context) {
	if err := c.ShouldBind(&models.AddConnectionInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	id,_:=strconv.Atoi(c.Query("id"));
	sourceId,_:=strconv.Atoi(c.Query("sourceId"));
	destinationId,_:=strconv.Atoi(c.Query("destinationId"));
	weight,_:=strconv.Atoi(c.Query("weight"));

	newConnection:=models.AddConnectionInput{
		Id:            id,
		SourceId:      sourceId,
		DestinationId: destinationId ,
		Weight:        weight,
	}
	err:=model.AddConnection(c,newConnection)
	if err!=nil{
		response.Error(c,"AddConnection 失败")
		return
	}
	response.Success(c,"ok","")

}


func DeleteNode(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteNodeInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	id,_:=strconv.Atoi(c.Query("id"));
	deleteNode:=models.DeleteNodeInput{
		Id:						id,
	}
	err:=model.DeleteNode(c,deleteNode)
	if err!=nil{
		response.Error(c,"DeleteNode 失败")
		return
	}

	response.Success(c,"ok","")
}

func DeleteConnection(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteConnectionInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	id,_:=strconv.Atoi(c.Query("id"));
	deleteNode:=models.DeleteConnectionInput{
		Id:						id,
	}
	err:=model.DeleteConnection(c,deleteNode)
	if err!=nil{
		response.Error(c,"DeleteConnection 失败")
		return
	}

	response.Success(c,"ok","")

}
func Modify(c *gin.Context) {
	if err := c.ShouldBind(&models.ModifyInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database

	response.Success(c,"ok","")
}