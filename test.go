//2025.11.13

/*// 每个文件都必须归属于一个包 此处包名为main
package main
//go中定义的变量或者import的包如果没有使用，编译不能通过
import "fmt"

//执行入口依旧是main函数
func main() {
	fmt.Println("Hello,Go!")
} //语句后不用加分号 自动添加
//在命令行中对其编译 可以用go build test.go 得到test.exe 运行exe程序 具有移植性

//转义字符\n回车 \t制表位 \r回车，但是不换行，会将开头内容挤掉进行覆写
//dos常用操作指令 mkdr*/
//函数或方法首字母大写表示可导出，小写的不能导出

//2025.11.14 变量声明

/*func main(){
//golang三种变量使用方法
//第一种
var i int
fmt.Println("i=",i);
//第二种 自动识别类型
var num = 10.11
fmt.Println("num=",num);
//第三种
name :="tom";//:=声明并赋值 等价于var name string;name = "tom";
fmt.Println("name=",name);
//一次性定义多个相同类型或不同类型变量
var a1,a2,a3 int;
fmt.Println("n1=",a1,"n2=",a2,"n3=",a3);
var n1,name2,n3 = 100,"tom",888;
fmt.Println("n1=",n1,"name=",name2,"n3=",n3)
//在全局变量中也可以
	// var(
	// 	a1 = 1;a2 = 2;a3 = 3;
	// )
//同一变量的数据值可不断变化但是类型范围不变 不能由int到float
}*/

//2025.11.15 数据类型 %T表示type %v表示值的默认格式

// 数值型 int int8:8bit即1Byte int16 int32 int64 uint无符号整数 全取正
// rune 有符号 与int32同 表示一个Unicode码
// 浮点数 float32 float64 相当于float与double
// 无专门字符型，用byte来保存单个字母字符 故不能保存汉字 因为go里采用utf8 其中汉字占三个字节
// 字符串是基本数据类型string 传统字符串由字符组成 go中字符串由字节组成
// 字符串不可变 除了双引号 还有反引号，以原生形式输出，不用担心转义字符
// 字符串的拼接
// bool类型 值为false或true 而不是0或1 占一个字节
// 基本数据类型的转换 不能自动转换 需要强制转换
//基本数据类型与string间的相互转换
// 指针 数组 结构体 管道 函数 切片 函数 map集合
/*package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main() {
	var n1 int64 = 10
	fmt.Printf("n1的数据类型是%T,占用的字节数是%d", n1, unsafe.Sizeof(n1))
	fmt.Println()
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, ",c2=", c2) //输出的对应字符的ASCII码值
	fmt.Printf("c1=%c,c2=%c\n", c1, c2)  //printf格式化输出对应字符
	//若是输出汉字“北” byte会溢出
	var c3 int = '北'
	fmt.Printf("c3=%c\n", c3)
	//字符串反引号防止歧义
	str := `
	//原有注释
	`
	fmt.Println("str的内容是：", str);
	//字符串拼接
	var str1 string = "hello" + "world" + "iam" + //此处换行 +号要留在上一行
	"coming!";
	str1 += "haha";
	fmt.Println("str1拼接后为",str1);
	//基本数据类型的转换 大转小容易溢出
	var i int32 = 100;
	var n float32 = float32(i);//只改变i的值的数据类型 不改变i本身的数据类型
	fmt.Printf("i=%v,%T n=%v,%T\n",i,i,n,n);//%v表示值的默认格式
	//基本数据类型转string ①fmt.Sprintf("%参数",表达式) ②strconv包里的函数
	var num1 int = 99;
	var num2 float64 = 23.456;
	var b bool = true;
	var myChar byte = 'h';
	var changeStr string ;
	//方法①
	changeStr = fmt.Sprintf("%d",num1);//将原有%d的num1的值 转化为string类型存储到str中
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	changeStr = fmt.Sprintf("%f",num2);
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	changeStr = fmt.Sprintf("%t",b);
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	changeStr = fmt.Sprintf("%c",myChar);
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	//方法②
	changeStr = strconv.FormatInt(int64(num1),10);
	//strconv.FormatInt() 需要 int64 类型参数 10表示从10进制开始转换
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	//上述int的转换 也可用strconv.Itoa(num1)
	changeStr = strconv.FormatFloat(num2,'f',2,64);//f是格式 2表示小数位数保留2位 64表示小数是float64
	fmt.Printf("changeStr type is %T,str = %q\n",changeStr,changeStr);
	//string转基本数据类型 使用strconv包里的函数 必须转换有效，否则转换结果为0
	var strChange string = "true";
	var newb bool;
	newb,_ = strconv.ParseBool(strChange);
	//strconv.ParseBool()会返回两个值，一个bool，一个error,它是一个值，用_将error置空
	fmt.Printf("newb type is %T,newb = %v\n",newb,newb);
	var strChange2 string = "12345";
	var newn1 int64
	newn1,_ = strconv.ParseInt(strChange2,10,64);//64为转换结果的位数限制 若超出则输出err
	fmt.Printf("newn1 type is %T,newn1 = %v\n",newn1,newn1);
}
//print的不同类型及其作用
*/

// 指针
package main

func main() {

}
