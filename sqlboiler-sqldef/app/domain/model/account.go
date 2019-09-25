package model

// Account model
type Account struct {
	ID string
	AccountID string
}

// NewPost creates a new Post
func (a *Account) NewPost(body string) *Post {
	return &Post{
		AuthorID: a.ID,
		Body: body,
	}
}
