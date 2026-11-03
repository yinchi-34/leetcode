package solution

// GroupAnagrams groups the strings in strs that are anagrams of each other.
// Each string is keyed by its letter counts, a [26]int array; Go arrays are
// comparable, so the array itself can be a map key. Within a group, strings
// keep their input order; the order of the groups is unspecified.
// strs must contain only lowercase English letters.
func GroupAnagrams(strs []string) [][]string {
	groups := map[[26]int][]string{}
	for _, s := range strs {
		var cnt [26]int
		for i := 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}
		groups[cnt] = append(groups[cnt], s)
	}

	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}
