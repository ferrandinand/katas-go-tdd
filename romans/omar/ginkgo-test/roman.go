package roman 

func Convert(number int) string {
  var RomanNumber string
  if number == 1 {
    RomanNumber = "I"
   } 
  if number == 2 {
    RomanNumber = "II"
   }
  if number == 3 {
  RomanNumber = "III"
  }
  return RomanNumber
} 

func main () {
}
