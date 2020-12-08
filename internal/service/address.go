package service

import (
	"encoding/json"
	"go-user/internal/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AddressResponse struct {
	Status  int
	Message string
	Data    []model.Address
}

func GetAddressUser(userID int) ([]model.Address, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Load env failed")
	}

	addressService := os.Getenv("ADDRESS_SERVICE")

	res, err := http.Get(addressService + "/user_address?userID=" + strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return UserAddressResponse(res.Body)
}

func UserAddressResponse(body io.ReadCloser) ([]model.Address, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var addressRes AddressResponse
	json.Unmarshal(data, &addressRes)

	return addressRes.Data, nil
}
