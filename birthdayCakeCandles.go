/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
    var max int32 =candles[0]
    for i:=0;i<len(candles);i++{
        if candles[i]>max{
            max = candles[i]
        }
    }
  var count int32=0
  for i:=0;i<len(candles);i++{
      if candles[i] == max{
          count++
      }
  }
return count
}
