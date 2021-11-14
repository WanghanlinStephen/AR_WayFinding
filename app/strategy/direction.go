package strategy

import (
	"math"
	"strconv"
)

func Radians(x float64) float64 {
	return x/math.Pi
}

func GetAngle(lon1Str,lat1Str,lon2Str,lat2Str string) (string,float64){
	y1,_:= strconv.ParseFloat(lon1Str, 64)
	y2,_:= strconv.ParseFloat(lon2Str, 64)
	x1,_:= strconv.ParseFloat(lat1Str, 64)
	x2,_:= strconv.ParseFloat(lat2Str, 64)
	angle:=math.Atan(math.Abs(y2-y1)/math.Abs(x2-x1))*(180/math.Pi)
	direction:="North"
	//第一象限
	if x2>x1 && y2>y1{
		if angle>45{
			direction="North"
		}else{
			direction="East"
		}
	}
	//第二象限
	if x2<x1 && y2>y1{
		if angle>45{
			direction="North"
		}else{
			direction="West"
		}
	}
	//第三象限
	if x2<x1 && y2<y1{
		if angle>45{
			direction="South"
		}else{
			direction="West"
		}
	}
	//第四象限
	if x2>x1 && y2<y1{
		if angle>45{
			direction="South"
		}else{
			direction="East"
		}
	}
	return direction,angle
}
//func GetAngle(lon1Str,lat1Str,lon2Str,lat2Str string) float64{
//	lon1,_:= strconv.ParseFloat(lon1Str, 64)
//	lon2,_:= strconv.ParseFloat(lon2Str, 64)
//	lat1,_:= strconv.ParseFloat(lat1Str, 64)
//	lat2,_:= strconv.ParseFloat(lat2Str, 64)
//	var numerator = math.Sin(Radians(lon2-lon1)) * math.Cos(Radians(lat2))
//	var denominator = math.Cos(Radians(lat1)) * math.Sin(Radians(lat2))- math.Sin(Radians(lat1)) * math.Cos(Radians(lat2)) * math.Cos(Radians(lon2 - lon1))
//	var x = math.Atan2(math.Abs(numerator), math.Abs(denominator))
//	var result = x
//	// 右象限
//	if lon2 > lon1 {
//		// 第一象限
//		if lat2 > lat1{
//			result = x
//		} else if lat2 < lat1{ // 第四象限
//			result = math.Pi - x
//		}else{
//			result = math.Pi / 2 // 在正x轴上
//		}
//	}else if lon2 < lon1{
//		if lat2 > lat1{// 第二象限
//			result = 2 * math.Pi - x
//		}else if lat2 < lat1{ // 第三象限
//			result = math.Pi + x
//		}else{
//			result = math.Pi * 3 / 2 // 在负x轴上
//		}
//	}else // 同一经度
//	{
//		if lat2 > lat1{ // 在正y轴上
//			result = 0
//		}else if lat2 < lat1{
//			result = math.Pi // 在负y轴中
//		}else{
//			fmt.Println("不应该在同一个地方!")
//		}
//	}
//	return result * 180 / math.Pi
//}
