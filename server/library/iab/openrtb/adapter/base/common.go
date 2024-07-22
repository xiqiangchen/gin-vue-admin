package base

func DeepCopyStringSlice(original []string) (copiedSlice []string) {
	if original == nil {
		return
	}
	copiedSlice = make([]string, len(original))
	for i, v := range original {
		copiedSlice[i] = v
	}
	return copiedSlice
}
func DeepCopyByte(original []byte) (copiedByte []byte) {
	if original == nil {
		return
	}
	copiedByte = make([]byte, len(original))
	for i, v := range original {
		copiedByte[i] = v
	}
	return copiedByte
}
func DeepCopyInt(original []int) (copiedInt []int) {
	if original == nil {
		return
	}
	copiedInt = make([]int, len(original))
	for i, v := range original {
		copiedInt[i] = v
	}
	return copiedInt
}
