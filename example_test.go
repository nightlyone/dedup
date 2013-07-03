package dedup

import (
	"fmt"
)

func ExampleVisit() {
	dups := new(Seen)
	seen := dups.Visit("hello")
	fmt.Println(seen, dups.VisitedKeys())
	seen = dups.Visit("world")
	fmt.Println(seen, dups.VisitedKeys())
	seen = dups.Visit("world")
	fmt.Println(seen, dups.VisitedKeys())
	// Output: false [hello]
	// false [hello world]
	// true [hello world]
}
