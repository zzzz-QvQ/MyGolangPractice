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

//2025.11.16 指针 运算符 系统输入 进制转换

// 获取变量地址用& 指取指和引用；*是解引用 即获取指针类型所指向的值
// 指针变量存的是一个地址 地址指向的空间才是值
// 例如 var ptr *int = &i,ptr是一个int型指针变量 本身的值是&i
// 值类型和指针数据类型要对应 值类型包括 int\float\bool\string\数组\结构体
// 值类型以外，引用类型有 指针/slice/map/chan/interface
// 值类型变量直接存储值，内存通常在栈中分配；引用类型变量存储地址，内存在堆上分配，若没有变量来引用这个地址，其空间就会被GC回收
/*package main

import (
	"fmt"
)

func main() {
	var num int = 1
	var add *int = &num
	fmt.Printf("num的地址为：%v\n", add)
	num = *add + 1
	fmt.Printf("num+1的值为：%d\n", num)
	var ptr *int
	ptr = &num
	*ptr = 10
	fmt.Printf("num的新值为：%d\n", num)
	//运算符部分
	//%取模运算公式 a%b=a-a/b*b
	//自增或自减运算只能单独占一行再赋值 因此a=i++是错误的 可以用a+=1来描述自增 也不存在++a了
	//关系运算符结果是true和false；逻辑运算符&&、||都具有短路的功能 以下为测试
	var logi int = 10
	if logi < 9 && logical() {
		fmt.Println("ok")
	}
	if logi > 9 || logical() {
		fmt.Println("ok")
	}
	//赋值运算符 <<=左移后赋值 &=按位与后赋值 ^=按位异或后赋值 |=按位或后复制
	//如何不使用中间变量对两变量进行交换
	var a, b int = 10, 20
	//法①   法③ 按位异或
	a, b = b, a
	fmt.Println("a=", a, "b=", b)
	//法②
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a=", a, "b=", b)
	//go不支持三元运算符
}
func logical() bool {
	fmt.Println("logical show")
	return true
}*/
//下划线_是空标识符，可代表任意标识符，但其对应的值会被忽略，故仅能当占位符使用，不能作标识符
//变量名、函数名、常量名首字母大写是公有的，类似于public，可以被其他包访问；若小写，则只可以在本包使用
//导入包用go.mod init

// 系统输入
/*package main

import (
	"fmt"
)

func main() {
	//输入 fmt.scanf()/fmt.scanln() 前者是格式化输入 后者一个个输入
	//scanln自动消耗换行符 故scanf后的第一个scanln不工作
	var name string
	var age byte
	var score float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入成绩")
	fmt.Scanln(&score)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)
	fmt.Println("输入结果为：姓名：",name,"年龄：",age,"成绩：",score,"是否通过：",isPass);
	//以上等同于
	fmt.Scanf("%s,%d,%f,%t",&name,&age,&score,&isPass);
	fmt.Println("输入结果为：姓名：",name,"年龄：",age,"成绩：",score,"是否通过：",isPass);
}*/
//进制转换与位运算
//8进制以0开头 16进制以0x开头
//2进制转8进制，从右边每3位转为10进制为8进制的一位；16进制就是每四位
//计算机运算时都以补码方式运算 结果再恢复为原码
//异或：相异为1 故结合以上两条，-2^2=-4

//2025.11.17 流程控制 循环 跳出循环

//分支
/*package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入年龄：")
	var age int
	fmt.Scanln(&age)
	if age > 18 { //if后条件不用括弧；就算只有一条语句，也要大括号
		fmt.Println("已成年")
	} else if age == 18 { //此处不能换行
		fmt.Println("正好成年")
	} else { //多分支中else不是必须的
		fmt.Println("未成年")
	}

	//switch分支结构
	//switch中每个case分支唯一，case关键字后可以加多个表达式；表达式必须和switch语句数据类型一致
	// 匹配项后不用加break；最后有一个default
	//switch穿透：fallthrough 若是在case后增加fallthrough 则会执行下一个case语句 默认穿透一层
	var key int;
	fmt.Println("请输入一个月份");
	fmt.Scanln(&key);
	switch key {
	case 1,3,5,7,8,10,12://此处switch类型是int 那么case后的表达式结果也应该是int
		fmt.Println("非腊月");
	case 4,6,9,11:
		fmt.Println("是腊月");
		fallthrough;//相当于其他语言不加break 执行下一个case的结果句子；
		// 例如此处输入4 输出：是腊月 是特殊腊月
	case 2:
		fmt.Println("是特殊腊月");
	default:
		fmt.Println("输入有误");
	}
}*/

// 循环部分 go中无while和do-while的语法
// for后不加括号
/*package main

import (
	"fmt"
)

func main() {
	fmt.Println("请输入循环次数：")
	var count int
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		fmt.Println("i love u")
	}
	//以上for中的内容也可以等价为 将初始化和迭代放在不同地方
	var j int = 0
	for j < count {
		fmt.Println("i dont aggre")
		j++
	}
	//第三种写法是for后直接跟{}，这是一个无限循环，需要配合break使用

	//for-range遍历数组和字符串 下展示遍历字符串的两种方式
	var str string = "hello world!大家"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	} //①老方法 将string看作数组，按照每个字节来遍历 所以不能正确识别中文
	//以上问题要通过str转为[]rune切片来解决
	str2 := []rune(str) //用新的切片类型的变量来接收
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}
	//②用range将str拆分为index下标和val值 是按字符来遍历 所以不会出现以上中文乱码
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}
}*/

// 多重循环控制 外层循环一次 内层循环全部
/*// 三个班每班5人 统计每个班平均分和所有班级平均分
package main

import (
	"fmt"
)

func main() {
	var avrePerClass, averClass float32
	var classNum, stuNum int = 3, 5
	var totalSum float32 = 0
	for i := 0; i < classNum; i++ {
		fmt.Printf("请依次输入第%d班的%d个学生的成绩\n", i+1, stuNum)
		var sum float32 = 0
		for j := 0; j < stuNum; j++ {
			fmt.Printf("第%d个学生的成绩是：\n", j+1)
			var tempClassScore float32
			fmt.Scanln(&tempClassScore)
			sum += tempClassScore
		}
		avrePerClass = sum / float32(stuNum)
		fmt.Println("该班平均成绩为：", avrePerClass)
		totalSum += sum
	}
	averClass = totalSum / (float32(classNum) * float32(stuNum))
	fmt.Println("全部平均成绩为：", averClass)
}*/
// 跳转控制语句 break、continue、goto
//break出现在多层嵌套中时，可通过label标签来指明要终止的是哪一层语句块
//continue也可通过label来指定
//goto语句一般不使用 若使用 也是通过label
// 随机数的生成 目前已摈弃seed 0-100的整数：rand.Intn(100)

// 2025.11.18 函数
package main

import (
	"fmt"
)

func main() {
	fmt.Println()
}
