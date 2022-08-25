package main

import (
	"app/log"
	"app/registry"
	"app/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")              //run the log
	host, port := "localhost", "4000" //hardcode the host and port
	serviceAdress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAdress

	ctx, err := service.Start( //call start the service
		context.Background(),
		host, port,
		r,
		log.RegisterHandlers,
	)
	if err != nil { //handle errors
		stlog.Fatal(err)
	}
	<-ctx.Done()                                 //when the ctx is done, it means it's been closed. Time to move on !!
	fmt.Println("Shutting down Log Service Yo!") //move on msg ;)
}
