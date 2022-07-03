package nostd

func isRunesEqual(r1 []rune, r2 []rune) bool {
	return string(r1) == string(r2)
}

func Contains(s string, substr string) bool {
	if substr == "" {
		return true
	}
	r1 := []rune(s)
	r2 := []rune(substr)
	for i := 0; i < len(r1)-len(r2); i++ {
		if isRunesEqual(r1[i:i+len(r2)], r2) {
			return true
		}
	}
	return false
}

func ReplaceAll(s string, old string, new string) string {
	s2 := ""
	r1 := []rune(s)
	if old == "" {
		for i, r := range r1 {
			if i > 0 {
				s2 += new + string(r)
			} else {
				s2 += string(r)
			}
		}
		return new + s2 + new
	}
	r2 := []rune(old)
	for i := 0; i < len(r1); i++ {
		if i < len(r1)-len(r2) && isRunesEqual(r1[i:i+len(r2)], r2) {
			s2 += new
			i += len(r2) - 1
		} else {
			s2 += string(r1[i])
		}
	}
	return s2
}

func Repeat(s string, count int) string {
	if s == "" || count < 1 {
		return ""
	}
	s2 := ""
	for i := 0; i < count; i++ {
		s2 += s
	}
	return s2
}

func Trim(s string, cutset string) string {
	if cutset == "" || s == "" {
		return s
	}
	trimLeft := func(s string, cutset string) (string, bool) {
		r1 := []rune(s)
		r2 := []rune(cutset)
		if len(r1) >= len(r2) && isRunesEqual(r1[0:len(r2)], r2) {
			return string(r1[len(r2):]), true
		}
		return s, false
	}
	trimRight := func(s string, cutset string) (string, bool) {
		r1 := []rune(s)
		r2 := []rune(cutset)
		if len(r1) >= len(r2) && isRunesEqual(r1[len(r1)-len(r2):], r2) {
			return string(r1[:len(r1)-len(r2)]), true
		}
		return s, false
	}
	var ok bool
	for {
		s, ok = trimLeft(s, cutset)
		if !ok {
			break
		}
	}
	for {
		s, ok = trimRight(s, cutset)
		if !ok {
			break
		}
	}
	return s
}

func Join(elems []string, sep string) string {
	s2 := ""
	for i, s := range elems {
		if i > 0 {
			s2 += sep + s
		} else {
			s2 += s
		}
	}
	return s2
}

func Split(s string, sep string) []string {
	if s == "" && sep == "" {
		return nil
	}
	if s == "" {
		return []string{s}
	}
	result := make([]string, 0)
	r1 := []rune(s)
	if sep == "" {
		for _, r := range r1 {
			result = append(result, string(r))
		}
		return result
	}
	if len(s) < len(sep) {
		return []string{s}
	}
	r2 := []rune(sep)
	last := 0
	for i := 0; i <= len(r1)-len(r2); i++ {
		if isRunesEqual(r1[i:i+len(r2)], r2) {
			result = append(result, string(r1[last:i]))
			i += len(r2) - 1
			last = i + 1
		}
	}
	result = append(result, string(r1[last:]))
	return result
}
