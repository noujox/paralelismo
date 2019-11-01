package main

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
	m := 3
	mp := make([][]bool, m)
	for i := 0; i < len(mp); i++ {
		mp[i] = make([]bool, n)
	}

	//muerte(mp, area{a: point{x: 0, y: 0}, b: point{x: n - 1, y: 3}})

}

//e.a.x 00    e.b.x 9      e.b.y 15
//se revisa un area del mapa buscando celulas
func muerte(mp [][]bool, e area) {
	for i := e.a.y; i < e.b.y; i++ {
		for j := e.a.x; j < e.b.x; j++ {
			if mp[i][j] {

			}
		}
	}
}

//comprabamos cada lado de una celula
func moore(mp [][]bool, i, j int) {
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
			mp[i][j] = false
		case con == 3 || con == 4:
			break
		case con > 4:
			mp[i][j] = false
		}
	} else {
		if con == 3 {
			mp[i][j] = true
		}
	}
}
