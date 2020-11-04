package main

import "fmt"

/*
有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。
如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
words = ["I","am","a","student","from","a","university","in","a","city"],
word1 = "a", word2 = "student"
输出：1
words.length <= 100000
链接：https://leetcode-cn.com/problems/find-closest-lcci/
 */

func findClosest(words []string, word1 string, word2 string) int {
	n1, n2 := -0x3f3f3f3f, -0x3f3f3f3f
	rec := 0x3f3f3f3f
	min := func(n1, n2 int) int {
		if n1 < n2 {
			return n1
		}
		return n2
	}
	for i, x := range words {
		if x == word1 {
			n1 = i
			rec = min(rec, n1 - n2)
		} else if x == word2 {
			n2 = i
			rec = min(rec, n2 - n1)
		}
	}
	return rec
}


func main() {
	var words = []string{"I","am","a","student","from","a","university","in","a","city"}
	var word1, word2 = "a", "student"
	fmt.Print(findClosest(words, word1, word2))
}

