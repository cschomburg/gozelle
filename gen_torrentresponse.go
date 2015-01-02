package gozelle

type TorrentResponse struct {
	Response struct {
		Group struct {
			CatalogueNumber string `json:"catalogueNumber"`
			CategoryId      int    `json:"categoryId"`
			CategoryName    string `json:"categoryName"`
			ID              int    `json:"id"`
			MusicInfo       struct {
				Artists []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"artists"`
				Composers []interface{} `json:"composers"`
				Conductor []interface{} `json:"conductor"`
				Dj        []interface{} `json:"dj"`
				Producer  []interface{} `json:"producer"`
				RemixedBy []interface{} `json:"remixedBy"`
				With      []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"with"`
			} `json:"musicInfo"`
			Name        string `json:"name"`
			RecordLabel string `json:"recordLabel"`
			ReleaseType int    `json:"releaseType"`
			Time        string `json:"time"`
			VanityHouse bool   `json:"vanityHouse"`
			WikiBody    string `json:"wikiBody"`
			WikiImage   string `json:"wikiImage"`
			Year        int    `json:"year"`
		} `json:"group"`
		Torrent struct {
			Description             string      `json:"description"`
			Encoding                string      `json:"encoding"`
			FileCount               int         `json:"fileCount"`
			FileList                string      `json:"fileList"`
			FilePath                string      `json:"filePath"`
			Format                  string      `json:"format"`
			FreeTorrent             bool        `json:"freeTorrent"`
			HasCue                  bool        `json:"hasCue"`
			HasLog                  bool        `json:"hasLog"`
			ID                      int         `json:"id"`
			Leechers                int         `json:"leechers"`
			LogScore                int         `json:"logScore"`
			Media                   string      `json:"media"`
			RemasterCatalogueNumber string      `json:"remasterCatalogueNumber"`
			RemasterRecordLabel     string      `json:"remasterRecordLabel"`
			RemasterTitle           string      `json:"remasterTitle"`
			RemasterYear            int         `json:"remasterYear"`
			Remastered              bool        `json:"remastered"`
			Scene                   bool        `json:"scene"`
			Seeders                 int         `json:"seeders"`
			Size                    int         `json:"size"`
			Snatched                int         `json:"snatched"`
			Time                    string      `json:"time"`
			UserId                  int         `json:"userId"`
			Username                interface{} `json:"username"`
		} `json:"torrent"`
	} `json:"response"`
	Status string `json:"status"`
}
