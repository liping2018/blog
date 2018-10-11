package utils

//异常捕捉
func Expel(code int, msg string, args ...interface{}) {
	data := make(map[string]interface{})
	panic(&R{code, msg, data})
}
