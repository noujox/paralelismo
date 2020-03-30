package main

import (
	"fmt"
	"math/rand"
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
	n := 1 //x
	m := 1 //y
	nGo := 1
	bloqueBool := false
	semilla := 1

	//se obtienen los datos de parametros
	for i, arg := range args {
		switch arg {
		case "-ng":
			nGo, _ = strconv.Atoi(args[i+1])
		case "-c":
			n, _ = strconv.Atoi(args[i+1])
		case "-r":
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

	//creado de mapa
	mp := make([][]bool, m)
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]bool, n)
	}
	//fmt.Println(semilla)
	populate(mp, semilla, area{a: point{x: 0, y: 0}, b: point{x: n, y: m}})

	e := calculateArea(bloqueBool, nGo, n, m)
	for i := 0; i < len(e); i++ {
		println("area: (", e[i].a.x, ",", e[i].a.y, ") (", e[i].b.x, ",", e[i].b.y, ")")
	}

	_ = ite
	/*
		for i := 0; i < ite; i++ {
			render(mp)
			muerte(mp, e[0])
		}
		render(mp)
	*/
}

func calculateArea(bloqueBool bool, chunks int, n, m int) []area {
	var e []area
	if bloqueBool {
		// cuadrado
	} else { //columna

		blocks := n / chunks
		rest := n % chunks
		if rest == 0 {
			fmt.Println("hola", blocks)
			for i := 0; i < chunks; i++ {
				e = append(e, area{
					a: point{x: (i * blocks), y: 0},
					b: point{x: ((i+1)*blocks - 1), y: m}})

			}
		} else {

			for i := 0; i < chunks; i++ {
				if i < rest {
					e = append(e, area{
						a: point{x: (i * (blocks + 1)), y: 0},
						b: point{x: ((i+1)*blocks + i), y: m}})
				} else {
					e = append(e, area{
						a: point{x: ((i + 1) * blocks), y: 0},
						b: point{x: ((i+1)*blocks + 1), y: m}})
				}
			}
		}
	}
	return e
}

func populate(mp [][]bool, sem int, e area) {
	s := rand.NewSource(42)
	r := rand.New(s)
	for i := 0; i < sem; i++ {
		x := (r.Intn(e.b.x-e.a.x) + e.a.x)
		y := (r.Intn(e.b.y-e.a.y) + e.a.y)
		mp[y][x] = true // falta comprobar si ya habia algo
	}
}

func render(mp [][]bool) {
	for i := range mp {
		for j, v := range mp[i] {
			if i == 4 && j == 3 {
				print("x ")
			} else {
				if v {
					print("■ ")
				} else {
					print("□ ")
				}
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
	n := len(mp[0]) - 1
	m := len(mp) - 1

	con := 0
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
	if i == 4 && j == 3 {
		fmt.Println(con)
	}

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
