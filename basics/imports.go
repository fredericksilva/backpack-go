package main

import (
   "fmt"
   "math"
)

func main(){
	fmt.Printf("Now you %g problems.", math.Nextafter(2,3))
}
