package dictionary

import (
	"aer/util/utility/dictionary-iso/data"
	"encoding/json"
)

const dataPath = "./data.json"

var file = data.MustAsset("data/data.json")

func GetAllCountries() ([]Country, error) {
	var countries []Country
	err := json.Unmarshal(file, &countries)
	if err != nil {
		return nil, err
	}
	return countries, nil
}

func GetCountryByAlpha3(a3 string) (Country, error) {
	// if len(ff) == 0 {
	// 	if err := Load(); err != nil {
	// 		return Country{}, err
	// 	}
	// }

	// for _, v := range ff {
	// 	if v.Alpha3 == a3 {
	// 		return v, nil
	// 	}
	// }
	// return Country{}, nil

}
