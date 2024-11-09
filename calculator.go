package main

import (
	"fmt"
)

var number, num1, num2 int

func run(numbers ...int) int {
	switch numbers[0] {
	case 1:
		return numbers[1] * numbers[2]
	case 2:
		if numbers[1] != 0 || numbers[2] != 0 {
			return numbers[1] / numbers[2]
		}
		fmt.Println("Error: Pembagian dengan nol")
		return 0
	case 3:
		return numbers[1] + numbers[2]
	case 4:
		return numbers[1] - numbers[2]
	default:
		return numbers[1] % numbers[2]
	}
}
func main() {
	fmt.Print("PROGRAM KALKULATOR SEDERHANA DI GOLANG\n1. OPERASI PERKALIAN\n2. OPERASI PEMBAGIAN\n3. OPERASI PENAMBAHAN\n4. OPERASI PENGURANGAN\n5. OPERASI MODULUS\n6. EXIT\nPilih salah satu angka di atas: ")

	fmt.Scan(&number)
	if number > 0 && number <= 6 {
		if number == 6 {
			fmt.Println("END")
			return
		}
		fmt.Print("Masukan Angka Pertama: ")
		fmt.Scanf("%d", &num1)
		fmt.Print("Masukan Angka Kedua: ")
		fmt.Scanf("%d", &num2)
		result := run(number, num1, num2)
		fmt.Println(result)

	} else {
		fmt.Print("Input tidak valid, silakan masukkan angka antara 1 hingga 6.")
	}
}
