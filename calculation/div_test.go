package calculation

import (
	"fmt"
	"testing"
)

func Test_Division(t *testing.T) {
	if i, e := Division(10, 5); i != 2 || e != nil {
		t.Error("Fail to pass")
	} else {
		t.Log("Pass")
	}
}

func Test_Division_Zero(t *testing.T) {
	if _, e := Division(10, 0); e == nil {
		t.Error("Should not go here.")
	} else {
		t.Log("Check zero correctly.")
	}
}

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(10, 2)
	}
}

func Example_Division() {
	fmt.Println(Division(10, 2))
	// Output: 5 <nil>

}
