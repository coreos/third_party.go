// +build ignore

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

// This program builds a project and is a copy of build.go. See
// github.com/philips/build.go
//
// $ go run build.go
//
// See the README file for more details.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

const (
	DefaultThirdParty = "third_party"
)

var (
	setup = flag.String("setup", "", "Do an initial project setup. Takes the base repo name as an argument.")
	pkg   = flag.String("package", "", "Name of the package to build")
)

type Package struct {
}

func thirdPartyDir() string {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}
	return path.Join(root, DefaultThirdParty)
}

func binDir() string {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}
	return path.Join(root, "bin")
}

func run(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	cmd.Env = append(os.Environ(),
		"GOPATH="+thirdPartyDir(),
		"GOBIN="+binDir(),
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
	go io.Copy(os.Stdout, stdout)
	go io.Copy(os.Stderr, stderr)
	cmd.Wait()
}

// setupProject does the initial setup of the third_party src directory
// including setting up the symlink to the cwd from the src directory.
func setupProject(root string, pkg string) {
	src := path.Join(thirdPartyDir(), "src", pkg)
	srcdir := path.Dir(src)

	os.MkdirAll(srcdir, 0777)

	rel, err := filepath.Rel(srcdir, root)
	if err != nil {
		log.Fatalf("creating relative third party path: %v", err)
	}

	err = os.Symlink(rel, src)
	if err != nil && os.IsExist(err) == false {
		log.Fatalf("creating project third party symlink: %v", err)
	}
}

func build(pkg string) {
	run("go", "build", pkg)
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}

	if *pkg == "" && *setup == "" {
		log.Fatalf("Package name is required")
	}

	if *setup != "" {
		setupProject(root, *setup)
		return
	}

	build(*pkg)
}
