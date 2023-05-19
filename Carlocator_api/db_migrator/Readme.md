## Migrations

This is an internal tool to do the migrations on the production and on the development. The latest dump would be in the [schema.sql](../schema.sql) 

### Config 
Loads the database credentials from the root folder i-e `Carlocator_api` and load the .env inside the tool. The migrations are inside [Carlocator_api/migrations](../migrations) 


### Command
```
$ pwd
  /development/carlocator/Carlocator_api
$ go run db_migrator/main.go
  -count int
        Number of migration to run
  -migration string
        up - For doing up migration, down - For doing down migration
```

#### Up migration
For doing the up migrations
```
$ go run db_migrator/main.go --migration=up count=1
```

#### Down migration
For doing the down migration
```
$ go run db_migrator/main.go --migration=down count=1
```