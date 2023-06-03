package services

import (
	"balancer/models"
	"balancer/repositories"
	"bytes"
	"net/http"
	"errors"
)

func CreateDomain(domain models.Domain) error {
	err := repositories.CreateDomain(domain)
	if err != nil {
		return err
	}

	// HTTP endpoint
	posturl := "http://172.0.0.1:8083/add-domain"

	// JSON body
	body := []byte(`{"domain":"` + domain.Label + `"}`)

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

func GetAllDomains() ([]models.Domain, error) {
	domains, err := repositories.GetAllDomains()
	if err != nil {
		return []models.Domain{}, err
	}

	return domains, nil
}

func DeleteDomain(domainId string) error {
	servers, _ := repositories.GetServersByDomain(domainId)
	if len(servers) != 0 {
		return errors.New("To delete domain, delete domain's servers first")
	}
	domain, err := repositories.GetDomainById(domainId)
	if err != nil {
		panic(err)
	}

	posturl := "http://172.0.0.1:8083/delete-domain"

	// JSON body
	body := []byte(`{"domain":"` + domain.Label + `"}`)

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

	err = repositories.DeleteDomain(domainId)
	if err != nil {
		return err
	}

	return nil
}
