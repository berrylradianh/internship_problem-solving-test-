package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"internship-gits/function"
)

func main() {
	fmt.Println("Pilih fungsi:")
	fmt.Println("1. A000124 of Sloanes OEIS")
	fmt.Println("2. Dense Ranking")
	fmt.Println("3. Balanced Bracket")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan pilihan: ")
	menu, _ := reader.ReadString('\n')
	menu = strings.TrimSpace(menu)

	switch menu {
	case "1":
		fmt.Print("Masukkan Jumlah output : ")
		qty, _ := reader.ReadString('\n')
		qty = strings.TrimSpace(qty)
		n, _ := strconv.Atoi(qty)
		fmt.Println(function.Soloanes(n))
	case "2":
		scores := []int{100, 100, 50, 40, 40, 20, 10}
		gitsScores := []int{5, 25, 50, 120}
		result := function.DenseRanking(scores, gitsScores)
		fmt.Println(result)
	case "3":

	default:
		fmt.Println("Input yang anda masukkan salah!")
	}
}
