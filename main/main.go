package main

import (
  "fmt"
  "reflect"
  "jashanpreet_singh_99/matrix"
  )

func main() {
  mat := matrix.RandomMatrix(6,6, 10)
  fmt.Println(reflect.TypeOf(mat))
  mat.Print()
  fmt.Println(matrix.Determinant(mat))
}
