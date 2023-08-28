package main

func longestCommonPrefix(strs []string) string {
	pref := strs[0]

	for i := range strs {
		if len(pref) > len(strs[i]) {
			pref = strs[i]
		}
	}
	for i := range strs {
		for j, c := range pref {
			if byte(c) != strs[i][j] {
				pref = pref[:j]
				break
			}
		}
	}
	return pref
}
