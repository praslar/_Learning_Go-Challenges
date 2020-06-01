package main

import "fmt"

func main() {
	fmt.Println(Solution([]int{-3, -14, -5, 7, 8, 42, 8, 3}))
}

func Solution(year []int) string {
	seasonDays := len(year) / 4
	max := 0
	season := 0
	// For each season
	for i := 0; i < len(year); i += seasonDays {
		amplitude := 0
		// maxT and minT of each season
		maxT := year[i]
		minT := year[i]
		for j := 0; j < seasonDays-1; j++ {
			if year[j+i+1] > maxT {
				maxT = year[i+j+1]
			}
			if year[i+j+1] < minT {
				minT = year[i+j+1]
			}
		}
		// Ampltude = maxT - minT
		amplitude = maxT - minT
		if amplitude > max {
			max = amplitude
			season = i
		}
	}
	// Return season have the max amplitude
	// i	== 0 || 3 || 6 || 9
	// days == 3 || 3 || 3 || 3
	// season  WI|| SP|| SU|| AU
	switch season / seasonDays {
	default:
		return "WINTER"
	case 1:
		return "SPRING"
	case 2:
		return "SUMMER"
	case 3:
		return "AUTUMN"
	}
}
