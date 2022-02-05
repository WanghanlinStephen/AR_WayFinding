package strategy

import (
	"math"
	"strconv"
)

func Radians(x float64) float64 {
	return x/math.Pi
}

//fixme:添加墙面角度
func GetAngle(lon1Str,lat1Str string,intersectionalAngle float64,lon2Str,lat2Str string) float64{
	y1,_:= strconv.ParseFloat(lon1Str, 64)
	y2,_:= strconv.ParseFloat(lon2Str, 64)
	x1,_:= strconv.ParseFloat(lat1Str, 64)
	x2,_:= strconv.ParseFloat(lat2Str, 64)
	angle:=math.Atan(math.Abs(y2-y1)/math.Abs(x2-x1))*(180/math.Pi)
	//direction:="North"
	isWallRight:=false;
	//第一象限
	if x2>x1 && y2>y1{
		//检查墙相对人的位置
		if intersectionalAngle-angle<0{
			isWallRight=true;
		}
		if isWallRight{
			angle=180-angle+intersectionalAngle
		}else {
			angle=intersectionalAngle-angle
		}
	}
	//第二象限
	if x2<x1 && y2>y1{
		//检查墙相对人的位置
		if intersectionalAngle<180-angle{
			isWallRight=true;
		}
		if isWallRight{
			angle=angle+intersectionalAngle
		}else {
			angle=angle+intersectionalAngle-180
		}
	}
	//第三象限
	if x2<x1 && y2<y1{
		//检查墙相对人的位置
		if intersectionalAngle>angle{
			isWallRight=true;
		}
		if isWallRight{
			//fixme:todo
			angle=180-(angle+180-intersectionalAngle)
		}else {
			angle=intersectionalAngle+180-angle
		}
	}
	//第四象限
	if x2>x1 && y2<y1{
		if 180-intersectionalAngle>angle{
			isWallRight=true;
		}
		if isWallRight{
			angle=intersectionalAngle+angle
		}else {
			angle=intersectionalAngle+angle-180
		}
	}
	return angle
}