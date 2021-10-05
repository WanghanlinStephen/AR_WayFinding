package structure

//fixme:测试初始化图像格式
func Run() error {
	graph:=NewGraph()
	//初始化图形点
	Nodes:=make([]Label,0)
	NodeTest1:=Label{
		Id:                   "1",
		NameEnglish:          "TestEnglish",
		NameChinese:          "TestChinese",
		NameChineseTradition: "TestChineseTradition",
		Latitude:             "11.11",
		Longitude:            "22.22",
	}
	NodeTest2:=Label{
		Id:                   "2",
		NameEnglish:          "TestEnglish",
		NameChinese:          "TestChinese",
		NameChineseTradition: "TestChineseTradition",
		Latitude:             "11.11",
		Longitude:            "22.22",
	}
	Nodes=append(Nodes,NodeTest1,NodeTest2)
	graph.AddNodes(Nodes)
	//初始化图形点关系
	graph.AddEdge(NodeTest1,NodeTest2)
	//输出图像关系
	graph.adjacentEdgesExample(NodeTest1)
	return nil
}