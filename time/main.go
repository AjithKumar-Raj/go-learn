package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	x := time.Now()

	fmt.Println(reflect.TypeOf(x), x)
}
