# Lesson 1

Three small programs, each with its own `main`, so run them one file at a time from this folder. You'll need Go 1.27 or newer (`go version` to check).

```sh
go run hello.go
# Hello, World!

go run greetings_manual.go Oleksandr
# Hello, Oleksandr

go run greetings_ai.go Oleksandr
# Hello, Oleksandr!
```

Running `go run .` here won't work: Go complains that `main` is redeclared, since all three files share the same package.

`greetings_manual.go` is the hand-written version, `greetings_ai.go` is the AI-assisted one. They do the same job, but the AI one is more defensive: it checks that a name was actually passed and, if not, prints a usage hint to stderr and exits with a non-zero code. The manual one just grabs `os.Args[1]` and panics with an index-out-of-range error when it isn't there. The AI code is also a bit more idiomatic: `username :=` instead of the spelled-out `var name string =`, and `Printf` instead of gluing strings together with `+`. On the flip side, it pulls in `path/filepath` just to show the program name in the usage line, which is more than a hello-world really needs. So the manual one is the minimum that works, and the AI one is closer to what you'd actually ship.

```sh
go run greetings_manual.go
# panic: runtime error: index out of range [1] with length 1

go run greetings_ai.go
# usage: greetings_ai <username>
```

Screenshots from the IDE and the Killercoda playground are in `screenshots/`.
