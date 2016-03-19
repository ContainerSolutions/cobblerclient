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

func distro_tests(c cobbler.Client) error {
	d := cobbler.Distro{
		Name:      "Test",
		Breed:     "Ubuntu",
		OSVersion: "trusty",
		Arch:      "x86_64",
		Kernel:    "/var/www/cobbler/ks_mirror/Ubuntu-14.04/install/netboot/ubuntu-installer/amd64/linux",
		Initrd:    "/var/www/cobbler/ks_mirror/Ubuntu-14.04/install/netboot/ubuntu-installer/amd64/initrd.gz",
	}

	fmt.Println("Creating Distro")
	newDistro, err := c.CreateDistro(d)
	if err != nil {
		return err
	}

	fmt.Printf("New Distro: %+v\n", newDistro)

	if newDistro.Name != "Test" {
		return fmt.Errorf("Distro name does not match.")
	}

	fmt.Println("Updating Distro")
	newDistro.Comment = "Update Test"
	if err := c.UpdateDistro(newDistro); err != nil {
		return err
	}

	fmt.Println("Deleting Distro")
	_, err = c.DeleteDistro("Test")
	if err != nil {
		fmt.Printf("\n%+v\n", err)
	}

	return nil
}

func main() {
	c := cobbler.NewClient(http.DefaultClient, config)
	_, err := c.Login()
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Printf("Token: %s\n", c.Token)

	if err := distro_tests(c); err != nil {
		fmt.Println(err)
	}

	/*
		systems, err := c.GetSystems()
		if err != nil {
			fmt.Printf("%+v", err)
		}

		fmt.Printf("%+v\n", systems)

		system, err := c.GetSystem("test")
		if err != nil {
			fmt.Printf("%+v", err)
		}

		fmt.Printf("%+v\n", system)

		eth0 := map[string]interface{}{
			"mac_address": "aa:bb:cc:dd:ee:ff",
			"static":      true,
		}

		s := cobbler.CreateSystemOpts{
			Comment:     "WTF",
			Name:        "Foobar",
			Profile:     "Ubuntu-14.04-x86_64",
			NameServers: []string{"8.8.8.8", "1.1.1.1"},
			PowerID:     "foo",
			Interfaces: map[string]interface{}{
				"eth0": eth0,
			},
		}

		ns, err := c.CreateSystem(s)
		if err != nil {
			fmt.Printf("\n%+v\n", err)
		}

		fmt.Printf("\n%+v\n", ns)

		neth0 := ns.Interfaces["eth0"].(map[string]interface{})
		fmt.Printf("\n%s\n", neth0["mac_address"])
	*/
}
