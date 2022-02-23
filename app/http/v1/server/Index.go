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
//Fetch Next Node
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

	shortestDistance,nextStep,_:=strategy.CyberPortMap.Dijkstra(source.NodeID(),destination.NodeID())
	nextNode:= strategy.CyberPortMap.NodeMap[nextStep]
	angle:=strategy.GetAngle(source.Longitude,source.Latitude,source.IntersectionalAngle,nextNode.Longitude,nextNode.Latitude)
	fmt.Printf("Source:%s to Destination%s with next step %s with a total weight %f, with an angle of %f",source,destination,nextStep,shortestDistance,angle)
	responseData := &models.SearchOutput{
		Angle:     angle,
	}
	response.Success(c,"ok",responseData)
}

func FetchPath(c *gin.Context) {
	if err := c.ShouldBind(&models.FetchPathInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}

	sourceId:=c.Query("source")
	destinationId:=c.Query("destination")
	source:= strategy.CyberPortMap.NodeMap[sourceId]
	destination:= strategy.CyberPortMap.NodeMap[destinationId]
	_,_,pathIds:=strategy.CyberPortMap.Dijkstra(source.NodeID(),destination.NodeID())
	path:=make([]models.Node,0)
	for _, pathId := range pathIds {
		node:= strategy.CyberPortMap.NodeMap[pathId]
		id,_ := strconv.Atoi(node.Id)
		latitude,_:=strconv.ParseFloat(node.Latitude, 64)
		longitude,_:=strconv.ParseFloat(node.Longitude, 64)
		path=append(path, models.Node{
			Id:                     id,
			NameEnglish:            node.NameEnglish,
			NameChinese:            node.NameChinese,
			NameTraditionalChinese: node.NameChineseTradition,
			Latitude:               latitude,
			Longitude:              longitude,
			IntersectionalAngle:    node.IntersectionalAngle,
		})
	}
	responseData := &models.FetchPathOutput{Path: path}
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
func GetNodes(c *gin.Context) {
	//fixme:第二个版本加上页面的概念

	nodeMap, err :=model.GetNodes();
	if err != nil {
		fmt.Println(err)
		return
	}

	nodes:=make([]models.Node,0)
	for _, node := range nodeMap {
		nodes=append(nodes,node);
	}
	responseData := &models.GetNodesOutput{
		Nodes:nodes,
	}
	response.Success(c,"ok",responseData)
}

func GetNodeId(c *gin.Context) {
	if err := c.ShouldBind(&models.Node{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
	}
	latitude,_:=strconv.ParseFloat(c.Query("latitude"), 64);
	longitude,_:=strconv.ParseFloat(c.Query("longitude"), 64);
	newNode:=models.Node{
		Latitude:               latitude,
		Longitude:              longitude,
	}
	nodeId,err:=model.GetNodeID(newNode)
	if err!=nil{
		response.Error(c,"GetNodeID 失败")
		return
	}
	responseData := &models.Node{
		Id: nodeId,
	}
	response.Success(c,"ok",responseData)

}
func GetConnections(c *gin.Context) {
	//fixme:第二个版本加上页面的概念
	connections, err :=model.GetConnectionsList();
	if err != nil {
		fmt.Println(err)
		return
	}
	responseData := &models.GetConnectionsOutput{
		Connections: connections,
	}
	response.Success(c,"ok",responseData)
}


//Admin Section
func AddNode(c *gin.Context) {

	if err := c.ShouldBind(&models.AddNodeInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	latitude,_:=strconv.ParseFloat(c.PostForm("latitude"), 64);
	longitude,_:=strconv.ParseFloat(c.PostForm("longitude"), 64);
	newNode:=models.AddNodeInput{
		NameEnglish:            c.PostForm("nameEnglish"),
		NameChinese:            c.PostForm("nameChinese"),
		NameTraditionalChinese: c.PostForm("nameTraditionalChinese"),
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
	sourceLatitude,_:=strconv.ParseFloat(c.PostForm("sourceLatitude"), 64);
	sourceLongitude,_:=strconv.ParseFloat(c.PostForm("sourceLongitude"), 64);
	destinationLatitude,_:=strconv.ParseFloat(c.PostForm("destinationLatitude"), 64);
	destinationLongitude,_:=strconv.ParseFloat(c.PostForm("destinationLongitude"), 64);
	weight,_:=strconv.Atoi(c.PostForm("weight"));
	//fixme:获取sourceID 和 destinationID
	sourceId,err := model.GetNodeID(models.Node{
		Latitude:               sourceLatitude,
		Longitude:              sourceLongitude,
	})
	if err!=nil{
		response.Error(c,"AddConnection 失败")
		return
	}
	destinationId,err := model.GetNodeID(models.Node{
		Latitude:               destinationLatitude,
		Longitude:              destinationLongitude,
	})
	if err!=nil{
		response.Error(c,"AddConnection 失败")
		return
	}

	newConnection:=models.AddConnectionInput{
		SourceId:      sourceId,
		DestinationId: destinationId ,
		Weight:        weight,
	}
	err=model.AddConnection(c,newConnection)
	if err!=nil{
		response.Error(c,"AddConnection 失败")
		return
	}
	response.Success(c,"ok","")
}

func Delete(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Delete Node
	latitude,_:=strconv.ParseFloat(c.PostForm("nodeLatitude"), 64);
	longitude,_:=strconv.ParseFloat(c.PostForm("nodeLongitude"), 64);
	nodeId,err:=model.GetNodeID(models.Node{
		Latitude:               latitude,
		Longitude:              longitude,
	})
	if err!=nil{
		response.Error(c,"GetNodeID 失败")
		return
	}

	deleteNode:=models.DeleteNodeInput{
		Id:			nodeId,
	}
	err=model.DeleteNode(c,deleteNode)
	if err!=nil{
		response.Error(c,"DeleteNode 失败")
		return
	}
	//get all connection ids
	var connectionIdList = make([]int,0)
	fmt.Println(strategy.ConnectionsList);
	for _,value := range strategy.ConnectionsList{
		if value.Source.Id!=nodeId && value.Destination.Id!=nodeId{
			continue;
		}
		connectionIdList=append(connectionIdList,value.Id);
	}
	//Delete Connections
	for _,value :=range connectionIdList{
		deleteConnection:=models.DeleteConnectionByIDInput{
			Id:						value,
		}
		err = model.DeleteConnectionByID(c, deleteConnection)
		if err!=nil{
			response.Error(c,"DeleteConnection 失败")
			return
		}
	}
	response.Success(c,"ok","")
}

// Delete Connection Only
func DeleteConnection(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteConnectionByIDInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	sourceLatitude,_:=strconv.ParseFloat(c.PostForm("sourceLatitude"), 64);
	sourceLongitude,_:=strconv.ParseFloat(c.PostForm("sourceLongitude"), 64);
	destinationLatitude,_:=strconv.ParseFloat(c.PostForm("destinationLatitude"), 64);
	destinationLongitude,_:=strconv.ParseFloat(c.PostForm("destinationLongitude"), 64);
	//fixme:获取sourceID 和 destinationID
	sourceId,err := model.GetNodeID(models.Node{
		Latitude:               sourceLatitude,
		Longitude:              sourceLongitude,
	})
	if err!=nil{
		response.Error(c,"DeleteConnection 失败")
		return
	}
	destinationId,err := model.GetNodeID(models.Node{
		Latitude:               destinationLatitude,
		Longitude:              destinationLongitude,
	})
	if err!=nil{
		response.Error(c,"DeleteConnection 失败")
		return
	}
	//Delete Connection
	err = model.DeleteConnectionByNode(c, models.DeleteConnectionByNodeInput{
		SourceId:      sourceId,
		DestinationId: destinationId,
	})
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