package luhn

import (
	"fmt"
	"testing"
)

func Test_ToNumber(t *testing.T) {
	if toNumber('0') != 0 {
		t.Fail()
	}
	if toNumber('1') != 1 {
		t.Fail()
	}
	if toNumber('2') != 2 {
		t.Fail()
	}
	if toNumber('3') != 3 {
		t.Fail()
	}
	if toNumber('4') != 4 {
		t.Fail()
	}
	if toNumber('5') != 5 {
		t.Fail()
	}
	if toNumber('6') != 6 {
		t.Fail()
	}
	if toNumber('7') != 7 {
		t.Fail()
	}
	if toNumber('8') != 8 {
		t.Fail()
	}
	if toNumber('9') != 9 {
		t.Fail()
	}
}

func Test_GenAndValidate(t *testing.T) {
	for i := 0; i < 100; i++ {
		var id, err = Gen("8", 10)
		if err != nil {
			t.Log(fmt.Sprintf("%d: fail generate", i))
			t.Fail()
		}
		if len(id) != 10 {
			t.Log(fmt.Sprintf("%d: not equal to 10", i))
			t.Fail()
		}
		s1 := string(id[0])
		s2 := string('8')
		if s1 != s2 {
			t.Log(fmt.Sprintf("%d: starting digit not correct", i))
			t.Fail()
		}

		if !Validate(id) {
			t.Log(fmt.Sprintf("%d: validate fail", i))
			t.Fail()
		}
	}
}
