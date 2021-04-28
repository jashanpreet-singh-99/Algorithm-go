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
  Value [][]float32
  Shape []int
}

var seeded bool // Which if seed is initialized
var rand_1 *rand.Rand // the rand object linked to seed

/*
Objective := generate a random matrix of given size
parameters :-
          r     : number of rows
          c     : number of columns
          limit : max value
return := Mat Objective
*/
func RandomMatrix(r int, c int, limit int) Mat {
  if seeded == false {
    x1 := rand.NewSource(time.Now().UnixNano())
    rand_1 = rand.New(x1)
    seeded = true
  }
  var temp [][]float32
  for i := 0; i < r; i++ {
    temp_r := make([]float32, c)
    for j := 0 ; j < c; j++ {
      temp_r[j] = float32(rand_1.Intn(limit))
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
      fmt.Printf("%4.2f ", mat.Value[i][j])
    }
    fmt.Printf("]\n")
  }
  fmt.Printf("Shape (%d,%d)\n", r,c)
}

/*
Objective := add scaler value from matrix
*/
func (mat Mat) AddScalar(s float32) Mat {
  var temp [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      temp_r[j] = mat.Value[i][j] + s
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, mat.Shape}
}

/*
Objective := add matrix from matrix
*/
func (mat Mat) AddMatrix(mat_2 Mat) Mat {
  if len(mat.Shape) != len(mat_2.Shape) {
    fmt.Println("Miss-matched dimmensions.")
    return mat
  } else {
    for dim := 0 ; dim < len(mat.Shape); dim++ {
      if mat.Shape[dim] != mat.Shape[dim] {
        fmt.Println("Miss-matched shape at : ", dim)
        return mat
      }
    }
  }
  var temp [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      temp_r[j] = mat.Value[i][j] + mat_2.Value[i][j]
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, mat.Shape}
}

/*
Objective := subtract scaler value from matrix
*/
func (mat Mat) SubtractScalar(s float32) Mat {
  var temp [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      temp_r[j] = mat.Value[i][j] - s
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, mat.Shape}
}

/*
Objective := subtract matrix from matrix
*/
func (mat Mat) SubtractMatrix(mat_2 Mat) Mat {
  if len(mat.Shape) != len(mat_2.Shape) {
    fmt.Println("Miss-matched dimmensions.")
    return mat
  } else {
    for dim := 0 ; dim < len(mat.Shape); dim++ {
      if mat.Shape[dim] != mat.Shape[dim] {
        fmt.Println("Miss-matched shape at : ", dim)
        return mat
      }
    }
  }
  var temp [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      temp_r[j] = mat.Value[i][j] - mat_2.Value[i][j]
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, mat.Shape}
}

func (mat Mat) TransposeMatrix() Mat {
  r := mat.Shape[0]
  c := mat.Shape[1]
  var temp [][]float32
  for i := 0; i < c; i++ {
    temp_r := make([]float32, r)
    for j := 0 ; j < r; j++ {
        temp_r[j] = mat.Value[j][i]
    }
    temp = append(temp, temp_r)
  }
  return  Mat{temp,[]int{mat.Shape[1], mat.Shape[0]}}
}

/*
Objective := Multiply scaler value with matrix
*/
func (mat Mat) MultiplyScalar(s float32) Mat {
  var temp [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      temp_r[j] = mat.Value[i][j] * s
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, mat.Shape}
}

/*
Objective := Find inverse of a matrix if possible
*/
func (mat Mat) InverseMatrix() Mat {
  if mat.Shape[0] != mat.Shape[1] {
    fmt.Println("Not a Square Matrix. Skipping.")
    return mat
  }
  adj_mat,determinant :=  AdjointMatrix(mat)
  if determinant == 0 {
    fmt.Println("Determinant of the Matrix is 0. Skipping.")
    return mat
  }
  adj_mat = adj_mat.MultiplyScalar(1/determinant)
  return adj_mat
}

func Determinant(mat Mat) float32 {
  if mat.Shape[0] == 2 && mat.Shape[1] == 2 {
    return mat.Value[0][0]*mat.Value[1][1] - mat.Value[0][1]*mat.Value[1][0]
  }
  var deter float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    a,temp_mat := SplitMatrix(mat, 0, i)
    if (i % 2) != 0 {
      a = -a
    }
    deter += a * Determinant(temp_mat)
  }
  return deter
}

func AdjointMatrix(mat Mat) (Mat,float32) {
  var deter_mat float32 = 0
  var adj_mat [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    adj_mat_r := make([]float32, mat.Shape[1])
    for j:= 0; j < mat.Shape[1]; j ++ {
      a,temp_mat := SplitMatrix(mat, i, j)
      deter := Determinant(temp_mat)
      var r_m float32 = 1
      if (i % 2) != 0 {
        r_m = -1
      }
      var c_m float32 = 1
      if (j % 2) != 0 {
        c_m = -1
      }
      adj_mat_r[j] = r_m * c_m * deter
      if i == 0 {
        deter_mat += r_m * c_m * a * deter
      }
    }
    adj_mat = append(adj_mat, adj_mat_r)
  }
  semi_adj_mat := Mat{adj_mat, mat.Shape}
  return semi_adj_mat.TransposeMatrix(), deter_mat
}


func SplitMatrix(mat Mat, row int,col int) (float32, Mat) {
  var temp [][]float32
  for i := 0; i < mat.Shape[0]; i++ {
    if i == row {
      continue
    }
    count_j := 0
    temp_r := make([]float32, mat.Shape[1])
    for j := 0 ; j < mat.Shape[1]; j++ {
      if j == col {
        continue
      }
      temp_r[count_j] = mat.Value[i][j]
      count_j += 1
    }
    temp = append(temp, temp_r)
  }
  return mat.Value[row][col],Mat{temp, []int{mat.Shape[0]-1, mat.Shape[1] -1}}
}
