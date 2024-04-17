package main
import "fmt"

func main(){

	// operator aritmatika
	var value = ((2 * 5) + 12) % 3;
	fmt.Printf("value : %d \n", value)

	// operator perbandingan
	var isEqual = (value == 1)
	var isPositive = value > 2
	var isNegative = value < 0
	
	fmt.Println(isEqual, isPositive, isNegative)

	// operator logika || && !
	fmt.Printf("or : %t \n", (isEqual || isPositive) )
	fmt.Printf("and : %t \n", (isPositive && isNegative))
	fmt.Printf("negasi : %t \n", (!isNegative))
}