package bytesearch

//Go implementation of Boyer–Moore–Horspool string search algorithm, but for byte slices
//http://stackoverflow.com/questions/16252518/boyer-moore-horspool-algorithm-for-all-matches-find-byte-array-inside-byte-arra

//IndexOfAll Find all occurrences of a byte pattern within a byte slice
func IndexOfAll(value []byte, pattern []byte) []int {

	valueLength := len(value)
	patternLength := len(pattern)

	indexes := make([]int, 0)

	if valueLength == 0 || patternLength == 0 || (patternLength > valueLength) {
		return indexes
	}

	badCharacters := make([]int, 256)

	for i := 0; i < 256; i++ {
		badCharacters[i] = patternLength
	}

	lastPatternByte := patternLength - 1

	for i := 0; i < lastPatternByte; i++ {
		badCharacters[pattern[i]] = lastPatternByte - i
	}

	//Beginning
	index := 0

	for index <= (valueLength - patternLength) {
		for i := lastPatternByte; value[(index+i)] == pattern[i]; i-- {
			if i == 0 {
				indexes = append(indexes, index)
				break
			}

		}

		index += badCharacters[value[(index+lastPatternByte)]]
	}

	return indexes
}
