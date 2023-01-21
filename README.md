# Contest

Boilerplate for testing solutions for coding contests

## What is this?

1. You love to solve tasks from coding contests
2. You write a solution for a task
3. You want to test your solution against some inputs
4. You don't want to manually run your program and enter all the data after each code change
5. Download this boilerplate
6. Write your code in `main.go`
7. Put your inputs and expected outputs in `/tests`
8. Run `make`

```bash
$ make
go test -v ./...
=== RUN   Test_run
=== RUN   Test_run/1
--- PASS: Test_run (0.00s)
--- PASS: Test_run/1 (0.00s)
PASS
ok      contest 0.001s
```