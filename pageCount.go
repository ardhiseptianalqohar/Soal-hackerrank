
/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) int32 {
    var front,back int32
    front=p/2
    if n%2==0 && p%2==1{
        back=(n-p)/2 +1    
    } else {
        back=(n-p)/2 
    }
    
    if front<back{
        return int32(front)
    } else {
        return int32(back)
    }

}
