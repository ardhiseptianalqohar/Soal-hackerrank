// Complete the maximizingXor function below.
func maximizingXor(l int32, r int32) int32 {
var xor,maxsum int32 = 0,0
for i := l; i <= r; i++{
    for j:=l;j<=r;j++{
      xor=i^j
      if xor>maxsum{
          maxsum=xor
      } 
    }
}
return maxsum
}
