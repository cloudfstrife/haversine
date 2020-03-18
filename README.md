# haversine

haversine is a Go library which implements the haversine formula.

![haversine](https://latex.codecogs.com/gif.download?haversin%28%5Cdisplaystyle%20%5Cfrac%7Bd%7D%7BR%7D%29%3Dhaversin%28%5Cphi2-%5Cphi1%29+cos%28%5Cphi1%29cos%28%5Cphi2%29haversin%28%5Clambda2-%5Clambda1%29)

## USAGE

```
package main

import (
	"fmt"

	"github.com/cloudfstrife/haversine"
)

func main() {
	point1 := haversine.Point{Latitude: 22.54587746, Longitude: 114.12873077}
	point2 := haversine.Point{Latitude: 23.0, Longitude: 115.0}

	distance := haversine.Distance(point1, point2)
	fmt.Println(distance)
}
```