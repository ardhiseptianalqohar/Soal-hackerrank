// Complete the hurdleRace function below.
func hurdleRace(k int32, height []int32) int32 {
n:=len(height)
max:=height[0]
for i:=0 ; i<n ;i++{
    if height[i]>max{
        max=height[i]
    }
}
if max>k{
    return (max-k)
} else {
    return 0
}

}
