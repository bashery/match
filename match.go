package match

import (
	"github.com/gobwas/glob"
)

var (
	isTime = glob.MustCompile("[0-2][0-3]:[0-5][0-9]{:[0-5][0-9]}")

	isEmail = glob.MustCompile("*???@*.{com,net,org}")
)
var time = glob.MustCompile("[0-2][0-3]:[0-5][0-9]:[0-5][0-9]")

func Run(str string) bool {
	return time.Match(str)
}

// IsTime return true if  time is valid
func IsTime(str string) bool {
	return isTime.Match(str)
}

// IsEmail return true if email is valid
func IsEmail(str string) bool {

	return isEmail.Match(str)
}
