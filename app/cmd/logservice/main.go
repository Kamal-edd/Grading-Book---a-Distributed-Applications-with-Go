package main

import (
	"app/log"
	"app/service"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")              //run the log
	host, port := "localhost", "4000" //hardcode the host and port
	ctx, err := service.Start(        //call start the service
		context.Background(),
		"Log Service Yo!",
		host, port,
		log.RegisterHandlers,
	)
	if err != nil { //handle errors
		stlog.Fatal(err)
	}
	<-ctx.Done()                                 //when the ctx is done, it means it's been closed. Time to move on !!
	fmt.Println("Shutting down Log Service Yo!") //move on msg ;)
}
