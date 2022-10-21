package match

import (
	"testing"
)

func Test_IsTime(t *testing.T) {
	for _, time := range tTime {
		if IsTime(time) != true {
			t.Errorf("%s is not a true time:", time)

		}
	}

	for _, time := range fTime {
		if IsTime(time) != false {
			t.Errorf("%s is not a false time", time)
		}
	}
}

var tTime = []string{
	"09:45",
	"9:45",
	"23:59:59",
	"00:00:00",
	"12:12:12",
	"00:56:56",
	"12:13:12",
	"12:12:01",
	"23:45",
	"9:00am",
	"9:00 A.M.",
	"9:00 pm",
}

var fTime = []string{
	"19:45 pm",
	"19:45 am",
	"9:60",
	"pm23:45",
	"9:00:am",
	"9am",
	"19:00 A.M.",
	"9:00:pm",
}

/*
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
*/
