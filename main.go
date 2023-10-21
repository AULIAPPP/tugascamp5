package main

import (
	"fmt"
)

func main() {
	var (
		username string
		password string
		balance  float64
		income   float64
		expenses float64
	)

	// Loop utama menu
	for {
		fmt.Println("=== Manajemen Pengeluaran ===")
		fmt.Println("1. Login")
		fmt.Println("2. Info Saldo")
		fmt.Println("3. Info Riwayat")
		fmt.Println("4. Info Pemasukan")
		fmt.Println("5. Info Pengeluaran")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih menu (1-6): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&username)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&password)
			// Validasi login di sini (misalnya, sesuai dengan basis data)
			// Jika berhasil, berikan akses ke menu lain
			fmt.Println("Login berhasil!")
		case 2:
			fmt.Printf("Saldo Anda: %.2f\n", balance)
		case 3:
			// Tampilkan riwayat transaksi di sini
		case 4:
			fmt.Printf("Total pemasukan: %.2f\n", income)
		case 5:
			fmt.Printf("Total pengeluaran: %.2f\n", expenses)
		case 6:
			fmt.Println("Terima kasih! Sampai jumpa lagi.")
			return
		default:
			fmt.Println("Pilihan menu tidak valid. Silakan coba lagi.")
		}
	}
}
