package response

// 在controller中断言service返回的error类型，如果是自定义类型，返回自定义错误的msg
type CusError struct {
	Msg string
}

func (err *CusError) Error() string {
	return err.Msg
}
