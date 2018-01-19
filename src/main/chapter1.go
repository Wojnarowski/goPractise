package main 

import "fmt"

//字符串
//var frenchHello string  // 声明变量为字符串的一般方法
//var emptyString string = ""  // 声明了一个字符串变量，初始化为空字符串

func chapter1() {
//	no, yes, maybe := "no", "yes", "maybe"  // 简短声明，同时声明多个变量
//	japaneseHello := "Konichiwa"  // 同上
//	frenchHello = "Bonjour"  // 常规赋值
//	fmt.Println(no);
//	fmt.Println(yes);
//	fmt.Println(maybe);
//	fmt.Println(japaneseHello);
//	fmt.Println(frenchHello);
	
	//设置字符串
	//s := "hello"
	//c := []byte(s)  // 将字符串 s 转换为 []byte 类型
	//c[0] = 'c'
	//s2 := string(c)  // 再转换回 string 类型
	//fmt.Printf("%s\n", s2)
	
	
//	s:="hello"
//	m:="world"
//	a:=s+m;
//	fmt.Println(a);

//	s := "hello"
//	s = "c" + s[2:] // 字符串虽不能更改，但可进行切片操作
//	fmt.Printf("%s\n", s)

	//多行字符串
//	m := `hello
//	world`
//	fmt.Printf("%s\n", m)
	//数组
	var arr [10]int  // 声明了一个int类型的数组
	arr[0] = 42      // 数组下标是从0开始的
	arr[1] = 13      // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0])  // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0
	
		
}