package helper

// Ternary 三元运算符泛型
// condition 为 true 返回 trueVal,否则返回 falseVal
func Ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
