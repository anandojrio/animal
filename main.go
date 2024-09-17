package main

import "fmt"

func main() {

	fmt.Printf("Ko je najbolja riba na svetu")
	var ime string
	fmt.Scan(&ime)
	fmt.Println(ime)
	for i := 0; i < 30; i++ {
		if ime == "Jana" || ime == "jana" || ime == "vasina sestra" {
			fmt.Println("Tacan odgovor!!!")
			break
		}
		fmt.Println("Ne tacan odgovor, ona je blize nego sto mislis")
		fmt.Scan(&ime)

	}
}
