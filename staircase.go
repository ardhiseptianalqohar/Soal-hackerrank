// Complete the staircase function below.
func staircase(n int32) {
var i,j,k int32
for i=0;i<n;i++{
    for k=0;k<n-1-i;k++{
        fmt.Printf(" ")
        }
    for j=0;j<=i;j++{
        fmt.Printf("#")
        }
    fmt.Printf("\n")
    }

}
