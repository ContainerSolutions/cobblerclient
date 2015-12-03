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

func TestCreateSnippet(t *testing.T) {
	expectedReq, err := utils.Fixture("create-snippet-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("create-snippet-res.xml")
	utils.FailOnError(t, err)

	s := Snippet{
		Name: "some-snippet",
		Body: "sample content",
	}

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response
	cobblerClient := NewClient(hc, config)
	cobblerClient.token = "kirby123"
	ok, err := cobblerClient.CreateSnippet(&s)
	utils.FailOnError(t, err)

	if !ok {
		t.Errorf("true expected but got false")
	}
}

func TestCreateSnippetWithError(t *testing.T) {
	expectedReq, err := utils.Fixture("create-snippet-req.xml")
	utils.FailOnError(t, err)

	response, err := utils.Fixture("create-snippet-res-err.xml")
	utils.FailOnError(t, err)

	s := Snippet{
		Name: "some-snippet",
		Body: "sample content",
	}

	hc := utils.NewStubHTTPClient(t)
	hc.Expected = expectedReq
	hc.Response = response
	cobblerClient := NewClient(hc, config)
	cobblerClient.token = "kirby123"
	ok, err := cobblerClient.CreateSnippet(&s)
	utils.FailOnError(t, err)

	if ok {
		t.Errorf("false expected but got true")
	}
}
