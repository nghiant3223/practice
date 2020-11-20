package alphabetwordlist

var ans []string

const letterCount = 3

var letters = []uint8{'a', 'b', 'c'}

func alphabetWordList(k, n int) []string {
	alphabetWordListUtil(k, n, "")
	return ans
}

func alphabetWordListUtil(k, n int, cur string) {
	lenCur := len(cur)
	if lenCur > k {
		return
	}
	if lenCur == k && len(ans) < n {
		ans = append(ans, cur)
		return
	}
	for i := 0; i < letterCount; i++ {
		if cur == "" || cur[lenCur-1] != letters[i] {
			alphabetWordListUtil(k, n, cur+string(letters[i]))
		}
	}
}
