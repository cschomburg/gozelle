package gozelle

type IndexResponse struct {
	Response struct {
		Authkey       string `json:"authkey"`
		ID            int    `json:"id"`
		Notifications struct {
			Messages        int  `json:"messages"`
			NewAnnouncement bool `json:"newAnnouncement"`
			NewBlog         bool `json:"newBlog"`
			Notifications   int  `json:"notifications"`
		} `json:"notifications"`
		Passkey   string `json:"passkey"`
		Username  string `json:"username"`
		Userstats struct {
			Class         string  `json:"class"`
			Downloaded    int     `json:"downloaded"`
			Ratio         float64 `json:"ratio"`
			Requiredratio float64 `json:"requiredratio"`
			Uploaded      int     `json:"uploaded"`
		} `json:"userstats"`
	} `json:"response"`
	Status string `json:"status"`
}
