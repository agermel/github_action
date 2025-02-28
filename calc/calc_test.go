// 能有效避免与主包混淆，同时依然能够测试 calc 包中的功能
package calc_test

import (
	"action/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := calc.Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected to be 3, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, 7, 14},
		{"wrong", 2, 9, 1},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := calc.Multiple(c.A, c.B); ans != c.Expected {
				t.Errorf("%d*%d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}
