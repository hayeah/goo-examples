```
go run golang.org/x/tools/cmd/gonew@latest \
     github.com/hayeah/goo-examples/cli github.com/hayeah/hello
```

# NOTE

GOPROXY may cache the old version of the module, use GOPROXY=direct to bypass the cache:

```
GOPROXY=direct go install github.com/hayeah/goo-examples/cli@latest
```

Or remove the cache:

```
rm -rf `go env GOPATH`/pkg/mod/github.com/hayeah/goo-examples
```
