/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
    n:=len(a)
    max:=0
    for i:=0;i<n;i++{
        count:=0
        for j:=0;j<n;j++{
            if a[i]-a[j]==1 || a[i]-a[j]==0 {
                count++
            }
            if count>max{
                max = count
            } 
        }
    }
return int32(max)
}
