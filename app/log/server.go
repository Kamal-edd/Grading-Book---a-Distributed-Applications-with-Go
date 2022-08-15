package log

import (
	"io/ioutil"
	stlog "log" //alias to avoid confusion
	"net/http"
	"os"
)

var log *stlog.Logger //a custom logger to handle the logging for our app

type fileLog string //a custom writer which will handle the actlual writing
//this log service will take post requests that are coming in and write their
//content to the log through an endpoit that we'll specify later

func (fl fileLog) Write(data []byte) (int, error) {
	//this is a write method of fileLog, that accepts a slice of bytes and returns
	//an int and an error
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	//open the log file
	if err != nil { //handling error
		return 0, err
	}
	defer f.Close()      //close the file when done
	return f.Write(data) //return the results writing that data mentioned in
	//the definition to our file
}
func Run(destination string) { //this function will set our logger to a new log,
	//give it the destinaton, no prefix, and standard time flags
	log = stlog.New(fileLog(destination), "", stlog.LstdFlags)
}

//now we register our http endpoints
func RegisterHandlers() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		//add an http function to "/log"
		msg, err := ioutil.ReadAll(r.Body) //read the requests body into msg and cache errors
		if err != nil || len(msg) == 0 {   //handle errors and empty equests
			w.WriteHeader(http.StatusBadRequest) //write bad requst stat
			return
		}
		write(string(msg)) //or write msg
	})
}

func write(message string) { //the func to write the msg
	log.Printf("%v\n", message)
}
