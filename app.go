package main

// App - Application
type App struct {
	Port         string
	DataFileName string
}

// NewApp - create new instance
func NewApp(port, dataFileName string) *App {
	a := &App{Port: port, DataFileName: dataFileName}
	return a
}
