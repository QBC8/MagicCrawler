package main

import (
	"context"
	"fmt"
	"github.com/QBC8-Team7/MagicCrawler/pkg/db/sqlc"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/QBC8-Team7/MagicCrawler/config"
	"github.com/QBC8-Team7/MagicCrawler/internal/server"
	"github.com/QBC8-Team7/MagicCrawler/pkg/db"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("could not read config file: ", err)
	}

	ctx := context.Background()

	dbUri := db.GetDbUri(cfg)
	dbConn, err := db.GetDBConnection(ctx, dbUri)

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal("could not close connection:", err)
		}
	}(dbConn, ctx)

	dbQueries := sqlc.New(dbConn)

	s := server.NewServer(cfg, dbQueries)

	go func() {
		fmt.Println("Bot Server Started...")
		s.Serve()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down bot...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	s.Bot.Stop()

	<-ctx.Done()

	s.Logger.Info("Bot exited gracefully.")

}
