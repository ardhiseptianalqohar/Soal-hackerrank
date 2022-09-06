// Complete the saveThePrisoner function below.
func saveThePrisoner(n int32, m int32, s int32) int32 {
    if ((s+m-1)%n==0){
        return n
    } else{
        return ((s+m-1)%n)    
    }
    
}
