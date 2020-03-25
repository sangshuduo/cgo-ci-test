package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/taosdata/driver-go/taosSql"
)

func main() {
	fmt.Println("This is an example GO project uses CGO, go module, and some popular CI services and unit test frameworks.")

	fmt.Println("add new import taosSql to call CGO module")
}
