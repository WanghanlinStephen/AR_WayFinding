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
	if _, err := Db.Exec("insert into node set id=?, name_english=?, name_chinese=?,name_traditional_chinese=?,longitude=?,latitude=?,intersectional_angle=?",node.Id,node.NameEnglish,node.NameChinese,node.NameTraditionalChinese,node.Longitude,node.Latitude,node.IntersectionalAngle); err != nil {
		return err
	}
	return nil
}


func DeleteNode(c *gin.Context,node models.DeleteNodeInput) error{
	if _, err := Db.Exec("delete from node where id=?",node.Id); err != nil {
		return err
	}
	return nil
}

func DeleteConnection(c *gin.Context,node models.DeleteConnectionInput) error{
	if _, err := Db.Exec("delete from connection where id=?",node.Id); err != nil {
		return err
	}
	return nil
}

func AddConnection(c *gin.Context,connection models.AddConnectionInput) error{
	if _, err := Db.Exec("insert into connection (id,source,destination,weight) VALUES (?,?,?,?);",connection.Id,connection.SourceId,connection.DestinationId,connection.Weight); err != nil {
		return err
	}
	return nil
}
//
//func UpdateNode(c *gin.Context,node models.ModifyInput) {
//	if _, err := Db.Exec("update into node set id=?, name_english=?, name_chinese=?,name_traditional_chinese=?,longitude=?,latitude=?",node.Id,node.NameEnglish,node.NameChinese,node.NameTraditionalChinese,node.Longitude,node.Latitude); err != nil {
//		panic(err)
//	}
//	c.AbortWithStatus(204)
//}