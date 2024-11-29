package piscine

func StringToIntSlice(str string) []int {
	if str != "" {
		arr := []int{}
		for _, c := range str {
			arr = append(arr, int(c))
		}
		return arr
	}
	return nil
}
