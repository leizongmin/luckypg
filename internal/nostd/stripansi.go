package nostd

// import "regexp"
// const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"
// var re = regexp.MustCompile(ansi)
// func StripAnsi(str string) string {
// 	return re.ReplaceAllString(str, "")
// }

// StripAnsi 去除 ANSI 控制字符
// 以 \x1b[ 开头，以 m 结尾
func StripAnsi(s string) string {
	runes := []rune(s)
	result := make([]rune, 0)
	isAnsiStart := false
	for i, r := range runes {
		if isAnsiStart {
			if r == 'm' {
				isAnsiStart = false
				continue
			}
		} else {
			if r == '\x1b' {
				if i < len(runes)-1 && runes[i+1] == '[' {
					isAnsiStart = true
					continue
				}
			} else {
				result = append(result, r)
			}
		}
	}
	return string(result)
}
