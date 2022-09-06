// Complete the findDigits function below.
func findDigits(n int32) int32 {
    defer func() {
     if r := recover(); r != nil {
        //fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
    }
  }()
var r,q int32
var count int32 =0
for q=n;q!=0;{
    r=q%10
    if r==0{
        q/=10
        continue
    }
    if n%r==0{
        count++
    }
    q/=10
}
return count
}
