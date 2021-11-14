package model

type Node struct {
	Id int
	NameEnglish string
	NameChinese string
	NameTraditionalChinese string
	Latitude float64
	Longitude float64
}
func GetNodes() (nodeMap map[int]Node, err error) {
	node := Node{}
	var result = make(map[int]Node)
	rows,err := Db.Query("select id,name_english,name_chinese,name_traditional_chinese,latitude,longitude from node")
	for rows.Next(){
		err = rows.Scan(&node.Id,&node.NameEnglish,&node.NameChinese,&node.NameTraditionalChinese,&node.Latitude,&node.Longitude)
		checkErr(err)
		result[node.Id] = node

	}
	return result, nil
}
