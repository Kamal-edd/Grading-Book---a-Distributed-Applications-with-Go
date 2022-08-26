package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/Kamal-edd/Grading-Book-app/grades"
	"github.com/Kamal-edd/Grading-Book-app/log"
	"github.com/Kamal-edd/Grading-Book-app/registry"
	"github.com/Kamal-edd/Grading-Book-app/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	} else {
		stlog.Println(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
