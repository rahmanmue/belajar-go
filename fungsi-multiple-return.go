package main
import (
	"fmt"
	"math"
)

func main(){
	var diameter float64 = 15.55

	area, circumFerence := calculate(diameter)

	fmt.Printf("area : %.2f, circumFerence: %.2f \n", area, circumFerence)

	area2, circumFerence2 := calculate2(33.4)
	fmt.Printf("area : %.3f, circumFerence: %.3f \n", area2, circumFerence2)

}

// jika multiple return return type nya tulis semua
func calculate(d float64) (float64, float64){
	area := math.Pi * math.Pow(d / 2, 2)
	circumFerence := math.Pi * d

	return area, circumFerence
}

// return nya bisa didiefinisikan diawal 
// type datanya sudah didefinisikan tidak perlu :=
func calculate2(d float64) (area2 float64, circumFerence2 float64 ){
	area2 = math.Pi * math.Pow(d / 2, 2)
	circumFerence2 = math.Pi * d

	return
}