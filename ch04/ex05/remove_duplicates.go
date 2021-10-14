package remove_duplicates

func removeDuplicates(strs []string) []string {
	w := 0
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}
