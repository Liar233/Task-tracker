package app

type App struct {
	config *ApplicationConfig
}

func (app *App) Bootstrap() error {

	return nil
}

func (app *App) Run() error {

	return nil
}

func (app *App) Stop() error {

	return nil
}

func NewApp(config *ApplicationConfig) *App {

	return &App{
		config: config,
	}
}
