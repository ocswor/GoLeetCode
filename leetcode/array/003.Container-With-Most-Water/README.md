#[#Container With Most Water](https://leetcode-cn.com/problems/container-with-most-water/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china)   

##题目
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

###中文
给出一个非负整数数组 a1，a2，a3，…… an，每个整数标识一个竖立在坐标轴 x 位置的一堵高度为 ai 的墙，选择两堵墙，和 x 轴构成的容器可以容纳最多的水。


###解题思路

双指针移动法

求容纳更多的水。盛水的多少，取决于宽度 和高度的变化

思路上有一个干扰项就是宽度不固定，从数组首位遍历到数组末尾，移动的过程中宽度也在变化，这个导致思维有点难以聚集。

第一步还是要抽象化，建立模型
位置i  位置j   对应的高度分别是h[i] h[j]
那么 水槽的实际储水的最大面积是有由两个高度中的短板决定的。那么水槽的 面积公式 = min(h[i],h[j])×(j−i)。

那么如何解出i 和 j的值呢
一种暴力解法 
  for i:=0;i< len;i++{
      for j:=0;j<len;j++{
          area = min(h[i],h[j])×(j−i)
      }
  }

如果不用暴力解法 那就找规律了。

暴力解法是移动一个位置，固定一个位置。
那么移动两个位置，是不是更优解？

[4,8,6,2,5,4,8,1,7]
* 在每一个状态下，无论长板或短板收窄 1 格，都会导致水槽 底边宽度 -1：
   h[i] h[j] i j               面积
   4     7   0 8              4 * 8            // 水槽的面积是由短板决定的，所以 短板面积记录一下
   如果左指针右移 那么舍弃的就是 4 8 ;4 6; 4 2 ;4 5 ;4 4;4 8;4 1; 如果 h(i) < h(j)
  则相当于消去了 {S(i, j - 1), S(i, j - 2), ... , S(i, i + 1)}S(i,j−1),S(i,j−2),...,S(i,i+1) 状态集合。
  而这些消去的状态的面积一定<= S(i,j) //也可以称为裁剪搜索空间
  
代码都很简单关键是要理解下面几句话
>短板往中间走时直接挪到下一个比短板高的位置才有可能面积更大。
   
>面积取决于短板。①因此即使长板往内移动时遇到更长的板，矩形的面积也不会改变；遇到更短的板时，面积会变小。

这句话很有意思 就是面积是由短板决定的，及时遇到更长的高度的板也没用，因为短板决定面积。这个很有意思。遇到更短的面积也只会更小。

很有意思的题目,短板决定面积，盛水问题。

> ②因此想要面积变大，只能让短板往内移动(因为移动方向固定了)，当然也有可能让面积变得更小，但只有这样才存在让面积变大的可能性

基于条件①所以只有一种可能了么？ 只能移动短板了么？貌似是的。

   

