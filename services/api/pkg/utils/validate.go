package utils

import "regexp"

var (
	mobileRegexString = "^1\\d{10}$"
	emailRegexString  = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
)

var (
	emailRegex  = regexp.MustCompile(emailRegexString)
	mobileRegex = regexp.MustCompile(mobileRegexString)
)

// ValidateIsEmail 验证是否邮箱
func ValidateIsEmail(s string) bool {
	return emailRegex.MatchString(s)
}

// ValidateIsMobile 验证是否手机号
func ValidateIsMobile(s string) bool {
	return mobileRegex.MatchString(s)
}
