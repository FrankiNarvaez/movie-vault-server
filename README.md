# Server of movie vault

## Dependencies

```bash
go get -u github.com/gin-gonic/gin
go get github.com/jmoiron/sqlx
go get github.com/lib/pq
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
