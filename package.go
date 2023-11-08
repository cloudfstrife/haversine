// Package haversine is a Go library which implements the haversine formula.
//
// The below example shows how to use this package
//
//	package main
//
//	import (
//		"fmt"
//
//		"github.com/cloudfstrife/haversine"
//	)
//
//	func main() {
//		point1 := haversine.Point{Latitude: 22.54587746, Longitude: 114.12873077}
//		point2 := haversine.Point{Latitude: 23.0, Longitude: 115.0}
//
//		distance := haversine.Distance(point1, point2)
//		fmt.Println(distance)
//	}
package haversine
