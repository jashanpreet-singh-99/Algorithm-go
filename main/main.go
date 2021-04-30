package main

import (
  "fmt"
  "time"
  "jashanpreet_singh_99/matrix"
  )

func main() {
  startTime := time.Now()
  fmt.Println("Start Time : ", startTime)
  mat := matrix.RandomMatrix(10,10, 10)
  mat.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
  mat_1 := mat.InverseMatrix()
  mat_1.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
  f_mat := mat.MultiplyMatrix(mat_1)
  f_mat.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
}
