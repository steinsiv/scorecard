// Copyright 2020 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"fmt"
	"os"
	"sort"

	docs "github.com/ossf/scorecard/v3/docs/checks"
)

func main() {
	if len(os.Args) != 2 {
		// nolint: goerr113
		panic(fmt.Errorf("usage: %s filename", os.Args[0]))
	}
	yamlFile := os.Args[1]

	m, err := docs.Read()
	if err != nil {
		panic(err)
	}
	checks := m.GetChecks()
	keys := make([]string, 0, len(checks))
	for _, v := range checks {
		keys = append(keys, v.GetName())
	}
	sort.Strings(keys)

	f, err := os.Create(yamlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(`
<!-- Do not edit this file manually! Edit checks.yaml instead. --> 
# Check Documentation

This page describes each Scorecard check in detail, including scoring criteria,
remediation steps to improve the score, and an explanation of the risks
associated with a low score. The checks are continually changing and we welcome
community feedback. If you have ideas for additions or new detection techniques,
please [contribute](CONTRIBUTING.md)!
`)
	if err != nil {
		panic(err)
	}
	for _, k := range keys {
		_, err := f.WriteString("## " + k + " \n\n")
		if err != nil {
			panic(err)
		}
		c, err := m.GetCheck(k)
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(c.GetDescription() + " \n\n")
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString("**Remediation steps**\n")
		if err != nil {
			panic(err)
		}
		for _, r := range c.GetRemediation() {
			_, err = f.WriteString("- " + r + "\n")
			if err != nil {
				panic(err)
			}
		}
		_, err = f.WriteString("\n")
		if err != nil {
			panic(err)
		}
	}
}
