package baidutranslatefree

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var i = "320305.131321201"

// n 函数实现位操作
func n(r int64, o string) int64 {
	for t := 0; t < len(o)-2; t += 3 {
		e := rune(o[t+2])
		var eNum int64
		if e >= 'a' {
			eNum = int64(e - 87)
		} else {
			eNum = int64(e - '0')
		}

		if o[t+1] == '+' {
			r = int64(uint64(r) >> eNum)
		} else {
			r = r << eNum
		}

		if o[t] == '+' {
			r = (r + eNum) & 4294967295
		} else {
			r ^= eNum
		}
	}
	return r
}

// a 函数实现主要逻辑
func a(r string) string {
	// 匹配代理对
	re := regexp.MustCompile(`[\x{D800}-\x{DBFF}][\x{DC00}-\x{DFFF}]`)
	t := re.FindAllString(r, -1)

	// 字符串截取逻辑
	if t == nil {
		a := len(r)
		if a > 30 {
			r = r[:10] + r[a/2-5:a/2+5] + r[len(r)-10:]
		}
	} else {
		C := strings.Split(r, `[\x{D800}-\x{DBFF}][\x{DC00}-\x{DFFF}]`)
		var u []string
		for h := 0; h < len(C); h++ {
			if C[h] != "" {
				u = append(u, strings.Split(C[h], "")...)
			}
			if h != len(C)-1 {
				u = append(u, t[h])
			}
		}
		if len(u) > 30 {
			r = strings.Join(u[:10], "") + strings.Join(u[len(u)/2-5:len(u)/2+5], "") + strings.Join(u[len(u)-10:], "")
		}
	}

	// 处理 token
	l := i
	m := strings.Split(l, ".")
	S, _ := strconv.ParseInt(m[0], 10, 64)
	s, _ := strconv.ParseInt(m[1], 10, 64)

	// 字符编码转换
	var c []int
	for F := 0; F < len(r); F++ {
		p := int(r[F])
		if p < 128 {
			c = append(c, p)
		} else if p < 2048 {
			c = append(c, p>>6|192)
			c = append(c, p&63|128)
		} else {
			if (p&64512) == 55296 && F+1 < len(r) && (int(r[F+1])&64512) == 56320 {
				p = 65536 + ((1023 & p) << 10) + (1023 & int(r[F+1]))
				F++
				c = append(c, p>>18|240)
				c = append(c, p>>12&63|128)
			} else {
				c = append(c, p>>12|224)
				c = append(c, p>>6&63|128)
			}
			c = append(c, p&63|128)
		}
	}

	// 位运算和字符串处理
	w := S
	A := "+-a^+6"
	b := "+-3^+b+-f"
	for D := 0; D < len(c); D++ {
		w += int64(c[D])
		w = n(w, A)
	}
	w = n(w, b)
	w ^= s

	// 处理负数
	if w < 0 {
		w = (2147483647 & w) + 2147483648
	}

	// 取模并格式化
	w %= 1e6
	return fmt.Sprintf("%d.%d", w, w^S)
}
