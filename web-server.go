package main

import (
	"context"
	"html/template"
	"net/http"

	cmtlog "github.com/cometbft/cometbft/libs/log"
	nm "github.com/cometbft/cometbft/node"
)

var nodeTemplate *template.Template

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
	node     *nm.Node
}

type NodeData struct {
	NodeID string
	Status string
}

func NewWebServer(app *KVStoreApplication, httpPort string, logger cmtlog.Logger, node *nm.Node) *WebServer {
	mux := http.NewServeMux()

	server := &WebServer{
		app:      app,
		httpAddr: ":" + httpPort,
		server: &http.Server{
			Addr:    ":" + httpPort,
			Handler: mux,
		},
		logger: logger,
		node:   node,
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
	nodeTemplate, err := template.ParseFiles("templates/node.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data := NodeData{
		NodeID: string(webserver.node.NodeInfo().ID()),
		Status: "online",
	}
	err = nodeTemplate.ExecuteTemplate(w, "node", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
