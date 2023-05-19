package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	migrator "capregsoft.com/carlocator/db_migrator/migrator"

	"capregsoft.com/carlocator/db_migrator/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %s", err)
	}
	file, err := os.OpenFile("schema.sql", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	cfg.Writer = file
	dns := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.DbUser, cfg.DbPassword, cfg.Host, cfg.Port, cfg.DbName, cfg.SslMode)
	log.Println(dns)
	conn, err := sqlx.Connect("postgres", dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		log.Println("Connected")
	}
	defer conn.Close()
	migrations(conn)

	dumpErr := migrator.Dump(&cfg)
	if dumpErr != nil {
		log.Fatal(dumpErr)
	}
}

func migrations(conn *sqlx.DB) {
	migration := flag.String("migration", "", "up - For doing up migration, down - For doing down migration")
	count := flag.Int("count", 0, "Number of migrations to run")
	flag.Parse()
	if *count <= 0 {
		*count++
	}
	migrator := migrator.ProvideMigrator(conn.DB, *migration, *count)
	_, err := migrator.RunMigrations()
	if err != nil {
		log.Fatal(err)
	}
}
