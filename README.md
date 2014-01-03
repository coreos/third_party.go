# path.go

path.go is a self contained single file that helps with bumping third party repos and managing the GOPATH.

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

### Bumping a package

```
go run path.go bump github.com/goraft/raft
```

### Bumping all packages

This will walk all of the packages installed in `third_party/src` and run the
`bump` subcommand on it if it is a valid go pkg.

```
go run path.go bump-all
```

## Testing path.go

```
./test
```
