package repositories

import (
	"balancer/models"
	"log"
	"strconv"

	"github.com/google/uuid"
)

func CreateServer(server models.Server) error {
	_, err := DB.Query("INSERT INTO `servers`(`id`, `domain_id`, `address`, `label`, `weight`) VALUES ('" + uuid.New().String() + "','" + server.DomainId + "','" + server.Address + "', '" + server.Label + "', '" + strconv.Itoa(server.Weight) + "')")
	if err != nil {
		return err
	}
	return nil
}

func GetServersByDomain(domainId string) ([]models.Server, error) {
	servers, err := DB.Query("SELECT * FROM `servers` WHERE domain_id='" + domainId + "'")
	if err != nil {
		return []models.Server{}, err
	}
	defer servers.Close()

	resp := []models.Server{}
	for servers.Next() {
		server := models.Server{}
		err := servers.Scan(&server.Id, &server.DomainId, &server.Address, &server.Label, &server.Weight)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, server)
	}

	err = servers.Err()
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func GetAllServers() ([]models.Server, error) {
	servers, err := DB.Query("SELECT * FROM `servers`")
	if err != nil {
		return []models.Server{}, err
	}
	defer servers.Close()

	resp := []models.Server{}
	for servers.Next() {
		server := models.Server{}
		err := servers.Scan(&server.Id, &server.Address, &server.Label, &server.Weight)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, server)
	}

	err = servers.Err()
	if err != nil {
		log.Fatal(err)
	}

	return resp, nil
}

func GetServerById(serverId string) (models.Server, error) {
	servers, err := DB.Query("SELECT * FROM `servers` WHERE id='" + serverId + "'")
	if err != nil {
		return models.Server{}, err
	}

	resp := []models.Server{}
	for servers.Next() {
		server := models.Server{}
		err := servers.Scan(&server.Id, &server.DomainId, &server.Address, &server.Label, &server.Weight)
		if err != nil {
			log.Fatal(err)
		}
		resp = append(resp, server)
	}

	err = servers.Err()
	if err != nil {
		log.Fatal(err)
	}
	return resp[0], nil
}

func DeleteServer(serverId string) error {
	_, err := DB.Query("DELETE FROM `servers` WHERE `id` ='" + serverId + "'")
	if err != nil {
		return err
	}
	return nil
}
