[![GoDoc](https://img.shields.io/badge/go.dev-reference-blue)](https://pkg.go.dev/github.com/chmike/emailaddress)
![Build](https://github.com/chmike/emailaddress/actions/workflows/workflow.yaml/badge.svg)
[![codecov](https://codecov.io/gh/chmike/emailaddress/branch/master/graph/badge.svg?token=9XNNVJXV1E)](https://codecov.io/gh/chmike/emailaddress)
[![Go Report](https://goreportcard.com/badge/github.com/chmike/emailaddress)](https://goreportcard.com/report/github.com/chmike/emailaddress)
![Status](https://img.shields.io/badge/status-stable-brightgreen.svg)
![release](https://img.shields.io/github/release/chmike/emailaddress/all.svg)

# emailaddress.Check()

This package contains two functions to check email addresses. 

An email address name must respect rules presented in [https://en.wikipedia.org/wiki/Email_address](https://en.wikipedia.org/wiki/Email_address).

The `emailaddress.Check` function ensures that the email address respect those rules. If not, it returns an error explaining the detected problem.
This function accepts emails without a domain like "root" which is valid on unix systems.

The `emailaddress.CheckWithDNS()` calls `Check` and also checks that the domain after the last @ is an existing domain accepting emails. With 
this call the email "root" is invalid since it has no domain accepting emails.

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
