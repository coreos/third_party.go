/*
Copyright 2013 Brandon Philips

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
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	TestDir = "test_dir"
)

func TestRemoveVcs(t *testing.T) {
	temp, _ := ioutil.TempDir(TestDir, "removeVcs")
	defer os.RemoveAll(temp)

	proj := "github.com/removevcs"
	p := path.Join(temp, proj, ".git")

	os.MkdirAll(p, 0755)

	removeVcs(path.Join(temp, proj))
	info, _ := os.Stat(p)

	if info != nil {
		t.Log(".git still exists")
		t.Fail()
	}
}
