package instagram

// Post represents a single Instagram post.
type Post struct {
	ID   string `json:"-"`
	Text string `json:"text"`
	URL  string `json:"url"`
	SRC  string `json:"src"`
	Like int    `json:"like"`
}

// PostMap represents a map of Instagram post.
type PostMap map[string]Post
