package main

import "test_internals/pkg"

func main() {
	_, _ = funcToTest(false)
	_, _ = pkg.FuncToTestFromPkg(false)
}
