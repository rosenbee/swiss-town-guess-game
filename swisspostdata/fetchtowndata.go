package towndata

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var HelloTest = "This is just a test"

type TownInfo struct {
	CantonCode string `json:"success"`
	Name       string `json:"statusCode"`
}

type SwissPostTownInfoResponse struct {
	Records []SwissPostTownInfoResponseRecords `json:"records"`
}

type SwissPostTownInfoResponseRecords struct {
	Record SwissPostTownInfoResponseRecord `json:"record"`
}

type SwissPostTownInfoResponseRecord struct {
	Fields SwissPostTownInfoResponseFields `json:"fields"`
}

type SwissPostTownInfoResponseFields struct {
	GemeindeName string `json:"gemeindename"`
	Kanton       string `json:"kanton"`
}

// GetTown gets townInfo of specified bfsnr from the post.ch api.
// If no townInfo was found for specified bfsnr, it will return nil.
func GetTown(bfsnr int, apiKey string) (townInfo *TownInfo, err error) {

	url := "https://swisspost.opendatasoft.com/api/v2/catalog/datasets/politische-gemeinden_v2/records?select=bfsnr%2Cgemeindename%2Ckanton&where=bfsnr%3D"
	url += fmt.Sprintf("%d", bfsnr)
	url += "&limit=10&offset=0&timezone=UTC&apikey="
	url += apiKey

	tr := &http.Transport{}

	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("Could not retrieve town data from swisspost")
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New("Could not establish a connection")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Could not read body from request")
	}

	var responseStruct SwissPostTownInfoResponse

	err = json.Unmarshal(body, &responseStruct)
	if err != nil {
		return nil, errors.New("Could not unmarshal the JSON structure")
	}

	if len(responseStruct.Records) > 1 {
		return nil, errors.New("More than one record was found. This should not happen!")
	}

	if len(responseStruct.Records) == 0 {
		// No record was found. This can happen and is not an error!
		return nil, nil
	}

	var result TownInfo

	// map to return value
	// Through previous validations, we made sure that there is always one record available
	result.CantonCode = responseStruct.Records[0].Record.Fields.Kanton
	result.Name = responseStruct.Records[0].Record.Fields.GemeindeName

	return &result, nil
}
