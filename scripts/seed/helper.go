package main

import "time"

func mergeFloat64SliceMap(ms ...map[string][]float64) map[string][]float64 {
	res := map[string][]float64{}
	for _, m := range ms {
	srcMap:
		for k, v := range m {
			// Check if (k,v) was added before:
			for _, v2 := range res[k] {
				if float64InSlice(v2, v) {
					continue srcMap
				}
			}
			res[k] = append(res[k], v...)
		}
	}
	return res
}

func float64InSlice(a float64, list []float64) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getCurrYear() int {
	t := time.Now()
	return t.Year()
}

func getCurrSeason() string {
	t := time.Now()
	month := int(t.Month())
	if month >= 1 && month <= 3 {
		return "spring"
	}
	if month >= 4 && month <= 6 {
		return "summer"
	}
	if month >= 7 && month <= 9 {
		return "fall"
	}
	return "winter"
}

func getLastSeason() string {
	season := getCurrSeason()
	switch s := season; s {
	case "spring":
		return "winter"
	case "summer":
		return "spring"
	case "fall":
		return "summer"
	default:
		return "fall"
	}
}

func isLastSeason(year int, season string) bool {
	currYear := getCurrYear()
	currSeason := getCurrSeason()

	lastYear := currYear
	lastSeason := getLastSeason()

	if currSeason == "spring" {
		lastYear = currYear - 1
	}

	return year == lastYear && season == lastSeason
}
