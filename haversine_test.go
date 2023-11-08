package haversine

import (
	"fmt"
	"testing"
)

var testCases = map[string]struct {
	Angle  float64
	Radian float64
}{
	"General":   {Angle: 57.29577951, Radian: 1.00},
	"Angle-90":  {Angle: 90.00000000, Radian: 1.57079633},
	"Angle-180": {Angle: 180.00000000, Radian: 3.14159265},
}

var distTestCases = map[string]struct {
	point1 Point
	point2 Point
	dist   float64
}{
	"General": {point1: Point{Longitude: 106.486654, Latitude: 29.490295}, point2: Point{Longitude: 106.581515, Latitude: 29.615467}, dist: 16.67090},
	"Dist-0":  {point1: Point{Longitude: 0.0, Latitude: 0.0}, point2: Point{Longitude: 0.0, Latitude: 0.0}, dist: 0.0},
}

func TestAngleToRadian(t *testing.T) {
	for name, testCase := range testCases {
		v := AngleToRadian(testCase.Angle)
		if fmt.Sprintf("%.5f", v) != fmt.Sprintf("%.5f", testCase.Radian) {
			t.Errorf("%s Faild : want : %.5f got : %.5f", name, testCase.Radian, v)
		}
	}
}

func TestRadianToAngle(t *testing.T) {
	for name, testCase := range testCases {
		v := RadianToAngle(testCase.Radian)
		if fmt.Sprintf("%.5f", v) != fmt.Sprintf("%.5f", testCase.Angle) {
			t.Errorf("%s Faild : want : %.5f got : %.5f", name, testCase.Angle, v)
		}
	}
}

func TestDistance(t *testing.T) {
	for name, testCase := range distTestCases {
		v := Distance(testCase.point1, testCase.point2)
		if fmt.Sprintf("%.5f", v) != fmt.Sprintf("%.5f", testCase.dist) {
			t.Errorf("%s Faild : want : %.5f got : %.5f", name, testCase.dist, v)
		}
	}
}
