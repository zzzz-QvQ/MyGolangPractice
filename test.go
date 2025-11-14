//2025.11.13

// 每个文件都必须归属于一个包 此处包名为main
package main

//go中定义的变量或者import的包如果没有使用，编译不能通过
import "fmt"

//执行入口依旧是main函数
func main() {
	fmt.Println("Hello,Go!")
} //语句后不用加分号 自动添加
//在命令行中对其编译 可以用go build test.go 得到test.exe 运行exe程序 具有移植性

//转义字符 \t制表位 \r回车，但是不换行，会将开头内容挤掉进行覆写
//dos常用操作指令 mkdr

//2025.11.14
