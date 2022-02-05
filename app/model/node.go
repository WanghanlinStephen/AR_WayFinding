package model

import "pro/app/models"

func GetNodes() (nodeMap map[int]models.Node, err error) {
	//fixme:等待优化
	node := models.Node{}
	var result = make(map[int]models.Node)
	rows,err := Db.Query("select id,name_english,name_chinese,name_traditional_chinese,latitude,longitude,intersectional_angle from node")
	for rows.Next(){
		err = rows.Scan(&node.Id,&node.NameEnglish,&node.NameChinese,&node.NameTraditionalChinese,&node.Latitude,&node.Longitude,&node.IntersectionalAngle)
		checkErr(err)
		result[node.Id] = node
	}
	return result, nil
}
