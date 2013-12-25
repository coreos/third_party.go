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
	"log"
	"os"
)

var (
	setup = flag.Bool("setup", false, "Do an initial project setup")
)

func setupProject(root string) {
	os.Mkdir("third_party", 0777)
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	root, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get the current working directory: %v", err)
	}

	if *setup {
		setupProject(root)
	}

}
