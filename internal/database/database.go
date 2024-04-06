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

var DB *gorm.DB

func DBConnection() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Construir el DSN (Data Source Name) a partir de las variables de entorno
	DSN := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT")

	// Configurar las opciones del GORM
	dbConfig := &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Umbral para las consultas lentas
				LogLevel:                  logger.Info, // Establece el nivel de log a Info
				IgnoreRecordNotFoundError: true,        // Ignorar el error ErrRecordNotFound para el logger
				Colorful:                  false,       // Desactiva el color
			},
		),
	}

	// Conectar a la base de datos
	DB, err = gorm.Open(postgres.Open(DSN), dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB.AutoMigrate(models.User{}, models.Task{})

	log.Println("Database connected")
}
