package contracts

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateServerRequest struct {
	DomainId string `json:"domain_id"`
	ServerAddress string `json:"server_address"`
	ServerLabel   string `json:"server_label"`
	ServerWeight  string `json:"server_weight"`
}

type DeleteServerRequest struct {
	ServerId string `json:"server_id"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CreateDomainRequest struct {
	DomainAddress string `json:"domain_address"`
	DomainLabel   string `json:"domain_label"`
}

type DeleteDomainRequest struct {
	DomainId string `json:"domain_id"`
}

type PeersManagerRequest struct{
	DomainId string `json:"domain_id"`
}
