package utils

// @function: Pointer
// @description: 获取变量的指针
// @param:  in T
// @return: *T
func Pointer[T any](in T) (out *T) {
	return &in
}
