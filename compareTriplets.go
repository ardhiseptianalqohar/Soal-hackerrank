// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
    var alice,bob int32
    result := []int32{}
    for i:=0;i<3;i++{
        if a[i]>b[i] {
            alice+=1
        }
        if b[i] > a[i] {
            bob+=1
        }
        }
        result = append(result,alice,bob)    
        return result
    }
