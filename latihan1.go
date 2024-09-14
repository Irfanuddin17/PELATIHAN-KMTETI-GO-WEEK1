package main

import "fmt"

func main() {
	for {
		fmt.Println("Masukan nilai celcius dengan bilangan desimal, contoh :15.00")
		fmt.Println("opsi 1 (farenheit) , opsi 2 (kelvin) , opsi 3 (Reamur) , opsi 4 (Berhenti)")
		fmt.Println("============================================")

		var c float32
		fmt.Print("celcius=")
		fmt.Scan(&c)

		var counter int
		fmt.Print("Pilih Opsi=")
		fmt.Scan(&counter)

		f := (9 / 5 * c) + 32.00
		k := c + 273.15
		r := 4.00 / 5.00 * c

		switch counter {
		case 1:
			fmt.Println("Farenheit=", f)
			fmt.Println("============================================")
		case 2:
			fmt.Println("Kelvin=", k)
			fmt.Println("============================================")
		case 3:
			fmt.Println("Reamur=", r)
			fmt.Println("============================================")
		}
		if counter == 4 {
			fmt.Println("Terima Kasih")
			break
		}
	}
}
