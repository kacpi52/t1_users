package data_pattern

import "sync"

type UserCredentialsCollection struct {
	Collection []UserCredentials
	sync.Mutex
}
type UserCredentials struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
}

type UserCollection struct {
	Results []UserDetails `json:"results"`
	Info    Info          `json:"info"`
}

type UserDetails struct {
	Gender     string   `json:"gender"`
	Name       Name     `json:"name"`
	Location   Location `json:"location"`
	Email      string   `json:"email"`
	Login      Login    `json:"login"`
	DOB        DateAge  `json:"dob"`
	Registered DateAge  `json:"registered"`
	Phone      string   `json:"phone"`
	Cell       string   `json:"cell"`
	ID         ID       `json:"id"`
	Picture    Picture  `json:"picture"`
	Nat        string   `json:"nat"`
}

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Location struct {
	Street      Street     `json:"street"`
	City        string     `json:"city"`
	State       string     `json:"state"`
	Country     string     `json:"country"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Coordinate struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

type DateAge struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

type ID struct {
	Name  string  `json:"name"`
	Value *string `json:"value"`
}

type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}
