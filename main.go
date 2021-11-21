/*
Copyright © 2021 Venkatesh Macha

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
	"log"

	"github.com/mvenkatesh431/todo/cmd"
	"github.com/mvenkatesh431/todo/db"
)

func main() {
	// Initialize the BoltDB
	checkErr(db.Initialize())

	// Init Cobra and start
	cmd.Execute()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}
