package model

type Connection struct {
	Id int
	Source Node
	Destination Node
	Time int
}

func GetConnectionsMap() (connectionMap map[int]Connection, err error) {
	connection := Connection{}
	var result = make(map[int]Connection)
	rows,err := Db.Query("select c.id,n1.id,n1.name_english,n2.id,n2.name_english,c.weight from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id")
	for rows.Next() {
		err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Time)
		checkErr(err)
		result[connection.Id] = connection
	}
	return result, nil
}


func GetConnectionsList() (connectionMap []Connection, err error) {
	connection := Connection{}
	var result = make([]Connection,0)
	rows,err := Db.Query("select c.id,n1.id,n1.name_english,n2.id,n2.name_english,c.weight from connection AS c,node AS n1,node As n2 where c.source=n1.id and c.destination=n2.id")
	for rows.Next() {
		err = rows.Scan(&connection.Id,&connection.Source.Id,&connection.Source.NameEnglish,&connection.Destination.Id,&connection.Destination.NameEnglish,&connection.Time)
		checkErr(err)
		result=append(result,connection)
	}
	return result, nil
}


