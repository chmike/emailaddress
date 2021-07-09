package emailaddress

import (
	"strings"
	"testing"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		s  string
		ok bool
	}{
		// 0
		{"", false},
		{"foo\x15@example.com", false},
		{"(comment foo@example.com", false},
		{"(comment )", false},
		{`"foo \\ \" `, false},
		// 5
		{`"foo \\ \""@`, false},
		{"foo@", false},
		{"foo(right comment", false},
		{"foo(right comment)", true},
		{"root", true},
		// 10
		{"foo(right comment)@example.com", true},
		{"foo@example.com@gmail.com", true},
		{`"root"x`, false},
		{"@example.com", false},
		{strings.Repeat("a", 64), true},
		// 15
		{strings.Repeat("a", 65), false},
		{"root.", false},
		{".root", false},
		{"ro..ot", false},
		{"root<", false},
		// 20
		{"{f_o-o}@example.com", true},
		{"0123045@example.com@example.com", true},
		{"0123045@example@example.com", true},
	}
	for i, test := range tests {
		err := Check(test.s)
		if err != nil && test.ok {
			t.Errorf("%3d expected nil error, got %s", i, err)
		} else if err == nil && !test.ok {
			t.Errorf("%3d expected error, got nil", i)
		}
	}
}

func TestCheckWithDNS(t *testing.T) {
	if err := CheckWithDNS(""); err == nil {
		t.Error("unexpected nil error")
	}
	if err := CheckWithDNS("toto"); err == nil {
		t.Error("unexpected nil error")
	}
	if err := CheckWithDNS("toto@example.com"); err == nil {
		t.Error("unexpected nil error")
	}

	if err := CheckWithDNS("toto@gmail.com"); err != nil {
		t.Error("unexpected error:", err)
	}
}
