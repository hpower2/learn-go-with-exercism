package scrabble

import "strings"


func Score(word string) int {
  newWord := strings.ToLower(word)
  words := map[int][]string{
    1 : {"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"},
    2 : {"d", "g"},
    3 : {"b", "c", "m", "p"},
    4 : {"f", "h", "v", "w", "y"},
    5 : {"k"},
    8 : {"j", "x"},
    10 : {"q", "z"},
  }
  var counter int
  for _, val := range newWord {
    for key, elements := range words{
      for _, elm := range elements{
        if string(val) == elm{
          counter += key
        }
      } 
    }
  }
  return counter
}
