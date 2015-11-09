package cobblerclient

import (
	"io"
)

type System struct {
	Id            string
	cobblerClient *Client
	SystemConfig
}

type SystemConfig struct {
	Name        string
	Profile     string
	Hostname    string
	Nameservers string
	Network     NetworkConfig
}

type NetworkConfig struct {
	Mac     string
	DNSName string
	Ip      string
	Netmask string
	Gateway string
}

func (c *Client) CreateSystem(config SystemConfig) (*System, error) {
	id, err := NewSystemId(c)
	if err != nil {
		return nil, err
	}

	system := System{cobblerClient: c}
	system.SystemConfig = config
	system.Id = id

	_, err = system.SetName(config.Name)
	if err != nil {
		return nil, err
	}

	_, err = system.SetProfile(config.Profile)
	if err != nil {
		return nil, err
	}

	_, err = system.SetHostname(config.Hostname)
	if err != nil {
		return nil, err
	}

	_, err = system.SetNameservers(config.Nameservers)
	if err != nil {
		return nil, err
	}

	_, err = system.SetNetwork(config.Network)
	if err != nil {
		return nil, err
	}

	_, err = system.Save()
	if err != nil {
		return nil, err
	}

	return &system, nil
}

// Requests Cobbler to create a new system.
// Returns the newly created system if it was successfully created.
// Returns an error otherwise.
func NewSystemId(c *Client) (string, error) {
	body := tplNewSystem(c.token)
	res, err := c.post(body)
	if err != nil {
		return "", err
	}

	return systemIDFromResponse(res)
}

// Saves the current state of the system.
// Returns true if the save was successful, or false if it was not.
// Returns an error if anything went wrong
func (s *System) Save() (bool, error) {
	reqBody := tplSaveSystem(s.Id, s.cobblerClient.token)
	res, err := s.cobblerClient.post(reqBody)
	if err != nil {
		return false, err
	}

	return boolFromResponse(res)
}

// Requests Cobbler to create a new system and sets the newly created system's id
// into the `system` instace.
// Returns an error in case anything goes wrong.
func (s *System) SetId() error {
	body := tplNewSystem(s.cobblerClient.token)
	res, err := s.cobblerClient.post(body)
	if err != nil {
		return err
	}

	id, err := systemIDFromResponse(res)
	if err != nil {
		return err
	}

	s.Id = id

	return nil
}

func (s *System) SetName(name string) (bool, error) {
	body := tplSetSystemName(s.Id, name, s.cobblerClient.token)
	return s.modify(body)
}

func (s *System) SetProfile(profile string) (bool, error) {
	body := tplSetSystemProfile(s.Id, profile, s.cobblerClient.token)
	return s.modify(body)
}

func (s *System) SetHostname(hostname string) (bool, error) {
	body := tplSetSystemHostname(s.Id, hostname, s.cobblerClient.token)
	return s.modify(body)
}

func (s *System) SetNameservers(nameservers string) (bool, error) {
	body := tplSetSystemNameservers(s.Id, nameservers, s.cobblerClient.token)
	return s.modify(body)
}

func (s *System) SetNetwork(config NetworkConfig) (bool, error) {
	body := tplSetSystemNetwork(s.Id, config, s.cobblerClient.token)
	return s.modify(body)
}

// Requests Cobbler to modify an existing system.
func (s *System) modify(body io.Reader) (bool, error) {
	_, err := s.cobblerClient.post(body)

	if err != nil {
		return false, err
	}

	return true, nil
}

// Given a Cobbler's response for creating a new system this will
// return the newly created system's id or an error if anything goes wrong.
func systemIDFromResponse(body []byte) (string, error) {
	return findXPath("//param/value/string", body)
}

func (c *Client) DeleteSystem(name string) (bool, error) {
	reqBody := tplDeleteSystem(name, c.token)
	result, err := c.post(reqBody)
	if err != nil {
		return false, err
	}

	return boolFromResponse(result)
}
