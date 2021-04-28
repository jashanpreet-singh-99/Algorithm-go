package main

import (
  "fmt"
  "reflect"
  "jashanpreet_singh_99/matrix"
  )

func main() {
  mat := matrix.RandomMatrix(2,3)
  fmt.Println(reflect.TypeOf(mat))
  mat.Print()
  mat = mat.TransposeMatrix()
  mat.Print()
  mat_2 := mat.AddMatrix(mat)
  mat_2.Print()
  mat_2 = mat_2.AddScalar(10)
  mat_2.Print()
}
