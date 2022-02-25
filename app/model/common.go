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
	if _, err := Db.Exec("insert into node (name_english,name_chinese,name_traditional_chinese,latitude,longitude,intersectional_angle) VALUES (?,?,?,?,?,?);",node.NameEnglish,node.NameChinese,node.NameTraditionalChinese,node.Longitude,node.Latitude,node.IntersectionalAngle); err != nil {
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

func AddConnection(c *gin.Context,connection models.AddConnectionInput) error{
	if _, err := Db.Exec("insert into connection (source,destination,weight) VALUES (?,?,?);",connection.SourceId,connection.DestinationId,connection.Weight); err != nil {
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

func GetMaps()([]models.Map,error){
	rows,err := Db.Query("select * from map")
	mapInstance := models.Map{}
	var maps = make([]models.Map,0)

	for rows.Next(){
		err = rows.Scan(&mapInstance.Id,&mapInstance.Name,&mapInstance.Url)
		checkErr(err)
		maps=append(maps,mapInstance)
	}
	return maps, nil
}

func GetMapByName(name string)(models.Map,error){
	rows,err := Db.Query("select * from map where map.name=?",name)
	mapInstance := models.Map{}
	for rows.Next(){
		err = rows.Scan(&mapInstance.Id,&mapInstance.Url,&mapInstance.Name)
		checkErr(err)
	}
	return mapInstance,nil
}