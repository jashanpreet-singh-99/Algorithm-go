package main

import (
  "fmt"
  "reflect"
  "jashanpreet_singh_99/matrix"
  )

func main() {
  mat := matrix.RandomMatrix(3,3, 10)
  fmt.Println(reflect.TypeOf(mat))
  mat.Print()
  mat = mat.InverseMatrix()
  mat.Print()
}
