package main

import (
  "fmt"
  "reflect"
  "jashanpreet_singh_99/matrix"
  )

func main() {
  mat := matrix.RandomMatrix(3,3)
  fmt.Println(reflect.TypeOf(mat))
  mat.Print()
  mat.Add(10)
  mat.Print()
  mat_2 := matrix.RandomMatrix(3,3)
  mat_2.print()
  mat.Add(mat_2)
  mat.print()
}
