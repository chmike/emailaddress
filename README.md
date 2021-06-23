# emailaddress.Check()

This package contains a single function that checks the syntactic validity of an email address.

An email address name must respect rules presented in [https://en.wikipedia.org/wiki/Email_address](https://en.wikipedia.org/wiki/Email_address).

The `emailaddress.Check` function ensures that the email address respect those rules. If not, it returns an error explaining the detected problem.

## Prerequisites

The package has no prerequisites. It depends on github.com/chmike/domain@v1.0.0 to verify domain names.

## Installation

To install or update this package use the instruction:

```bash
go get -u "github.com/chmike/emailaddress"
```

## Usage examples

The `Check` function can be used to check the validity of host or domain names.

```go
import "github.com/chmike/emailaddress"

. . . 

emailAddress := "foo@example.com"
if err := emailaddress.Check(emailAddress); if err != nil {
    log.Fatalf("invalid email address '%s': %v", emailAddress, err)
}
```
