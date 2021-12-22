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
	//第一象限
	if x2>x1 && y2>y1{
		angle=intersectionalAngle-angle
	}
	//第二象限
	if x2<x1 && y2>y1{
		angle=180-(intersectionalAngle+angle)
	}
	//第三象限
	if x2<x1 && y2<y1{
		angle=angle+(180-intersectionalAngle)
	}
	//第四象限
	if x2>x1 && y2<y1{
		angle=angle+intersectionalAngle
	}
	return angle
}