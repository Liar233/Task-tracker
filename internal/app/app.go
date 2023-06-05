package app

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Liar233/Task-tracker/internal/app/server"
	"github.com/Liar233/Task-tracker/internal/app/server/action"
	"github.com/Liar233/Task-tracker/internal/app/storage"
	"github.com/gorilla/mux"
)

type App struct {
	config  *ApplicationConfig
	taskRep storage.TaskRepositoryInterface
	server  server.HttpServerAdapterInterface
}

func (app *App) Bootstrap() {

	app.taskRep = storage.NewTaskPostgresRepository(
		app.config.DBHost,
		app.config.DBName,
		app.config.DBUser,
		app.config.DBPassword,
		app.config.DBPort,
	)

	app.server = server.NewHttpServerAdapter(app.config.HttpHost, app.config.HttpPort)

	router := mux.NewRouter()

	router.Use(server.ProtocolMiddleware)
	router.Handle("/exec", action.NewExecAction(app.taskRep)).Methods(http.MethodPost)

	app.server.SetHandler(router)
}

func (app *App) Run() (exitErr error) {

	var err error

	if err = app.taskRep.Open(); err != nil {

		return err
	}

	go func() {

		if err = app.server.ListenAndServe(); err != nil {

			exitErr = err

			return
		}
	}()

	return app.Stop()
}

func (app *App) Stop() error {

	sigintChan := make(chan os.Signal, 1)

	defer close(sigintChan)

	signal.Notify(sigintChan, syscall.SIGINT, syscall.SIGTERM)

	_ = <-sigintChan

	app.server.Close()

	return app.taskRep.Close()
}

func NewApp(config *ApplicationConfig) *App {

	return &App{
		config: config,
	}
}
