package main

import (
	"context"
	"net/http"

	cmtlog "github.com/cometbft/cometbft/libs/log"
)

type QueryRequest struct {
	Key string `json:"key"`
}
type QueryResponse struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Error string `json:"error,omitempty"`
}

type WebServer struct {
	app      *KVStoreApplication
	httpAddr string
	server   *http.Server
	logger   cmtlog.Logger
}

func NewWebServer(app *KVStoreApplication, httpPort string, logger cmtlog.Logger) *WebServer {
	mux := http.NewServeMux()

	server := &WebServer{
		app:      app,
		httpAddr: ":" + httpPort,
		server: &http.Server{
			Addr:    ":" + httpPort,
			Handler: mux,
		},
		logger: logger,
	}

	mux.HandleFunc("/", server.handleRoot)

	return server
}

func (webserver *WebServer) Start() error {
	webserver.logger.Info("Starting HTTP web server", "addr", webserver.httpAddr)
	go func() {
		if err := webserver.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			webserver.logger.Error("HTTP server error: ", "err", err)
		}
	}()
	return nil
}

func (webserver *WebServer) Shutdown(ctx context.Context) error {
	webserver.logger.Info("Shutting down web server")
	return webserver.server.Shutdown(ctx)
}

func (webserver *WebServer) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	html := "<h1>Hello from CometBFT Node Web Server</h1>"
	w.Write([]byte(html))
}
