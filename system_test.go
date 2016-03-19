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
 * NOTE: We're skipping the testing of CREATE, UPDATE, DELETE methods for now because
 *       the current implementation of the StubHTTPClient does not allow
 *       buffered mock responses so as soon as the method makes the second
 *       call to Cobbler it'll fail.
 *       This is a system test, so perhaps we can run Cobbler in a Docker container
 *       and take it from there.
 */
