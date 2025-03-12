package utils

// Pointer 获取变量的指针
// @param:  in T
// @return: out *T
func Pointer[T any](in T) (out *T) {
	return &in
}
