package match

import (
	"strconv"
	"strings"

	"github.com/gobwas/glob"
)

var (
	isEmail = glob.MustCompile("*????@*.{com,net,org}")

	isTime = glob.MustCompile("{[0-2],}[0-9]:[0-5][0-9]{:[0-5][0-9],,am, A.M., pm")
)

func IsTime(str string) bool {
	if strings.HasSuffix(str, "m") || strings.HasSuffix(str, "M.") {
		hour, err := strconv.Atoi(strings.Split(str, ":")[0])
		if err != nil || hour > 11 {
			return false
		}
	}
	return isTime.Match(str)
}

// IsEmail return true if email is valid
func IsEmail(str string) bool {
	return isEmail.Match(str)
}
