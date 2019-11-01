package main

import "fmt"

type point struct {
	x int
	y int
}
type area struct {
	a point
	b point
}

func main() {

	n := 5
	m := 5
	mp := make([][]bool, m)
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]bool, n)
	}

	mp[1][2] = true
	mp[3][2] = true
	mp[2][1] = true
	mp[2][2] = true
	mp[2][3] = true
	render(mp)
	muerte(mp, area{a: point{x: 0, y: 0}, b: point{x: n - 1, y: m - 1}})
	render(mp)
	muerte(mp, area{a: point{x: 0, y: 0}, b: point{x: n - 1, y: m - 1}})
	render(mp)
	muerte(mp, area{a: point{x: 0, y: 0}, b: point{x: n - 1, y: m - 1}})
	render(mp)
	muerte(mp, area{a: point{x: 0, y: 0}, b: point{x: n - 1, y: m - 1}})
	render(mp)

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
	fmt.Println("")
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

	if mp[i][j] {
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
