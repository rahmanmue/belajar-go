package main
import "fmt"

func main(){

	// ini slice karena [] bukan [...][2]

	var club_epl = []string{
		"liverpool", 
		"mu", 
		"man city", 
		"arsenal", 
		"chelsea",
	}

	fmt.Println(club_epl)


	var new_slice1 = club_epl[:3]
	var new_slice2 = club_epl[2:]
	var new_slice3 = club_epl[:]

	// slice tipe data reference
	club_epl[1] = "aston villa"
	
	fmt.Println(new_slice1)
	fmt.Println(new_slice2)
	fmt.Println(new_slice3)

	// cap merupakan kapasitas maximum
	// jika mulai dari index 0 maka capasitas sama dengan len nya
	// jika bukan mulai dari 0 maka sesuai sisanya
	fmt.Println(len(club_epl), cap(new_slice2), cap(new_slice1))

	// append
	var a_club = []string{"nu", "west ham", "mu"}
	var add_to_a_club = append(a_club, "brighton")
	fmt.Println(add_to_a_club)
	
	// jika len dan cap nya sama jika di append menghasilkan elemen baru dan referensi baru
	// sedangkan jika len < cap maka jika di append elemen baru maka akan merubah array lama nya juga
	// karena mereferensi ke yang lama

	fmt.Println(cap(a_club))
	fmt.Println(len(a_club))

	b_club := a_club[0:2]
	fmt.Println("len b_club :", len(b_club), "cap b_club :", cap(b_club)) 
	c_club := append(b_club, "ini pasti mereplace array utama")
	fmt.Println(a_club, c_club)


	// copy yang di copy akan sebanyak jumlah dst atau len dst nya
	dst := make([]string, 3)
	fmt.Println(dst)
	src := []string{"banana", "apple", "orange", "watermelon" }
	// copy src ke dct, n hanya mengembalikan nilai yang dicopynya
	n := copy(dst, src)
	fmt.Println(n)
	fmt.Println(dst)

	// jika ada isi nya otomatis mereplace tpi yang di copy sesuai len dst1
	dst1 := []string{"potato", "potato", "potato"}
	src1 := []string{"orange", "melon", "apple", "pear"}
	n1 := copy(dst1, src1)
	fmt.Println(n1)
	fmt.Println(dst1)

	// akses elemen dengan 3 index
	// [0:1:1] index terakhir menentukan cap nya
	products := []string{"infinix", "samsung", "lenovo"}
	productsA := products[0:2]
	productsB := products[0:2:2]

	fmt.Println(products, len(products), "cap :",cap(products))
	fmt.Println(productsA, len(productsA), "cap :",cap(productsA))
	fmt.Println(productsB, len(productsB), "cap :",cap(productsB))

}