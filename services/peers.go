package services

import (
	"balancer/models"
	"balancer/repositories"
	"bytes"
	"net/http"
	"strconv"
	"fmt"
)

func CreateServer(server models.Server) error {
	err := repositories.CreateServer(server)
	if err != nil {
		return err
	}

	// HTTP endpoint
	posturl := "http://172.18.0.1:8083/add-peer"
	domain, _ := repositories.GetDomainById(server.DomainId)
	// JSON body
	body := []byte(`{"address":"` + server.Address + `", "weight":"` + strconv.Itoa(server.Weight) + `", "domain":"` + domain.Label + `"}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "text/html")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	return nil
}

func GetAllServers() ([]models.Server, error) {
	servers, err := repositories.GetAllServers()
	if err != nil {
		return []models.Server{}, err
	}

	return servers, nil
}

func GetServersByDomain(domainId string) ([]models.Server, error) {
	servers, err := repositories.GetServersByDomain(domainId)
	if err != nil {
		return []models.Server{}, err
	}
	return servers, nil
}

func DeleteServer(serverId string) error {
	fmt.Println(serverId)
	server, err := repositories.GetServerById(serverId)
	if err != nil {
		panic(err)
	}

	posturl := "http://172.18.0.1:8083/delete-peer"

	// JSON body
	body := []byte(`{"address":"` + server.Address +`"}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "text/html")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	err = repositories.DeleteServer(serverId)
	if err != nil {
		return err
	}

	return nil
}
