package errorp

import "fmt"

//exercise err
//func Errorf(format string, a ...any) error {...}
type ErrSqrt float64

/*
方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error()
调用该方法应返回 "cannot Sqrt negative number: -2"。

注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。
可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
*/
func (er ErrSqrt) Error() string {

	/*
		细节：
		fmt.Sprint 会尝试调用 e.Error() 来获取 e 的字符串表示。
		如果 e.Error() 内部又调用了 fmt.Sprint(e)，
		那么 fmt.Sprint 会再次调用 e.Error()，如此往复，形成无限循环
		可以强制转换er类型

	*/
	return fmt.Sprint("cannot Sqrt negative number:", float64(er))

}
