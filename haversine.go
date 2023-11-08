package haversine

import (
	"math"
)

// EarthRadius  is the radius of the earth in kilometre
const EarthRadius = 6371

// Point location point
type Point struct {
	//Longitude longitude of location
	Longitude float64
	//Latitude latitude of location
	Latitude float64
}

/*
Distance return the Distance between two point in kilometre.

Parameters:

point1 : first location point

point2 : second location point
*/
func Distance(point1, point2 Point) float64 {

	lat1 := AngleToRadian(point1.Latitude)
	lon1 := AngleToRadian(point1.Longitude)
	lat2 := AngleToRadian(point2.Latitude)
	lon2 := AngleToRadian(point2.Longitude)

	vLon := math.Abs(lon1 - lon2)
	vLat := math.Abs(lat1 - lat2)

	return 2 * EarthRadius * math.Asin(math.Sqrt(HaverSin(vLat)+math.Cos(lat1)*math.Cos(lat2)*HaverSin(vLon)))
}

// AngleToRadian convert angle to radian
func AngleToRadian(angle float64) float64 {
	return angle * math.Pi / 180
}

// RadianToAngle convert radian to angle
func RadianToAngle(radian float64) float64 {
	return radian * 180.0 / math.Pi
}

// HaverSin haversin
func HaverSin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
