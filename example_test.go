package dedup

import (
	"fmt"
)

func ExampleVisit() {
	dups := new(Seen)
	seen := dups.Visit("hello")
	fmt.Println(seen, dups.Known("hello"), dups.VisitedKeys())
	fmt.Println(seen, dups.Known("world"), dups.VisitedKeys())
	seen = dups.Visit("world")
	fmt.Println(seen, dups.Known("world"), dups.VisitedKeys())
	seen = dups.Visit("world")
	fmt.Println(seen, dups.Known("world"), dups.VisitedKeys())
	// Output: false true [hello]
	// false false [hello]
	// false true [hello world]
	// true true [hello world]
}
