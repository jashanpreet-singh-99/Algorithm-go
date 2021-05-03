package matrix

import (
  "time"
  "math/rand"
  "fmt"
)

/*
Matrix object
parameters :-
        Value : contain all the values of each element
        Shape : store info about number of olumns and rows
*/
type Mat struct{
  Value [][]float32
  Shape []int
}

type MatrixError struct {
  Message string
}

func (err *MatrixError) Error() string {
  return err.Message
}

var seeded bool // Which if seed is initialized
var rand_1 *rand.Rand // the rand object linked to seed

/*
Objective := generate a random matrix of given size

parameters :-
          r     : number of rows
          c     : number of columns
          limit : max value
return :- Mat Objective
*/
func GRandomMatrix(r int, c int, limit int) Mat {
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
Objective := generate a Identity matrix of given size

parameters :-
          s     : number of columns or rows
return :- Mat Objective
*/
func GIdentityMatrix(s int) Mat {
  var temp [][]float32
  for i := 0; i < s; i++ {
    temp_r := make([]float32, s)
    temp_r[i] = 1
    temp = append(temp, temp_r)
  }
  return Mat{temp, []int{s,s}}
}

/*
Objective := print the matrix in readable format

parameters :-
       mat : the matrix to print
*/
func (mat Mat) GPrint() {
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
Objective := add scaler value from matrix.

parameters :-
        s  : the scalar value to add.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GAddScalar(s float32) Mat {
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

parameters :-
        mat : the matrix to add.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GAddMatrix(mat_2 Mat) (Mat,error) {
  if len(mat.Shape) != len(mat_2.Shape) {
    return mat,&MatrixError{"Miss-matched dimmensions."}
  } else {
    for dim := 0 ; dim < len(mat.Shape); dim++ {
      if mat.Shape[dim] != mat_2.Shape[dim] {
        return mat,&MatrixError{("Miss-matched shape.")}
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
  return Mat{temp, mat.Shape},nil
}

/*
Objective := subtract scaler value from matrix

parameters :-
        s   : the scalar to subtract.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GSubtractScalar(s float32) Mat {
  return mat.GAddScalar(-1 * s)
}

/*
Objective := subtract matrix from matrix

parameters :-
        mat : the matrix to subtract.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GSubtractMatrix(mat_2 Mat) (Mat,error) {
  if len(mat.Shape) != len(mat_2.Shape) {
    return mat,&MatrixError{"Miss-matched dimmensions."}
  } else {
    for dim := 0 ; dim < len(mat.Shape); dim++ {
      if mat.Shape[dim] != mat_2.Shape[dim] {
        return mat,&MatrixError{("Miss-matched shape.")}
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
  return Mat{temp, mat.Shape},nil
}

/*
Objective := return the transpose of the given matrix
*/
func (mat Mat) GTransposeMatrix() Mat {
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

parameters :-
        s : the scalar to multiply.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GMultiplyScalar(s float32) Mat {
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
Objective := Multiply scaler value with matrix

parameters :-
        mat : the mat to multiply with.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GMultiplyMatrix(mat_2 Mat) (Mat,error) {
  if mat.Shape[1] != mat_2.Shape[0] {
    return mat,&MatrixError{("Miss-matched shape.")}
  }
  Shape := []int{ mat.Shape[1], mat_2.Shape[0]}
  var temp [][]float32
  for i := 0 ; i < Shape[0]; i++ {
    temp_r := make([]float32, Shape[1])
    for j := 0 ; j < Shape[1]; j++ {
      var p_val float32 = 0
      for k := 0; k < Shape[0]; k++ {
        p_val += mat.Value[i][k] * mat_2.Value[k][j]
      }
      temp_r[j] = p_val
    }
    temp = append(temp, temp_r)
  }
  return Mat{temp, Shape},nil
}

/*
Objective := Divide scaler value with matrix

parameters :-
        s : the scalar to divide.
return :-
        mat : the modified matrix.
*/
func (mat Mat) GDivideScalar(s float32) (Mat,error) {
  if s == 0 {
    return mat,&MatrixError{"Divede By zero error."}
  }
  return mat.GMultiplyScalar(1/s),nil
}

/*
Objective := Find inverse of a matrix if possible
*/
func (mat Mat) GInverseMatrix() (Mat,error) {
  if mat.Shape[0] != mat.Shape[1] {
    return mat,&MatrixError{"Not a Square matrix."}
  }
  adj_mat,determinant :=  GAdjointMatrix(mat)
  if determinant == 0 {
    return mat,&MatrixError{"The determinant of the matrix is zero."}
  }
  adj_mat = adj_mat.GMultiplyScalar(1/determinant)
  return adj_mat,nil
}

/*
Objective := Find the determinent of the Matrix
normally split the matrix until its 2 X 2 and then recursively call back.
parameters :-
       mat : matrix on which to operate.
return :-
      deter : dterminant of the matrix float32 value.
*/
func GDeterminant(mat Mat) float32 {
  if mat.Shape[0] == 2 && mat.Shape[1] == 2 {
    return mat.Value[0][0]*mat.Value[1][1] - mat.Value[0][1]*mat.Value[1][0]
  }
  var deter float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    a,temp_mat := GSplitMatrix(mat, 0, i)
    if (i % 2) != 0 {
      a = -a
    }
    deter += a * GDeterminant(temp_mat)
  }
  return deter
}

/*
Objective := Find the Adjoint Matrix

parameters :-
       mat : matrix on which to operate.
return :-
      mat   : adjoint matrix
      deter : dterminant of the matrix float32 value.
*/
func GAdjointMatrix(mat Mat) (Mat,float32) {
  var deter_mat float32 = 0
  var adj_mat [][]float32
  for i := 0 ; i < mat.Shape[0]; i++ {
    adj_mat_r := make([]float32, mat.Shape[1])
    for j:= 0; j < mat.Shape[1]; j ++ {
      a,temp_mat := GSplitMatrix(mat, i, j)
      deter := GDeterminant(temp_mat)
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
  return semi_adj_mat.GTransposeMatrix(), deter_mat
}

/*
Objective := Split the matrix into sub matrix based on the placement of the cofficient.

parameters :-
       mat : matrix on which to operate.
       row : the row placement of the cofficient.
       col : the column placement of the cofficient.
return :-
      a   : the cofficient of the caller matrix
      mat : the sub matrix, after removing the row and column of the cofficient.
*/
func GSplitMatrix(mat Mat, row int,col int) (float32, Mat) {
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
