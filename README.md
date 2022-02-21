# Go Crash Course

## References
- [Learnerbly Resource: (Video) Go / Goland crash course](https://app.learnerbly.com/resources/817ae5eb-2792-4e39-a70a-1fc910f49e8a/) also available on [Traversy Media (creator's) Youtube channel](https://www.youtube.com/watch?v=SqrbIlUwR0U)
- [How to write Go Code](https://go.dev/doc/code)

## Errors

> main.go:7:2: no required module provides package github.com/avin-dsilva-bazaarvoice/go_crash_course/03_packages/strutil: go.mod file not found in current directory or any parent directory; see 'go help modules'   

https://stackoverflow.com/questions/66894200/error-message-go-go-mod-file-not-found-in-current-directory-or-any-parent-dire
```bash
go env -w GO111MODULE=off
```
