package match

import (
	"testing"
)

func TestIsTime(t *testing.T) {
	for _, time := range trueTime {
		if IsTime(time) != true {
			t.Errorf("%s is not a time", time)

		}
	}
	for _, time := range falseTime {
		if IsTime(time) != false {
			t.Errorf("%s is not a time", time)
		}
	}
}

func BenchmarkIsTime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, time := range trueTime {
			IsTime(time)

		}
		for _, time := range falseTime {
			IsTime(time)
		}
	}
}

func TestIsEmail(t *testing.T) {
	for _, email := range trueEmails {
		if IsEmail(email) != true {
			t.Errorf("%s is not a email", email)

		}
	}

	for _, email := range falseEmails {
		if IsEmail(email) != false {
			t.Errorf("%s is not a email", email)
		}
	}
}
