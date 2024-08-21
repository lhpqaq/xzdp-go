package utils

import (
	"regexp"
)

func ValidateMobile(phone string) bool {
	var re = regexp.MustCompile(`^1[3-9]\d{9}$`)
	return re.MatchString(phone)
}
