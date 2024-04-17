package main
import "fmt"

func main(){
	var nilai = 80

	if nilai < 50 && nilai > 0 {
		fmt.Println("D")
	} else if nilai > 50 && nilai < 75 {
		fmt.Println("C")
	} else if nilai > 75 && nilai < 90 {
		fmt.Println("B")
	} else if nilai > 90 && nilai < 100 {
		fmt.Println("A")
	} else {
		fmt.Println("Masukan rentang nilai dari 0 - 100")
	}

	// variabel temporary pada if

	if percent := nilai * 10 /100; percent > 1 {
		fmt.Printf("nilai : %d \n", percent)
	} else if percent < 1 {
		fmt.Printf("nilai : %d \n", percent)
	}

	// seleksi switch cas

	var point = 90

	switch point{

	case 80:
		fmt.Println("nilai 80")
	case 70:
		fmt.Println("nilai 70")
	case 90:
		

	}

	// case banyak kondisi

	var score = 10;

	switch score{
	case 5:
		fmt.Println("kecil")
	case 8,9,10:
		fmt.Println("luar biasa")
	default:
		fmt.Println("lebih dari 10 kurang dari 5 ini default tapi")
	}


	// default kurung kurawal
	var score2 = 11;

	switch score2{
	case 5:
		fmt.Println("kecil")
	case 8,9,10:
		fmt.Println("luar biasa")
	default:
		{
			fmt.Println("ini default")
			fmt.Println("lebih dari 10 kurang dari 5")
		}
		
	}


	// switch case gaya if switch nya pas awal ga dimention
	// mention nya di case

	var klasemen_liverpool = 1;

	switch {
	case (klasemen_liverpool == 1):
		fmt.Println("juara dongg..")
	case (klasemen_liverpool < 4) && (klasemen_liverpool > 1):
		fmt.Println("masih masuk ucl")
	default:
		{
			fmt.Println("masuk uel")
			fmt.Println("lagi bapuk")
		}
	}


	// fallthrough melanjutkan ekseskusi ke case selanjutnya jika benar
	// hanya dieksekusi satu case untuk satu fallthrough

	var score3 = 9;

	switch{
	case score3 == 9 :
		fmt.Println("ini 9")
		fallthrough
	case score3 > 1:
		fmt.Println("lebih besar dari 1")
		fallthrough
	case score3 < 10:
		fmt.Println("lebih kecil dari 10")
	default:
		fmt.Println("bukan di rentang 1 - 10")
	}


	// campur if else switch case bersarang

	score4 := 10;

	if score4 > 10{
		switch {
		case score4 == 11:
			fmt.Printf("Ini 11")
		}
	}else {
		switch score4{
		case 10, 11, 12:
			fmt.Println("Ini nilai 2 digit")

		}
	}




}




