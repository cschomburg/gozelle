// Copyright (C) 2014 Constantin Schomburg <me@cschomburg.com>
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Package gozelle provides access to the Gazelle API.
package gozelle

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
)

var (
	ErrNotLoggedIn    = errors.New("Not logged in")
	ErrLoginFailed    = errors.New("Login failed")
	ErrInvalidAuthkey = errors.New("Invalid authkey")
	ErrInvalidPasskey = errors.New("Invalid passkey")
)

// A client handles the connection to the API.
type Client struct {
	host string
	http *http.Client

	authkey string
	passkey string
}

// New creates a new Gazelle API client.
func New() *Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &Client{
		http: &http.Client{
			Jar: jar,
		},
	}
}

// NewWithClient creates a new Gazelle API client with a custom HTTP client.
// The HTTP client must support cookie storage.
func NewWithClient(http *http.Client) *Client {
	return &Client{
		http: http,
	}
}

// Do makes a generic request to the API with the specified arguments and
// unmarshals the JSON response into resp.
func (c *Client) Do(action string, v url.Values, resp interface{}) error {
	if c.passkey == "" {
		return ErrNotLoggedIn
	}
	u, err := url.Parse(c.host)
	if err != nil {
		return err
	}

	u.Path = "ajax.php"
	if v == nil {
		v = url.Values{}
	}
	v.Set("action", action)
	u.RawQuery = v.Encode()

	r, err := c.http.Get(u.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		return errors.New("unexpected status: " + r.Status)
	}

	return json.NewDecoder(r.Body).Decode(&resp)
}

// Login connects to the Gazelle host with username and password. It returns
// user stats on successful login, or an error message otherwise.
func (c *Client) Login(host, user, pass string) (r IndexResponse, err error) {
	c.host = strings.TrimRight(host, "/")

	v := url.Values{}
	v.Set("username", user)
	v.Set("password", pass)
	resp, err := c.http.PostForm(host+"/login.php", v)
	if err != nil {
		return r, err
	}

	if resp.StatusCode != 200 {
		return r, errors.New("unexpected status: " + resp.Status)
	}

	switch page := resp.Request.URL.Path; page {
	case "/index.php":
		break
	case "/login.php":
		return r, ErrLoginFailed
	default:
		return r, errors.New("unexpected login page: " + page)
	}

	c.passkey = "temp"
	if r, err = c.Index(); err != nil {
		return r, err
	}
	c.authkey = r.Response.Authkey
	c.passkey = r.Response.Passkey
	return r, nil
}

// Index retrieves a few user stats.
func (c *Client) Index() (IndexResponse, error) {
	var r IndexResponse
	err := c.Do("index", nil, &r)
	return r, err
}

// User retrieves stats of a specific user.
func (c *Client) User(id int) (UserResponse, error) {
	var r UserResponse
	v := url.Values{}
	v.Set("id", strconv.Itoa(id))
	err := c.Do("user", v, &r)
	return r, err
}

// TorrentsFilter represents a set of search filters.
type TorrentsFilter struct {
	TagList   string
	TagsType  string
	OrderBy   string
	OrderWay  string
	FilterCat string

	Freetorrent bool
	Vanityhouse bool
	Scene       bool
	HasLog      bool

	ReleaseType     string
	Media           string
	Format          string
	Encoding        string
	ArtistName      string
	FileList        string
	GroupName       string
	RecordLabel     string
	CatalogueNumber string
	Year            int

	RemasterTitle           string
	RemasterYear            string
	RemasterRecordLabel     string
	RemasterCatalogueNumber string
}

// SearchTorrents retrieves torrents based on a search query. If filters is not nil,
// the search is narrowed down.
func (c *Client) SearchTorrents(search string, page int, filters *TorrentsFilter) (TorrentSearchResponse, error) {
	var r TorrentSearchResponse
	v := url.Values{}
	if filters != nil {
		structToValues(filters, v)
	}
	v.Set("searchstr", search)
	if page != 0 {
		v.Set("page", strconv.Itoa(page))
	}
	err := c.Do("browse", v, &r)
	return r, err
}

// Artist retrieves info of a specific artist, either by providing an exact id
// greater than 0 or a non-empty artist name.
func (c *Client) Artist(id int, name string) (ArtistResponse, error) {
	var r ArtistResponse
	v := url.Values{}
	if id != 0 {
		v.Set("id", strconv.Itoa(id))
	}
	if name != "" {
		v.Set("artistname", name)
	}
	err := c.Do("artist", v, &r)
	return r, err
}

// Torrent retrieves info of a specific torrent, either by providing an exact id
// greater than 0 or a non-empty torrent hash.
func (c *Client) Torrent(id int, hash string) (TorrentResponse, error) {
	var r TorrentResponse
	v := url.Values{}
	if id != 0 {
		v.Set("id", strconv.Itoa(id))
	}
	if hash != "" {
		v.Set("hash", hash)
	}
	err := c.Do("torrent", v, &r)
	return r, err
}

// TorrentGroup retrieves info of a specific torrent group, either by providing
// an exact id greater than 0 or a non-empty torrent hash.of a torrent in the group.
func (c *Client) TorrentGroup(id int, hash string) (TorrentGroupResponse, error) {
	var r TorrentGroupResponse
	v := url.Values{}
	if id != 0 {
		v.Set("id", strconv.Itoa(id))
	}
	if hash != "" {
		v.Set("hash", hash)
	}
	err := c.Do("torrentgroup", v, &r)
	return r, err
}

// Downloadlink generates a download link for the specified torrent file that can
// be used externally. The link includes personal authentication data.
//
// Unspecified or improper keys cause an error.
func Downloadlink(host, authkey, passkey string, torrentid int) (string, error) {
	u, err := url.Parse(host)
	if err != nil {
		return "", err
	}
	if len(authkey) != 32 {
		return "", ErrInvalidAuthkey
	}
	if len(passkey) != 32 {
		return "", ErrInvalidPasskey
	}

	u.Path = "torrents.php"
	v := url.Values{}
	v.Set("action", "download")
	v.Set("id", strconv.Itoa(torrentid))
	v.Set("authkey", authkey)
	v.Set("torrent_pass", passkey)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

// Downloadlink generates a download link for the specified torrent file that can
// be used externally. The link includes personal authentication data.
func (c *Client) Downloadlink(id int) (string, error) {
	if c.passkey == "" {
		return "", ErrNotLoggedIn
	}
	return Downloadlink(c.host, c.authkey, c.passkey, id)
}

// Download fetches a torrent file. The file body is returned as an ReadCloser
// and can be zero length if the torrent id / authentication is wrong.
func (c *Client) Download(id int) (io.ReadCloser, error) {
	u, err := c.Downloadlink(id)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Get(u)
	if resp.StatusCode != 200 {
		return nil, errors.New("unexpected status: " + resp.Status)
	}
	return resp.Body, err
}
