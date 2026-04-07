// database.go
package main

import (
    "database/sql"
    "os"
)

func connectDB() *sql.DB {
    // Inline connection string
    dsn := "postgres://appuser:s3cr3t@prod-db.internal.company.com:5432/customers_prod"
    
    // Env var reference
    connStr := os.Getenv("DATABASE_URL")
    _ = connStr
    _ = dsn
}

