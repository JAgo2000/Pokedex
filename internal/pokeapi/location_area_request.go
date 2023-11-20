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

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache HIT!!")
		locationAreasResp := LocationAreasResp{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
		}
		return locationAreasResp, nil
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

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	c.cache.Add(fullURL, dat)
	return locationAreasResp, nil
}

// ///////////////////////////////////////////////////////////////////////////////////////////
func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache HIT!!")
		locationArea := LocationArea{}
		err := json.Unmarshal(dat, &locationArea)
		if err != nil {
			return LocationArea{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil) //rep=>repuest
	if err != nil {
		return LocationArea{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	resp, err := c.httpClient.Do(req) //resp=>reponse
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close() //closes the respObject when function closes

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("Bad Statuscode: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err //LocationArea{} gives back a emty LocationArea
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(dat, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, dat)
	return locationArea, nil
}
