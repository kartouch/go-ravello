package ravello

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const baseURL = "https://cloud.ravellosystems.com/api/v1"

//Handler takes care of requesting the data to Ravello
//Data is returned in a struct
func handler(verb string, endpoint string, data []byte) (body []byte, err error) {
	b := base64.StdEncoding.EncodeToString([]byte(os.Getenv("RAVELLO_USER") + ":" + os.Getenv("RAVELLO_PWD")))
	req, err := http.NewRequest(verb, baseURL+endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %v", b))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(req.URL, " ", resp.Status)
	defer resp.Body.Close()

	return
}
