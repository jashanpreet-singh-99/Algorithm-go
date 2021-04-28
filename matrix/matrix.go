package matrix

import (
  "time"
  "math/rand"
  "fmt"
)

var seeded bool // Which if seed is initialized
var rand_1 *rand.Rand // the rand object linked to seed

/*
Objective := generate a random matrix of given size
parameters :-
          r : number of rows
          c : number of columns
return := [][]int
*/
func RandomMatrix(r int, c int) [][]int {
  if seeded == false {
    x1 := rand.NewSource(time.Now().UnixNano())
    rand_1 = rand.New(x1)
    seeded = true
  }
  var temp [][]int
  for i := 0; i < r; i++ {
    temp_r := make([]int, c)
    for j := 0 ; j < c; j++ {
      temp_r[j] = rand_1.Intn(100)
    }
    temp = append(temp, temp_r)
  }
  return temp
}

/*
Objective := print the matrix in readable format
parameters :-
          mat : [][]int the mat to print
*/
func Mprint(mat [][]int) {
  r := len(mat)
  c := len(mat[0])
  for i := 0; i < r; i++ {
    fmt.Println(mat[i])
  }
  fmt.Printf("Shape (%d,%d)\n", r,c)
}
