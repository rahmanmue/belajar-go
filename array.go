package main
import "fmt"

func main(){
	var names [3]string;
	names[0] = "monkey"
	names[1] = "d"
	names[2] = "luffy"
	
	fmt.Println(names[0], names[1], names[2])

	// jika melebihi alokasi awal akan error
	// names[3]="ez"

	// inisialisasi array langsung jumlah elemennya, 
	// bukan diitung dari 0
	var fruits = [4]string{"apel", "pisang", "pepaya", "kiwi"}
	fmt.Println(len(fruits))
	fmt.Println(fruits) 

	// cara vertikal
	var club = [4]string{
		"liverpool",
		"madrid",
		"mu",
		"barca",
	}

	fmt.Println(club)

	// inilisialisasi tanpa jumlah elemen
	var numbers = [...]int{1,2,3,4,5,6,7,7,8}
	fmt.Println(numbers)

	// dibacanya array dengan 2 baris berisi 3 kolom
	var numbers2 = [2][3]int{{1,2,3}, {4,5,6}}
	fmt.Println(numbers2)

	var number3 = [3][2] int{[2]int{2,2}, [2]int{1,5}, [2]int{5,5}}
	fmt.Println(number3)
	
	for i := 0; i<len(club); i++{
		fmt.Println(club[i])
	}

	for i,v := range fruits{
		fmt.Println(i, "=", v)
	}

	for _, fruit := range fruits{
		fmt.Println(fruit)
	}

	// jika butuh index nya
	for i := range club{
		fmt.Println(i)
	}

	// alokasi array dengan make
	var stackBelajar = make([]string, 2)
	stackBelajar[0] = "frontend"
	stackBelajar[1] = "backend"

	fmt.Println(stackBelajar)



}