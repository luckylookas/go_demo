# go demo for TwinGo

##prerequisites

* have go 1.11+ installed
* set env variables:
    * GOROOT to the directory of your go runtime
    * GOPATH to any other directory (go tools you install will be installed here, as well as any go code you write without using modules)
* run ```go mod tidy``` in the ```go_demo``` directory

##ide

* IntelliJ has a golang Plugin that works reasonably well, but is discontinued because
* Jetbrains published GoLand as dedicated ide for Go
* Visual Studio Code also supports Go

##running, building, testing

- ```go run [any_neccessary_files.go,...]``` will run main.main() found in the files supplied
- ```go build [any_neccessary_files.go,...]``` will build a binary for your architecture
- ```env GOOS=linux | GOARCH=arm  GOarm=5 go build [any_neccessary_files.go,...]``` will build a binary for linux | running on arm v5
- ```env GOOS=linux | GOARCH=amd64 | go build [any_neccessary_files.go,...]``` will build a binary for linux | running on amd64 |
- ```go o test [-bench=.]``` will run any tests [and benchmarks] in the current directory (go recognizes tests by the ```_test``` suffix)

##supported architectures

| GOOS | GOARCH |
| ----|-----|
|android | arm |
|darwin | 386 |
|darwin |	amd64 |
|darwin |	arm |
|darwin |	arm 64 |
|dragonfly |	amd64 |
|freebsd |	386 |
|freebsd |	amd64 |
|freebsd |	arm |
|linux |	386 |
|linux |	amd64 |
|linux |	arm |
|linux |	arm 64 |
|linux |	ppc64 |
|linux |	ppc64le |
|linux |	mips |
|linux |	mipsle |
|linux |	mips64 |
|linux |	mips64le |
|netbsd |	386 |
|netbsd |	amd64 |
|netbsd |	arm |
|openbsd |	386 |
|openbsd |	amd64 |
|openbsd |	arm |
|plan9 |	386 |
|plan9 |	amd64 |
|solaris |	amd64 |
|windows |	386 |
|windows |	amd64 |

##initializing your own projects

```
$> mkdir myproject
$> go mod init mymodulename
```

then simply create .go files, and import github dependencies.
```go run | build | mod tidy | mod download``` will download any added dependencies and pin their version to the current ones.

