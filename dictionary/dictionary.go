package dictionary

import (
	"encoding/json"
	"os"
)

const dataPath = "data/data.json"

var ff []Country

func Load() error {
	b, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &ff)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCountries() ([]Country, error) {
	if len(ff) == 0 {
		if err := Load(); err != nil {
			return nil, err
		}
	}
	return ff, nil
}

func GetCountryByAlpha3(a3 string) (Country, error) {
	if len(ff) == 0 {
		if err := Load(); err != nil {
			return Country{}, err
		}
	}

	for _, v := range ff {
		if v.Alpha3 == a3 {
			return v, nil
		}
	}

	
}
