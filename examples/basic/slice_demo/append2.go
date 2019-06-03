package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	s1 := arr1[1:2]
	s1 = append(s1, 6, 7, 8)

	fmt.Println("slice1:", s1)
	fmt.Println("arr1:", arr1)

	arr2 := [5]int{1,2 ,3 ,4, 5}
	s2 := arr2[2:3]
	s2 = append(s2, 6, 7, 8)
	fmt.Println("slice2:", s2)
	fmt.Println("arr2:", arr2)
}

/*
output:
slice1: [2 6 7 8]
arr1: [1 2 6 7 8]
slice2: [3 6 7 8]
arr2: [1 2 3 4 5]

从上述输出结果来看，arr1被修改，arr2保持了原样。slice1和slice2看似相同的操作，为啥会对arr1和arr2分别产生不同的影响呢？

slice对象有三个属性：ptr、len、cap，其中ptr表示执行分配的数组的指针，len表示切片长度，cap表示切片容量。
调用make来创建的slice语法：make([]int, len, cap)。可以看到，一个slice接收一个指定类型，一个指定长度和一个可选的容量参数。make方法调用后，它背后其实一样分配了一个指定类型的数组，并返回一个slice引用指向该数组。当cap参数未指定时，cap与len值相等。
使用切片操作创建slice语法：slice1 := arr[1:4]。对slice进行切片操作，并不会新创建一个对象（分配内存），而只是在原来slice的基础上移动指针位置。如上面的例子。
调用append的时候会先判断是否需要扩容，扩容也即是创建新的slice，其实也就是开辟了一个新的内存空间，并返回了指向新地址的指针，并将旧的slice的元素值，copy到新创建的slice，最后slice重新指向新分配的slice，这也就是本文开头的例子里arr2的值没有变化的原因。

当原slice容量足够，不需要进行扩容时，那对slice元素的追加，都是发生在原slice里的（数组里），所以，原数组被“悄悄”改变了。这也解释了，为什么arr1的状态被改变了的原因。
 */
