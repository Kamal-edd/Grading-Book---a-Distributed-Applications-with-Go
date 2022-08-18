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
	if res.StatusCode != 200 {
		return fmt.Errorf("failed to register, registry service"+
			"responds with code %v", res.StatusCode)
	}
	return nil

}
