# GOMAIL

Gomail will send you a mail with Content-Type: text/plain using gomail.

## Pre Request

Please enable [less secure app access](https://myaccount.google.com/lesssecureapps) from your account

## Installation

Use the [go get](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) to install gomail.

```bash
go get gopkg.in/mail.v2
```

## Config

Create a config.go with fromAdd, password, and toAdd 

```go
// const.go
package main

// Sender data.
const fromAdd = "from@gmail.com"
const password = "pass"

// To Addresses
var toAdd = "to@gmail.com"
}
```