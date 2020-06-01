package main

import "fmt"

func main() {
	fmt.Println(Solution("nice", "nicer"))
}
func Solution(s1 string, s2 string) string {
	if s1 == s2 {
		return "NOTHING"
	}
	// ADD 1 character
	if (len(s2) - len(s1)) == 1 {
		add := true
		for i := 0; i < len(s1); i++ {
			if s2[i] != s1[i] {
				add = false
			}
		}
		if add {
			return "ADD " + string(s2[len(s2)-1])
		}
	}
	//  MOVE, CHANGE
	if len(s1) == len(s2) {
		//Get differnt pair
		difS1, difS2 := []byte{}, []byte{}
		// If s1 and s2 have only 1 different pair of characters
		// return CHANGE
		for i := 0; i < len(s2); i++ {
			if s2[i] != s1[i] {
				// Get different pair of characters
				difS1 = append(difS1, s1[i])
				difS2 = append(difS2, s2[i])
				for j := i + 1; j < len(s1); j++ {
					// Check MOVE able
					if s1[i] == s2[j] {
						// Try to change if possible
						sTemp := []byte(s1)
						temp := sTemp[i]
						for k := i; k < j; k++ {
							sTemp[k] = sTemp[k+1]
						}
						sTemp[j] = temp
						// If changed string == s2
						// Return MOVE
						if string(sTemp) == s2 {
							return "MOVE " + string(s1[i])
						}
					}

				}

			}
		}
		if len(difS1) == len(difS2) && len(difS1) == 1 {
			return "CHANGE " + string(difS1[0]) + " " + string(difS2[0])
		}
		return "IMPOSSIBLE"
	}
	return "IMPOSSIBLE"
}
