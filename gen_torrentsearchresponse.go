package gozelle

type TorrentSearchResponse struct {
	Response struct {
		CurrentPage int `json:"currentPage"`
		Pages       int `json:"pages"`
		Results     []struct {
			Artist      string   `json:"artist"`
			Bookmarked  bool     `json:"bookmarked"`
			GroupId     int      `json:"groupId"`
			GroupName   string   `json:"groupName"`
			GroupTime   string   `json:"groupTime"`
			GroupYear   int      `json:"groupYear"`
			MaxSize     int      `json:"maxSize"`
			ReleaseType string   `json:"releaseType"`
			Tags        []string `json:"tags"`
			Torrents    []struct {
				Artists []struct {
					Aliasid int    `json:"aliasid"`
					ID      int    `json:"id"`
					Name    string `json:"name"`
				} `json:"artists"`
				CanUseToken             bool   `json:"canUseToken"`
				EditionId               int    `json:"editionId"`
				Encoding                string `json:"encoding"`
				FileCount               int    `json:"fileCount"`
				Format                  string `json:"format"`
				HasCue                  bool   `json:"hasCue"`
				HasLog                  bool   `json:"hasLog"`
				IsFreeleech             bool   `json:"isFreeleech"`
				IsNeutralLeech          bool   `json:"isNeutralLeech"`
				IsPersonalFreeleech     bool   `json:"isPersonalFreeleech"`
				Leechers                int    `json:"leechers"`
				LogScore                int    `json:"logScore"`
				Media                   string `json:"media"`
				RemasterCatalogueNumber string `json:"remasterCatalogueNumber"`
				RemasterTitle           string `json:"remasterTitle"`
				RemasterYear            int    `json:"remasterYear"`
				Remastered              bool   `json:"remastered"`
				Scene                   bool   `json:"scene"`
				Seeders                 int    `json:"seeders"`
				Size                    int    `json:"size"`
				Snatches                int    `json:"snatches"`
				Time                    string `json:"time"`
				TorrentId               int    `json:"torrentId"`
				VanityHouse             bool   `json:"vanityHouse"`
			} `json:"torrents"`
			TotalLeechers int  `json:"totalLeechers"`
			TotalSeeders  int  `json:"totalSeeders"`
			TotalSnatched int  `json:"totalSnatched"`
			VanityHouse   bool `json:"vanityHouse"`
		} `json:"results"`
	} `json:"response"`
	Status string `json:"status"`
}
