package main
import "fmt"

func main(){
	const firstName string = "Muhammad";
	const lastName = "Rahman";

	// jika Println pke koma otomatis ada spasi di sela nilai
	fmt.Println(firstName, lastName);

	// Print tidak
	fmt.Print(firstName, " ", lastName, "\n");

	// multi konstanta
	const (
		square = "lingkaran"
		pi float32 = 3.14
		isRounded bool = true
		radius = 5
	)

	fmt.Println(square, pi, isRounded, radius)

	// b jika kosong akan merujuk ke piA
	const (
		piA = 3.14
		b
	)

	fmt.Println(piA, b)

	const(
		today string = "senin"
		sekarang
		isToday = true
	)

	fmt.Println(today, sekarang, isToday);

	// multiple konstanta
	const sebelas, duabelas = 11, 12
	const senin, selasa string = "senin", "selasa";
	
	fmt.Println(sebelas, duabelas);
	fmt.Println(senin, selasa);

}