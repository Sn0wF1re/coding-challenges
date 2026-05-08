
# Go Challenges

This folder contains Go implementations used for personal practice and learning. Expect experiments, drafts, and short notes alongside runnable examples. See the top-level `README.md` for repository-wide purpose and conventions.

Layout

```text
golang/
├─ README.md
├─ go.mod
└─ <challenge-name>/
   ├─ source files
   ├─ test files
   ├─ testdata/
   └─ README.md
```


Current challenge folders

- `ccwc-tool`

Existing challenge: `ccwc-tool`

`ccwc-tool` is a small `wc`-style command line tool implemented in Go. Files include:

- `ccwc-tool/ccwc.go`
- `ccwc-tool/test.txt`

Supported options include `-c` (byte count), `-l` (line count), `-w` (word count), and `-m` (rune/character count). Without flags it prints lines, words, and bytes.

Run examples (from the `golang` folder)

```bash
go run ./ccwc-tool/ccwc.go test.txt
go run ./ccwc-tool/ccwc.go -l test.txt
```

Development

- Format: `go fmt ./...`
- Test: `go test ./...`
- Module is declared in `go.mod` — update module path if you fork the repo.

Add a new Go challenge

1. Create a new challenge folder in kebab-case.
2. Use `package main` and keep a clear `main.go` or other entry points for runnable examples.
3. Add tests (`*_test.go`) and a brief `README.md` showing how to run and test the challenge.

Example

```text
golang/
└─ prime-sieve/
   ├─ main.go
   ├─ main_test.go
   ├─ testdata/
   └─ README.md
```
