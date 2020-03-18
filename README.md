# haversine

haversine is a Go library which implements the haversine formula.

$$
haversin(\displaystyle \frac{d}{R})=haversin(\phi2-\phi1)+cos(\phi1)cos(\phi2)haversin(\lambda2-\lambda1)
$$

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