package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func getEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}

func main() {
	// --help command line option
	helpFlag := flag.Bool("help", false, "display help message")
	flag.Parse()

	if *helpFlag {
		fmt.Println(`This program requires the following environment variables to be set:
- PGHOST: Database host
- PGPORT: Database port
- PGDATABASE: Database name
- PGUSER: Database user
- PGPASSWORD: Database password
- PGSSLMODE: Database SSL mode
- DEBUG: (optional) Set to "true" to print debug messages`)
		return
	}

	waiting := true
	for waiting {
		fmt.Println("Trying to connect to DB...")
		connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s connect_timeout=3",
			getEnvOrFail("PGHOST"), getEnvOrFail("PGPORT"), getEnvOrFail("PGDATABASE"), getEnvOrFail("PGUSER"), getEnvOrFail("PGPASSWORD"), getEnvOrFail("PGSSLMODE"))

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Failed to open a DB connection: %v", err)
		}

		if err := db.Ping(); err != nil {
			fmt.Println("Waiting on DB...")
			if os.Getenv("DEBUG") == "true" {
				fmt.Println(err)
			}
			time.Sleep(2 * time.Second)
		} else {
			waiting = false
			db.Close()
		}
	}

	fmt.Println("DB responded!")
}

