//实践检验真理
package main

import(
	"fmt"
)

//1.1 找出程序的错误
// func main() {

// 	slice := []int{0,1,2,3}
// 	m := make(map[int]*int)

// 	for key,val := range slice {
// 		m[key] = &val
// 	}

//    for k,v := range m {
// 	   fmt.Println(k,"->",*v)
//    }
// }

//错误结果
// 0 -> 3
// 1 -> 3
// 2 -> 3
// 3 -> 3


//正确的写法
// func main() {

// 	slice := []int{0,1,2,3}
// 	m := make(map[int]*int)

// 	for key,_ := range slice {
// 		m[key] = &slice[key]
// 	}

//    for k,v := range m {
// 	   fmt.Println(k,"->",*v)//map的遍历是无序的
//    }
// }

//本题知识点
//参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，
//而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，
//所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.


//1,2 Go指针

// func main() {
//     a := 13
//     b := &a
//     fmt.Printf("a:%d ptr:%p\n", a, &a) // a:10 ptr:0xc00001a078
//     fmt.Printf("b:%p type:%T\n", b, b) // b:0xc00001a078 type:*int
//     fmt.Println(&b)                    // 0xc00000e018
//     fmt.Println(*b)                    // 0xc00000e018
// }



type student struct {
    id   int
    name string
    age  int
}

func changeValue(ce []student) {
    //切片是引用传递，是可以改变值的
    ce[1].age = 999
}

func addValue(ce []student)(res []student) {
    res = append(ce, student{3, "xiaowang", 56})
    return res
}
func main() {
    var ce []student  //定义一个切片类型的结构体
    ce = []student{
        student{1, "xiaoming", 22},
        student{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	changeValue(ce)
	fmt.Println(ce)
    res := addValue(ce)
    fmt.Println(res)
}



//1.3
//new 和make的区别
//指针 结构体 数组 切片
//先介绍基础用法 然后是底层原理
//go web项目实战 
//基础数据结构