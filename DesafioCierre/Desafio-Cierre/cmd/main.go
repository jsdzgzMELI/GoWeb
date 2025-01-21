package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/handler"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/repository"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/repository/loader"
	"github.com/jsdzgzMELI/Desafio-Cierre/internal/service"
)

func main() {
	// env
	os.Setenv("DB_FILE", "/Users/judiazgutier/Documents/GoWeb/DesafioCierre/Desafio-Cierre/tickets.csv")
	// ...

	// application
	// - config
	cfg := &ConfigAppDefault{
		ServerAddr: os.Getenv("SERVER_ADDR"),
		DbFile:     os.Getenv("DB_FILE"),
	}
	app := NewApplicationDefault(cfg)

	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// ConfigAppDefault represents the configuration of the default application
type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db := loader.NewLoaderTicketCSV(a.dbFile)
	// if err != nil {
	// 	return
	// }
	repo, lastId, err := db.Load()
	if err != nil {
		return err
	}

	rp := repository.NewRepositoryTicketMap(repo, lastId)
	// service ...
	ps := service.NewServiceTicketDefault(rp)
	// handler ...
	hd := handler.NewHandlerTicketDefault(ps)
	// handler ...
	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})
	(*a).rt.Get("/ticket", hd.GetHttp)
	(*a).rt.Get("/ticket/{id}", hd.GetByIdHttp)
	(*a).rt.Get("/ticket/getByCountry/{country}", hd.GetCountryHttp)
	(*a).rt.Get("/ticket/getAverage/{country}", hd.GetProportionHttp)
	(*a).rt.Post("/ticket", hd.AddHttp)
	(*a).rt.Delete("/ticket/{id}", hd.DeleteHttp)
	(*a).rt.Patch("/ticket/{id}", hd.PatchHttp)
	(*a).rt.Put("/ticket/{id}", hd.UpdateHttp)

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
