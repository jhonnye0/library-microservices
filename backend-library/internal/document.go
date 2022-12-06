package internal

type Document struct {
	ThumbnailUrl     string   `json:"thumbnailUrl"`
	ShortDescription string   `json:"shortDescription"`
	LongDescription  string   `json:"longDescription"`
	Title            string   `json:"title"`
	Authors          []string `json:"author"`
	Categories       []string `json:"categories"`
	Watermark        string   `json:"watermark,omitempty"`
	ISBN             string   `json:"isbn"`
	PageCount        int      `json:"pageCount"`
	PublishedDate    struct {
		Date string `json:"$date"`
	} `json:"publishedDate"`
	Status string `json:"status"`
}

type Filter struct {
	Key string `json:"key"`

	// If value is empty, just return everything but sorted with the key
	Value string `json:"value,omitempty"`
}

type Status string

const (
	Pending    Status = "Pending"
	Started    Status = "Started"
	InProgress Status = "InProgress"
	Finished   Status = "Finished"
	Failed     Status = "Failed"
)
