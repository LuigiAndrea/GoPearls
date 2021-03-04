package fastsearch

//hash table saved in memory
var hashTable map[int]int

//build hash table
func init() {
	hashTable = make(map[int]int)
	for i := 0; i < len(data); i++ {
		hashTable[data[i]] = i //saved position
	}
}

// 9.5.10
// fast search with fix size array with 100 unique elements
func fastHashingSearch(v int) int {
	if _, ok := hashTable[v]; ok {
		return hashTable[v]
	}
	return -1
}
