package gorgonnx

import "strconv"

var uniq int

func getUniqNodeName(prefix string) string {
	uniq++
	return prefix + strconv.Itoa(uniq)
}
