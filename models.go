package main

// The client model.
type Client struct {
	Login      string
	LoginAlias string
	Email      string
	ID         int
}

// The clients model.
type Clients struct {
	Results []*Client
}
