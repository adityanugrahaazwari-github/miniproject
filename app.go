package main

import (
	"fmt"
)

type vehicle struct {
	idVehicle string
	arrival   string
	departure string
	duration  string
	cost      float32
}

func main() {
	// start := time.Now()

	var operation int = 3
	for operation > 2 || operation < 0 {
		fmt.Print("\033[H\033[2J")
		showMenu()
		fmt.Scan(&operation)

		if operation == 1 {
			postVehicle()

		} else if operation == 2 {
			fmt.Print("Pengaturan: ")
		} else {
			fmt.Print("Tidak ada silakan pilih menu")
		}
	}
}

func showMenu() {
	fmt.Println("1. Input Data")
	fmt.Println("2. Pengaturan")
	fmt.Println("0. Keluar")
	fmt.Print("Silakan pilih menu: ")
}

func postVehicle() {
	var vehicle vehicle
	fmt.Print("No. Plat Kendaraan: ")
	fmt.Scan(&vehicle.idVehicle)

	fmt.Println(vehicle.idVehicle)
}

// func countCost() {
// 	duration := time.Since(start)
// 	fmt.Println("The call took ", duration.Minutes())
// 	fmt.Println(1000 * duration.Minutes())
// }
