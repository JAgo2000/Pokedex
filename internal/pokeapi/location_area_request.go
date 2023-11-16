package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil) //rep=>repuest
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	resp, err := c.httpClient.Do(req) //resp=>reponse
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	defer resp.Body.Close() //closes the respObject when function closes

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("Bad Statuscode: %v", resp.StatusCode)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	return locationAreasResp, nil
}
