package main
import "fmt"

func main(){
	// tipe data numerik, string, boolean

	//1. numerik non desimal
	// uint bilangan cacah only positif 2^ -> 8,16,32,64 
	// int bilangan bulat negatif dan positif 2^ -> 8,16,32,64 dibagi 2 negatif to positif
	// byte sama dengan uint8

	var positiveNumber uint8 = 89;

	// secara otomatis negativeNumber tipe data nya int32
	var negativeNumber = -1234566;

	fmt.Println(positiveNumber, negativeNumber);


	// 2. numerik desimal
	// float32, dan float64
	var decimalNumber = 2.45;

	// .2f untuk menentukan digit dibelakang titik
	fmt.Printf("decimal number %f \n", decimalNumber);
	fmt.Printf("decimal number %.2f \n", decimalNumber);

	// 3. boolean
	var trueVal bool = true;
	falseVal := false;
	fmt.Println(trueVal, falseVal);
	fmt.Printf("data %t \n", trueVal);

	// 4. String kekuatan `` karakter \n tidak diescape
	message := 
	`Halo nama saya "Muhammad Rahman".
saat ini saya sedang belajar golang,
	di enigmacamp`;

	fmt.Println(message);


	// 5. nilai nil dan zero value
	// nil berbeda dengan zero value
	// "" merupakan zero value string
	// 0 zero value int
	// 0.0 zero value numerik desimal
	// false zero value boolean

	// nil benar2 kosong atau nilai kosong dipakai untuk
	// pointer, tipe data func, slice, map, channel, any (interface kosong)

}