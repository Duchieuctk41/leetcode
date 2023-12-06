package main

// solution 1: use map
func findDifferences(oldStrings, newStrings []string) (onlyInOld, onlyInNew []string) {
	oldMap := make(map[string]bool)
	for _, v := range oldStrings {
		oldMap[v] = true
	}
	for _, v := range newStrings {
		if _, ok := oldMap[v]; ok {
			delete(oldMap, v)
		} else {
			onlyInNew = append(onlyInNew, v)
		}
	}
	for k := range oldMap {
		onlyInOld = append(onlyInOld, k)
	}
	return
}
