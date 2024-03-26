package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Menyimpan data-data yang diperlukan.
	var team01, team02 string
	team01, team02 = "Lumba-Lumba", "Koala"
	var scoreTeam01 = [3]int{96, 108, 89}
	var scoreTeam02 = [3]int{88, 91, 110}

	fmt.Printf("===== %s vs %s =====", team01, team02)

	fmt.Println()
	fmt.Println("[ SCORES ]")
	// Mencetak skor dengan format.
	printTeamAndScore(team01, scoreTeam01[:])
	printTeamAndScore(team02, scoreTeam02[:])

	// Menghitung rata-rata nilai tim.
	var avgTeam01 = countAverage(scoreTeam01[:])
	var avgTeam02 = countAverage(scoreTeam02[:])

	fmt.Println()
	fmt.Println("[ AVERAGE ]")
	// Mencetak rata-rata dengan format.
	printTeamAndAvg(team01, avgTeam01)
	printTeamAndAvg(team02, avgTeam02)

	fmt.Println()
	fmt.Println("[ RESULT ]")
	// Mencari pemenangnya.
	winner := getWinner(avgTeam01, avgTeam02)
	if winner == 0 {
		fmt.Println("Draw !!!")
	} else if winner == 1 {
		fmt.Printf("%s wins !!!", team01)
	} else {
		fmt.Printf("%s wins !!!", team02)
	}

}

func countAverage(data []int) int {
	temp := 0
	for i := 0; i < len(data); i++ {
		temp += data[i]
	}
	return temp / len(data)
}

func printTeamAndScore(team string, scores []int) {
	var result string = ""

	result += team
	result += " : "

	for i := 0; i < len(scores); i++ {
		result += strconv.Itoa(scores[i])
		result += " "
	}

	fmt.Println(result)
}

func printTeamAndAvg(team string, avg int) {
	fmt.Printf("%s : %d", team, avg)
	fmt.Println()
}

/*
	Fungsi ini mengembalikan 1 jika tim pertama yang menang,
	mengembalikan 2 jika tim kedua yang menang,
	dan mengembalikan 0 jika kedua tim seri.
*/
func getWinner(avg01 int, avg02 int) int {
	if avg01 == avg02 {
		return 0
	} else if avg01 > avg02 {
		return 1
	} else {
		return 2
	}
}
