package gozelle

type ArtistResponse struct {
	Response struct {
		Body                 string `json:"body"`
		HasBookmarked        bool   `json:"hasBookmarked"`
		ID                   int    `json:"id"`
		Image                string `json:"image"`
		Name                 string `json:"name"`
		NotificationsEnabled bool   `json:"notificationsEnabled"`
		Requests             []struct {
			Bounty     int    `json:"bounty"`
			CategoryId int    `json:"categoryId"`
			RequestId  int    `json:"requestId"`
			TimeAdded  string `json:"timeAdded"`
			Title      string `json:"title"`
			Votes      int    `json:"votes"`
			Year       int    `json:"year"`
		} `json:"requests"`
		SimilarArtists []interface{} `json:"similarArtists"`
		Statistics     struct {
			NumGroups   int `json:"numGroups"`
			NumLeechers int `json:"numLeechers"`
			NumSeeders  int `json:"numSeeders"`
			NumSnatches int `json:"numSnatches"`
			NumTorrents int `json:"numTorrents"`
		} `json:"statistics"`
		Tags []struct {
			Count int    `json:"count"`
			Name  string `json:"name"`
		} `json:"tags"`
		Torrentgroup []struct {
			GroupCatalogueNumber string   `json:"groupCatalogueNumber"`
			GroupId              int      `json:"groupId"`
			GroupName            string   `json:"groupName"`
			GroupRecordLabel     string   `json:"groupRecordLabel"`
			GroupVanityHouse     bool     `json:"groupVanityHouse"`
			GroupYear            int      `json:"groupYear"`
			HasBookmarked        bool     `json:"hasBookmarked"`
			ReleaseType          int      `json:"releaseType"`
			Tags                 []string `json:"tags"`
			Torrent              []struct {
				Encoding            string `json:"encoding"`
				FileCount           int    `json:"fileCount"`
				Format              string `json:"format"`
				FreeTorrent         bool   `json:"freeTorrent"`
				GroupId             int    `json:"groupId"`
				HasCue              bool   `json:"hasCue"`
				HasFile             int    `json:"hasFile"`
				HasLog              bool   `json:"hasLog"`
				ID                  int    `json:"id"`
				Leechers            int    `json:"leechers"`
				LogScore            int    `json:"logScore"`
				Media               string `json:"media"`
				RemasterRecordLabel string `json:"remasterRecordLabel"`
				RemasterTitle       string `json:"remasterTitle"`
				RemasterYear        int    `json:"remasterYear"`
				Remastered          bool   `json:"remastered"`
				Scene               bool   `json:"scene"`
				Seeders             int    `json:"seeders"`
				Size                int    `json:"size"`
				Snatched            int    `json:"snatched"`
				Time                string `json:"time"`
			} `json:"torrent"`
		} `json:"torrentgroup"`
		VanityHouse bool `json:"vanityHouse"`
	} `json:"response"`
	Status string `json:"status"`
}
