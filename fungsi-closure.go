package main
import (
	"fmt"
)

// closure disimpan sebagai variabel

func main(){
	var getMinMax = func(n []int) (int, int){
		var min, max int
		for i, v :=range n{
			switch {
			case i == 0:
				min, max = v, v
			case v > max:
				max = v
			case v < min:
				min = v
			}
		}
		return min, max
	}

	numbers := []int{2,3,55,44,2,33,44,567,3,4,3,1}
	min, max := getMinMax(numbers)
	fmt.Println(min, max)


}