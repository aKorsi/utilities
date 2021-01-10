package utilities

func IranNationalCodeValidator(input string) bool {
	for i := range input {
		if input[i] < '0' || input[i] > '9' {
			return false
		}
	}
	check := int(input[9] - '0')
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(input[i]-'0') * (10 - i)
	}
	sum %= 11
	return (sum < 2 && check == sum) || (sum >= 2 && check+sum == 11)
}
