package main

import (
	"fmt"
	"strings"
)

/*
二进制数转字符串。给定一个介于0和1之间的实数（如0.72），类型为double，打印它的二进制表达式。
如果该数字不在0和1之间，或者无法精确地用32位以内的二进制表示，则打印“ERROR”。
输入：0.625
输出："0.101"
输入：0.1
输出："ERROR"
提示：0.1无法被二进制准确表示
32位包括输出中的"0."这两位。
十进制小数转二进制小数，应用常规方法乘二取整即可求解
链接：https://leetcode-cn.com/problems/bianry-number-to-string-lcci
 */

func printBin(num float64) string {
	ret := []string{"0", "."}
	for i := 30; num != 0 && i > 0; i-- {
		num *= 2
		if num >= 1 {
			ret = append(ret, "1")
			num -= 1
		} else {
			ret = append(ret, "0")
		}
	}
	if num == 0 {
		return strings.Join(ret, "")
	} else {
		return "ERROR"
	}
}

func main() {
	num := 0.625
	num = 0.1
	ret := printBin(num)
	fmt.Println(ret)
}