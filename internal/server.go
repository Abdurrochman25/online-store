package internal

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Abdurrochman25/online-store/internal/config"
	"github.com/Abdurrochman25/online-store/internal/middleware"
	"github.com/Abdurrochman25/online-store/internal/router"
	"github.com/Abdurrochman25/online-store/pkg/common"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() {
	conf := config.NewConfig()

	s := config.NewServer(conf)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	if err := s.NewDB(ctx); err != nil {
		cancel()
		log.Fatalf("Failed to initialize database; error: %v", err)
	}
	cancel()

	s.Fiber = fiber.New(fiber.Config{
		Immutable: true,
	})

	s.Fiber.Use(logger.New())

	s.Fiber.Use(middleware.RequestIDMiddleware())

	s.Fiber.Use(middleware.DefaultHeaderAcceptLanguage())

	// Setting Localize
	s.Fiber.Use(middleware.LocalizeMiddleware())

	// Setting Validator
	s.Fiber.Use(middleware.ValidatorMiddleware())

	// Setting Authentication
	customMiddleware := middleware.NewAuthConfig(&conf)
	s.Fiber.Use(customMiddleware.CustomAuthentication())

	// Setting Authorization
	s.Authorization = middleware.NewAuthzConfig().CustomAuthorization()

	s.Router = &config.Router{
		Routes: []fiber.Router{
			s.Fiber.Get("/", s.Authorization.RequiresPermissions([]string{"welcome:read"}),
				func(c *fiber.Ctx) error {
					return c.SendString(fiberi18n.MustLocalize(c, "http.200_read"))
				}),
		},
	}

	router.AttachAllRoutes(s)

	// Custom 404 Handler
	s.Fiber.Use(func(c *fiber.Ctx) error {
		return common.NewHTTPHandler().NotFound(c, "http.404_url_not_found")
	})
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

		log.Infof("Shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Warn("Timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
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

				log.Info("Cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Warn("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Info("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
