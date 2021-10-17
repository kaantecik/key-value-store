package main

import (
	"fmt"
	"github.com/kaantecik/key-value-store/internal/logging"
	"net"
	"net/http"
	"os"

	"github.com/kaantecik/key-value-store/internal/config"
	"github.com/kaantecik/key-value-store/internal/entities"
	"github.com/kaantecik/key-value-store/internal/router"
)

func main() {
	host, hostExist := os.LookupEnv("HOST")
	if !hostExist {
		host = config.DefaultHost
	}

	port, portExist := os.LookupEnv("PORT")
	if !portExist {
		port = config.DefaultPort
	}

	listenAddr := net.JoinHostPort(host, port)

	c := entities.NewCache(&entities.CacheOptions{})

	http.Handle("/api/cache/set", router.Set(c))
	http.Handle("/api/cache/get", router.Get(c))
	http.Handle("/api/cache/flush", router.Flush(c))

	fmt.Printf("Listening on %s. Check %s directory for logs.", listenAddr, config.LogPath)
	logging.HttpLogger.Infof("Listening on %s", listenAddr)
	logging.HttpLogger.Fatalln(http.ListenAndServe(listenAddr, nil))
}
