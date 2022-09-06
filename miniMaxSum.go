// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
var min,max int32 = arr[0],arr[0]
var sum int64 =0
for i:=0;i<len(arr);i++{
    if arr[i]>max {
        max = arr[i]
    }
    if arr[i]<min {
        min = arr[i]
    }
    sum+=int64(arr[i])
}

fmt.Println(sum-int64(max),sum-int64(min))
}
