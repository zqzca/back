package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/zqzca/back/controllers"
	"github.com/zqzca/back/controllers/chunks"
	"github.com/zqzca/back/controllers/dashboard"
	"github.com/zqzca/back/controllers/files"
	"github.com/zqzca/back/controllers/p2p"
	"github.com/zqzca/back/controllers/sessions"
	"github.com/zqzca/back/controllers/thumbnails"
	"github.com/zqzca/back/controllers/users"
	"github.com/zqzca/back/lib"
	"github.com/zqzca/back/scp"
	"github.com/zqzca/echo"

	"github.com/rsc/letsencrypt"
	"github.com/zqzca/echo/engine"
	"github.com/zqzca/echo/engine/standard"
	"github.com/zqzca/echo/middleware"
)

//----------
// Handlers
//----------

// func sshServer() {
// 	s := scp.NewSCPServer()
// 	s.ListenAndServe()
// }

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func redirect() {
	http.ListenAndServe(":3001", http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			http.Redirect(w, req, "https://"+req.Host+req.RequestURI, http.StatusMovedPermanently)
		},
	))
}

var livereload *bool

func main() {
	secure := flag.Bool("secure", false, "Enable HTTPS")
	livereload = flag.Bool("livereload", false, "Enable Live Reload")
	cdnURL := flag.String("cdn", "", "CDN URL")

	fmt.Println(cdnURL)

	flag.Parse()

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} - ${latency_human}, rx=${rx_bytes}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	// servers other static files
	// e.ServeDir("/assets", "assets")
	// e.ServeFile("/", "assets/index.html")
	// e.ServeFile("/favicon.ico", "assets/favicon.ico")

	// api := e.Group("/api")
	v1 := e.Group("/api/v1")

	// Route
	// e.Get("/chunk/status", controllers.ChunkStatus)

	db, err := lib.Connect()

	if err != nil {
		fmt.Printf("Failed to connect to db")
		return
	}

	log := logrus.New()
	log.Level = logrus.DebugLevel
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{}

	deps := controllers.Dependencies{
		Fs:     afero.NewOsFs(),
		Logger: log,
		DB:     db,
	}

	e.Get("/", dashboard.AppIndex)

	// Files
	files := &files.FileController{deps}
	e.Get("/d/:slug", files.Download)
	v1.Get("/check/:hash", files.Status)
	v1.Get("/files", files.Index)
	v1.Get("/files/:slug", files.Read)
	v1.Get("/files/:slug/data", files.Download)
	v1.Post("/files", files.Create)
	v1.Get("/files/:slug/process", files.Process)

	// Thumbnail
	thumbnails := thumbnails.ThumbnailsController{deps}
	v1.Get("/thumbnails/:id", thumbnails.Download)

	// Chunks
	chunks := chunks.ChunkController{deps}
	v1.Post("/files/:file_id/chunks/:chunk_id/:hash", chunks.Write)

	// Users
	users := users.UsersController{deps}
	v1.Post("/users", users.Create)
	v1.Get("/username/valid", users.ValidateUsername)
	v1.Get("/users/:id", users.Read)

	// Sessions
	sessions := sessions.SessionsController{deps}
	v1.Post("/sessions", sessions.Create)

	// P2P
	v1.Get("/p2p/signaling", standard.WrapHandler(http.HandlerFunc(p2p.Signaling())))
	v1.Get("/p2p/:id", p2p.Join)
	v1.Post("/p2p/:id", p2p.Answer)

	// Dashboard
	dash := dashboard.DashboardController{deps}
	v1.Get("/dashboard", dash.Index)

	// r := api.Group("/users")
	// r.Use(JWTAuth())
	// r.Get("/:id", controllers.UserGet)
	// e.Patch("/users/:id", updateUser)
	// e.Delete("/users/:id", deleteUser)

	// e.ServeFile("/signin", "assets/signin.html")
	// e.ServeFile("/*", "assets/index.html")
	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:   "public",
	// 	Browse: false,
	// 	Index:  "index.html",
	// 	HTML5:  true,
	// }))

	e.Static("/assets", "assets")
	e.Get("/*", dashboard.AppIndex)

	var s *standard.Server

	bindAddr := ":3001"

	if *secure == true {
		fmt.Println("Running in Secure Mode")
		var m letsencrypt.Manager
		if err := m.CacheFile("certs/letsencrypt.cache"); err != nil {
			log.Fatal(err)
		}

		cfg := &tls.Config{
			GetCertificate: m.GetCertificate,
		}

		config := engine.Config{
			Address:   bindAddr,
			TLSConfig: cfg,
		}

		s = standard.WithConfig(config)
		http2.ConfigureServer(s.Server, &http2.Server{})
		deps.Info("Listening for HTTP2 connections", "addr", bindAddr)
		go redirect()
	} else {
		deps.Info("Listening for HTTP connections", "addr", bindAddr)
		s = standard.New(bindAddr)
	}

	scpServer := scp.SCPServer{}
	scpServer.DB = deps.DB
	scpServer.Logger = deps.Logger
	scpServer.CertPath = "certs/scp.rsa"

	go scpServer.ListenAndServe()

	// Start server
	e.Run(s)
}
