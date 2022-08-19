package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterService(r Registration) error {
	buf := new(bytes.Buffer)    //create a buffer
	enc := json.NewEncoder(buf) //create json encoder
	err := enc.Encode(r)        //encode registration
	if err != nil {             //handle encoring error
		return err
	}
	res, err := http.Post(ServiceURL, "application/json", buf)
	//post the encoded registration to the service
	if err != nil { //handle posting error
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register, registry service "+
			"of URL %v"+
			" responds with code %v", ServiceURL, res.StatusCode)
	}
	return nil
}
func ShutdownService(ServiceURL string) error {
	req, err := http.NewRequest(http.MethodDelete,
		ServiceURL,
		bytes.NewBuffer([]byte(ServiceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registsy "+
			"service responds with code %v", res.StatusCode)
	}
	return err
}
