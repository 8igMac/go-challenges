package main

func IsUnique(str string) bool {
  // Method1: Use map.
	// dict := make(map[rune]bool)
	// for _, c := range str {
	// 	_, exist := dict[c]
	// 	if exist {
	// 		return false
	//   } else {
	//     dict[c] = true
	//   }
	// }
	// return true

  // Method2: If I can't use any datastructure.
	var dict [256]bool
	for _, c := range str {
    if dict[c] {
      return false
    } else {
      dict[c] = true
    }
	}
  return true
}
