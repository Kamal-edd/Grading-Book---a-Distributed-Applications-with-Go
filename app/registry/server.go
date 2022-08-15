package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type ServiceName string

type Registration struct {
	ServiceName ServiceName
	ServiceURL  string
}

const (
	LogService = ServiceName("LogService")
)

const ServerPort = ":3000"
const ServiceURL = "http://localhost" + ServerPort + "/services"

type registry struct { //this struct type will contain a slice of registrations with a mutex
	//so that we can controle writing so it
	registration []Registration
	mutex        *sync.Mutex
}

func (r *registry) add(reg Registration) error { //this method will attempt to write to a
	r.mutex.Lock()
	r.registration = append(r.registration, reg)
	r.mutex.Unlock()
	return nil
}

var reg = registry{registration: make([]Registration, 0), mutex: new(sync.Mutex)}

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println("Request Recieved")
	switch req.Method {
	case http.MethodPost:
		dec := json.NewDecoder(req.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
