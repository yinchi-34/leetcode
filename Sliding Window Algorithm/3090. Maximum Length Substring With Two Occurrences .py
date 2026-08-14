func maximumLengthSubstring(s string) int {
      n := len(s)
      res := 0
      left := 0
      count := make([]int, 26)

      for right :=0; right < n; right ++ {
        ch := s[right] - 'a'
        count[ch] ++

        for count[ch] > 2{
            ch2 := s[left] - 'a'
            count[ch2] --
            left ++
        }
        length := right - left + 1
        if length > res{
            res = length
        }
      }
      return res

    //   // 每一次子查询都需要重新统计
    //   // 相当于while
    //   for left := 0; left < n; left++ {
    //     // 全是字母，python用ord，go可以直接相减
    //     count := make([]int, 26)
    //     for right := left; right < n; right ++ {
    //         ch := s[right] - 'a'
    //         count[ch] ++

    //         if count[ch] > 2 {
    //             break
    //         }
    //         // 所以先计算当前最长长度，以及更新最长长度
    //         length := right - left + 1
    //         if length > res {
    //             res = length
    //         }
    //     }
    //   }
    //   return res
}