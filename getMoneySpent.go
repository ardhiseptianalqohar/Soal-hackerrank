/*
 * Complete the getMoneySpent function below.
 */
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
    var sum,maxsum,flag int32=0,0,0
    for i:=0;i<len(keyboards);i++{
        for j:=0;j<len(drives);j++{
          sum=keyboards[i]+drives[j]
          if sum>maxsum && sum<=b{
              maxsum=sum
              flag=1
          }  
        }
    }
    if flag==1{
        return maxsum
    }
    return -1
}
