// Copyright (C) 2014 Constantin Schomburg <me@cschomburg.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Example commandline client for the Gazelle API.
package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/xconstruct/gozelle"
)

var (
	host = flag.String("host", "https://what.cd", "Gazelle tracker website")
	user = flag.String("user", "", "Your username")
	pass = flag.String("pass", "", "Your password")

	search   = flag.String("search", "", "Searches for a torrent")
	download = flag.Int("download", 0, "Downloads the specified torrent id")
)

func main() {
	flag.Parse()

	if *user == "" || *pass == "" {
		log.Fatal("No username and password specified!")
	}

	goz := gozelle.New()
	_, err := goz.Login(*host, *user, *pass)
	if err != nil {
		log.Fatal(err)
	}

	if *search != "" {
		r, err := goz.SearchTorrents(*search, 0, nil)
		if err != nil {
			log.Fatal(err)
		}
		for _, group := range r.Response.Results {
			fmt.Printf("%s - %s (%d)\n", group.Artist, group.GroupName, group.GroupYear)

			for _, t := range group.Torrents {
				fmt.Printf("    [%d] %s %s\n", t.TorrentId, t.Format, t.Encoding)
			}
		}
	}
	if *download > 0 {
		u, err := goz.Downloadlink(*download)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(u)
	}
}
