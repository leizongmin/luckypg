package nostd

func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	const radix = 10
	nums := []rune("0123456789")
	sign := ""
	if i < 0 {
		sign = "-"
		i = -i
	}
	result := ""
	for i > 0 {
		rem := i % radix
		result = string(nums[rem]) + result
		i = i / radix
	}
	return sign + result
}
