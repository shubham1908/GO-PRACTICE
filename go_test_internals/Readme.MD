# Test Commands Docs

'Go test' automates testing the packages named by the import paths. It prints a summary of the test results in the format:
```
        ok   archive/tar   0.011s
        FAIL archive/zip   0.022s
        ok   compress/gzip 0.033s
        ...
```
followed by detailed output for each failed package.

'Go test' recompiles each package along with any files with names matching the file pattern "*_test.go". These additional files can contain test functions, benchmark functions, fuzz tests and example functions. 

* Each listed package causes the execution of a separate test binary.
* **Files whose names begin with "_" (including "_test.go") or "." are ignored.**
* Test files that declare a package with the suffix "_test" will be compiled as a separate package, and then linked and run with the main test binary.
* **The go tool will ignore a directory named "testdata", making it available to hold ancillary data needed by the tests.**
* As part of building a test binary, go test runs go vet on the package and its test source files to identify significant problems. If go vet finds any problems, go test reports those and does not run the test binary. Only a high-confidence subset of the default go vet checks are used. That subset is: atomic, bool, buildtags, directive, errorsas, ifaceassert, nilfunc, printf, and stringintconv.
* To disable the running of go vet, use the `-vet=off` flag. To run all checks, use the `-vet=all` flag.

#### Go test runs in two different modes
* The first, called local directory mode, occurs when go test is invoked with no package arguments (for example, 'go test' or 'go test -v'). In this mode, go test compiles the package sources and tests found in the current directory and then runs the resulting test binary. In this mode, caching (discussed below) is disabled. After the package test finishes, go test prints a summary line showing the test status ('ok' or 'FAIL'), package name, and elapsed time.

* The second, called package list mode, occurs when go test is invoked with explicit package arguments (for example 'go test math', 'go test ./...', and even 'go test .'). In this mode, go test compiles and tests each of the packages listed on the command line. If a package test passes, go test prints only the final 'ok' summary line. If a package test fails, go test prints the full test output. If invoked with the -bench or -v flag, go test prints the full output even for passing package tests, in order to display the requested benchmark results or verbose logging. After the package tests for all of the listed packages finish, and their output is printed, go test prints a final 'FAIL' status if any package test has failed.

> * In package list mode only, go test caches successful package test results to avoid unnecessary repeated running of tests. When the result of a test can be recovered from the cache, go test will redisplay the previous output instead of running the test binary again. When this happens, go test prints '(cached)' in place of the elapsed time in the summary line.
> * To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1.

### Running "test print string, assertion failed" test
```console
➜  go_test_internals git:(main) ✗ go test                                                              
2024/05/26 22:39:53 printing something
--- FAIL: TestAdd (0.00s)
    --- FAIL: TestAdd/test_print_string,_assertion_failed (0.00s)
        print_string_test.go:13: 
                Error Trace:    /Users/shubham/Coding/GO-PRACTICE/go_test_internals/print_string_test.go:13
                Error:          Not equal: 
                                expected: "test passed"
                                actual  : ""
                            
                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -test passed
                                +
                Test:           TestAdd/test_print_string,_assertion_failed
FAIL
exit status 1
FAIL    test_internals  0.643s
```

### Running "test addition of 2 numbers, assertion success" test
* Running normally
```console
➜  go_test_internals git:(main) ✗ go test 
2024/05/26 22:41:19 printing something
PASS
ok      test_internals  0.852s
```
* Result will not be cached because it is mode 1 (note that it has printed the log)

```console
➜  go_test_internals git:(main) ✗ go test 
2024/05/26 22:49:06 printing something
PASS
ok      test_internals  0.172s
```

