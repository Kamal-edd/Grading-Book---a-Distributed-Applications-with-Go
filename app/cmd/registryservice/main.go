package main

import (
	"app/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/sevices", &registry.RegistryService{}) //call the regestry service
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server            //create a server object
	srv.Addr = registry.ServerPort //set it's address

	go func() { //a routine to start our srv
		log.Println(srv.ListenAndServe()) //start up and, call the Listen and Serve on that Server
		//if that returns, it means we have an error fromtrying to start up a server
		cancel() //and so we'll cancel
	}()
	go func() { //a routine that will allow us to cancel
		fmt.Printf("Regisry service started. Press any key to stop.\n") //print a msg
		var s string                                                    //create a variable
		fmt.Scanln(&s)                                                  //scan into it
		srv.Shutdown(ctx)                                               //shutdown the srv ctx
		cancel()                                                        //cancel
	}()
	<-ctx.Done()
	fmt.Println("Registry service has shut down")
}
