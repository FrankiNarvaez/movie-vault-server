# Server of movie vault

## Dependencies

```bash
go get github.com/gin-gonic/gin
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

## Script

To execute the script, in root path execute:

```bash
go run cmd/script [options]
```

### options

- `-m | --migrate` Create tables
- `-s | --seed` Insert data
- `-h | --help` Show help menu
