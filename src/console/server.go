package console

import (
	"go-notif/src/config"
	"go-notif/src/controller"
	"go-notif/src/database"
	"go-notif/src/repository"
	"go-notif/src/router"
	"go-notif/src/service"
	"go-notif/src/subscribe"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  "Start running the server",
	Run:   server,
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func server(cmd *cobra.Command, args []string) {
	// Initiate DB
	MysqlDB := database.InitDB()
	db, err := MysqlDB.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	redisConn := database.NewRedisConn(config.RedisHost())
	defer redisConn.Close()

	// Create Echo instance
	httpServer := echo.New()
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(middleware.CORS())

	// Depedency Injection
	notificationRepository := repository.NewNotificationRepository(MysqlDB)
	notificationSubscribe := subscribe.NewNotificationSubscribe(redisConn, notificationRepository)
	notificationService := service.NewNotificationService(redisConn)
	notificationController := controller.NewNotificationController(notificationService)

	router.RouteService(httpServer.Group("/api"), notificationController)

	// Graceful Shutdown
	// Catch Signal
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)
		defer close(sigChan)

		<-sigChan
		log.Info("Received termination signal, initiating graceful shutdown...")
		cancel()
	}()

	// Start http server
	go func() {
		log.Info("Starting server...")
		if err := httpServer.Start(":" + config.Port()); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting the server: %v", err)
		}
	}()

	// Start Redis Subscriber
	go func() {
		log.Info("Starting Redis subscriber...")
		notificationSubscribe.ReceiveNotif() // Start the Redis subscriber loop
	}()

	// Shutting down any connection and server
	<-ctx.Done()
	log.Info("Shutting down server...")

	notificationSubscribe.Close()

	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Info("Server gracefully shut down")
}
