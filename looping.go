package main
import "fmt"

func main(){
	// di go perulangan hanya for

	// 1. for umumnya
	for i:=0; i< 2; i++ {
		fmt.Println(i, "biasa")
	}

	// 2. for hanya kondisi
	var a = 10
	for a < 13{
		fmt.Println(a, "pke kondisi")
		a++
	}


	// 3. for tanpa argumen
	var b = 20
	for{
		if b == 23{
			break
		}

		fmt.Println(b, "tanpa argumen")
		b++
	}

	// 4. for - range digunakan untuk looping data gabungan

	// string
	var word string = "ini kata"

	for i, v := range word{
		fmt.Printf("index : %d, value: %d \n", i, v)
	}

	// array
	var ys = [5]int{1,11,23,44,65}

	for _, v := range ys{
		fmt.Printf("value : %d \n", v)
	}

	// slice array dari ys
	var yz = ys[3:]
	for _, v := range yz{
		fmt.Println(v)
	}

	// map
	var kvs = map[byte]int{'a':0, 'b':1, 'c':3}
	for k, v := range kvs{
		fmt.Println("key", k, "value", v)
	}

	// break dan continue 

	for i := 1; i<=10; i++{
		if i == 5{
			fmt.Println("Continue nih")
			continue
		}

		fmt.Println(i)

		if i == 8 {
			fmt.Println("Break nih")
			break
		}
	} 

	// perulangan bersarang
	for a:=0; a<5; a++{
		for b:=a; b<5; b++{
			fmt.Print(b, " ")
		}

		fmt.Println()
		
	}

	// pemanfaatan label dalam perulangan
	// penggunaan ini akan menghentikan perulangan dan tidak dilanjutkan lagi
	outerloop:
	for a:=0; a<5; a++{
		for b:=a; b<5; b++{
			if b == 4 {
				break outerloop
			}
			fmt.Printf("[%d] [%d] \n", a,b )
		}		
		
	}

	// tanpa label
	fmt.Println("tanpa label :")
	
	for a:=0; a<5; a++{
		for b:=a; b<5; b++{
			if b == 4 {
				break 
			}
			fmt.Printf("[%d] [%d] \n", a,b )
		}		
		
	}





}