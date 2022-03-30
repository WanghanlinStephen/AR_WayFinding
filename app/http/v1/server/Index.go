package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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
	//fixme:补充上返回的楼数量+最近的楼梯口,不进行处理x
	if source.MapId != destination.MapId{
		response.Error(c,"destination.mapId != source.mapId")
		return
	}
	shortestDistance,nextStep,_:=strategy.CyberPortMap.Dijkstra(source.NodeID(),destination.NodeID(),source.MapId)
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
	mapId := source.MapId
	//source & destination floor doesn't match
	closetestStaircaseId := ""
	destinationFloor := 0
	pathIds :=make([]string,0)
	if source.MapId != destination.MapId{
		//Fetch Floor ID:
		mapInstance,err := model.GetMapById(strconv.Itoa(destination.MapId))
		if err!=nil{
			response.Error(c,"FetchMapByFilter 失败")
			return
		}
		destinationFloor = mapInstance.Floor
		//Fetch nearest staircase and return the path
		//Step1:fetch all nodes with tag: staircase = true
		staircaseList:=make([]string,0)
		for _, node := range strategy.CyberPortMap.NodeMap {
			if node.MapId != source.MapId{
				continue
			}
			if node.IsStaircase == 0 {
				continue
			}
			staircaseList=append(staircaseList,node.Id)
		}
		//Step2:find the closet staircase nearby
		closetestStaircaseDistance := math.MaxFloat32
		for _,staircaseId := range staircaseList {
			shortestDistance,_,shorestFullPath:=strategy.CyberPortMap.Dijkstra(source.NodeID(),staircaseId,source.MapId)
			if shortestDistance < closetestStaircaseDistance{
				closetestStaircaseId = staircaseId
				pathIds = shorestFullPath
				closetestStaircaseDistance = shortestDistance
			}
		}
	}else{
		_,_,pathIds=strategy.CyberPortMap.Dijkstra(source.NodeID(),destination.NodeID(),mapId)
	}
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
	responseData := &models.FetchPathOutput{
		Path: path,
		IsSameFloor:source.MapId == destination.MapId,
		DestinationId:closetestStaircaseId,
		Floor:destinationFloor,
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

func GetNodesByMapId(c *gin.Context) {
	if err := c.ShouldBind(&models.GetNodesByMapId{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	id := c.Query("id")
	mapIdStr,_:= strconv.Atoi(id)
	nodes:=make([]models.Node,0)
	for _, node := range strategy.CyberPortMap.NodeMap {
		if node.MapId != mapIdStr{
			continue
		}
		nodeId , _ := strconv.Atoi(node.Id)
		nodeLatitude, _ := strconv.ParseFloat(node.Latitude, 64)
		nodeLongitude, _ := strconv.ParseFloat(node.Longitude, 64)
		nodes=append(nodes,models.Node{
			Id:                     nodeId,
			NameEnglish:            node.NameEnglish,
			NameChinese:            node.NameChinese,
			NameTraditionalChinese: node.NameChineseTradition,
			Latitude:               nodeLatitude,
			Longitude:              nodeLongitude,
			IntersectionalAngle:    node.IntersectionalAngle,
			IsStaircase:            node.IsStaircase,
		})
	}
	responseData := &models.GetNodesByMapIdOutput{
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

func GetConnectionsByMapId(c *gin.Context) {
	if err := c.ShouldBind(&models.GetNodesByMapId{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	id := c.Query("id")
	mapId,_:= strconv.Atoi(id)
	connections, err :=model.GetConnectionsListByMapId(mapId)
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
	//todo:refresh map
	strategy.Initialization()
	response.Success(c,"ok","")
}




func AddConnection(c *gin.Context) {
	//if err := c.ShouldBind(&models.AddConnectionInput{}); err != nil {
	//	fmt.Println(err.Error())
	//	response.Error(c, "参数错误")
	//	return
	//}
	//Modify Database
	sourceLatitude,_:=strconv.ParseFloat(c.PostForm("sourceLatitude"), 64)
	sourceLongitude,_:=strconv.ParseFloat(c.PostForm("sourceLongitude"), 64)
	destinationLatitude,_:=strconv.ParseFloat(c.PostForm("destinationLatitude"), 64)
	destinationLongitude,_:=strconv.ParseFloat(c.PostForm("destinationLongitude"), 64)
	weight,_:=strconv.ParseFloat(c.PostForm("weight"),64)
	mapId,_:=strconv.Atoi(c.PostForm("mapId"))
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
		MapId:		   mapId,
	}
	err=model.AddConnection(c,newConnection)
	if err!=nil{
		response.Error(c,"AddConnection 失败")
		return
	}
	//todo:refresh map
	strategy.Initialization()
	response.Success(c,"ok","")
}


func AddStaircase(c *gin.Context) {
	if err := c.ShouldBind(&models.AddStaircaseInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	mapId,_ := strconv.Atoi(c.PostForm("mapId"))
	latitude:=c.PostForm("latitude")
	longitude:=c.PostForm("longitude")
	nodeIdStr := ""
	//fixme:此时map没有更新,无法找到进行写操作
	for _, node := range strategy.CyberPortMap.NodeMap {
		if node.MapId != mapId{
			continue
		}
		if node.Latitude != latitude && node.Longitude != longitude {
			continue
		}
		nodeIdStr = node.Id
	}
	nodeId,_:= strconv.Atoi(nodeIdStr)
	//change is_staircase = 1
	err :=model.AddIsStaircase(mapId,nodeId)
	if err!=nil{
		response.Error(c,"AddStaircaseEntry 失败")
		return
	}
	response.Success(c,"ok","")
}

func AddMap(c *gin.Context) {
	if err := c.ShouldBind(&models.AddMapInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Modify Database
	floor, _ := strconv.Atoi(c.PostForm("Floor"))
	newMap:=models.AddMapInput{
		Url:   c.PostForm("Url"),
		Name:  c.PostForm("Name"),
		Floor: floor,
	}
	err:=model.AddMap(c,newMap)
	if err!=nil{
		response.Error(c,"AddMap 失败")
		return
	}
	response.Success(c,"ok","")
}

//func AddEmergent(c *gin.Context) {
//	if err := c.ShouldBind(&models.AddEmergentInput{}); err != nil {
//		fmt.Println(err.Error())
//		response.Error(c, "参数错误")
//		return
//	}
//	mapId,_ := strconv.Atoi(c.PostForm("mapId"))
//	latitude,_:=strconv.ParseFloat(c.PostForm("nodeLatitude"), 64)
//	longitude,_:=strconv.ParseFloat(c.PostForm("nodeLongitude"), 64)
//	nodeId,err:=model.GetNodeID(models.Node{
//		Latitude:               latitude,
//		Longitude:              longitude,
//	})
//	if err!=nil{
//		response.Error(c,"GetNodeID 失败")
//		return
//	}
//	//add node information to map
//	err =model.AddEmergentEntry(mapId,nodeId)
//	if err!=nil{
//		response.Error(c,"AddEmergentEntry 失败")
//		return
//	}
//	response.Success(c,"ok","")
//}

func Delete(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	//Delete Node
	//fixme:改造map使用
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
	//todo:refresh map
	strategy.Initialization()
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
	//todo:refresh map
	strategy.Initialization()
	response.Success(c,"ok","")
}

func DeleteMap(c *gin.Context) {
	if err := c.ShouldBind(&models.DeleteMapByNameAndFloorInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}

	floor, _ := strconv.Atoi(c.PostForm("floor"))
	mapInstance:=models.DeleteMapByNameAndFloorInput{
		Name:  c.PostForm("Name"),
		Floor: floor,
	}
	err := model.DeleteMapByNameAndFloor(c, mapInstance)
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


//no input required
func FetchMaps(c *gin.Context){
	maps,err := model.GetMaps()
	if err!=nil{
		response.Error(c,"FetchMaps 失败")
		return
	}
	responseData := &models.GetMapsOutput{
		Maps:     maps,
	}
	response.Success(c,"ok",responseData)
}
func FetchMapNames(c *gin.Context){
	maps,err := model.GetMaps()
	if err!=nil{
		response.Error(c,"FetchMaps 失败")
		return
	}
	mapNamesList := make([]string,0)
	mapNameMap := make(map[string]bool)

	for _,value := range maps{
		if _, ok := mapNameMap[value.Name]; ok {
			continue
		}
		mapNamesList=append(mapNamesList,value.Name)
		mapNameMap[value.Name]=true
	}

	responseData := &models.GetMapNamesOutput{
		Names:     mapNamesList,
	}
	response.Success(c,"ok",responseData)
}
func FetchMapByIdFilter(c *gin.Context){
	if err := c.ShouldBind(&models.GetMapByIdInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	id := c.Query("id")
	mapInstance,err := model.GetMapById(id)
	if err!=nil{
		response.Error(c,"FetchMapByIDFilter 失败")
		return
	}
	responseData := &models.GetMapByIdOutput{
		Map:     mapInstance,
	}
	response.Success(c,"ok",responseData)
}

func FetchMapByNameFilter(c *gin.Context){
	if err := c.ShouldBind(&models.GetMapByNameInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	name := c.Query("name")
	mapInstanceList,err := model.GetMapsByName(name)
	if err!=nil{
		response.Error(c,"FetchMapByNameFilter 失败")
		return
	}
	responseData := &models.GetMapByNameOutput{
		Map:     mapInstanceList,
	}
	response.Success(c,"ok",responseData)

}

func FetchMapIdByNodeId(c *gin.Context){
	if err := c.ShouldBind(&models.GetMapIdByNodeIdInput{}); err != nil {
		fmt.Println(err.Error())
		response.Error(c, "参数错误")
		return
	}
	nodeId := c.Query("id")
	mapId := 0
	for _, node := range strategy.CyberPortMap.NodeMap{
		if node.Id == nodeId{
			mapId = node.MapId
		}
	}

	responseData := &models.GetMapIdByNodeIdOutput{
		Id:     mapId,
	}
	response.Success(c,"ok",responseData)
}