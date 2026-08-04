package pokeapi

import (
	"encoding/json"
	"io"
)

func (c *Client) GetData(url *string) (Response, error) {

	cachedResults, ok := c.cache.Get(*url)

	results := Response{}

	if ok {
		err := json.Unmarshal(cachedResults, &results)
		if err != nil {
			return Response{}, err
		}
		return results, nil
	}

	res, err := c.httpClient.Get(*url)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, err
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		return Response{}, err
	}

	c.cache.Add(*url, data)

	return results, nil

}

func (c *Client) GetPokemon(area string) (PokemonList, error) {

	const baseUrl = "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + area + "/"
	cachedResults, ok := c.cache.Get(url)

	results := PokemonList{}

	if ok {
		err := json.Unmarshal(cachedResults, &results)
		if err != nil {
			return PokemonList{}, err
		}
		return results, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return PokemonList{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonList{}, err
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		return PokemonList{}, err
	}

	c.cache.Add(url, data)

	return results, nil

}

func (c *Client) CatchPokemon(name string) (Pokemon, error) {

	const baseUrl = "https://pokeapi.co/api/v2/pokemon/"
	url := baseUrl + name + "/"
	cachedResults, ok := c.cache.Get(url)

	results := Pokemon{}

	if ok {
		err := json.Unmarshal(cachedResults, &results)
		if err != nil {
			return Pokemon{}, err
		}
		return results, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &results)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return results, nil

}
