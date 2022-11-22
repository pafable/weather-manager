# Weather Manager

The purpose of this project is to explore creating commandline tools using Go.

```
go run main.go temp --degree <int> 
```

```
go run main.go weather --location <string>
```

### Building a binary

MacOS arm64

```
GOOS=darwin GOARCH=arm64 go build -o wm
```

MacOS x86_64

```
GOOS=darwin GOARCH=amd64 go build -o wm
```

Linux x86_64

```
GOOS=linux GOARCH=amd64 go build -o wm
```

Windows x86_64

```
GOOS=windows GOARCH=amd64 go build -o wm
```