# golang_target_sum

You are given an integer array `nums` and an integer `target`.

You want to build an **expression** out of nums by adding one of the symbols `'+'` and `'-'` before each integer in nums and then concatenate all the integers.

- For example, if `nums = [2, 1]`, you can add a `'+'` before `2` and a `'-'` before `1` and concatenate them to build the expression `"+2-1"`.

Return the number of different **expressions** that you can build, which evaluates to `target`.

## Examples

**Example 1:**

```
Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3

```

**Example 2:**

```
Input: nums = [1], target = 1
Output: 1

```

**Constraints:**

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `1000 <= target <= 1000`

## 解析

給定一個正整數陣列 nums, 還有一個整數 target

要求寫一個演算法 來找出所有可能透過 + 或是 - 把所有 nums 組成 target 的方法數

直覺的去思考

每個 nums[i] 有 + 或 - 兩種選擇

所以可以透過 decision tree 畫出所有可能

如下圖：

![](https://i.imgur.com/raRKzUE.png)

透過 DFS 走訪所有 leaves 如果 走到 leaves 的 total == target 則把方法數+ 1

這樣的話因為有 len(nums) 個元素 所以有 $2^n$ , n = len(nums)

所以時間複雜度是 O($2^n$)

而再走訪的過程可以發現 其實有些結點是重複出現的

所以可以利用 cache的方式來避免運算重複的值

因為已經算過的值不用重複運算

所以時間複雜度只要考量可能出現的不同值

因為每個值最多是 sum(nums) 到 -sum(nums) 也就是 介於 -t ~ +t , 假設 t= sum(nums)

所以時間複雜度就是 O(t * n) , 其中 n = len(nums) 

要使用動態規劃來減少運算

因為 target 可能是 正數或是複數

所以不同於以往 需要關注的可能有正值或是負值

而思考一下所有 nums 所可能組成最大最小值 會是 [-t, t],  where  t= sums(nums)

最基礎的排除法 假設 target 不在 這個範圍內 則方法數一定是 0 因為無法組成

為了要讓程式方便去做

所以可以把 amount 做 shift 從 [-t, t] 對應到 [0, 2t]

結束的位置則是 從 0 - len(nums)-1

![](https://i.imgur.com/iCFjm2d.png)

定義 dp[end][shiftedTotal]為 從 0 到 end 位置組成 shiftedTotal - sum(nums)的方法數

因為 dp[0][sum(nums)+nums[0]]  因為組成 nums[0] 的在 0 之前的方法數剛好只有 1 個

dp[0][sum(nums)+nums[0]] = 1

同樣的 dp[0][sum(nums)-nums[0]] = 1

對每個 dp[i-1][sum] > 0 代表這個 

代表sum + nums[i], sum- num[i] 可以由 前i 個組成

所以需要把 dp[i][sum+nums[i]] 還有 dp[i][sum-nums[i]]  最累加

每次更新 i 的 row 需要往前看前一個 row 非零的值做累加

而所就是 dp[len(nums)-1][sum(nums) + target]

因為需要 把每個可能值都loop 一遍 所以時間複雜度一樣是 O(n*t) , n 是陣列長度 , t代表 nums 的和

空間複雜度也是 O(n*t)

從更新的關係式可以發現每次只要根據前一次的值來做更新

所以其實可以只保留前一次 還有建立下一次的值代表只需要 O(t) 的空間複雜度 

## 程式碼
```go
package sol

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	if abs(target) > sum {
		return 0
	}
	nLen := len(nums)
	dp := make([][]int, nLen)
	for row := range dp {
		dp[row] = make([]int, 2*sum+1)
	}
	dp[0][sum+nums[0]] = 1
	dp[0][sum-nums[0]] += 1
	for end := 1; end < nLen; end++ {
		for total := -sum; total <= sum; total++ {
			if dp[end-1][total+sum] > 0 {
				dp[end][total+nums[end]+sum] += dp[end-1][total+sum]
				dp[end][total-nums[end]+sum] += dp[end-1][total+sum]
			}
		}
	}
	return dp[nLen-1][sum+target]
}

```
## 困難點

1. 要理解怎麼去累加出所要求出的target 
2. 這次的動態規劃問題出現了正數負數 amount 需要使用平移的方式來做處理
3. 動態規劃的方式並不夠直觀，遞迴關係需要跟位置與累積的和做處理

## Solve Point

- [x]  先計算出所有 nums 的和用來推算整個動態規劃的上下界
- [x]  建立矩陣 sum(nums)  by len(nums) 整數矩陣 dp 用來暫存 中間所有已算過的結果
- [x]  透過關係式算出所有 dp[i][sum+total]
- [x]  回傳 dp[len(nums)-1][sum(nums)+target]