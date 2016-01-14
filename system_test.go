/*
Copyright 2015 Container Solutions

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cobblerclient

import (
	"testing"

	"github.com/ContainerSolutions/go-utils"
)

func TestGetSystems(t *testing.T) {
	c := createStubHTTPClient(t, "get-systems-req.xml", "get-systems-res.xml")
	systems, err := c.GetSystems()
	utils.FailOnError(t, err)

	if len(systems) != 1 {
		t.Errorf("Wrong number of systems returned.")
	}
}

func TestGetSystem(t *testing.T) {
	c := createStubHTTPClient(t, "get-system-req.xml", "get-system-res.xml")
	system, err := c.GetSystem("test")
	utils.FailOnError(t, err)

	if system.Name != "test" {
		t.Errorf("Wrong system returned.")
	}
}

/*
 * NOTE: We're skipping the testing of this method for now because
 *       the current implementation of the StubHTTPClient does not allow
 *       buffered mock responses so as soon as the method makes the second
 *       call to Cobbler it'll fail.
 *       This is a system test, so perhaps we can run Cobbler in a Docker container
 *       and take it from there.
 */

/*
func TestCreateSystem(t *testing.T) {
	t.Skip()
	sysConfig := SystemConfig{
		Name:        "blah",
		Profile:     "some-profile",
		Hostname:    "blahhost",
		Nameservers: "8.8.8.8 8.8.4.4",
		Network: NetworkConfig{
			Mac:     "01:02:03:04:05:06",
			DNSName: "blah",
			Ip:      "1.2.3.4",
			Netmask: "255.255.255.0",
			Gateway: "4.3.2.1",
		},
	}
	hc := utils.NewStubHTTPClient(t)
	hc.ShouldVerify = false
	c := NewClient(hc, config)

	_, err := c.CreateSystem(sysConfig)
	utils.FailOnError(t, err)
}

func TestNewSystemId(t *testing.T) {
	expectedReq, err := utils.Fixture("new-system-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("new-system-res.xml")
	utils.FailOnError(t, err)

	expectedId := "___NEW___system::abc123=="
	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response
	cobblerClient := NewClient(hc, config)
	cobblerClient.token = "securetoken99"
	id, err := NewSystemId(&cobblerClient)
	utils.FailOnError(t, err)

	if id != expectedId {
		t.Errorf("%s expected; got %s", expectedId, id)
	}
}

func TestSetSystemId(t *testing.T) {
	expectedReq, err := utils.Fixture("new-system-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("new-system-res.xml")
	utils.FailOnError(t, err)

	expectedId := "___NEW___system::abc123=="
	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response
	cobblerClient := NewClient(hc, config)
	cobblerClient.token = "securetoken99"
	system := System{cobblerClient: &cobblerClient}
	err = system.SetId()
	utils.FailOnError(t, err)

	if system.Id != expectedId {
		t.Errorf("%s expected; got %s", expectedId, system.Id)
	}
}

func TestSaveSystem(t *testing.T) {
	expectedReq, err := utils.Fixture("save-system-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("save-system-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "securetoken99"
	system := System{}
	system.cobblerClient = &c
	system.Id = "___NEW___system::abc123=="

	ok, err := system.Save()
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}

func TestSetSystemName(t *testing.T) {
	expectedReq, err := utils.Fixture("set-system-name-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("set-system-name-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "securetoken99"
	system := System{cobblerClient: &c}
	system.Id = "___NEW___system::foobar123=="

	ok, err := system.SetName("mytestsystem")
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}

func TestSetSystemProfile(t *testing.T) {
	expectedReq, err := utils.Fixture("set-system-profile-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("set-system-profile-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "abc123"
	system := System{cobblerClient: &c}
	system.Id = "___NEW___system::c3po"

	ok, err := system.SetProfile("centos7-x86_64")
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}

func TestSetSystemHostname(t *testing.T) {
	expectedReq, err := utils.Fixture("set-system-hostname-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("set-system-hostname-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "abc123"
	system := System{cobblerClient: &c}
	system.Id = "___NEW___system::c3po"

	ok, err := system.SetHostname("blahhost")
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}

func TestSetSystemNameservers(t *testing.T) {
	expectedReq, err := utils.Fixture("set-system-nameservers-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("set-system-nameservers-res.xml")
	utils.FailOnError(t, err)

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "securetoken99"
	system := System{cobblerClient: &c}
	system.Id = "___NEW___system::foobar123=="

	ok, err := system.SetNameservers("8.8.8.8 8.8.4.4")
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}
func TestSetSystemNetwork(t *testing.T) {
	expectedReq, err := utils.Fixture("set-system-network-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("set-system-network-res.xml")
	utils.FailOnError(t, err)

	networkConfig := NetworkConfig{
		Mac:     "01:02:03:04:05:06",
		DNSName: "deathstar",
		Ip:      "1.2.3.4",
		Netmask: "255.255.255.0",
		Gateway: "4.3.2.1",
	}
	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response

	c := NewClient(hc, config)
	c.token = "abc123=="
	system := System{cobblerClient: &c}
	system.Id = "___NEW___system::abc123=="

	ok, err := system.SetNetwork(networkConfig)
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected; got false")
	}
}

*/
