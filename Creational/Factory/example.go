package main

import "fmt"

func main() {
	mag1, _ := NewPublication("magazine", "Tyme", 50, "The Tymes")
	mag2, _ := NewPublication("magazine", "Lyfe", 40, "Lyfe Inc")
	news1, _ := NewPublication("newspaper", "The Herald", 60, "Heralders")
	news2, _ := NewPublication("newspaper", "The Standard", 30, "Standarders")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub iPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------------------\n")
}
