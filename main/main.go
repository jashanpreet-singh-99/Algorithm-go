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
  mat_1 := matrix.IdentityMatrix(3,)
  mat_1.Print()
  f_mat := mat.MultiplyMatrix(mat_1)
  f_mat.Print()
}
