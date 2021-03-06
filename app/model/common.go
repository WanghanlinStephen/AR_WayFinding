package model

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"pro/app/models"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func AddNode(c *gin.Context,node models.AddNodeInput) error{
	if _, err := Db.Exec("insert into node (name_english,name_chinese,name_traditional_chinese,latitude,longitude,intersectional_angle) VALUES (?,?,?,?,?,?);",node.NameEnglish,node.NameChinese,node.NameTraditionalChinese,node.Latitude,node.Longitude,node.IntersectionalAngle); err != nil {
		return err
	}
	return nil
}

func AddMap(c *gin.Context,mapInstance models.AddMapInput) error{
	if _, err := Db.Exec("insert into map (url,name,floor) VALUES (?,?,?);",mapInstance.Url,mapInstance.Name,mapInstance.Floor); err != nil {
		return err
	}
	return nil
}

func DeleteNode(c *gin.Context,node models.DeleteNodeInput) error{
	//fixme：search layer
	if _, err := Db.Exec("delete from node where id=?",node.Id); err != nil {
		return err
	}
	return nil
}

func DeleteConnectionByID(c *gin.Context,node models.DeleteConnectionByIDInput) error{
	//fimxe: search layer
	if _, err := Db.Exec("delete from connection where id=?",node.Id); err != nil {
		return err
	}
	return nil
}

func DeleteConnectionByNode(c *gin.Context,node models.DeleteConnectionByNodeInput) error{
	if _, err := Db.Exec("delete from connection where source=? and destination=?",node.SourceId,node.DestinationId); err != nil {
		return err
	}
	if _, err := Db.Exec("delete from connection where source=? and destination=?",node.DestinationId,node.SourceId); err != nil {
		return err
	}

	return nil
}

func DeleteMapByNameAndFloor(c *gin.Context,mapInstance models.DeleteMapByNameAndFloorInput) error {
	if _, err := Db.Exec("delete from map where name=? and floor=?",mapInstance.Name,mapInstance.Floor); err != nil {
		return err
	}
	return nil
}
func AddConnection(c *gin.Context,connection models.AddConnectionInput) error{
	if _, err := Db.Exec("insert into connection (source,destination,weight,map_id) VALUES (?,?,?,?);",connection.SourceId,connection.DestinationId,connection.Weight,connection.MapId); err != nil {
		return err
	}
	return nil
}


func GetNodeID(node models.Node) (int,error) {
	//fixme:等待优化
	var nodeId int
	rows,err := Db.Query("select id from node AS n where n.latitude=? and n.longitude=?",node.Latitude,node.Longitude)
	for rows.Next(){
		err = rows.Scan(&nodeId)
		checkErr(err)
	}
	return nodeId, nil
}

func GetBuildingNameByMapId(mapId int) (string,error) {
	var buildingName string
	rows,err := Db.Query("select name from map where id=?",mapId)
	for rows.Next(){
		err = rows.Scan(&buildingName)
		checkErr(err)
	}
	return buildingName, nil
}

func AddIsStaircase(mapId int ,nodeId int) error {
	if _, err := Db.Exec("update node set is_staircase= ? where id = ? ",1,nodeId); err != nil {
		return err
	}
	return nil
}

func GetMaps()([]models.Map,error){
	rows,err := Db.Query("select * from map")
	mapInstance := models.Map{}
	var maps = make([]models.Map,0)

	for rows.Next(){
		err = rows.Scan(&mapInstance.Id,&mapInstance.Url,&mapInstance.Name,&mapInstance.Floor)
		checkErr(err)
		maps=append(maps,mapInstance)
	}
	return maps, nil
}

func GetMapById(id string)(models.Map,error){
	rows,err := Db.Query("select * from map where map.id=?",id)
	mapInstance := models.Map{}
	for rows.Next(){
		err = rows.Scan(&mapInstance.Id,&mapInstance.Url,&mapInstance.Name,&mapInstance.Floor)
		checkErr(err)
	}
	return mapInstance,nil
}


func GetMapsByName(name string)([]models.Map,error){
	rows,err := Db.Query("select * from map where map.name=?",name)
	mapInstance := models.Map{}
	var mapInstanceList = make([]models.Map,0)
	for rows.Next(){
		err = rows.Scan(&mapInstance.Id,&mapInstance.Url,&mapInstance.Name,&mapInstance.Floor)
		checkErr(err)
		mapInstanceList=append(mapInstanceList,mapInstance)
	}
	return mapInstanceList,nil
}
