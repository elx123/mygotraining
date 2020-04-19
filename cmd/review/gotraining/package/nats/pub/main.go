// Copyright 2012-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

// NOTE: Can test with demo servers.
// nats-pub -s demo.nats.io <subject> <msg>
// nats-pub -s demo.nats.io:4443 <subject> <msg> (TLS version)

func main() {

	// Connect Options.
	opts := []nats.Option{nats.Name("NATS Sample Publisher")}

	// Connect to NATS
	nc, err := nats.Connect("nats://192.168.99.100:30009", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj, msg := "test", []byte("test123")

	nc.Publish(subj, msg)

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	}
}
