// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
    var dcata,dcatb int32
    dcata=z-x
    dcatb=z-y
    if dcata<0{
        dcata=-dcata
    }
    if dcatb<0{
        dcatb=-dcatb
    }
    if dcata<dcatb{
        return "Cat A"
    } else if dcatb<dcata {
        return "Cat B"
    } else {
        return "Mouse C"
    }

}
