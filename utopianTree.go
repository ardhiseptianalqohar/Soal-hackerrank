// Complete the utopianTree function below.
func utopianTree(n int32) int32 {
    var h int32=1
    for i:=0;i<int(n);i++{
        if i%2==0{
            h*=2
        } else {
            h+=1
        }
    }
    return h

}
