package utils

func UpdateResponseMap(respMap map[string][][]byte, key string, value []byte) {
	if _, ok := respMap[key]; ok {
		respMap[key] = append(respMap[key], value)
	} else {
		respMap[key] = [][]byte{value}
	}
}
