package matrix

import (
  "time"
  "math/rand"
  "fmt"
)

type Scalar_Vector interface {
  Add()
}

type Mat struct{
  Value [][]int
  Shape []int
}

var seeded bool // Which if seed is initialized
var rand_1 *rand.Rand // the rand object linked to seed

/*
Objective := generate a random matrix of given size
parameters :-
          r : number of rows
          c : number of columns
return := Mat Objective
*/
func RandomMatrix(r int, c int) Mat {
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
  return Mat{temp, []int{r,c}}
}

/*
Objective := print the matrix in readable format
*/
func (mat Mat) Print() {
  r := mat.Shape[0]
  c := mat.Shape[1]
  for i := 0; i < r; i++ {
    fmt.Printf("[ ")
    for j := 0; j < c; j++ {
      fmt.Printf("%02d ", mat.Value[i][j])
    }
    fmt.Printf("]\n")
  }
  fmt.Printf("Shape (%d,%d)\n", r,c)
}

/*
Objective := add scaler value to matrix
*/
func (mat Mat) Add(s int) {
  for i := 0 ; i < mat.Shape[0]; i++ {
    for j := 0 ; j < mat.Shape[1]; j++ {
      mat.Value[i][j] = mat.Value[i][j] + s
    }
  }
}

/*
Objective := add scaler value to matrix
*/
func (mat Mat) Add(mat_2 Mat) {
  if len(mat.Shape) != len(mat_2.Shape) {
    fmt.Println("Miss-matched dimmensions.")
    return
  } else {
    for dim := 0 ; dim < len(mat.Shape); dim++ {
      if mat.Shape[dim] != mat.Shape[dim] {
        fmt.Println("Miss-matched shape at : ", dim)
        return
      }
    }
  }
  for i := 0 ; i < mat.Shape[0]; i++ {
    for j := 0 ; j < mat.Shape[1]; j++ {
      mat.Value[i][j] = mat.Value[i][j] + mat_2.Value[i][j]
    }
  }
}
