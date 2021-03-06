package ravello

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const baseURL = "https://cloud.ravellosystems.com/api/v1"

//Handler takes care of requesting the data to Ravello
func handler(method string, endpoint string, data []byte) (body []byte, err error) {
	req, err := http.NewRequest(method, baseURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.SetBasicAuth(os.Getenv("RAVELLO_USER"), os.Getenv("RAVELLO_PWD"))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Connection", "close")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if os.Getenv("RAVELLO_DEBUG") == "1" {
		log.Println(resp)
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	return
}
