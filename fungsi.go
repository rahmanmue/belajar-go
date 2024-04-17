package main
import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main(){
	names := []string{"John", "Wick"}
	printMessage("halo", names)

	var randomVal int

	randomVal = randomWithRange(2,10)
	fmt.Println(randomVal)

}

func printMessage(message string, arr []string){
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// cara penulisan func jika ada param dan return
// func nameOfFunc(paramA type, paramB type) returnType
// lebih singkat jika type param nya sama func nameOfFunc(paramA, paramB type) returnType


func randomWithRange(min, max int) int{
	return randomizer.Int()%(max-min+1) + min
}