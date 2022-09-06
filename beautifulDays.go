// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
    var x,count int32
    count=0
    for x=i;x<=j;x++{
        if (x-reverse(x))%k==0{
            count++
        }
        
    }
return count
}
func reverse(x int32) int32{
    var rev int32=0
    var q int32
    for q=x;q!=0;{
        r:=q%10
        rev=rev*10 + r    
        q/=10
    }
    
    return rev  
}
