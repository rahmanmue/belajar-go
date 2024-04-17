package main
import (
	"fmt"
	"strings"
)


// variadic fungsi yaitu fungsi dengan parameter sejenis yang tak terbatas

func main(){

	average := calculate(2,3,4,5,6,3,2,5,6,7,4,53,22)

	// printF menampilkan nilai
	fmt.Printf("Printf Rata-rata : %.2f \n", average)

	// Sprintf mengembalikan nilai berupa string
	msg := fmt.Sprintf("Sprintf Rata-rata : %.2f", average)
	fmt.Println(msg)


	// mengisi parameter dengan slice
	numbers2 := []int{2,5,6,7,8,22,44,66,77,88,99,77}
	average2 := calculate(numbers2...)
	fmt.Println(average2)

	// 2parameter campur biasa dan variadic
	yourHobbies("Rahman", "joggging", "badminton", "reading book")

	// atau bisa
	hobbies := []string{"watching movie", "watching football", "like liverpool"}
	yourHobbies("Mrahman", hobbies...)



}

// fungsi variadic
func calculate(numbers ...int) (float64){
	total := 0

	for _, v := range numbers{
		total += v 
	}

	// float64() merupakan casting

	average := float64(total) / float64(len(numbers))
	return average
}

// pengisian parameter variadic dan biasa, syarat variadic diakhir
func yourHobbies(name string, hobbies... string){
	hobby := strings.Join(hobbies, ", ")
	fmt.Println("hello my name", name)
	fmt.Println("my hobby is:", hobby)
}