### Running tests from another pkg 
```console
➜  go_test_internals git:(main) ✗ go test ./pkg
ok      test_internals/pkg      0.634s
```
* Running again will give the cached result (note that it hasn't printed the log)

```console
➜  go_test_internals git:(main) ✗ go test ./pkg
ok      test_internals/pkg      (cached)
```

* Running with `-v` flag
> Verbose output: log all tests as they are run. Also print all text from `Log` and `Logf` calls even if the test succeeds.

* mode 1
```console
➜  go_test_internals git:(main) ✗ go test -v
=== RUN   TestAdd
=== RUN   TestAdd/test_print_string,_assertion_success
2024/05/26 22:51:16 printing something
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/test_print_string,_assertion_success (0.00s)
PASS
ok      test_internals  0.363s
```
 Result will not be cached
```console
➜  go_test_internals git:(main) ✗ go test -v
=== RUN   TestAdd
=== RUN   TestAdd/test_print_string,_assertion_success
2024/05/26 22:51:33 printing something
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/test_print_string,_assertion_success (0.00s)
PASS
ok      test_internals  0.145s
```

* mode 2 
```console
➜  go_test_internals git:(main) ✗ go test ./pkg -v
=== RUN   TestAdd
=== RUN   TestAdd/test_print_string,_assertion_success
2024/05/26 22:52:33 printing something
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/test_print_string,_assertion_success (0.00s)
PASS
ok      test_internals/pkg      0.299s
```
Result will be cached
```console
=== RUN   TestAdd
=== RUN   TestAdd/test_print_string,_assertion_success
2024/05/26 22:52:33 printing something
--- PASS: TestAdd (0.00s)
    --- PASS: TestAdd/test_print_string,_assertion_success (0.00s)
PASS
ok      test_internals/pkg      (cached)
```

### Go test with flags

#### -coverprofile cover.out
Write a coverage profile to the file after all tests have passed.

* Running with test - "test print string, assertion success"
```console
➜  go_test_internals git:(main) ✗ go test . -coverprofile cover.out
ok      test_internals  0.614s  coverage: 50.0% of statements
```
cover.out file content - 
```console
mode: set
test_internals/main.go:5.13,8.2 2 0
test_internals/print_string.go:8.48,9.14 1 1
test_internals/print_string.go:9.14,11.3 1 0
test_internals/print_string.go:12.2,13.27 2 1
```

* Running with multiple pkgs
```console
➜  go_test_internals git:(main) ✗ go test . ./pkg -coverprofile cover.out 
ok      test_internals  0.386s  coverage: 50.0% of statements
ok      test_internals/pkg      0.891s  coverage: 75.0% of statements
```
cover.out file content - 
```console
mode: set
test_internals/main.go:5.13,8.2 2 0
test_internals/print_string.go:8.48,9.14 1 1
test_internals/print_string.go:9.14,11.3 1 0
test_internals/print_string.go:12.2,13.27 2 1
test_internals/pkg/print_string_pkg.go:8.55,9.14 1 1
test_internals/pkg/print_string_pkg.go:9.14,11.3 1 0
test_internals/pkg/print_string_pkg.go:12.2,13.27 2 1
```

> [!Note]
> To view the code covered and not covered in the browser, you can use the `go tool cover -html=cover.out` command.

#### -cpuprofile cpu.out
Write a CPU profile to the specified file before exiting.
Writes test binary as -c would.

> will dig it later

#### -memprofile mem.out
Write an allocation profile to the file after all tests have passed.           Writes test binary as -c would.

> will dig it later

#### -c flag
used to compile the test binary without actually running the tests. This can be useful if you want to manually run the test binary or inspect it before execution. When you use the -c flag, the resulting test binary will be named after the package being tested by default.
If you want to specify the output file for the test binary, you can use the -o flag.

```console
➜  go_test_internals git:(main) ✗ go test . -c -o tests.test 
```
it will create a `tests.test` file

Running this file
```console
➜  go_test_internals git:(main) ✗ ./tests.test               
2024/05/26 23:07:07 printing something
PASS
```

if running this command with multiple pkgs -o must refer to a directory
```console
➜  go_test_internals git:(main) ✗ go test . ./pkg -c -o ./tests
with multiple packages, -o must refer to a directory or /dev/null
➜  go_test_internals git:(main) ✗ go test . ./pkg -c -o ./tests
```
it will create two files in the diretory tests
- `tests/pkg.test`
- `tests/test_internals.test`

#### -count flag
When 'go test' runs in package list mode, 'go test' caches successful package test results to avoid unnecessary repeated running of tests. To disable test caching, use any test flag or argument other than the cacheable flags. The idiomatic way to disable test caching explicitly is to use -count=1.

```console
➜  go_test_internals git:(main) ✗ go test ./pkg
ok      test_internals/pkg      0.622s
➜  go_test_internals git:(main) ✗ go test ./pkg                          
ok      test_internals/pkg      (cached)
➜  go_test_internals git:(main) ✗ go test ./pkg -count=1               
ok      test_internals/pkg      0.731s
➜  go_test_internals git:(main) ✗ go test ./pkg -count=1
ok      test_internals/pkg      0.118s
```

### go clean -testcache
used to remove the cached result of tests.
```console
➜  go_test_internals git:(main) ✗ go test ./pkg         
ok      test_internals/pkg      (cached)
➜  go_test_internals git:(main) ✗ go clean -testcache
➜  go_test_internals git:(main) ✗ go test ./pkg      
ok      test_internals/pkg      0.121s
```