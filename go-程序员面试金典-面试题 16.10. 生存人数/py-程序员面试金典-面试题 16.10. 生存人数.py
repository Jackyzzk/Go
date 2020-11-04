class Solution(object):
    """
给定 N 个人的出生年份和死亡年份，第 i 个人的出生年份为 birth[i]，
死亡年份为 death[i]，实现一个方法以计算生存人数最多的年份。
你可以假设所有人都出生于 1900 年至 2000 年（含 1900 和 2000 ）之间。如果一个人在某一年的任意时期处于生存状态，
那么他应该被纳入那一年的统计中。例如，生于 1908 年、死于 1909 年的人应当被列入 1908 年和 1909 年的计数。
如果有多个年份生存人数相同且均为最大值，输出其中最小的年份。
输入：
birth = {1900, 1901, 1950}
death = {1948, 1951, 2000}
输出： 1901
0 < birth.length == death.length <= 10000
birth[i] <= death[i]
链接：https://leetcode-cn.com/problems/living-people-lcci
    """
    def maxAliveYear(self, birth, death):
        """
        :type birth: List[int]
        :type death: List[int]
        :rtype: int
        """
        rec = [(birth[x], death[x]) for x in range(len(birth))]
        rec.sort(key=lambda x: (x[0], x[0] - x[1]))
        left, right = 1900, 2000
        count = 0
        ret = 0
        target = 0
        for x in rec:
            if left <= x[0] <= right:
                left = x[0]
                count += 1
                right = min(right, x[1])
            else:
                if count > ret:
                    target = left
                    ret = count
                count = 0
                left, right = 1900, 2000
        return target


def main():
    birth, death = [1900, 1901, 1950], [1948, 1951, 2000]
    birth, death = [1972,1908,1915,1957,1960,1948,1912,1903,1949,1977,1900,1957,1934,1929,1913,1902,1903,1901], \
                    [1997,1932,1963,1997,1983,2000,1926,1962,1955,1997,1998,1989,1992,1975,1940,1903,1983,1969]  # 1960
    test = Solution()
    ret = test.maxAliveYear(birth, death)
    print(ret)


if __name__ == '__main__':
    main()
