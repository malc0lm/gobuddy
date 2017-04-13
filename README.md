# gobuddy - An simple buddy memory allocator for Go

The buddy memory allocation technique is a memory allocation algorithm that divides memory into partitions to try to satisfy a memory request as suitably as possible. This system makes use of splitting memory into halves to try to give a best-fit.

This library is non-invasive. It just manage offset for your continuous memory block.

## Example Usage

```go
package main

import (
  "github.com/malc0lm/gobuddy"
  "fmt"
)

func main() {

	b := gobuddy.NewBuddySystem(16)
	a := b.Alloc(5)

	c := b.Alloc(3)
	b.Free(a)
	a = b.Alloc(9)
  // If there is no appropriate memory for allocation Alloc will return -1.
	fmt.Println(a)
	fmt.Println(c)
}
```
