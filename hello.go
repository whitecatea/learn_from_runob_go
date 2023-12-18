package main

import "fmt"

// func add(x int,y int){
// 	var z
// 	z = x + y
// 	return z
// }

func keyWord() {

	/*
		以下是有效标识符：
		mahesh	kumar	abc	move_name	a_123
		myname50	_temp	j	a23b9	retVal
	*/

	/*
		以下是无效的标识符：
		1ab（以数字开头）
		case（Go语言的关键字）
		a+b（运算符是不允许的）
	*/

	//关键字
	/*
		下面列举了 Go 代码中会使用到的 25 个关键字或保留字：
		break	default		func	interface	select
		case	defer		go		map			struct
		chan	else		goto	package		switch
		const	fallthrough	if		range		type
		continue	for		import	return		var

		除了以上介绍的这些关键字，Go语言还有36个预定义标识符：
		append	bool	btye	cap		close	complex	complex64	complex128	uint16
		copy	false	float32	float64	imag	int		int8		int16		uint32
		int32	int64	iota	len		make	new		nil			panic		uint64
		print	println	real	recover	string	true	uint		uint8		uintptr

		程序一般由关键字、常量、变量、运算符、类型和函数组成。

		程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

		程序中可能会使用到这些标点符号：.、,、;、: 和 …。
	*/
}

func space() {
	// 单行注释
	/*
		Author by Gelen Steven
		我是多行注释
	*/

	//Go语言的字符串连接可以通过+实现：
	fmt.Println("Google" + "Runoob")

	//Go语言的空格
	var x int = 4
	const Pi float64 = 3.14159265358979323846

	//在运算符和操作数之间要使用空格能让程序更易阅读：
	// var apple = "apple"
	// var oranges = "oranges"
	// var fruit string
	// fruit=apple+oranges
	//在变量与运算符间加入空格，程序看起来更加美观，如：
	// fruit = apple + oranges

	//在关键字和表达式之间要使用空格。
	if x > 0 {
		//do something
		var result float64 = Pi * 5
		fmt.Println(result)
	}

	//在函数调用时，函数名和左边等号之间要使用空格，参数之间也要使用空格。
	// result := add(2, 3)

	//格式化字符串
	/*
		Go 语言中使用 fmt.Sprintf 或 fmt.Printf 格式化字符串并赋值给新串：

		Sprintf 根据格式化参数生成格式化的字符串并返回该字符串。
		Printf 根据格式化参数生成格式化的字符串并写入标准输出。
	*/
	// %d 表示整型数字， %s 表示字符串
	var stockcode = 123
	var enddate = "2023-11-21"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println((target_url))
}

func dataType() {

	//Go语言数据类型
	/*
		布尔型		布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
		数字类型	整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
		字符串类型	字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
		派生类型	包括：
					(a) 指针类型（Pointer）
					(b) 数组类型
					(c) 结构化类型(struct)
					(d) Channel 类型
					(e) 函数类型
					(f) 切片类型
					(g) 接口类型（interface）
					(h) Map 类型
	*/

	//数字类型
	/*
		uint8	无符号 8 位整型 (0 到 255)
		uint16	无符号 16 位整型 (0 到 65535)
		uint32	无符号 32 位整型 (0 到 4294967295)
		uint64	无符号 64 位整型 (0 到 18446744073709551615)
		int8	有符号 8 位整型（-128 到127）
		int16	有符号 16 位整型（-32768 到 32767）
		int32	有符号 32 位整型（-2147483648 到 2147483647）
		int64	有符号 64 位整型（-9223372036854775808 到 9223372036854775807）

		浮点型
		float32		IEEE-754 32位浮点型数
		float64		IEEE-754 64位浮点型数
		complex64	32 位实数和虚数
		complex128	64 位实数和虚数
	*/

	//其他数字类型
	/*
		byte	类似uint8
		rune	类型int32
		uint	32位或64位
		int		与uint一样大小
		uintptr	无符号整型，用于存放一个指针
	*/
}

func variable() {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	/***第一种，指定变量类型，如果没有初始化，则变量默认为零值***/
	//声明一个变量并初始化
	var aa = "RUNOOB"
	fmt.Println(aa)

	//没有初始化就为零值
	var bb int
	fmt.Println(bb)

	//bool 零值为false
	var cc bool
	fmt.Println(cc)

	var i int
	var f float64
	var bbb bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, bbb, s)

	/***第二种，根据值自行判定变量类型***/
	//var v_name = value

	var d = true
	fmt.Println(d)

	/***第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：
	v_name := value

	例如：
	var intVal int
	intVal := 1	// 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
	直接使用下面的语句即可
	intVal := 1	// 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
	intVal := 1 相等于：
	var intVal int
	intVal = 1
	*/
	// 可以将 var f string = "Runoob" 简写为 f:="Runoob":
	fStr := "Runoob" // var f string = "Runoob"
	fmt.Println(fStr)

	/*多变量声明
	类型相同多个变量，非全局变量
	var vname1, vname2, vname3 type
	vname1, vname2, vname3 = v1, v2, v3

	var vname1, vname2, vname3 = v1, v2, v3	// 和 Python很像，不需要显式声明变量类型，自动推断
	vname1, vname2, vname3 := v1, v2, v3	//出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误

	// 这种因式分解关键字的用法一般用于声明全局变量
	var (
		vname1 v_type1
		vname2 v_type2
	)
	*/
	var x, y int
	var (
		a1 int
		b1 bool
	)

	var c1, d1 int = 1, 2
	var e1, f1 = 123, "hello"

	// 这种不带声明格式的只能在函数体中出现
	//g, h := 123, "hello"

	g, h := 123, "hello"
	fmt.Println(x, y, a1, b1, c1, d1, e1, f1, g, h)
}

func main() {
	/* 这是我第一个简单的程序 */
	// fmt.Println("Hello, World!")
	// fmt.Println("菜鸟教程：runoob.com")

	//关键字
	// keyWord()

	//注释和空格
	// space()

	//数据类型
	// dataType()

	//变量
	variable()
}
