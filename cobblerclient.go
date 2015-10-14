package cobblerclient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	xmlpath "gopkg.in/xmlpath.v2"
)

const bodyTypeXML = "text/xml"

type HTTPClient interface {
	Post(string, string, io.Reader) (*http.Response, error)
}

type Client struct {
	httpClient HTTPClient
	config     ClientConfig
	token      string
}

type ClientConfig struct {
	Url      string
	Username string
	Password string
}

func NewClient(httpClient HTTPClient, c ClientConfig) Client {
	return Client{
		httpClient: httpClient,
		config:     c,
	}
}

// Performs a login request to Cobbler using the credentials provided
// in the configuration in the initializer.
func (c *Client) Login() (bool, error) {
	credentials := loginCredentials{c.config.Username, c.config.Password}
	body := tplLogin(credentials)
	res, err := c.post(body)
	if err != nil {
		return false, err
	}

	token, err := tokenFromResponse(res)
	if err != nil {
		return false, nil
	}
	c.token = token

	return true, nil
}

// Sync the system.
// Returns true if the sync was successful, or false if it was not.
// Returns an error if anything went wrong
func (c *Client) Sync() (bool, error) {
	reqBody := tplSync(c.token)
	res, err := c.post(reqBody)
	if err != nil {
		return false, err
	}

	return boolFromResponse(res)
}

// Performs a POST request to the Cobbler server.
// Returns an HTTP Response and an error tuple. Either can be nil.
func (c *Client) post(req io.Reader) ([]byte, error) {
	res, err := c.httpClient.Post(c.config.Url, bodyTypeXML, req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = errorInCobbler(body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Cobbler's errors come in an XML document which contains
// both an error code and an error message. This snippet converts
// that document into a proper Go error object.
// For more detail on how that document looks like check
// `./fixtures/login-res-err.xml`
func errorInCobbler(body []byte) error {
	path := xmlpath.MustCompile("//member/value")
	rootNode, err := xmlpath.Parse(bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	found := 0
	numberOfNodesToFind := 2
	elements := make([]string, numberOfNodesToFind)
	iterator := path.Iter(rootNode)

	for iterator.Next() {
		if found < numberOfNodesToFind {
			elements[found] = iterator.Node().String()
			found++
		}
	}

	if found == 2 {
		return fmt.Errorf("error %s: %s", elements[0], elements[1])
	}

	return nil
}

// Given a Cobbler's API's XML document this will return the
// token as a string or an error if anything goes wrong parsing
// the document, etc.
func tokenFromResponse(body []byte) (string, error) {
	return findXPath("//param/value/string", body)
}

// Given Cobbler's response to a save system call this will return
// the boolean result of the save call, or false and an error if anything went wrong.
func boolFromResponse(body []byte) (bool, error) {
	result, err := findXPath("//param/value/boolean", body)

	if err != nil {
		return false, err
	}

	return (result == "1"), nil
}

// Find the given xpath in the given document.
func findXPath(xpath string, doc []byte) (string, error) {
	path := xmlpath.MustCompile(xpath)
	rootNode, err := xmlpath.Parse(bytes.NewBuffer(doc))
	if err != nil {
		return "", err
	}

	if value, ok := path.String(rootNode); ok {
		return value, nil
	}

	return "", fmt.Errorf("node blank or not found\n%s", rootNode.String())
}
