// Complete the circularArrayRotation function below.
func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
    n:=len(a)
    result:=[]int32{}
    for j:=0;j<int(k%int32(n));j++{
        temp:=a[n-1]
        for i:=n-1;i>0;i--{
            a[i]=a[i-1]
        }
        a[0]=temp  
    }
    for q:=0;q<len(queries);q++{
        result=append(result,a[queries[q]])
    }
return result
}
