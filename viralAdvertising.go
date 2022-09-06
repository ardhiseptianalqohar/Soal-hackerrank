// Complete the viralAdvertising function below.
func viralAdvertising(n int32) int32 {
var r,l,i int32=5,0,0
var sum int32=0
for i=0;i<n;i++{
    l=r/2
    sum+=l
    r=(r/2)*3
}
return sum
}
