/*
 * Copyright (c) 2017, 2020 ADLINK Technology Inc.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License 2.0 which is available at
 * http://www.eclipse.org/legal/epl-2.0, or the Apache License, Version 2.0
 * which is available at https://www.apache.org/licenses/LICENSE-2.0.
 *
 * SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
 *
 * Contributors:
 *   ADLINK zenoh team, <zenoh@adlink-labs.tech>
 */

package main

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
	znet "github.com/eclipse-zenoh/zenoh-go/net"
)

func main() {
	// --- Command line argument parsing --- --- --- --- --- ---
	var args struct {
		Path    string `default:"/zenoh/examples/go/stream/hello" arg:"-p" help:"the path representing the URI"`
		Locator string `arg:"-l" help:"The locator to be used to boostrap the zenoh session. By default dynamic discovery is used"`
		Msg     string `default:"Zenitude streamed from zenoh-net-go!" arg:"-m" help:"The quote associated with the welcoming resource"`
	}
	arg.MustParse(&args)

	// zenoh-net code  --- --- --- --- --- --- --- --- --- --- ---
	fmt.Println("Opening session...")
	s, err := znet.Open(&args.Locator, nil)
	if err != nil {
		panic(err.Error())
	}
	defer s.Close()

	fmt.Println("Declaring Publisher on '" + args.Path + "'...")
	pub, err := s.DeclarePublisher(args.Path)
	if err != nil {
		panic(err.Error())
	}
	defer s.UndeclarePublisher(pub)

	for idx := 0; idx < 100; idx++ {
		time.Sleep(1 * time.Second)
		str := fmt.Sprintf("[%4d] %s", idx, args.Msg)
		fmt.Printf("Streaming Data ('%s': '%s')...\n", args.Path, str)
		pub.StreamData([]byte(str))
	}
}
