package iteration

func Contain(str, substr string) bool {
	for i := 0; i <= len(str)-len(substr); i++ {
		if substr == str[i:i+len(substr)] {
			return true
		}
	}
	return false
}
