// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reverse_test

import (
	"fmt"

	"hello/reverse"
)

func ExampleString() {
	fmt.Println(reverse.String("hello"))
	fmt.Println(reverse.String("world"))
	fmt.Println(reverse.String("!"))
	// Output: olleh
	// dlrow
	// !
}
