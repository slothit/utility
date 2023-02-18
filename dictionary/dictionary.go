package dictionary

import (
	"encoding/json"
	"fmt"
	"github.com/slothit/utility/data"
)

var (
	countries []Country
)

func init() {
	_ = json.Unmarshal(data.MustAsset("data/data.json"), &countries)
	fmt.Println("1")
}

// GetAllCountries return all countries
func GetAllCountries() []Country {
	return countries
}

func GetCountryByAlpha3(a3 string) Country {
	for _, v := range countries {
		if v.Alpha3 == a3 {
			return v
		}
	}
	return Country{}
}
