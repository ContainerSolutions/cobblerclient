package main

import (
	"fmt"
	"net/http"

	cobbler "github.com/ContainerSolutions/cobblerclient"
)

var config = cobbler.ClientConfig{
	Url:      "http://localhost:25151",
	Username: "cobbler",
	Password: "password",
}

func main() {
	c := cobbler.NewClient(http.DefaultClient, config)
	_, err := c.Login()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	systems, err := c.GetSystems()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("%s", systems[0]["profile"])

}
