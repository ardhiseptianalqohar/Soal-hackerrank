// Complete the plusMinus function below.
func plusMinus(arr []int32) {
    n:=len(arr)
    
    var l,g,z float64
  //  var lt,gt,zz float64 
    for i:=0;i<n;i++{
        if arr[i]<0{
            l++
        } else if arr[i]>0 {
            g+=1
        } else {
            z++
        } 
    }
    fmt.Printf("%.6f\n",g/float64(n))
    fmt.Printf("%.6f\n",l/float64(n))
    fmt.Printf("%.6f",z/float64(n))
}
