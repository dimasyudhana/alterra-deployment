package helper

// isPasswordMatched membandingkan password dan confirm_password
func IsPasswordMatched(password string, confirm_password string) bool {
	return password == confirm_password
}
