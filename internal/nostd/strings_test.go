package nostd

import (
	"fmt"
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	s := "今天的天气真好，abc，万里无云678"
	test := func(s string, substr string) {
		fmt.Println(strings.Contains(s, substr) == Contains(s, substr), "\t|\t", strings.Contains(s, substr), "\t|\t", Contains(s, substr))
		assert(strings.Contains(s, substr) == Contains(s, substr))
	}
	test(s, "")
	test(s, "天气")
	test(s, "ab")
	test(s, "无云6")
	test(s, "很好")
	test(s, "xyz")
}

func TestRepeat(t *testing.T) {
	test := func(s string, count int) {
		fmt.Println(Repeat(s, count) == strings.Repeat(s, count), "\t|\t", strings.Repeat(s, count), "\t|\t", Repeat(s, count))
		assert(Repeat(s, count) == strings.Repeat(s, count))
	}
	test("a", 1)
	test("abc", 10)
	test("好c", 10)
	test("", 10)
}

func TestJoin(t *testing.T) {
	test := func(elems []string, sep string) {
		fmt.Println(strings.Join(elems, sep) == Join(elems, sep), "\t|\t", strings.Join(elems, sep), "\t|\t", Join(elems, sep))
		assert(strings.Join(elems, sep) == Join(elems, sep))
	}
	test([]string{}, "x")
	test([]string{"Y"}, "x")
	test([]string{"a", "b", "c", "大家好"}, "")
	test([]string{"a", "b", "c", "大家好"}, "-")
	test([]string{"a", "b", "c", "大家好"}, "hello你好")
	test([]string{"", "a", "b", "c", "大家好", ""}, "hello你好")
}

func TestReplaceAll(t *testing.T) {
	s := "今天的天气真好，abc，万里无云678"
	test := func(s string, old string, new string) {
		fmt.Println(strings.ReplaceAll(s, old, new) == ReplaceAll(s, old, new), "\t|\t", strings.ReplaceAll(s, old, new), "\t|\t", ReplaceAll(s, old, new))
		assert(strings.ReplaceAll(s, old, new) == ReplaceAll(s, old, new))
	}
	test(s, "xx", "yy")
	test(s, "云6", "yy")
	test(s, "云6", "")
	test(s, "真好", "很好")
	test(s, "", "x")
	test(s, "", "很好")
}

func TestTrim(t *testing.T) {
	test := func(s string, cutset string) {
		fmt.Println(strings.Trim(s, cutset) == Trim(s, cutset), "\t|\t", strings.Trim(s, cutset), "\t|\t", Trim(s, cutset))
		assert(strings.Trim(s, cutset) == Trim(s, cutset))
	}
	test("今天的天气真好，abc，万里无云678", "")
	test("今天的天气真好，abc，万里无云678", "\n")
	test("今天的天气真好，abc，万里无云678", " ")
	test("\n\n今天的天气真好，abc，万里无云678\n\n", "\n")
	test(" \n今天的天气真好，abc，万里无云678\n\n", "\n")
	test("  今天的天气真好，abc，万里无云678\n\n", " ")
	test("  今天的天气真好，abc，万里无云678      ", " ")
}

func TestSplit(t *testing.T) {
	test := func(s string, sep string) {
		r1 := strings.Split(s, sep)
		r2 := Split(s, sep)
		s1 := "\"" + strings.Join(r1, "\",\"") + "\""
		s2 := "\"" + strings.Join(r2, "\",\"") + "\""
		ok := len(r1) == len(r2) && s1 == s2
		fmt.Println(ok, "\t|\t", s1, "\t|\t", s2)
		assert(ok)
	}
	test("", "")
	test("", "\n")
	test("今天的天气真好", "")
	test("今天的天气真好", "x")
	test("今 天 的天 气真好", " ")
	test("今 天 的x天 x气真好", "x")
	test("的今 天 的x天的x的气真好的", "的")
}
