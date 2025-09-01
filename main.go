package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	sum := 0
	for i := 2; i < len(os.Args); i++ {
		n, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Printf("Lewati argumen bukan angka: %q (error: %v)\n", os.Args[i], err)
			continue
		}
		sum += n
	}

	fmt.Printf("Hello, %s!\n", name)
	if len(os.Args) > 2 {
		fmt.Printf("Jumlah angka: %d\n", sum)
	} else {
		fmt.Println("Tip: kamu bisa tambah angka setelah nama untuk dijumlahkan.")
		fmt.Println("Contoh: go run main.go Alice 3 5 7")
	}
}
