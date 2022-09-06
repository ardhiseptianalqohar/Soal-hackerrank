package main
import "fmt"

func simpleArraySum(ar []int32) int32 {
    size := len(ar)
    i:=0
    var sum int32 =0
    for ; i< size ; {
        sum = sum + ar[i]
        i++
    } 
    return sum
}

func main(){
  arr := []int32{1,2,3} //array elements
  sum := simpleArraySum(arr)
  fmt.Println(sum)
}
