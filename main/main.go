package main

import (
  "fmt"
  "time"
  "log"
  "jashanpreet_singh_99/matrix"
  "jashanpreet_singh_99/matrix_go"
  )

func main() {
  startTime := time.Now()
  fmt.Println("Start Time : ", startTime)
  mat := matrix_go.GRandomMatrix(10,10, 10)
  mat.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
  mat_2 := matrix.RandomMatrix(10,9, 10)
  mat_2.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
  f_mat,err := mat.DivideScalar(0)
  if err != nil {
    log.Fatal(err)
  }
  f_mat.Print()
  fmt.Println(" Execution Time : ", time.Since(startTime))
}
