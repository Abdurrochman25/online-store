package cmd

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/router"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Config config.Config
	DB     *sql.DB
	Fiber  *fiber.App
	Router *router.Router
}

func NewServer(config config.Config) *Server {
	return &Server{
		Config: config,
		DB:     nil,
		Fiber:  nil,
	}
}

func (s *Server) NewDB(ctx context.Context) error {
	db, err := sql.Open("postgres", s.Config.Database.ConnectionString())
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	s.DB = db

	log.Println("database successfully connected")
	return nil
}

func Init() {
	conf := config.NewConfig()

	s := NewServer(conf)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err := s.NewDB(ctx); err != nil {
		cancel()
		log.Fatalf("Failed to initialize database; error: %v", err)
	}
	cancel()

	s.Fiber = fiber.New(fiber.Config{
		Immutable: true,
	})

	s.Router = &router.Router{
		Routes: []fiber.Router{
			s.Fiber.Get("/", func(c *fiber.Ctx) error {
				return c.SendString("Hello World!")
			}),
		},
	}

	log.Fatal(s.Fiber.Listen(":3000"))
}

// Operation is a cleanup function on shutting down
type Operation func(ctx context.Context) error

// Shutdown waits for termination syscalls and doing clean up operations after received it
func Shutdown(ctx context.Context, timeout time.Duration, ops map[string]Operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("Shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("Timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("Cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
