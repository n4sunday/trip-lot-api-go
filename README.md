### Installation module

To tidy the Go module dependencies:

```bash
go mod tidy
```

### Complies

To compile and run the main application:

```bash
go run cmd/main.go
```

### Complies Hot Reload

To use Air for hot reloading:

### Installation

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

Add the following alias to your .zshrc (or equivalent) to use Air conveniently:

```
alias air='$(go env GOPATH)/bin/air
```

### Docker
Create local database
```bash
docker-compose up -d
```