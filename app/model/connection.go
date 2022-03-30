package model

import "pro/app/models"

func GetConnectionsMap() (connectionMap map[int]models.Connection, err error) {
	//fixme:等待优化
	sql:="select c.id,n1.id,n1.name_english,n1.name_chinese,n1.name_traditional_chinese,n1.latitude,n1.longitude,n2.id,n2.name_english,n2.name_chinese,n2.name_traditional_chinese,n2.latitude,n2.longitudec,n2.weight from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id"
	connection := models.Connection{}
	var result = make(map[int]models.Connection)
	rows,err := Db.Query(sql)
	for rows.Next() {
		err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Source.NameChinese,&connection.Source.NameTraditionalChinese,&connection.Source.Latitude,&connection.Source.Longitude,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Source.NameChinese,&connection.Source.NameTraditionalChinese,&connection.Source.Latitude,&connection.Source.Longitude,&connection.Time)
		checkErr(err)
		result[connection.Id] = connection
	}
	return result, nil
}


func GetConnectionsList() (connectionList []models.Connection, err error) {
	//fixme:等待优化
	sql:="select c.id,n1.id,n1.name_english,n1.name_chinese,n1.name_traditional_chinese,n1.latitude,n1.longitude,n1.intersectional_angle,n1.is_staircase,n2.id,n2.name_english,n2.name_chinese,n2.name_traditional_chinese,n2.latitude,n2.longitude,n2.intersectional_angle,n2.is_staircase,c.weight,c.map_id from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id"
	//sql:="select c.id,n1.id,n1.name_english,n2.id,n2.name_english,c.weight from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id"
	connection := models.Connection{}
	var result = make([]models.Connection,0)
	rows,err := Db.Query(sql)
	for rows.Next() {
		err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Source.NameChinese,&connection.Source.NameTraditionalChinese,&connection.Source.Latitude,&connection.Source.Longitude,&connection.Source.IntersectionalAngle,&connection.Source.IsStaircase,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Destination.NameChinese,&connection.Destination.NameTraditionalChinese,&connection.Destination.Latitude,&connection.Destination.Longitude,&connection.Destination.IntersectionalAngle,&connection.Destination.IsStaircase,&connection.Time,&connection.MapId)
		//err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Time)
		checkErr(err)
		result=append(result,connection)
	}
	return result, nil
}

func GetConnectionsListByMapId(mapId int) (connectionList []models.Connection, err error) {
	//fixme:等待优化
	//sql:="select c.id,n1.id,n1.name_english,n2.id,n2.name_english,c.weight from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id"
	connection := models.Connection{}
	var result = make([]models.Connection,0)
	rows,err := Db.Query("select c.id,n1.id,n1.name_english,n1.name_chinese,n1.name_traditional_chinese,n1.latitude,n1.longitude,n1.intersectional_angle,n1.is_staircase,n2.id,n2.name_english,n2.name_chinese,n2.name_traditional_chinese,n2.latitude,n2.longitude,n2.intersectional_angle,n2.is_staircase,c.weight,c.map_id from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id and c.map_id=?",mapId)
	for rows.Next() {
		err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Source.NameChinese,&connection.Source.NameTraditionalChinese,&connection.Source.Latitude,&connection.Source.Longitude,&connection.Source.IntersectionalAngle,&connection.Source.IsStaircase,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Destination.NameChinese,&connection.Destination.NameTraditionalChinese,&connection.Destination.Latitude,&connection.Destination.Longitude,&connection.Destination.IntersectionalAngle,&connection.Destination.IsStaircase,&connection.Time,&connection.MapId)
		//err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Time)
		checkErr(err)
		result=append(result,connection)
	}
	return result, nil
}

