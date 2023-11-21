package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ///////////////////////////////////////////////////////////////////////////////////////////
func (c *Client) GetPokemon(pokemon_str string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemon_str
	fullURL := baseURL + endpoint

	//check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		fmt.Println("cache HIT!!")
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil) //rep=>repuest
	if err != nil {
		return Pokemon{}, err //LocationAreasResp{} gives back a emty LocationAreasResp
	}
	resp, err := c.httpClient.Do(req) //resp=>reponse
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close() //closes the respObject when function closes

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("Bad Statuscode: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(fullURL, dat)
	return pokemon, nil
}
