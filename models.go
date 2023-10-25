package main

// The client model.
type Client struct {
	Login      string
	LoginAlias string
	Email      string
	ID         int
	URL        string
}

// The clients model.
type Clients struct {
	Results []*Client
}
