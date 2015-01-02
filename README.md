gozelle
=======

[![API Documentation](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](http://godoc.org/github.com/xconstruct/gozelle)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](http://opensource.org/licenses/MIT)

A Go package to interact with the JSON API of What.CD / the private tracker framework [Gazelle](https://whatcd.github.io/Gazelle/).

A list of all available methods can be found at the official [Gazelle documentation](https://github.com/WhatCD/Gazelle/wiki/JSON-API-Documentation). Currently only a limited set of functions for searching/downloading is implemented. For everything else there is the general `(c *Client) Do()`.

Install
-------

```
go get "github.com/xconstruct/gozelle"
```

Example
-------

```go
g := gozelle.New()
r, err := g.Login("https://what.cd", "user", "pass")
fmt.Println("Hi", r.Response.Username)

r, err = g.SearchTorrents("favorite artist", 0, &gozelle.TorrentsFilter{
	Format: "FLAC",
})

fmt.Println("download link:", r.Response.Results[0].Torrents[0].TorrentId)
```
