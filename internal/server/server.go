package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/joho/godotenv/autoload"

	"portfolio/internal/database"
	"portfolio/internal/models"
)

type Server struct {
	port int

	db             database.Service
	sessionManager *scs.SessionManager
	models         models.Models
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(NewServer.db.GetDB())
	NewServer.sessionManager = sessionManager
	NewServer.models = models.NewModels(NewServer.db.GetDB())

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
