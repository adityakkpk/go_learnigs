# go_learnigs
This repo contains all the code files which I have write while learning GO Lang

## Day-1
- Setting UP the Go environment and Installation
- To Run Go program in any of the folder outside we need to initialize `Go module` first using command:
```bash
go mod init <module_name>
```
- Some more commands: 
    1. For dependencies download 
    2. For all the packages/dependencies download
    3. Build a file (Cross Platform Executable Generator)
    3. Build a folder having main.go for a specific OS
```bash
go get <package_name>

go mod download

go build <file_name>

GOOS=windows GOARCH=amd64 go build .
```

- Some conventions about package names:
    - Package name should be in small case
    - Package names have to be as short as possible

- Functions (some of the example):
    
```
func add(x int, y int) int {
	return x + y
}

func add(x, y int) int {
    return x + y
}
```

- Next i am following this documentation : [Go](https://go.dev/tour)
