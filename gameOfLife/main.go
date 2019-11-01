package main

import (
	"fmt"
	"os"
	"strconv"
)

type point struct {
	x int
	y int
}
type area struct {
	a point
	b point
}

func main() {
	//  -ng NUM_GORUTINAS -r NUM_FILAS -c NUM_COLS -i GENERACIONES \path -m MET_PART -s SEMILLA
	args := os.Args

	ite := 1
	n := 1
	m := 1
	nGo := 1
	bloqueBool := false
	semilla := 1

	for i, arg := range args {
		switch arg {
		case "-ng":
			nGo, _ = strconv.Atoi(args[i+1])
		case "-r":
			n, _ = strconv.Atoi(args[i+1])
		case "-c":
			m, _ = strconv.Atoi(args[i+1])
		case "-i":
			ite, _ = strconv.Atoi(args[i+1])
		case "-m":
			if args[i+1] == "1" {
				bloqueBool = true
			}
		case "-s":
			semilla, _ = strconv.Atoi(args[i+1])
		}
	}

	mp := make([][]bool, m)
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]bool, n)
	}

	/* mp[1][2] = true
	mp[3][2] = true
	mp[2][1] = true
	mp[2][2] = true
	mp[2][3] = true */

}

func render(mp [][]bool) {
	for i := range mp {
		for _, j := range mp[i] {
			if j {
				print("■ ")
			} else {
				print("▫ ")
			}
		}
		print("\n")
	}
	print("\n")
}

//e.a.x 00    e.b.x 9      e.b.y 15
//se revisa un area del mapa buscando celulas
func muerte(mp [][]bool, e area) {
	//copia de la matris
	cmp := make([][]bool, len(mp))
	for i := range mp {
		cmp[i] = make([]bool, len(mp[i]))
		copy(cmp[i], mp[i])
	}

	for i := e.a.y; i < e.b.y; i++ {
		for j := e.a.x; j < e.b.x; j++ {
			mp[i][j] = moore(cmp, i, j)
		}
	}
}

//comprabamos cada lado de una celula
func moore(mp [][]bool, i, j int) bool {
	n := len(mp[0])
	m := len(mp)

	var con int
	if i != 0 && mp[i-1][j] { // 				↓
		con++
	}
	if i != 0 && j != 0 && mp[i-1][j-1] { // 	↙
		con++
	}
	if j != 0 && mp[i][j-1] { // 				←
		con++
	}
	if j != 0 && i != m && mp[i+1][j-1] { // 	↖
		con++
	}
	if i != m && mp[i+1][j] { // 				↑
		con++
	}
	if j != n && i != m && mp[i+1][j+1] { // 	↗
		con++
	}
	if j != n && mp[i][j+1] { //				→
		con++
	}
	if j != n && i != 0 && mp[i-1][j+1] { //	↘
		con++
	}
	// con CON cantidad que sucede...

	return reglas(mp[i][j], con)
}

func reglas(mp bool, con int) bool {
	if mp {
		switch {
		case con < 3:
			return false
		case con == 3 || con == 4:
			return true
		case con > 4:
			//fmt.Println(j+1, " ,", i+1, "muere")
			return false
		}
	} else {
		if con == 3 {
			return true
		}
		return false
	}
	fmt.Println("nunca saldre")
	return false
}
