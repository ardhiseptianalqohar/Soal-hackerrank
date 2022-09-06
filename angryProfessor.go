// Complete the angryProfessor function below.
func angryProfessor(k int32, a []int32) string {
    var count,i int32 =0,0
    for i=0;i<int32(len(a));i++{
        if a[i]<=0{
            count++
        }
    }
    if count>=k{
        return "NO"
    }
    return "YES"


}
