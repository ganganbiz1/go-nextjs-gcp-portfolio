package util

func StringPointer(s string) *string {
	return &s
}

func IntPointer(i int) *int {
	return &i
}

func MakeChunkSlice[T any](slice []T, chunckSize int) [][]T {
	var chuncks [][]T
	for len(slice) > 0 {
		end := chunckSize
		if end > len(slice) {
			end = len(slice)
		}
		chuncks = append(chuncks, slice[:end:end])
		slice = slice[end:]
	}
	return chuncks
}
