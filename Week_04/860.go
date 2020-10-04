/*
860. 柠檬水找零
https://leetcode-cn.com/problems/lemonade-change/description/
*/
//2020/10/04
package main

func main() {

}

//贪心
//如果5元，不找零
//如果10元，找5元，否则找不开
//如果20元，优先10+5，在找5*3，否则找不开
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five <= 0 {
				return false
			}
			five--
			ten++
		} else {
			if ten > 0 {
				if five <= 0 {
					return false
				}
				ten--
				five--
			} else {
				if five < 3 {
					return false
				}
				five -= 3
			}
		}
	}
	return true
}
