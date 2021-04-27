package matrix

import (
  "time"
  "math/rand"
  "fmt"
)
var seeded bool = false

func RandomMatrix(r int, c int) [][]int {
  var temp [][]int
  for i := 0; i < r; i++ {
    temp_r := make([]int, c)
    for j := 0 ; j < c; j++ {
      temp_r[j] = rand.Intn(100)
    }
    temp = append(temp, temp_r)
  }
  return temp
}

func CallMatrix() string {
  message := "matrix called."
  return message
}

func Mprint(mat [][]int) {
  if seeded {
    r := len(mat)
    c := len(mat[0])
    for i := 0; i < r; i++ {
      fmt.Println(mat[i])
    }
    fmt.Printf("Shape (%d,%d)\n", r,c)
  } else {
    source := rand.Seed(time.Now().UnixNano())
    seeded = true
    Mprint(mat)
  }
}
