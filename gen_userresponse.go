package gozelle

type UserResponse struct {
	Response struct {
		Avatar    string `json:"avatar"`
		Community struct {
			CollagesContrib int `json:"collagesContrib"`
			CollagesStarted int `json:"collagesStarted"`
			Groups          int `json:"groups"`
			Invited         int `json:"invited"`
			Leeching        int `json:"leeching"`
			PerfectFlacs    int `json:"perfectFlacs"`
			Posts           int `json:"posts"`
			RequestsFilled  int `json:"requestsFilled"`
			RequestsVoted   int `json:"requestsVoted"`
			Seeding         int `json:"seeding"`
			Snatched        int `json:"snatched"`
			TorrentComments int `json:"torrentComments"`
			Uploaded        int `json:"uploaded"`
		} `json:"community"`
		IsFriend bool `json:"isFriend"`
		Personal struct {
			Class        string `json:"class"`
			Donor        bool   `json:"donor"`
			Enabled      bool   `json:"enabled"`
			Paranoia     int    `json:"paranoia"`
			ParanoiaText string `json:"paranoiaText"`
			Passkey      string `json:"passkey"`
			Warned       bool   `json:"warned"`
		} `json:"personal"`
		ProfileText string `json:"profileText"`
		Ranks       struct {
			Artists    int `json:"artists"`
			Bounty     int `json:"bounty"`
			Downloaded int `json:"downloaded"`
			Overall    int `json:"overall"`
			Posts      int `json:"posts"`
			Requests   int `json:"requests"`
			Uploaded   int `json:"uploaded"`
			Uploads    int `json:"uploads"`
		} `json:"ranks"`
		Stats struct {
			Downloaded    int     `json:"downloaded"`
			JoinedDate    string  `json:"joinedDate"`
			LastAccess    string  `json:"lastAccess"`
			Ratio         string  `json:"ratio"`
			RequiredRatio float64 `json:"requiredRatio"`
			Uploaded      int     `json:"uploaded"`
		} `json:"stats"`
		Username string `json:"username"`
	} `json:"response"`
	Status string `json:"status"`
}
