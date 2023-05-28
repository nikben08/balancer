package models

type Server struct {
	Id      string
	DomainId string
	Address string
	Label   string
	Weight  int
}
