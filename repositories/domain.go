package repositories

import (
	"balancer/models"
	"log"
	"github.com/google/uuid"
)

func CreateDomain(domain models.Domain) error {
	_, err := DB.Query("INSERT INTO `domains`(`id`, `label`) VALUES ('" + uuid.New().String() + "','" + domain.Label + "')")
	if err != nil {
		return err
	}
	return nil
}

func GetAllDomains() ([]models.Domain, error) {
	domains, err := DB.Query("SELECT * FROM `domains`")
	if err != nil {
		return []models.Domain{}, err
	}
	defer domains.Close()

	resp := []models.Domain{}
	for domains.Next() {
		domain := models.Domain{}
		err := domains.Scan(&domain.Id, &domain.Label)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, domain)
	}

	err = domains.Err()
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func GetDomainById(domainId string) (models.Domain, error) {
	domains, err := DB.Query("SELECT * FROM `domains` WHERE id='" + domainId + "'")
	if err != nil {
		return models.Domain{}, err
	}

	resp := []models.Domain{}
	for domains.Next() {
		domain := models.Domain{}
		err := domains.Scan(&domain.Id, &domain.Label)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, domain)
	}

	err = domains.Err()
	if err != nil {
		log.Fatal(err)
	}
	return resp[0], nil
}

func WhetherDomainExists(domainLabel string) (bool, error) {
	domains, err := DB.Query("SELECT * FROM `domains` WHERE label='" + domainLabel + "'")
	if err != nil {
		return false, err
	}

	resp := []models.Domain{}
	for domains.Next() {
		domain := models.Domain{}
		err := domains.Scan(&domain.Id, &domain.Label)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, domain)
	}

	err = domains.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	if len(resp) == 0 {
		return false, nil
	}
	return true, nil
}

func DeleteDomain(domainId string) error {
	_, err := DB.Query("DELETE FROM `domains` WHERE `id` ='" + domainId + "'")
	if err != nil {
		return err
	}
	return nil
}
