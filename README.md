# path.go - self contained GOPATH helper

path.go is a self contained single file that helps with bumping third party repos and managing the GOPATH

## Usage

### Setup a repo

```
go run path.go setup github.com/philips/build.go/example
git add .
git commit -m "Initial commit"
```

### Build the project

```
go run path.go build -v github.com/philips/build.go/example
```

## Test

```
go test path_test.go
```
