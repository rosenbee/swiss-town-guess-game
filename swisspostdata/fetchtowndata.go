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
	BFSNr        int    `json:"bfsnr"`
	GemeindeName string `json:"gemeindename"`
	Kanton       string `json:"kanton"`
}

func GetTown(bfsnr int, apiKey string) (townInfo *TownInfo, err error) {

	url := "https://swisspost.opendatasoft.com/api/v2/catalog/datasets/politische-gemeinden_v2/records?select=bfsnr%2Cgemeindename%2Ckanton&where=bfsnr%3D"
	url += fmt.Sprintf("%d", bfsnr)
	url += "&limit=10&offset=0&timezone=UTC&apikey="
	url += apiKey

	/* perhaps needed
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	*/

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

	fmt.Println(string(body))

	var result SwissPostTownInfoResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New("Could not unmarshal the JSON structure")
	}

	fmt.Println(result)

	// TODO
	return nil, nil
}
