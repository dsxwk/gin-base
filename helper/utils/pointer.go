package utils

// To 泛型转换
func To[T any](v any) T {
	return v.(T)
}

// InArray 判断元素是否在数组中
func InArray[T comparable](item T, list []T) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
