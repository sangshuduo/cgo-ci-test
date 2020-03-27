package main

import (
	. "fmt"
	"github.com/sangshuduo/cgo_ci_helloworld/calculation"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/taosdata/driver-go/taosSql"
)

func main() {
	Println("This is an example GO project uses CGO, go module, and some popular CI services and unit test frameworks.")

	Println("add new import taosSql to call CGO module")

	Println(calculation.Division(8, 2))
}
