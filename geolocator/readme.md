# SMTP

SMTP will send you a mail with Content-Type: text/plain.


## Config

Create a config.go with fromAdd, password, and toAdd 

```go
// const.go
package main

// DBName Database name.
const DBName = "geolocation"

// Database collections.
var (
	PointCollection = "points"
	connString      = "mongodb+srv://<username>:<password>@cluster0-zzart.mongodb.net/test?retryWrites=true&w=majority"
)

```