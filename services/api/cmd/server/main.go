package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	connectcors "connectrpc.com/cors"
	_ "github.com/lib/pq"
	"github.com/rs/cors"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/arthur-fontaine/culty/services/api/gen/proto/media/v1/mediav1connect"
	"github.com/arthur-fontaine/culty/services/api/internal/repository/postgres"
	"github.com/arthur-fontaine/culty/services/api/internal/usecase/mediaservice"
)

func main() {
	mux := http.NewServeMux()

	db, err := sql.Open("postgres", os.Getenv("PSQL_DSN"))
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	mediaRepository := postgres.NewMediaRepository(db)
	mediaService := mediaservice.NewMediaService(mediaRepository)
	mediaPath, mediaHandler := mediav1connect.NewMediaServiceHandler(mediaService)
	mux.Handle(mediaPath, mediaHandler)

	log.Println("Listening on :8080")
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		withCORS((h2c.NewHandler(mux, &http2.Server{}))),
	)
}

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}
