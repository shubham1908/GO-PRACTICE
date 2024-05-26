# All about Go Vet

Vet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

Vet is normally invoked through the go command. This command vets the package in the current directory:
```
go vet
```
whereas this one vets the packages whose path is provided:
```
go vet my/project/...
```

> [!Note]
> Vet's exit code is non-zero for erroneous invocation of the tool or if a problem was reported, and 0 otherwise. Note that the tool does not check every possible problem and depends on unreliable heuristics, so it should be used as guidance only, not as a firm indicator of program correctness.

To list the available checks, run `go tool vet help`:

```
appends          check for missing values after append
asmdecl          report mismatches between assembly files and Go declarations
assign           check for useless assignments
atomic           check for common mistakes using the sync/atomic package
bools            check for common mistakes involving boolean operators
buildtag         check //go:build and // +build directives
cgocall          detect some violations of the cgo pointer passing rules
composites       check for unkeyed composite literals
copylocks        check for locks erroneously passed by value
defers           report common mistakes in defer statements
directive        check Go toolchain directives such as //go:debug
errorsas         report passing non-pointer or non-error values to errors.As
framepointer     report assembly that clobbers the frame pointer before saving it
httpresponse     check for mistakes using HTTP responses
ifaceassert      detect impossible interface-to-interface type assertions
loopclosure      check references to loop variables from within nested functions
lostcancel       check cancel func returned by context.WithCancel is called
nilfunc          check for useless comparisons between functions and nil
printf           check consistency of Printf format strings and arguments
shift            check for shifts that equal or exceed the width of the integer
sigchanyzer      check for unbuffered channel of os.Signal
slog             check for invalid structured logging calls
stdmethods       check signature of methods of well-known interfaces
stringintconv    check for string(int) conversions
structtag        check that struct field tags conform to reflect.StructTag.Get
testinggoroutine report calls to (*testing.T).Fatal from goroutines started by a test
tests            check for common mistaken usages of tests and examples
timeformat       check for calls of (time.Time).Format or time.Parse with 2006-02-01
unmarshal        report passing non-pointer or non-interface values to unmarshal
unreachable      check for unreachable code
unsafeptr        check for invalid conversions of uintptr to unsafe.Pointer
unusedresult     check for unused results of calls to some functions
```

For details and flags of a particular check, such as printf, run `go tool vet help printf`.

> [!NOTE]
> By default, all checks are performed. If any flags are explicitly set to true, only those tests are run. Conversely, if any flag is explicitly set to false, only those tests are disabled. Thus -printf=true runs the printf check, and -printf=false runs all checks except the printf check.

### Example

Running `go vet` on the pkg_vet
```console
➜  test_go_internals git:(main) ✗ go vet ./pkg_vet 
# test_internals/pkg_vet
# [test_internals/pkg_vet]
pkg_vet/vet.go:6:9: fmt.Sprintf format %d has arg "ram" of wrong type string
```

## `go vet` on `go test`
As part of building a test binary, go test runs go vet on the package and its test source files to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary. Only a high-confidence subset of the default go vet checks are used. That subset is: `atomic, bool, buildtags, directive, errorsas, ifaceassert, nilfunc, printf, and stringintconv`. You can see the documentation for these and other vet tests via `go doc cmd/vet`. To disable the running of go vet, use the `-vet=off` flag. To run all checks, use the `-vet=all` flag.

### Example

* Running vet_test.go file
```console
➜  test_go_internals git:(main) ✗ go test test_internals/pkg_vet    
# test_internals/pkg_vet
# [test_internals/pkg_vet]
pkg_vet/vet.go:6:9: fmt.Sprintf format %d has arg "ram" of wrong type string
FAIL    test_internals/pkg_vet [build failed]
FAIL
```

* Running with vet check turning off
```console
➜  test_go_internals git:(main) ✗ go test test_internals/pkg_vet --vet=off  
--- FAIL: TestPrintCheckVet (0.00s)
    --- FAIL: TestPrintCheckVet/testing_with_go_vet (0.00s)
        vet_test.go:12: 
                Error Trace:    /Users/shubham/Coding/GO-PRACTICE/test_go_internals/pkg_vet/vet_test.go:12
                Error:          Not equal: 
                                expected: "ram"
                                actual  : "%!d(string=ram)"
                            
                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -ram
                                +%!d(string=ram)
                Test:           TestPrintCheckVet/testing_with_go_vet
FAIL
FAIL    test_internals/pkg_vet  0.631s
FAIL
```