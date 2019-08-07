package bpkg

import _ "unsafe"

/*
//go:linkname local_name path/to/package.name 的含义，
把对path/to/package.name的引用映射到local_name上。本例中，当外部引用github.com/aaronize/.../apkg.Say时，
会调用本包（bpkg）下的sayIt函数。
 */

//go:linkname sayIt github.com/aaronize/grandonion/examples/basic/compile_directive/go_linkname/apkg.Say
func sayIt(name string) string {
	return "hello, " + name
}
