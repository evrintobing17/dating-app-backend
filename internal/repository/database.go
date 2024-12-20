package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/evrintobing17/dating-app-go/config"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sql.DB
}

func NewDatabase(cfg config.DatabaseConfig) (*Database, error) {
	Dbdriver := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open(Dbdriver, dsn)
	if err != nil {
		fmt.Println("Cannot connect to database")
		log.Fatal("Error: ", err)
		return nil, err
	}

	log.Printf("We are connected to the %s database\n", Dbdriver)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

func (db *Database) Close() error {
	return db.Conn.Close()
}

func RunMigrations(cfg config.DatabaseConfig) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Directory containing migration files
	// migrationDir := "./migrations"

	// files, err := ioutil.ReadDir(migrationDir)
	// if err != nil {
	// 	return fmt.Errorf("error reading migration directory: %v", err)
	// }

	// for _, file := range files {
	// 	if filepath.Ext(file.Name()) == ".sql" {
	// 		migrationPath := filepath.Join(migrationDir, file.Name())
	// 		query, err := ioutil.ReadFile(migrationPath)
	// 		if err != nil {
	// 			return fmt.Errorf("error reading migration file %s: %v", file.Name(), err)
	// 		}

	// 		log.Printf("Executing migration: %s", file.Name())
	// 		log.Printf("Executing query: %s", query)
	// 		if _, err := db.Exec(string(query)); err != nil {
	// 			return fmt.Errorf("error executing migration %s: %v", file.Name(), err)
	// 		}
	// 	}
	// }

	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            email VARCHAR(100) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            is_premium BOOLEAN DEFAULT FALSE
        );`,
		`CREATE TABLE IF NOT EXISTS swipes (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            profile_id INT NOT NULL,
            action VARCHAR(10) CHECK (action IN ('like', 'pass')),
            swiped_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
		`CREATE TABLE IF NOT EXISTS premium_purchases (
            id SERIAL PRIMARY KEY,
            user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            purchased_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`,
	}
	for _, query := range queries {
		log.Printf("Executing query: %s", query)
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("error executing query: %v", err)
		}
	}

	log.Println("Migrations completed successfully.")

	return nil
}
