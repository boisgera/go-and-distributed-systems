package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hashicorp/mdns"
)

func halt() {
	ch := make(chan struct{})
	<-ch
}

func Lookup(service string, entries chan<- *mdns.ServiceEntry, Timeout time.Duration) error {
	params := mdns.DefaultParams(service)
	params.Entries = entries
	params.Timeout = Timeout
	err := mdns.Query(params)
	return err
}

func main() {

	go func() {
		host, _ := os.Hostname()
		info := []string{"My awesome service"}
		service, err := mdns.NewMDNSService(host, "_foobar._tcp", "", "", 8080, nil, info)

		if err != nil {
			panic(err)
		}

		// Create the mDNS server, defer shutdown
		server, _ := mdns.NewServer(&mdns.Config{Zone: service})
		defer func() {
			fmt.Println("SHUTDOWN")
			server.Shutdown()
		}()

		halt()
	}()

	time.Sleep(time.Second)

	entriesCh := make(chan *mdns.ServiceEntry, 1000)
	defer close(entriesCh)

	go func() {
		for entry := range entriesCh {
			fmt.Printf("Got new entry: %v\n", entry)
		}
	}()

	// Start the lookup
	for {
		Lookup("_workstation._tcp", entriesCh, 10*time.Second)
	}

}
