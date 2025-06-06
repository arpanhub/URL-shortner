package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5/pgxpool"
)
var DB *pgxpool.Pool
 func ConnectDB(){
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbHost == "" || dbUser == "" || dbPass == "" || dbName == "" || dbSSLMode == "" {
		log.Fatal("Missing required database environment variables!")
	}
	Db_string := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbName, dbSSLMode)
	log.Println("Connecting to DB...")

	config, db_err := pgxpool.ParseConfig(Db_string);
	if db_err != nil{
		log.Fatal("Error parsing db string:",db_err)
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.HealthCheckPeriod = 5 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB,err = pgxpool.New(ctx,Db_string);
	if err != nil{
		log.Fatal("Error connecting to db:",err)
	}
	err = DB.Ping(ctx)
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}
	log.Println("✅ Successfully connected to PostgreSQL")

 }