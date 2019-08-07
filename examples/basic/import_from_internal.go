package main

import (
    internal1 "github.com/aaronize/grandonion/internal"
    //internal2 "github.com/aaronize/grandonion/examples/servers/internal"  // Use of internal package is not allowed
)

// An import of a path containing the element “internal” is disallowed
// if the importing code is outside the tree rooted at the parent of the “internal” directory.

// 即，internal是其所在目录的内部内容，只能在internal所在目录中使用。
// 在本例中，
// 只有grandonion/下的包可以import grandonion/internal包及其子包的内容，因此在本文件中可以import；
// grandonion/examples/servers/internal包中的内容无法被引用。

func main() {
    internal1.MyPrint()

    //internal2.MyPrint2()
}
