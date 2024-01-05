package user

type Notes struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	AuthorID string `json:"authorID"` // the author (the user who created this notes)
}
