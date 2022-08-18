package service

import (
	"app/registry"
	"context"
	"fmt"
	"log"
	"net/http"
)

// following is a function that will start the service
func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	//we'll pass it a context, a service name, host and port, then the register func
	registerHandlersFunc()                               //register the handler
	ctx = startService(ctx, reg.ServiceName, host, port) //create a new context
	err := registry.RegisterService(reg)                 //reach out to the registry servic
	//and register reg
	if err != nil { //handle errors
		return ctx, err
	}
	return ctx, nil //return that new ctx
}

func startService(ctx context.Context, serviceName registry.ServiceName, host, port string) context.Context {
	//the function that will start the service
	ctx, cancel := context.WithCancel(ctx) // another ctx derived from the recieved ctx
	//that will have a cancel defined on it
	var srv http.Server   //create a server object
	srv.Addr = ":" + port //set it's address

	go func() { //a routine to start our srv
		log.Println(srv.ListenAndServe()) //start up and, call the Listen and Serve on that Server
		//if that returns, it means we have an error fromtrying to start up a server
		cancel() //and so we'll cancel
	}()
	go func() { //a routine that will allow us to cancel
		fmt.Printf("%v started. Press any key to stop.\n", serviceName) //print a msg
		var s string                                                    //create a variable
		fmt.Scanln(&s)                                                  //scan into it
		srv.Shutdown(ctx)                                               //shutdown the srv ctx
		cancel()                                                        //cancel
	}()
	return ctx
}
