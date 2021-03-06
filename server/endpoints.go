// Package server serves up the endpoints cache via http
package server

import (
	"encoding/json"
	"io"
	"net/http"

	"fmt"

	"github.com/VEVO/kubernetes-pod-discovery/cache"
)

// Endpoints interface that specifies the endpoint routes available for our http server
type Endpoints interface {
	Root(http.ResponseWriter, *http.Request)
	LastUpdated(http.ResponseWriter, *http.Request)
}

// EndpointsServer is the object that we use to store access to the cache through our endpoints routes
type EndpointsServer struct {
	cache *cache.Endpoints
}

// NewEndpointsServer creates a new endpoints server and points to the specified cache
func NewEndpointsServer(endpointsCache *cache.Endpoints) Endpoints {
	return &EndpointsServer{
		cache: endpointsCache,
	}
}

// Root serves our root endpoints route
func (e *EndpointsServer) Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	endpoints, err := json.Marshal(*e.cache.GetEndpoints())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(endpoints))
}

// LastUpdated serves our last_updated endpoints route
func (e *EndpointsServer) LastUpdated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lastUpdated := e.cache.GetLastUpdated()
	response := fmt.Sprintf("{\"lastUpdated\": \"%s\"}", lastUpdated.UTC().Format("2006-01-02T15:04:05-0700"))
	io.WriteString(w, string(response))
}
