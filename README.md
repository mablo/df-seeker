# Duplicate files seeker
## Usage
```shell script
$ df-seeker
```
```shell script
$ df-seeker --help
  -p string
        Directory path. (default ".")
  -r    Recursive
  -rlimit uint
        Percent usage of soft ulimit. (default 90)
  -so string
        Sort order (asc, desc). (default "asc")
  -sp string
        Sort parameter (hash, size). (default "size")
```

## Build
```shell script
$ go build -o df-seeker main.go
```
