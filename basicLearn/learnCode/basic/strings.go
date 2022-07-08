package main

import (
	//"strings"
	"strconv"
	"strings"
)

func main() {
	/*
		str1:="abcdefg"
		str2:="ab"
		str3:="fg"
		ans1:=strings.HasPrefix(str1,str2)//str2是否是str1的前缀
		ans2:=strings.HasSuffix(str1,str3)//str3是否是str1的后缀
		ans3:=strings.Contains(str1,str2)//str1中是否包含str2
		strings.ToUpper(str1)
		strings.ToLower(str1)

		println(ans1,ans2,ans3)


		//将字符串以空格分开
		str:="hello world !"
		s:=strings.Fields(str)
		for _,ss := range(s){
			println(ss)
		}
	*/

	//以具体内容分开
	str := "abc|bcd|def"
	strAns := strings.Split(str, "|")
	for index, ans := range strAns {
		println(index, ans)
	}

	//将分开的数据拼接
	strAns1 := strings.Join(strAns, "-")
	println(strAns1)

	//整数转化成字符串
	num1 := 10
	str1 := strconv.Itoa(num1)    //将整型转字符串
	str2, _ := strconv.Atoi(str1) //将字符串转整型
	println(str1, str2)

	//将字符串转化成小数
	str3 := "10.1"
	num2, _ := strconv.ParseFloat(str3, 32)
	println(num2)

	strSpace := "  1234  "
	space := strings.TrimSpace(strSpace)
	println(space)
}
