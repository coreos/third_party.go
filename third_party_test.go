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
	"path/filepath"
	"testing"
)

const (
	TestPkg = "foo.com/bar/baz"
)

func TestSetupProject(t *testing.T) {
	temp, _ := ioutil.TempDir("", "setupProject")
	defer os.RemoveAll(temp)

	if err := os.Chdir(temp); err != nil {
		t.Log("error changing directory")
		t.Fail()
	}

	setupProject(TestPkg)

	pkgLinkPath := filepath.Join(thirdPartyDir(), "src", TestPkg)

	fileinfo, _ := os.Lstat(pkgLinkPath)
	if fileinfo.Mode()&os.ModeSymlink != os.ModeSymlink {
		t.Logf("%v not a symbolic link!", pkgLinkPath)
		t.Fail()
	}

	relLink, _ := os.Readlink(pkgLinkPath)

	linkDest := filepath.Join(filepath.Dir(pkgLinkPath), relLink)
	realTemp, _ := filepath.EvalSymlinks(temp)
	if linkDest != realTemp {
		t.Logf("incorrect symbolic link (should be %v but points to %v)", realTemp, linkDest)
		t.Fail()
	}
}

func TestRemoveVcs(t *testing.T) {
	temp, _ := ioutil.TempDir("", "removeVcs")

	defer os.RemoveAll(temp)

	proj := "github.com/removevcs"
	p := filepath.Join(temp, proj, ".git")

	os.MkdirAll(p, 0755)

	removeVcs(filepath.Join(temp, proj))
	info, _ := os.Stat(p)

	if info != nil {
		t.Log(".git still exists")
		t.Fail()
	}
}
