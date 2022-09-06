// Complete the counterGame function below.
func counterGame(n int64) string {
var pow2,count int64=0,0
var i int64=0
for i=0;;i++{
    pow2=2<<i
    if pow2>n{
        count+=1
        n=n-(pow2>>1)
        if n%2==0{
        break    
        } else if n==1 {
            break
        } else {
            pow2=0
            i=0
        }
        
    } else if pow2 == n {
        count+=1
        n=n/2
        break
    } else {
        continue
    }
}
for i=0;;i++{
    if n==1{
        break
    }
    n=n>>1
    count+=1
}
if count%2==1{
    return "Louise"
}
return "Richard"
}
