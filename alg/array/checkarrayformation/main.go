package checkarrayformation

func canFormArray(arr []int, pieces [][]int) bool {
	lookup := make(map[int][]int)
	for _, piece := range pieces {
		lookup[piece[0]] = piece
	}
	var ok bool
	count := len(arr)
	piece, ok := lookup[arr[0]]
	if !ok {
		return false
	}
	i, j := 0, 0
	for i < count {
		if j < len(piece) {
			if arr[i] != piece[j] {
				return false
			}
			j++
			i++
			continue
		}
		piece, ok = lookup[arr[i]]
		if !ok {
			return false
		}
		j = 0
	}
	return true
}
