# cloud-file-storage


## Migrations
**Library** - golang-migrate/migrate/v4  
**Location** - /migrations  
**Name pattern** - `<version>_<name>.<up/down>.sql`  
**Name example** - `001_init.down.sql`

### Action:
```bash copy 
go run cmd/migrate/main.go up
```

```bash copy 
go run cmd/migrate/main.go down
```