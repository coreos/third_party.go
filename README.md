# build.go - self container Go build helper


build.go is a self contained build system that handles bumping third party repos, gopath and building

## Usage

### Setup a repo

```
go run build.go --setup github.com/philips/build.go/example
git add .
git commit -m "Initial commit"
```

### Build the project

```
go build --package github.com/philips/build.go/example
```
