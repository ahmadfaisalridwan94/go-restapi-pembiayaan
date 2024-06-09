# Migration:

#### Make Migration Table:

```
migrate create -ext sql -dir src/database/migrations -seq create_tablename_table
```

#### Using Migration

```
migrate -path src/database/migrations -database mysql://username:password@tcp(localhost:3306)/pembiayaan up
```

#### Rollback Migration

```
migrate -path db/migrations -database mysql://username:password@tcp(localhost:3306)/mydatabase down 1
```
