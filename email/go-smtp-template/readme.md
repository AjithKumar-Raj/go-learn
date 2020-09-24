# GO SMTP TEMPLATE

Go SMTP temple will send you a mail with Content-Type: text/html.

## Pre Request

Please enable [less secure app access](https://myaccount.google.com/lesssecureapps) from your account

## Config

Create a config.go with fromAdd, password, and toAdd 

```go
// const.go
package main

// Sender data.
const fromAdd = "from@gmail.com"
const password = "pass"

// To Addresses
var toAdd = []string{
	"to@gmail.com",
}
```