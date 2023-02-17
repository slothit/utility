package dictionary

type Country struct {
	Alpha2       string       `json:"alpha2"`
	Alpha3       string       `json:"alpha3"`
	Numeric      string       `json:"numeric"`
	Currencies   []Currencies `json:"currencies"`
	DialingCodes []string     `json:"dialingCodes"`
	Languages    []Languages  `json:"languages"`
	Name         string       `json:"name"`
}

type Currencies struct {
	Alpha3   string `json:"alpha3"`
	Numeric  string `json:"numeric"`
	Decimals int    `json:"decimals"`
	Name     string `json:"name"`
}
type Languages struct {
	Alpha2 string `json:"alpha2"`
	Alpha3 string `json:"alpha3"`
	Name   string `json:"name"`
}
