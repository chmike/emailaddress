package emailaddress

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"unicode/utf8"

	"github.com/chmike/domain"
)

// Check returns an error if the email address is not syntactically valid.
func Check(email string) error {
	// Reference: RFC822 and https://en.wikipedia.org/wiki/Email_address
	for i := 0; i < len(email); i++ {
		if email[i] < ' ' || email[i] == 0x7F {
			return fmt.Errorf("invalid character at %d", utf8.RuneCountInString(email[:i]))
		}
	}
	// skip left comments
	for email != "" && email[0] == '(' {
		pos := strings.IndexByte(email, ')')
		if pos == -1 {
			return errors.New("unclosed comment in front of local part")
		}
		email = email[pos+1:]
	}
	if email == "" {
		return errors.New("email is empty")
	}
	// extract local part from email
	var localPart string
	var pos int
	if email[0] == '"' {
		pos = -1
		for i := 1; i < len(email); i++ {
			if email[i] == '"' {
				pos = i + 1
				break
			}
			if email[i] == '\\' && i < len(email) && (email[i+1] == '\\' || email[i+1] == '"') {
				i++
			}
		}
		if pos == -1 {
			return errors.New("unclosed double quoted local part")
		}
	} else if pos = strings.IndexAny(email, "(@"); pos == -1 {
		pos = len(email)
	}
	localPart = email[:pos]
	email = email[pos:]
	// skip right most comments if any
	for email != "" && email[0] == '(' {
		pos := strings.IndexByte(email, ')')
		if pos == -1 {
			return errors.New("unclosed comment at end of local part")
		}
		email = email[pos+1:]
	}
	if email != "" && email[0] != '@' {
		c, _ := utf8.DecodeRuneInString(email)
		return fmt.Errorf("expected '@' after local part, got '%s'", string(c))
	}
	// check domains
	for email != "" && email[0] == '@' {
		email = email[1:]
		pos := strings.IndexByte(email, '@')
		if pos == -1 {
			pos = len(email)
		}
		domainPart := email[:pos]
		email = email[pos:]
		if err := domain.Check(domainPart); err != nil {
			return fmt.Errorf("domain %s: %w", domainPart, err)
		}
	}
	// check local part
	email = localPart
	if email == "" {
		return errors.New("local part is empty")
	}
	if len(email) > 64 {
		return fmt.Errorf("expect local part length < 65, got %d", len(email))
	}
	if email[0] != '"' {
		if email[0] == '.' || localPart[len(email)-1] == '.' {
			return errors.New("local part can't start or end with '.'")
		}
		if strings.Contains(email, "..") {
			return errors.New("local part can't contain '..'")
		}
		for _, c := range localPart {
			if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' || c > 127 ||
				strings.IndexByte("!#$%&'*+-/=?^_`{|}~.", byte(c)) != -1) {
				return fmt.Errorf("invalid character '%s' in local part", string(c))
			}
		}
	}
	return nil
}

// CheckWithDNS checks the email validity and also checks that the domain after the last @
// is an existing domain that acceptes mails.
func CheckWithDNS(email string) error {
	if err := Check(email); err != nil {
		return err
	}
	pos := strings.LastIndexByte(email, '@')
	if pos == -1 {
		return errors.New("email without domain")
	}
	if _, err := net.LookupMX(email[pos+1:]); err != nil {
		return err
	}
	return nil
}
