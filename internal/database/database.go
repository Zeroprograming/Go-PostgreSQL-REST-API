package database

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/zeroprograming/go-postgresql-rest-api/internal/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global variable representing the database connection.
var DB *gorm.DB

// DBConnection establishes a connection to the database using the configuration provided in the environment variables.
func DBConnection() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construct the Data Source Name (DSN) from environment variables
	DSN := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT")

	// Configure GORM options
	dbConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,  // Threshold for slow queries
				LogLevel:                  logger.Error, // Set log level to Error
				IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for the logger
				Colorful:                  false,        // Disable color
			},
		),
	}

	// Connect to the database
	DB, err = gorm.Open(postgres.Open(DSN), dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate models
	DB.AutoMigrate(models.User{}, models.Task{})

	log.Println("Database connected")
}
