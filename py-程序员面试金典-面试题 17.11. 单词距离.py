class Solution(object):
    """
有个内含单词的超大文本文件，给定任意两个单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。
如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?
words = ["I","am","a","student","from","a","university","in","a","city"],
word1 = "a", word2 = "student"
输出：1
words.length <= 100000
链接：https://leetcode-cn.com/problems/find-closest-lcci
    """
    def findClosest(self, words, word1, word2):
        """
        :type words: List[str]
        :type word1: str
        :type word2: str
        :rtype: int
        """
        n1 = n2 = float('-inf')
        m = float('inf')
        for i, x in enumerate(words):
            if x == word1:
                n1 = i
                m = min(m, n1 - n2)
            elif x == word2:
                n2 = i
                m = min(m, n2 - n1)
        return m


def main():
    words, word1, word2 = ["I", "am", "a", "student", "from", "a", "university", "in", "a", "city"], "a", "student"
    test = Solution()
    ret = test.findClosest(words, word1, word2)
    print(ret)


if __name__ == '__main__':
    main()
