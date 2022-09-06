/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {
     var rdiag,ldiag int32 
     n := len(arr)
     for i:=0 ; i< n ; i++{
         for j:=0;j<n ; j++{
             if i==j {
                 rdiag += arr[i][j]
             }
             if i+j==n-1 {
                 ldiag += arr[i][j]
             }
         }
     }
     dif := ldiag - rdiag
     if dif<0 {
         return -dif
     }
return dif
}
