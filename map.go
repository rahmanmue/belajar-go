package main
import "fmt"

func main(){
	// cara penulisan map, jika tanpa {} maka dinisialisasi nil
	var chicken = map[string]int{}
	chicken["januari"] = 50;
	chicken["februari"] = 100;

	fmt.Println("chicken bulan 1 =", chicken["januari"])

	// inisialisasi langsung
	var chicken1 = map[string]int{"maret":50, "agustus":100}
	chicken2 := map[string]int{
		"desember": 100,
		"november": 40,
	}

	fmt.Println(chicken1["agustus"])
	fmt.Println(chicken2["desember"])

	// penulisan map bisa dengan make dan new 
	// khusus new hasil merupakan data pointer untuk mengambil nilai aslinya dengan *
	var chicken3 = make(map[string]int)
	chicken3["juli"] = 33
	fmt.Println(chicken3)
	var chicken4 = *new(map[string]int)
	fmt.Println(chicken4)

	// menampilkan dengan for
	chicken5 := map[string]int{
		"jan": 10,
		"feb": 20,
		"maret": 30,
		"apr": 40,
		"mei": 44,
		"juni": 50,
	}

	for k, v := range chicken5{
		fmt.Println(k, ":", v)
	}

	// menghapus item map
	fmt.Println(len(chicken5))
	delete(chicken5, "juni")
	fmt.Println(len(chicken5))

	// cek keberadaan key dan value
	var value, isExist = chicken5["juni"]

	if isExist{
		fmt.Println(value)
	}else{
		fmt.Println("Kosong")
	}

	// kombinasi slice dan map seperti obj dalam arr
	chicken6 := []map[string]int{
		map[string]int{"lahir":22, "meninggal":44},
		map[string]int{"lahir": 33, "meninggal": 100},
	}

	for _, chicken := range chicken6{
		fmt.Println(chicken["lahir"], chicken["meninggal"])
	}

	// penulisan dipersingkat
	chicken7 := []map[string]string{
		{"nama":"ayam putih", "gender":"jantan"},
		{"nama":"ayam kuning", "gender":"betina"},
	}

	for _, chicken := range chicken7{
		fmt.Println(chicken["nama"], chicken["gender"])
	}


}