// 2048 project main.go
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"

	//tm "github.com/buger/goterm"
	term "github.com/nsf/termbox-go"
)

type pole struct {
	liczba int
	zmiana bool
}

var plansza [4][4]pole

func wys_plansze() {

	for i := range plansza {
		fmt.Println("-----------------------------------------")
		fmt.Println("|         |         |         |         |")
		for j := range plansza[i] {
			var k float64 = 10
			temp := 0
			spacje := "          "

			for l := 1; k >= 1; l++ {
				k = float64(plansza[i][j].liczba) / math.Pow10(l)
				//fmt.Print(" ")
				//spacje.range = spacje.Length - 1
				temp++
			}
			parzysty := 0
			if temp%2 == 0 {
				parzysty = 1
			}
			stringtemp := spacje[:len(spacje)-temp-1]
			var stringtemp2 string = stringtemp[:len(stringtemp)/2]
			var stringtemp3 string = stringtemp[:(len(stringtemp)+parzysty)/2]
			if plansza[i][j].liczba != 0 {
				fmt.Print("|", stringtemp2, plansza[i][j].liczba, stringtemp3)
			} else {
				fmt.Print("|", stringtemp2, " ", stringtemp3)
			}
		}
		fmt.Print("|")
		fmt.Println()
		fmt.Println("|         |         |         |         |")
	}

	fmt.Println("-----------------------------------------")

}

func gra() {
	//wys_plansze()
	for {

		przesuniecie := false
		ev := term.PollEvent()
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
		switch ev.Key {
		case term.KeyArrowUp:
			for i := range plansza {
				if i == 0 {
					continue
				}
				for j := range plansza[i] {
					for k := 0; i-k >= 1; k++ {
						if plansza[i-k-1][j].liczba == 0 {
							plansza[i-k-1][j].liczba = plansza[i-k][j].liczba
							plansza[i-k][j].liczba = 0
							przesuniecie = true
						} else if plansza[i-k][j].liczba == plansza[i-k-1][j].liczba && !plansza[i-k][j].zmiana {
							plansza[i-k-1][j].liczba *= 2
							plansza[i-k][j].liczba = 0
							plansza[i-k-1][j].zmiana = true
							przesuniecie = true
						}
					}
				}
			}
		case term.KeyArrowDown:
			for i := 2; i >= 0; i-- {
				if i == 3 {
					continue
				}
				for j := range plansza[i] {
					for k := 0; i+k <= 2; k++ {
						if plansza[i+k+1][j].liczba == 0 {
							plansza[i+k+1][j].liczba = plansza[i+k][j].liczba
							plansza[i+k][j].liczba = 0
							przesuniecie = true
						} else if plansza[i+k][j].liczba == plansza[i+k+1][j].liczba && !plansza[i+k][j].zmiana {
							plansza[i+k+1][j].liczba *= 2
							plansza[i+k][j].liczba = 0
							plansza[i+k+1][j].zmiana = true
							przesuniecie = true
						}
					}
				}
			}
		case term.KeyArrowLeft:
			for i := range plansza {
				for j := range plansza[i] {
					if j == 0 {
						continue
					}
					for k := 0; j-k >= 1; k++ {
						if plansza[i][j-k-1].liczba == 0 {
							plansza[i][j-k-1].liczba = plansza[i][j-k].liczba
							plansza[i][j-k].liczba = 0
							przesuniecie = true
						} else if plansza[i][j-k].liczba == plansza[i][j-k-1].liczba && !plansza[i][j-k].zmiana {
							plansza[i][j-k-1].liczba *= 2
							plansza[i][j-k].liczba = 0
							plansza[i][j-k-1].zmiana = true
							przesuniecie = true
						}
					}
				}
			}

		case term.KeyArrowRight:
			for i := 3; i >= 0; i-- {
				for j := range plansza[i] {
					if j == 3 {
						continue
					}
					for k := 0; j+k <= 2; k++ {
						if plansza[i][j+k+1].liczba == 0 {
							plansza[i][j+k+1].liczba = plansza[i][j+k].liczba
							plansza[i][j+k].liczba = 0
							przesuniecie = true
						} else if plansza[i][j+k].liczba == plansza[i][j+k+1].liczba && !plansza[i][j+k].zmiana {
							plansza[i][j+k+1].liczba *= 2
							plansza[i][j+k].liczba = 0
							plansza[i][j+k+1].zmiana = true
							przesuniecie = true
						}
					}
				}
			}
		}
		zeruj_zmiany()
		if przesuniecie {
			dodaj_liczby()
		}
		wys_plansze()
		przegrana()
	}
}
func dodaj_liczby() {

	czworka := rand.Intn(10)
	var liczba int
	if czworka == 0 {
		liczba = 4
	} else {
		liczba = 2
	}

	for {
		tempx := rand.Intn(4)
		tempy := rand.Intn(4)
		if plansza[tempx][tempy].liczba == 0 {
			plansza[tempx][tempy].liczba = liczba
			return
		}
	}

}

func zeruj_zmiany() {
	for i := range plansza {
		for j := range plansza[i] {
			plansza[i][j].zmiana = false
		}

	}
}

func przegrana() {

	for i := range plansza {
		for j := range plansza[i] {
			if plansza[i][j].liczba == 0 {
				return
			}
		}
	}

	for i := range plansza {
		for j := range plansza[i] {
			if i > 0 {
				if plansza[i][j].liczba == plansza[i-1][j].liczba {
					return
				}
			}
			if i < 3 {
				if plansza[i][j].liczba == plansza[i+1][j].liczba {
					return
				}
			}
			if j > 0 {
				if plansza[i][j].liczba == plansza[i][j-1].liczba {
					return
				}
			}
			if j < 3 {
				if plansza[i][j].liczba == plansza[i][j+1].liczba {
					return
				}

			}
		}

	}
	fmt.Println("Nie mozna wykonac zadnego ruchu -koniec gry")
	os.Exit(3)
}
func main() {

	dodaj_liczby()
	dodaj_liczby()
	zeruj_zmiany()
	term.Init()

	gra()

}
