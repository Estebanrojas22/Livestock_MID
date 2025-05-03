package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

func Metodo_get_all(host, endpoint string) ([]byte, error) {
	url := beego.AppConfig.String(host) + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcessarJsonArreglos(datos []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	err := json.Unmarshal(datos, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func Metodo_post(hots_servicio, endpoint string, data []byte) ([]byte, error) {
	url := beego.AppConfig.String(hots_servicio)+endpoint
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcessarJson(datos []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(datos, &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func Metodo_put(nombre_servicio string, id string, data []byte) ([]byte, error) {

	baseURL := beego.AppConfig.String(nombre_servicio)

	url := fmt.Sprintf("%s/%s", baseURL, id)

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ConvertInterfaceToSliceMap(input interface{}) ([]map[string]interface{}, error) {
	// Afirmar que es un slice de interface{}
	list, ok := input.([]interface{})
	if !ok {
		return nil, fmt.Errorf("no es un []interface{}")
	}

	// Convertir cada elemento a map[string]interface{}
	var result []map[string]interface{}
	for i, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("elemento %d no es un map[string]interface{}", i)
		}
		result = append(result, m)
	}

	return result, nil
}
