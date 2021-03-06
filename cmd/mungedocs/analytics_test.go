/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalytics(t *testing.T) {
	var cases = []struct {
		in  string
		out string
	}{
		{`aoeu`, `aoeu


` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`},
		{`aoeu


[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
`, `aoeu


` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`},
		{`aoeu

` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`, `aoeu


` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`},
		{`aoeu


[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()



` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`, `aoeu


` + beginMungeTag("GENERATED_ANALYTICS") + `
[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/path/to/file-name.md?pixel)]()
` + endMungeTag("GENERATED_ANALYTICS") + `
`},
	}
	for _, c := range cases {
		out, err := checkAnalytics("path/to/file-name.md", []byte(c.in))
		assert.NoError(t, err)
		if string(out) != c.out {
			t.Errorf("Expected \n\n%v\n\n but got \n\n%v\n\n", c.out, string(out))
		}
	}
}
