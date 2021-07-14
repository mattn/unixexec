# unixexec

Same as os/exec but does not look current directory for Windows.

## Usage

```go
p, err := unixexec.LookPath("git")
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
