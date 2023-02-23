package utils

// 定义filter函数
func Filter[TElement any](nums []TElement, f func(TElement) bool) []TElement {
	var result []TElement
	for _, num := range nums {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}
