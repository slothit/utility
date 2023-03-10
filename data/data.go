// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/data.json
package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataDataJson = []byte(`[{
	"alpha2": "MW",
	"alpha3": "MWI",
	"numeric": "454",
	"currencies": [{
		"alpha3": "MWK",
		"numeric": "454",
		"decimals": 2,
		"name": "Malawian kwacha"
	}],
	"dialing_codes": ["+265"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "ny",
		"alpha3": "nya",
		"name": "Chewa"
	}],
	"name": "Malawi"
}, {
	"alpha2": "SC",
	"alpha3": "SYC",
	"numeric": "690",
	"currencies": [{
		"alpha3": "SCR",
		"numeric": "690",
		"decimals": 2,
		"name": "Seychelles rupee"
	}],
	"dialing_codes": ["+248"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Seychelles"
}, {
	"alpha2": "TV",
	"alpha3": "TUV",
	"numeric": "798",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+688"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Tuvalu"
}, {
	"alpha2": "GL",
	"alpha3": "GRL",
	"numeric": "304",
	"currencies": [{
		"alpha3": "DKK",
		"numeric": "208",
		"decimals": 2,
		"name": "Danish krone"
	}],
	"dialing_codes": ["+299"],
	"languages": [{
		"alpha2": "kl",
		"alpha3": "kal",
		"name": "Greenlandic"
	}],
	"name": "Greenland"
}, {
	"alpha2": "GS",
	"alpha3": "SGS",
	"numeric": "239",
	"currencies": [{
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}],
	"dialing_codes": [],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "South Georgia And The South Sandwich Islands"
}, {
	"alpha2": "MU",
	"alpha3": "MUS",
	"numeric": "480",
	"currencies": [{
		"alpha3": "MUR",
		"numeric": "480",
		"decimals": 2,
		"name": "Mauritian rupee"
	}],
	"dialing_codes": ["+230"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Mauritius"
}, {
	"alpha2": "TW",
	"alpha3": "TWN",
	"numeric": "158",
	"currencies": [{
		"alpha3": "TWD",
		"numeric": "901",
		"decimals": 2,
		"name": "New Taiwan dollar"
	}],
	"dialing_codes": ["+886"],
	"languages": [{
		"alpha2": "zh",
		"alpha3": "zho",
		"name": "Chinese"
	}],
	"name": "Taiwan"
}, {
	"alpha2": "CG",
	"alpha3": "COG",
	"numeric": "178",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+242"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "ln",
		"alpha3": "lin",
		"name": "Lingala"
	}],
	"name": "Republic Of Congo"
}, {
	"alpha2": "KW",
	"alpha3": "KWT",
	"numeric": "414",
	"currencies": [{
		"alpha3": "KWD",
		"numeric": "414",
		"decimals": 3,
		"name": "Kuwaiti dinar"
	}],
	"dialing_codes": ["+965"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Kuwait"
}, {
	"alpha2": "SG",
	"alpha3": "SGP",
	"numeric": "702",
	"currencies": [{
		"alpha3": "SGD",
		"numeric": "702",
		"decimals": 2,
		"name": "Singapore dollar"
	}],
	"dialing_codes": ["+65"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "zh",
		"alpha3": "zho",
		"name": "Chinese"
	}, {
		"alpha2": "ms",
		"alpha3": "msa",
		"name": "Malay"
	}, {
		"alpha2": "ta",
		"alpha3": "tam",
		"name": "Tamil"
	}],
	"name": "Singapore"
}, {
	"alpha2": "TH",
	"alpha3": "THA",
	"numeric": "764",
	"currencies": [{
		"alpha3": "THB",
		"numeric": "764",
		"decimals": 2,
		"name": "Thai baht"
	}],
	"dialing_codes": ["+66"],
	"languages": [{
		"alpha2": "th",
		"alpha3": "tha",
		"name": "Thai"
	}],
	"name": "Thailand"
}, {
	"alpha2": "BY",
	"alpha3": "BLR",
	"numeric": "112",
	"currencies": [{
		"alpha3": "BYR",
		"numeric": "974",
		"decimals": 0,
		"name": "Belarusian ruble"
	}],
	"dialing_codes": ["+375"],
	"languages": [{
		"alpha2": "be",
		"alpha3": "bel",
		"name": "Belarusian"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Belarus"
}, {
	"alpha2": "EC",
	"alpha3": "ECU",
	"numeric": "218",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+593"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}, {
		"alpha2": "qu",
		"alpha3": "que",
		"name": "Quechua"
	}],
	"name": "Ecuador"
}, {
	"alpha2": "GI",
	"alpha3": "GIB",
	"numeric": "292",
	"currencies": [{
		"alpha3": "GIP",
		"numeric": "292",
		"decimals": 2,
		"name": "Gibraltar pound"
	}],
	"dialing_codes": ["+350"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Gibraltar"
}, {
	"alpha2": "GM",
	"alpha3": "GMB",
	"numeric": "270",
	"currencies": [{
		"alpha3": "GMD",
		"numeric": "270",
		"decimals": 2,
		"name": "Gambian dalasi"
	}],
	"dialing_codes": ["+220"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Gambia"
}, {
	"alpha2": "KN",
	"alpha3": "KNA",
	"numeric": "659",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 869"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Saint Kitts And Nevis"
}, {
	"alpha2": "MR",
	"alpha3": "MRT",
	"numeric": "478",
	"currencies": [{
		"alpha3": "MRO",
		"numeric": "478",
		"decimals": 0,
		"name": "Mauritanian ouguiya"
	}],
	"dialing_codes": ["+222"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Mauritania"
}, {
	"alpha2": "CX",
	"alpha3": "CXR",
	"numeric": "162",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+61"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Christmas Island"
}, {
	"alpha2": "HT",
	"alpha3": "HTI",
	"numeric": "332",
	"currencies": [{
		"alpha3": "HTG",
		"numeric": "332",
		"decimals": 2,
		"name": "Haitian gourde"
	}, {
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+509"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "ht",
		"alpha3": "hat",
		"name": "Haitian"
	}],
	"name": "Haiti"
}, {
	"alpha2": "OM",
	"alpha3": "OMN",
	"numeric": "512",
	"currencies": [{
		"alpha3": "OMR",
		"numeric": "512",
		"decimals": 3,
		"name": "Omani rial"
	}],
	"dialing_codes": ["+968"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Oman"
}, {
	"alpha2": "TF",
	"alpha3": "ATF",
	"numeric": "260",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": [],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "French Southern Territories"
}, {
	"alpha2": "EG",
	"alpha3": "EGY",
	"numeric": "818",
	"currencies": [{
		"alpha3": "EGP",
		"numeric": "818",
		"decimals": 2,
		"name": "Egyptian pound"
	}],
	"dialing_codes": ["+20"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Egypt"
}, {
	"alpha2": "MC",
	"alpha3": "MCO",
	"numeric": "492",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+377"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Monaco"
}, {
	"alpha2": "VE",
	"alpha3": "VEN",
	"numeric": "862",
	"currencies": [{
		"alpha3": "VEF",
		"numeric": "937",
		"decimals": 2,
		"name": "Venezuelan bol??var fuerte"
	}],
	"dialing_codes": ["+58"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Venezuela, Bolivarian Republic Of"
}, {
	"alpha2": "AM",
	"alpha3": "ARM",
	"numeric": "051",
	"currencies": [{
		"alpha3": "AMD",
		"numeric": "51",
		"decimals": 2,
		"name": "Armenian dram"
	}],
	"dialing_codes": ["+374"],
	"languages": [{
		"alpha2": "hy",
		"alpha3": "hye",
		"name": "Armenian"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Armenia"
}, {
	"alpha2": "CF",
	"alpha3": "CAF",
	"numeric": "140",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+236"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "sg",
		"alpha3": "sag",
		"name": "Sango"
	}],
	"name": "Central African Republic"
}, {
	"alpha2": "KG",
	"alpha3": "KGZ",
	"numeric": "417",
	"currencies": [{
		"alpha3": "KGS",
		"numeric": "417",
		"decimals": 2,
		"name": "Kyrgyzstani som"
	}],
	"dialing_codes": ["+996"],
	"languages": [{
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Kyrgyzstan"
}, {
	"alpha2": "TM",
	"alpha3": "TKM",
	"numeric": "795",
	"currencies": [{
		"alpha3": "TMT",
		"numeric": "934",
		"decimals": 2,
		"name": "Turkmenistani manat"
	}],
	"dialing_codes": ["+993"],
	"languages": [{
		"alpha2": "tk",
		"alpha3": "tuk",
		"name": "Turkmen"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Turkmenistan"
}, {
	"alpha2": "BD",
	"alpha3": "BGD",
	"numeric": "050",
	"currencies": [{
		"alpha3": "BDT",
		"numeric": "50",
		"decimals": 2,
		"name": "Bangladeshi taka"
	}],
	"dialing_codes": ["+880"],
	"languages": [{
		"alpha2": "bn",
		"alpha3": "ben",
		"name": "Bengali"
	}],
	"name": "Bangladesh"
}, {
	"alpha2": "ID",
	"alpha3": "IDN",
	"numeric": "360",
	"currencies": [{
		"alpha3": "IDR",
		"numeric": "360",
		"decimals": 2,
		"name": "Indonesian rupiah"
	}],
	"dialing_codes": ["+62"],
	"languages": [{
		"alpha2": "id",
		"alpha3": "ind",
		"name": "Indonesian"
	}],
	"name": "Indonesia"
}, {
	"alpha2": "BN",
	"alpha3": "BRN",
	"numeric": "096",
	"currencies": [{
		"alpha3": "BND",
		"numeric": "96",
		"decimals": 2,
		"name": "Brunei dollar"
	}],
	"dialing_codes": ["+673"],
	"languages": [{
		"alpha2": "ms",
		"alpha3": "msa",
		"name": "Malay"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Brunei Darussalam"
}, {
	"alpha2": "FM",
	"alpha3": "FSM",
	"numeric": "583",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+691"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Micronesia, Federated States Of"
}, {
	"alpha2": "MM",
	"alpha3": "MMR",
	"numeric": "104",
	"currencies": [{
		"alpha3": "MMK",
		"numeric": "104",
		"decimals": 0,
		"name": "Myanma kyat"
	}],
	"dialing_codes": ["+95"],
	"languages": [{
		"alpha2": "my",
		"alpha3": "mya",
		"name": "Burmese"
	}],
	"name": "Myanmar"
}, {
	"alpha2": "PY",
	"alpha3": "PRY",
	"numeric": "600",
	"currencies": [{
		"alpha3": "PYG",
		"numeric": "600",
		"decimals": 0,
		"name": "Paraguayan guaran??"
	}],
	"dialing_codes": ["+595"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Paraguay"
}, {
	"alpha2": "RO",
	"alpha3": "ROU",
	"numeric": "642",
	"currencies": [{
		"alpha3": "RON",
		"numeric": "946",
		"decimals": 2,
		"name": "Romanian new leu"
	}],
	"dialing_codes": ["+40"],
	"languages": [{
		"alpha2": "ro",
		"alpha3": "ron",
		"name": "Moldavian"
	}],
	"name": "Romania"
}, {
	"alpha2": "UM",
	"alpha3": "UMI",
	"numeric": "581",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "United States Minor Outlying Islands"
}, {
	"alpha2": "VN",
	"alpha3": "VNM",
	"numeric": "704",
	"currencies": [{
		"alpha3": "VND",
		"numeric": "704",
		"decimals": 0,
		"name": "Vietnamese dong"
	}],
	"dialing_codes": ["+84"],
	"languages": [{
		"alpha2": "vi",
		"alpha3": "vie",
		"name": "Vietnamese"
	}],
	"name": "Viet Nam"
}, {
	"alpha2": "CK",
	"alpha3": "COK",
	"numeric": "184",
	"currencies": [{
		"alpha3": "NZD",
		"numeric": "554",
		"decimals": 2,
		"name": "New Zealand dollar"
	}],
	"dialing_codes": ["+682"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "mi",
		"alpha3": "mri",
		"name": "Maori"
	}],
	"name": "Cook Islands"
}, {
	"alpha2": "CW",
	"alpha3": "CUW",
	"numeric": "531",
	"currencies": [{
		"alpha3": "ANG",
		"numeric": "532",
		"decimals": 2,
		"name": "Netherlands Antillean guilder"
	}],
	"dialing_codes": ["+599"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Curacao"
}, {
	"alpha2": "ST",
	"alpha3": "STP",
	"numeric": "678",
	"currencies": [{
		"alpha3": "STD",
		"numeric": "678",
		"decimals": 0,
		"name": "S??o Tom?? and Pr??ncipe dobra"
	}],
	"dialing_codes": ["+239"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Sao Tome and Principe"
}, {
	"alpha2": "TZ",
	"alpha3": "TZA",
	"numeric": "834",
	"currencies": [{
		"alpha3": "TZS",
		"numeric": "834",
		"decimals": 2,
		"name": "Tanzanian shilling"
	}],
	"dialing_codes": ["+255"],
	"languages": [{
		"alpha2": "sw",
		"alpha3": "swa",
		"name": "Swahili"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Tanzania, United Republic Of"
}, {
	"alpha2": "AO",
	"alpha3": "AGO",
	"numeric": "024",
	"currencies": [{
		"alpha3": "AOA",
		"numeric": "973",
		"decimals": 2,
		"name": "Angolan kwanza"
	}],
	"dialing_codes": ["+244"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Angola"
}, {
	"alpha2": "CH",
	"alpha3": "CHE",
	"numeric": "756",
	"currencies": [{
		"alpha3": "CHF",
		"numeric": "756",
		"decimals": 2,
		"name": "Swiss franc"
	}, {
		"alpha3": "CHE",
		"numeric": "947",
		"decimals": 2,
		"name": "WIR Euro (complementary currency)"
	}, {
		"alpha3": "CHW",
		"numeric": "948",
		"decimals": 2,
		"name": "WIR Franc (complementary currency)"
	}],
	"dialing_codes": ["+41"],
	"languages": [{
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "it",
		"alpha3": "ita",
		"name": "Italian"
	}, {
		"alpha2": "rm",
		"alpha3": "roh",
		"name": "Romansh"
	}],
	"name": "Switzerland"
}, {
	"alpha2": "PF",
	"alpha3": "PYF",
	"numeric": "258",
	"currencies": [{
		"alpha3": "XPF",
		"numeric": "953",
		"decimals": 0,
		"name": "CFP franc"
	}],
	"dialing_codes": ["+689"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "French Polynesia"
}, {
	"alpha2": "SZ",
	"alpha3": "SWZ",
	"numeric": "748",
	"currencies": [{
		"alpha3": "SZL",
		"numeric": "748",
		"decimals": 2,
		"name": "Swazi lilangeni"
	}],
	"dialing_codes": ["+268"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "ss",
		"alpha3": "ssw",
		"name": "Swati"
	}],
	"name": "Swaziland"
}, {
	"alpha2": "BZ",
	"alpha3": "BLZ",
	"numeric": "084",
	"currencies": [{
		"alpha3": "BZD",
		"numeric": "84",
		"decimals": 2,
		"name": "Belize dollar"
	}],
	"dialing_codes": ["+501"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Belize"
}, {
	"alpha2": "JM",
	"alpha3": "JAM",
	"numeric": "388",
	"currencies": [{
		"alpha3": "JMD",
		"numeric": "388",
		"decimals": 2,
		"name": "Jamaican dollar"
	}],
	"dialing_codes": ["+1 876"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Jamaica"
}, {
	"alpha2": "MA",
	"alpha3": "MAR",
	"numeric": "504",
	"currencies": [{
		"alpha3": "MAD",
		"numeric": "504",
		"decimals": 2,
		"name": "Moroccan dirham"
	}],
	"dialing_codes": ["+212"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Morocco"
}, {
	"alpha2": "MF",
	"alpha3": "MAF",
	"numeric": "663",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+590"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Saint Martin"
}, {
	"alpha2": "NA",
	"alpha3": "NAM",
	"numeric": "516",
	"currencies": [{
		"alpha3": "NAD",
		"numeric": "516",
		"decimals": 2,
		"name": "Namibian dollar"
	}, {
		"alpha3": "ZAR",
		"numeric": "710",
		"decimals": 2,
		"name": "South African rand"
	}],
	"dialing_codes": ["+264"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Namibia"
}, {
	"alpha2": "NE",
	"alpha3": "NER",
	"numeric": "562",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+227"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Niger"
}, {
	"alpha2": "BW",
	"alpha3": "BWA",
	"numeric": "072",
	"currencies": [{
		"alpha3": "BWP",
		"numeric": "72",
		"decimals": 2,
		"name": "Botswana pula"
	}],
	"dialing_codes": ["+267"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "tn",
		"alpha3": "tsn",
		"name": "Tswana"
	}],
	"name": "Botswana"
}, {
	"alpha2": "IN",
	"alpha3": "IND",
	"numeric": "356",
	"currencies": [{
		"alpha3": "INR",
		"numeric": "356",
		"decimals": 2,
		"name": "Indian rupee"
	}],
	"dialing_codes": ["+91"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "hi",
		"alpha3": "hin",
		"name": "Hindi"
	}],
	"name": "India"
}, {
	"alpha2": "NF",
	"alpha3": "NFK",
	"numeric": "574",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+672"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Norfolk Island"
}, {
	"alpha2": "NU",
	"alpha3": "NIU",
	"numeric": "570",
	"currencies": [{
		"alpha3": "NZD",
		"numeric": "554",
		"decimals": 2,
		"name": "New Zealand dollar"
	}],
	"dialing_codes": ["+683"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Niue"
}, {
	"alpha2": "RU",
	"alpha3": "RUS",
	"numeric": "643",
	"currencies": [{
		"alpha3": "RUB",
		"numeric": "643",
		"decimals": 2,
		"name": "Russian rouble"
	}],
	"dialing_codes": ["+7", "+7 3", "+7 4", "+7 8"],
	"languages": [{
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Russian Federation"
}, {
	"alpha2": "SB",
	"alpha3": "SLB",
	"numeric": "090",
	"currencies": [{
		"alpha3": "SBD",
		"numeric": "90",
		"decimals": 2,
		"name": "Solomon Islands dollar"
	}],
	"dialing_codes": ["+677"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Solomon Islands"
}, {
	"alpha2": "AE",
	"alpha3": "ARE",
	"numeric": "784",
	"currencies": [{
		"alpha3": "AED",
		"numeric": "784",
		"decimals": 2,
		"name": "United Arab Emirates dirham"
	}],
	"dialing_codes": ["+971"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "United Arab Emirates"
}, {
	"alpha2": "GF",
	"alpha3": "GUF",
	"numeric": "254",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+594"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "French Guiana"
}, {
	"alpha2": "NI",
	"alpha3": "NIC",
	"numeric": "558",
	"currencies": [{
		"alpha3": "NIO",
		"numeric": "558",
		"decimals": 2,
		"name": "Nicaraguan c??rdoba"
	}],
	"dialing_codes": ["+505"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Nicaragua"
}, {
	"alpha2": "PT",
	"alpha3": "PRT",
	"numeric": "620",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+351"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Portugal"
}, {
	"alpha2": "KI",
	"alpha3": "KIR",
	"numeric": "296",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+686"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Kiribati"
}, {
	"alpha2": "LK",
	"alpha3": "LKA",
	"numeric": "144",
	"currencies": [{
		"alpha3": "LKR",
		"numeric": "144",
		"decimals": 2,
		"name": "Sri Lankan rupee"
	}],
	"dialing_codes": ["+94"],
	"languages": [{
		"alpha2": "si",
		"alpha3": "sin",
		"name": "Sinhala"
	}, {
		"alpha2": "ta",
		"alpha3": "tam",
		"name": "Tamil"
	}],
	"name": "Sri Lanka"
}, {
	"alpha2": "BT",
	"alpha3": "BTN",
	"numeric": "064",
	"currencies": [{
		"alpha3": "INR",
		"numeric": "356",
		"decimals": 2,
		"name": "Indian rupee"
	}, {
		"alpha3": "BTN",
		"numeric": "64",
		"decimals": 2,
		"name": "Bhutanese ngultrum"
	}],
	"dialing_codes": ["+975"],
	"languages": [{
		"alpha2": "dz",
		"alpha3": "dzo",
		"name": "Dzongkha"
	}],
	"name": "Bhutan"
}, {
	"alpha2": "GT",
	"alpha3": "GTM",
	"numeric": "320",
	"currencies": [{
		"alpha3": "GTQ",
		"numeric": "320",
		"decimals": 2,
		"name": "Guatemalan quetzal"
	}],
	"dialing_codes": ["+502"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Guatemala"
}, {
	"alpha2": "ES",
	"alpha3": "ESP",
	"numeric": "724",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+34"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Spain"
}, {
	"alpha2": "QA",
	"alpha3": "QAT",
	"numeric": "634",
	"currencies": [{
		"alpha3": "QAR",
		"numeric": "634",
		"decimals": 2,
		"name": "Qatari riyal"
	}],
	"dialing_codes": ["+974"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Qatar"
}, {
	"alpha2": "BE",
	"alpha3": "BEL",
	"numeric": "056",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+32"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}],
	"name": "Belgium"
}, {
	"alpha2": "GA",
	"alpha3": "GAB",
	"numeric": "266",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+241"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Gabon"
}, {
	"alpha2": "NG",
	"alpha3": "NGA",
	"numeric": "566",
	"currencies": [{
		"alpha3": "NGN",
		"numeric": "566",
		"decimals": 2,
		"name": "Nigerian naira"
	}],
	"dialing_codes": ["+234"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Nigeria"
}, {
	"alpha2": "NP",
	"alpha3": "NPL",
	"numeric": "524",
	"currencies": [{
		"alpha3": "NPR",
		"numeric": "524",
		"decimals": 2,
		"name": "Nepalese rupee"
	}],
	"dialing_codes": ["+977"],
	"languages": [{
		"alpha2": "ne",
		"alpha3": "nep",
		"name": "Nepali"
	}],
	"name": "Nepal"
}, {
	"alpha2": "NZ",
	"alpha3": "NZL",
	"numeric": "554",
	"currencies": [{
		"alpha3": "NZD",
		"numeric": "554",
		"decimals": 2,
		"name": "New Zealand dollar"
	}],
	"dialing_codes": ["+64"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "New Zealand"
}, {
	"alpha2": "TL",
	"alpha3": "TLS",
	"numeric": "626",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+670"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Timor-Leste, Democratic Republic of"
}, {
	"alpha2": "ZM",
	"alpha3": "ZMB",
	"numeric": "894",
	"currencies": [{
		"alpha3": "ZMW",
		"numeric": "967",
		"decimals": 2,
		"name": "Zambian kwacha"
	}],
	"dialing_codes": ["+260"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Zambia"
}, {
	"alpha2": "DK",
	"alpha3": "DNK",
	"numeric": "208",
	"currencies": [{
		"alpha3": "DKK",
		"numeric": "208",
		"decimals": 2,
		"name": "Danish krone"
	}],
	"dialing_codes": ["+45"],
	"languages": [{
		"alpha2": "da",
		"alpha3": "dan",
		"name": "Danish"
	}],
	"name": "Denmark"
}, {
	"alpha2": "GW",
	"alpha3": "GNB",
	"numeric": "624",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+245"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Guinea-bissau"
}, {
	"alpha2": "HK",
	"alpha3": "HKG",
	"numeric": "344",
	"currencies": [{
		"alpha3": "HKD",
		"numeric": "344",
		"decimals": 2,
		"name": "Hong Kong dollar"
	}],
	"dialing_codes": ["+852"],
	"languages": [{
		"alpha2": "zh",
		"alpha3": "zho",
		"name": "Chinese"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Hong Kong"
}, {
	"alpha2": "LU",
	"alpha3": "LUX",
	"numeric": "442",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+352"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}, {
		"alpha2": "lb",
		"alpha3": "ltz",
		"name": "Letzeburgesch"
	}],
	"name": "Luxembourg"
}, {
	"alpha2": "MD",
	"alpha3": "MDA",
	"numeric": "498",
	"currencies": [{
		"alpha3": "MDL",
		"numeric": "498",
		"decimals": 2,
		"name": "Moldovan leu"
	}],
	"dialing_codes": ["+373"],
	"languages": [{
		"alpha2": "ro",
		"alpha3": "ron",
		"name": "Moldavian"
	}],
	"name": "Moldova"
}, {
	"alpha2": "AF",
	"alpha3": "AFG",
	"numeric": "004",
	"currencies": [{
		"alpha3": "AFN",
		"numeric": "971",
		"decimals": 2,
		"name": "Afghan afghani"
	}],
	"dialing_codes": ["+93"],
	"languages": [{
		"alpha2": "ps",
		"alpha3": "pus",
		"name": "Pashto"
	}],
	"name": "Afghanistan"
}, {
	"alpha2": "ER",
	"alpha3": "ERI",
	"numeric": "232",
	"currencies": [{
		"alpha3": "ERN",
		"numeric": "232",
		"decimals": 2,
		"name": "Eritrean nakfa"
	}],
	"dialing_codes": ["+291"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "ti",
		"alpha3": "tir",
		"name": "Tigrinya"
	}],
	"name": "Eritrea"
}, {
	"alpha2": "VU",
	"alpha3": "VUT",
	"numeric": "548",
	"currencies": [{
		"alpha3": "VUV",
		"numeric": "548",
		"decimals": 0,
		"name": "Vanuatu vatu"
	}],
	"dialing_codes": ["+678"],
	"languages": [{
		"alpha2": "bi",
		"alpha3": "bis",
		"name": "Bislama"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Vanuatu"
}, {
	"alpha2": "CR",
	"alpha3": "CRI",
	"numeric": "188",
	"currencies": [{
		"alpha3": "CRC",
		"numeric": "188",
		"decimals": 2,
		"name": "Costa Rican colon"
	}],
	"dialing_codes": ["+506"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Costa Rica"
}, {
	"alpha2": "IT",
	"alpha3": "ITA",
	"numeric": "380",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+39"],
	"languages": [{
		"alpha2": "it",
		"alpha3": "ita",
		"name": "Italian"
	}],
	"name": "Italy"
}, {
	"alpha2": "JP",
	"alpha3": "JPN",
	"numeric": "392",
	"currencies": [{
		"alpha3": "JPY",
		"numeric": "392",
		"decimals": 0,
		"name": "Japanese yen"
	}],
	"dialing_codes": ["+81"],
	"languages": [{
		"alpha2": "ja",
		"alpha3": "jpn",
		"name": "Japanese"
	}],
	"name": "Japan"
}, {
	"alpha2": "SM",
	"alpha3": "SMR",
	"numeric": "674",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+378"],
	"languages": [{
		"alpha2": "it",
		"alpha3": "ita",
		"name": "Italian"
	}],
	"name": "San Marino"
}, {
	"alpha2": "BR",
	"alpha3": "BRA",
	"numeric": "076",
	"currencies": [{
		"alpha3": "BRL",
		"numeric": "986",
		"decimals": 2,
		"name": "Brazilian real"
	}],
	"dialing_codes": ["+55"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Brazil"
}, {
	"alpha2": "MZ",
	"alpha3": "MOZ",
	"numeric": "508",
	"currencies": [{
		"alpha3": "MZN",
		"numeric": "943",
		"decimals": 2,
		"name": "Mozambican metical"
	}],
	"dialing_codes": ["+258"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Mozambique"
}, {
	"alpha2": "RW",
	"alpha3": "RWA",
	"numeric": "646",
	"currencies": [{
		"alpha3": "RWF",
		"numeric": "646",
		"decimals": 0,
		"name": "Rwandan franc"
	}],
	"dialing_codes": ["+250"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "rw",
		"alpha3": "kin",
		"name": "Kinyarwanda"
	}],
	"name": "Rwanda"
}, {
	"alpha2": "CD",
	"alpha3": "COD",
	"numeric": "180",
	"currencies": [{
		"alpha3": "CDF",
		"numeric": "976",
		"decimals": 2,
		"name": "Congolese franc"
	}],
	"dialing_codes": ["+243"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "ln",
		"alpha3": "lin",
		"name": "Lingala"
	}, {
		"alpha2": "kg",
		"alpha3": "kon",
		"name": "Kongo"
	}, {
		"alpha2": "sw",
		"alpha3": "swa",
		"name": "Swahili"
	}],
	"name": "Democratic Republic Of Congo"
}, {
	"alpha2": "GH",
	"alpha3": "GHA",
	"numeric": "288",
	"currencies": [{
		"alpha3": "GHS",
		"numeric": "936",
		"decimals": 2,
		"name": "Ghanaian cedi"
	}],
	"dialing_codes": ["+233"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Ghana"
}, {
	"alpha2": "GY",
	"alpha3": "GUY",
	"numeric": "328",
	"currencies": [{
		"alpha3": "GYD",
		"numeric": "328",
		"decimals": 2,
		"name": "Guyanese dollar"
	}],
	"dialing_codes": ["+592"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Guyana"
}, {
	"alpha2": "LS",
	"alpha3": "LSO",
	"numeric": "426",
	"currencies": [{
		"alpha3": "LSL",
		"numeric": "426",
		"decimals": 2,
		"name": "Lesotho loti"
	}, {
		"alpha3": "ZAR",
		"numeric": "710",
		"decimals": 2,
		"name": "South African rand"
	}],
	"dialing_codes": ["+266"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "st",
		"alpha3": "sot",
		"name": "Sotho, Southern"
	}],
	"name": "Lesotho"
}, {
	"alpha2": "MH",
	"alpha3": "MHL",
	"numeric": "584",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+692"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "mh",
		"alpha3": "mah",
		"name": "Marshallese"
	}],
	"name": "Marshall Islands"
}, {
	"alpha2": "MK",
	"alpha3": "MKD",
	"numeric": "807",
	"currencies": [{
		"alpha3": "MKD",
		"numeric": "807",
		"decimals": 0,
		"name": "Macedonian denar"
	}],
	"dialing_codes": ["+389"],
	"languages": [{
		"alpha2": "mk",
		"alpha3": "mkd",
		"name": "Macedonian"
	}],
	"name": "Macedonia, The Former Yugoslav Republic Of"
}, {
	"alpha2": "MT",
	"alpha3": "MLT",
	"numeric": "470",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+356"],
	"languages": [{
		"alpha2": "mt",
		"alpha3": "mlt",
		"name": "Maltese"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Malta"
}, {
	"alpha2": "SL",
	"alpha3": "SLE",
	"numeric": "694",
	"currencies": [{
		"alpha3": "SLL",
		"numeric": "694",
		"decimals": 0,
		"name": "Sierra Leonean leone"
	}],
	"dialing_codes": ["+232"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Sierra Leone"
}, {
	"alpha2": "SO",
	"alpha3": "SOM",
	"numeric": "706",
	"currencies": [{
		"alpha3": "SOS",
		"numeric": "706",
		"decimals": 2,
		"name": "Somali shilling"
	}],
	"dialing_codes": ["+252"],
	"languages": [{
		"alpha2": "so",
		"alpha3": "som",
		"name": "Somali"
	}],
	"name": "Somalia"
}, {
	"alpha2": "AU",
	"alpha3": "AUS",
	"numeric": "036",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+61"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Australia"
}, {
	"alpha2": "SN",
	"alpha3": "SEN",
	"numeric": "686",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+221"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Senegal"
}, {
	"alpha2": "WS",
	"alpha3": "WSM",
	"numeric": "882",
	"currencies": [{
		"alpha3": "WST",
		"numeric": "882",
		"decimals": 2,
		"name": "Samoan tala"
	}],
	"dialing_codes": ["+685"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "sm",
		"alpha3": "smo",
		"name": "Samoan"
	}],
	"name": "Samoa"
}, {
	"alpha2": "AX",
	"alpha3": "ALA",
	"numeric": "248",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+358"],
	"languages": [{
		"alpha2": "sv",
		"alpha3": "swe",
		"name": "Swedish"
	}],
	"name": "??land Islands"
}, {
	"alpha2": "BA",
	"alpha3": "BIH",
	"numeric": "070",
	"currencies": [{
		"alpha3": "BAM",
		"numeric": "977",
		"decimals": 2,
		"name": "Bosnia and Herzegovina convertible mark"
	}],
	"dialing_codes": ["+387"],
	"languages": [{
		"alpha2": "bs",
		"alpha3": "bos",
		"name": "Bosnian"
	}, {
		"alpha2": "cr",
		"alpha3": "cre",
		"name": "Cree"
	}, {
		"alpha2": "sr",
		"alpha3": "srp",
		"name": "Serbian"
	}],
	"name": "Bosnia \u0026 Herzegovina"
}, {
	"alpha2": "DZ",
	"alpha3": "DZA",
	"numeric": "012",
	"currencies": [{
		"alpha3": "DZD",
		"numeric": "12",
		"decimals": 2,
		"name": "Algerian dinar"
	}],
	"dialing_codes": ["+213"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Algeria"
}, {
	"alpha2": "MY",
	"alpha3": "MYS",
	"numeric": "458",
	"currencies": [{
		"alpha3": "MYR",
		"numeric": "458",
		"decimals": 2,
		"name": "Malaysian ringgit"
	}],
	"dialing_codes": ["+60"],
	"languages": [{
		"alpha2": "ms",
		"alpha3": "msa",
		"name": "Malay"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Malaysia"
}, {
	"alpha2": "YT",
	"alpha3": "MYT",
	"numeric": "175",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+262"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Mayotte"
}, {
	"alpha2": "LC",
	"alpha3": "LCA",
	"numeric": "662",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 758"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Saint Lucia"
}, {
	"alpha2": "ME",
	"alpha3": "MNE",
	"numeric": "499",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+382"],
	"languages": [{
		"alpha2": "",
		"alpha3": "mot",
		"name": "Montenegrin"
	}],
	"name": "Montenegro"
}, {
	"alpha2": "SD",
	"alpha3": "SDN",
	"numeric": "729",
	"currencies": [{
		"alpha3": "SDG",
		"numeric": "938",
		"decimals": 2,
		"name": "Sudanese pound"
	}],
	"dialing_codes": ["+249"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Sudan"
}, {
	"alpha2": "SI",
	"alpha3": "SVN",
	"numeric": "705",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+386"],
	"languages": [{
		"alpha2": "sl",
		"alpha3": "slv",
		"name": "Slovenian"
	}],
	"name": "Slovenia"
}, {
	"alpha2": "SY",
	"alpha3": "SYR",
	"numeric": "760",
	"currencies": [{
		"alpha3": "SYP",
		"numeric": "760",
		"decimals": 2,
		"name": "Syrian pound"
	}],
	"dialing_codes": ["+963"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Syrian Arab Republic"
}, {
	"alpha2": "TT",
	"alpha3": "TTO",
	"numeric": "780",
	"currencies": [{
		"alpha3": "TTD",
		"numeric": "780",
		"decimals": 2,
		"name": "Trinidad and Tobago dollar"
	}],
	"dialing_codes": ["+1 868"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Trinidad And Tobago"
}, {
	"alpha2": "AG",
	"alpha3": "ATG",
	"numeric": "028",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 268"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Antigua And Barbuda"
}, {
	"alpha2": "BJ",
	"alpha3": "BEN",
	"numeric": "204",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+229"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Benin"
}, {
	"alpha2": "CV",
	"alpha3": "CPV",
	"numeric": "132",
	"currencies": [{
		"alpha3": "CVE",
		"numeric": "132",
		"decimals": 0,
		"name": "Cape Verde escudo"
	}],
	"dialing_codes": ["+238"],
	"languages": [{
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Cabo Verde"
}, {
	"alpha2": "EE",
	"alpha3": "EST",
	"numeric": "233",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+372"],
	"languages": [{
		"alpha2": "et",
		"alpha3": "est",
		"name": "Estonian"
	}],
	"name": "Estonia"
}, {
	"alpha2": "IS",
	"alpha3": "ISL",
	"numeric": "352",
	"currencies": [{
		"alpha3": "ISK",
		"numeric": "352",
		"decimals": 0,
		"name": "Icelandic kr??na"
	}],
	"dialing_codes": ["+354"],
	"languages": [{
		"alpha2": "is",
		"alpha3": "isl",
		"name": "Icelandic"
	}],
	"name": "Iceland"
}, {
	"alpha2": "TN",
	"alpha3": "TUN",
	"numeric": "788",
	"currencies": [{
		"alpha3": "TND",
		"numeric": "788",
		"decimals": 3,
		"name": "Tunisian dinar"
	}],
	"dialing_codes": ["+216"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Tunisia"
}, {
	"alpha2": "FR",
	"alpha3": "FRA",
	"numeric": "250",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+33"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "France"
}, {
	"alpha2": "LA",
	"alpha3": "LAO",
	"numeric": "418",
	"currencies": [{
		"alpha3": "LAK",
		"numeric": "418",
		"decimals": 0,
		"name": "Lao kip"
	}],
	"dialing_codes": ["+856"],
	"languages": [{
		"alpha2": "lo",
		"alpha3": "lao",
		"name": "Lao"
	}],
	"name": "Lao People's Democratic Republic"
}, {
	"alpha2": "PW",
	"alpha3": "PLW",
	"numeric": "585",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+680"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Palau"
}, {
	"alpha2": "GD",
	"alpha3": "GRD",
	"numeric": "308",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+473"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Grenada"
}, {
	"alpha2": "UA",
	"alpha3": "UKR",
	"numeric": "804",
	"currencies": [{
		"alpha3": "UAH",
		"numeric": "980",
		"decimals": 2,
		"name": "Ukrainian hryvnia"
	}],
	"dialing_codes": ["+380"],
	"languages": [{
		"alpha2": "uk",
		"alpha3": "ukr",
		"name": "Ukrainian"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Ukraine"
}, {
	"alpha2": "AZ",
	"alpha3": "AZE",
	"numeric": "031",
	"currencies": [{
		"alpha3": "AZN",
		"numeric": "944",
		"decimals": 2,
		"name": "Azerbaijani manat"
	}],
	"dialing_codes": ["+994"],
	"languages": [{
		"alpha2": "az",
		"alpha3": "aze",
		"name": "Azerbaijani"
	}],
	"name": "Azerbaijan"
}, {
	"alpha2": "KZ",
	"alpha3": "KAZ",
	"numeric": "398",
	"currencies": [{
		"alpha3": "KZT",
		"numeric": "398",
		"decimals": 2,
		"name": "Kazakhstani tenge"
	}],
	"dialing_codes": ["+7", "+7 6", "+7 7"],
	"languages": [{
		"alpha2": "kk",
		"alpha3": "kaz",
		"name": "Kazakh"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Kazakhstan"
}, {
	"alpha2": "LR",
	"alpha3": "LBR",
	"numeric": "430",
	"currencies": [{
		"alpha3": "LRD",
		"numeric": "430",
		"decimals": 2,
		"name": "Liberian dollar"
	}],
	"dialing_codes": ["+231"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Liberia"
}, {
	"alpha2": "SK",
	"alpha3": "SVK",
	"numeric": "703",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+421"],
	"languages": [{
		"alpha2": "sk",
		"alpha3": "slk",
		"name": "Slovak"
	}],
	"name": "Slovakia"
}, {
	"alpha2": "BL",
	"alpha3": "BLM",
	"numeric": "652",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+590"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Saint Barth??lemy"
}, {
	"alpha2": "GN",
	"alpha3": "GIN",
	"numeric": "324",
	"currencies": [{
		"alpha3": "GNF",
		"numeric": "324",
		"decimals": 0,
		"name": "Guinean franc"
	}],
	"dialing_codes": ["+224"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Guinea"
}, {
	"alpha2": "JE",
	"alpha3": "JEY",
	"numeric": "832",
	"currencies": [{
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}],
	"dialing_codes": ["+44"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Jersey"
}, {
	"alpha2": "TK",
	"alpha3": "TKL",
	"numeric": "772",
	"currencies": [{
		"alpha3": "NZD",
		"numeric": "554",
		"decimals": 2,
		"name": "New Zealand dollar"
	}],
	"dialing_codes": ["+690"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Tokelau"
}, {
	"alpha2": "VG",
	"alpha3": "VGB",
	"numeric": "092",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 284"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Virgin Islands (British)"
}, {
	"alpha2": "ZA",
	"alpha3": "ZAF",
	"numeric": "710",
	"currencies": [{
		"alpha3": "ZAR",
		"numeric": "710",
		"decimals": 2,
		"name": "South African rand"
	}],
	"dialing_codes": ["+27"],
	"languages": [{
		"alpha2": "af",
		"alpha3": "afr",
		"name": "Afrikaans"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "nr",
		"alpha3": "nbl",
		"name": "Ndebele, South"
	}, {
		"alpha2": "so",
		"alpha3": "som",
		"name": "Somali"
	}, {
		"alpha2": "ts",
		"alpha3": "tso",
		"name": "Tsonga"
	}, {
		"alpha2": "ve",
		"alpha3": "ven",
		"name": "Venda"
	}, {
		"alpha2": "xh",
		"alpha3": "xho",
		"name": "Xhosa"
	}, {
		"alpha2": "zu",
		"alpha3": "zul",
		"name": "Zulu"
	}],
	"name": "South Africa"
}, {
	"alpha2": "DE",
	"alpha3": "DEU",
	"numeric": "276",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+49"],
	"languages": [{
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}],
	"name": "Germany"
}, {
	"alpha2": "GG",
	"alpha3": "GGY",
	"numeric": "831",
	"currencies": [{
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}],
	"dialing_codes": ["+44"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Guernsey"
}, {
	"alpha2": "IR",
	"alpha3": "IRN",
	"numeric": "364",
	"currencies": [{
		"alpha3": "IRR",
		"numeric": "364",
		"decimals": 0,
		"name": "Iranian rial"
	}],
	"dialing_codes": ["+98"],
	"languages": [{
		"alpha2": "fa",
		"alpha3": "fas",
		"name": "Persian"
	}],
	"name": "Iran, Islamic Republic Of"
}, {
	"alpha2": "MN",
	"alpha3": "MNG",
	"numeric": "496",
	"currencies": [{
		"alpha3": "MNT",
		"numeric": "496",
		"decimals": 2,
		"name": "Mongolian tugrik"
	}],
	"dialing_codes": ["+976"],
	"languages": [{
		"alpha2": "mn",
		"alpha3": "mon",
		"name": "Mongolian"
	}],
	"name": "Mongolia"
}, {
	"alpha2": "MX",
	"alpha3": "MEX",
	"numeric": "484",
	"currencies": [{
		"alpha3": "MXN",
		"numeric": "484",
		"decimals": 2,
		"name": "Mexican peso"
	}, {
		"alpha3": "MXV",
		"numeric": "979",
		"decimals": 2,
		"name": "Mexican Unidad de Inversion (UDI) (funds code)"
	}],
	"dialing_codes": ["+52"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Mexico"
}, {
	"alpha2": "PS",
	"alpha3": "PSE",
	"numeric": "275",
	"currencies": [{
		"alpha3": "JOD",
		"numeric": "400",
		"decimals": 3,
		"name": "Jordanian dinar"
	}, {
		"alpha3": "EGP",
		"numeric": "818",
		"decimals": 2,
		"name": "Egyptian pound"
	}, {
		"alpha3": "ILS",
		"numeric": "376",
		"decimals": 2,
		"name": "Israeli new shekel"
	}],
	"dialing_codes": ["+970"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Palestinian Territory, Occupied"
}, {
	"alpha2": "SX",
	"alpha3": "SXM",
	"numeric": "534",
	"currencies": [{
		"alpha3": "ANG",
		"numeric": "532",
		"decimals": 2,
		"name": "Netherlands Antillean guilder"
	}],
	"dialing_codes": ["+1 721"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Sint Maarten"
}, {
	"alpha2": "FO",
	"alpha3": "FRO",
	"numeric": "234",
	"currencies": [{
		"alpha3": "DKK",
		"numeric": "208",
		"decimals": 2,
		"name": "Danish krone"
	}],
	"dialing_codes": ["+298"],
	"languages": [{
		"alpha2": "fo",
		"alpha3": "fao",
		"name": "Faroese"
	}, {
		"alpha2": "da",
		"alpha3": "dan",
		"name": "Danish"
	}],
	"name": "Faroe Islands"
}, {
	"alpha2": "LB",
	"alpha3": "LBN",
	"numeric": "422",
	"currencies": [{
		"alpha3": "LBP",
		"numeric": "422",
		"decimals": 0,
		"name": "Lebanese pound"
	}],
	"dialing_codes": ["+961"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "hy",
		"alpha3": "hye",
		"name": "Armenian"
	}],
	"name": "Lebanon"
}, {
	"alpha2": "PM",
	"alpha3": "SPM",
	"numeric": "666",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+508"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Saint Pierre And Miquelon"
}, {
	"alpha2": "BG",
	"alpha3": "BGR",
	"numeric": "100",
	"currencies": [{
		"alpha3": "BGN",
		"numeric": "975",
		"decimals": 2,
		"name": "Bulgarian lev"
	}],
	"dialing_codes": ["+359"],
	"languages": [{
		"alpha2": "bg",
		"alpha3": "bul",
		"name": "Bulgarian"
	}],
	"name": "Bulgaria"
}, {
	"alpha2": "BS",
	"alpha3": "BHS",
	"numeric": "044",
	"currencies": [{
		"alpha3": "BSD",
		"numeric": "44",
		"decimals": 2,
		"name": "Bahamian dollar"
	}],
	"dialing_codes": ["+1 242"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Bahamas"
}, {
	"alpha2": "FI",
	"alpha3": "FIN",
	"numeric": "246",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+358"],
	"languages": [{
		"alpha2": "fi",
		"alpha3": "fin",
		"name": "Finnish"
	}, {
		"alpha2": "sv",
		"alpha3": "swe",
		"name": "Swedish"
	}],
	"name": "Finland"
}, {
	"alpha2": "GE",
	"alpha3": "GEO",
	"numeric": "268",
	"currencies": [{
		"alpha3": "GEL",
		"numeric": "981",
		"decimals": 2,
		"name": "Georgian lari"
	}],
	"dialing_codes": ["+995"],
	"languages": [{
		"alpha2": "ka",
		"alpha3": "kat",
		"name": "Georgian"
	}],
	"name": "Georgia"
}, {
	"alpha2": "NL",
	"alpha3": "NLD",
	"numeric": "528",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+31"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Netherlands"
}, {
	"alpha2": "GP",
	"alpha3": "GLP",
	"numeric": "312",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+590"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Guadeloupe"
}, {
	"alpha2": "IL",
	"alpha3": "ISR",
	"numeric": "376",
	"currencies": [{
		"alpha3": "ILS",
		"numeric": "376",
		"decimals": 2,
		"name": "Israeli new shekel"
	}],
	"dialing_codes": ["+972"],
	"languages": [{
		"alpha2": "he",
		"alpha3": "heb",
		"name": "Hebrew"
	}, {
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Israel"
}, {
	"alpha2": "LY",
	"alpha3": "LBY",
	"numeric": "434",
	"currencies": [{
		"alpha3": "LYD",
		"numeric": "434",
		"decimals": 3,
		"name": "Libyan dinar"
	}],
	"dialing_codes": ["+218"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Libya"
}, {
	"alpha2": "UY",
	"alpha3": "URY",
	"numeric": "858",
	"currencies": [{
		"alpha3": "UYU",
		"numeric": "858",
		"decimals": 2,
		"name": "Uruguayan peso"
	}, {
		"alpha3": "UYI",
		"numeric": "940",
		"decimals": 0,
		"name": "Uruguay Peso en Unidades Indexadas (URUIURUI) (funds code)"
	}],
	"dialing_codes": ["+598"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Uruguay"
}, {
	"alpha2": "CN",
	"alpha3": "CHN",
	"numeric": "156",
	"currencies": [{
		"alpha3": "CNY",
		"numeric": "156",
		"decimals": 2,
		"name": "Chinese yuan"
	}],
	"dialing_codes": ["+86"],
	"languages": [{
		"alpha2": "zh",
		"alpha3": "zho",
		"name": "Chinese"
	}],
	"name": "China"
}, {
	"alpha2": "GB",
	"alpha3": "GBR",
	"numeric": "826",
	"currencies": [{
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}],
	"dialing_codes": ["+44"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "kw",
		"alpha3": "cor",
		"name": "Cornish"
	}, {
		"alpha2": "ga",
		"alpha3": "gle",
		"name": "Irish"
	}, {
		"alpha2": "gd",
		"alpha3": "gla",
		"name": "Gaelic"
	}, {
		"alpha2": "cy",
		"alpha3": "cym",
		"name": "Welsh"
	}],
	"name": "United Kingdom"
}, {
	"alpha2": "KE",
	"alpha3": "KEN",
	"numeric": "404",
	"currencies": [{
		"alpha3": "KES",
		"numeric": "404",
		"decimals": 2,
		"name": "Kenyan shilling"
	}],
	"dialing_codes": ["+254"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "sw",
		"alpha3": "swa",
		"name": "Swahili"
	}],
	"name": "Kenya"
}, {
	"alpha2": "KP",
	"alpha3": "PRK",
	"numeric": "408",
	"currencies": [{
		"alpha3": "KPW",
		"numeric": "408",
		"decimals": 0,
		"name": "North Korean won"
	}],
	"dialing_codes": ["+850"],
	"languages": [{
		"alpha2": "ko",
		"alpha3": "kor",
		"name": "Korean"
	}],
	"name": "Korea, Democratic People's Republic Of"
}, {
	"alpha2": "SE",
	"alpha3": "SWE",
	"numeric": "752",
	"currencies": [{
		"alpha3": "SEK",
		"numeric": "752",
		"decimals": 2,
		"name": "Swedish krona/kronor"
	}],
	"dialing_codes": ["+46"],
	"languages": [{
		"alpha2": "sv",
		"alpha3": "swe",
		"name": "Swedish"
	}],
	"name": "Sweden"
}, {
	"alpha2": "TO",
	"alpha3": "TON",
	"numeric": "776",
	"currencies": [{
		"alpha3": "TOP",
		"numeric": "776",
		"decimals": 2,
		"name": "Tongan pa??anga"
	}],
	"dialing_codes": ["+676"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Tonga"
}, {
	"alpha2": "BF",
	"alpha3": "BFA",
	"numeric": "854",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+226"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Burkina Faso"
}, {
	"alpha2": "CM",
	"alpha3": "CMR",
	"numeric": "120",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+237"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Cameroon"
}, {
	"alpha2": "HR",
	"alpha3": "HRV",
	"numeric": "191",
	"currencies": [{
		"alpha3": "HRK",
		"numeric": "191",
		"decimals": 2,
		"name": "Croatian kuna"
	}],
	"dialing_codes": ["+385"],
	"languages": [{
		"alpha2": "hr",
		"alpha3": "hrv",
		"name": "Croatian"
	}],
	"name": "Croatia"
}, {
	"alpha2": "IE",
	"alpha3": "IRL",
	"numeric": "372",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+353"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "ga",
		"alpha3": "gle",
		"name": "Irish"
	}],
	"name": "Ireland"
}, {
	"alpha2": "TJ",
	"alpha3": "TJK",
	"numeric": "762",
	"currencies": [{
		"alpha3": "TJS",
		"numeric": "972",
		"decimals": 2,
		"name": "Tajikistani somoni"
	}],
	"dialing_codes": ["+992"],
	"languages": [{
		"alpha2": "tg",
		"alpha3": "tgk",
		"name": "Tajik"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Tajikistan"
}, {
	"alpha2": "AD",
	"alpha3": "AND",
	"numeric": "020",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+376"],
	"languages": [{
		"alpha2": "ca",
		"alpha3": "cat",
		"name": "Catalan"
	}],
	"name": "Andorra"
}, {
	"alpha2": "CY",
	"alpha3": "CYP",
	"numeric": "196",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+357"],
	"languages": [{
		"alpha2": "el",
		"alpha3": "ell",
		"name": "Greek, Modern (1453-)"
	}, {
		"alpha2": "tr",
		"alpha3": "tur",
		"name": "Turkish"
	}],
	"name": "Cyprus"
}, {
	"alpha2": "KY",
	"alpha3": "CYM",
	"numeric": "136",
	"currencies": [{
		"alpha3": "KYD",
		"numeric": "136",
		"decimals": 2,
		"name": "Cayman Islands dollar"
	}],
	"dialing_codes": ["+1 345"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Cayman Islands"
}, {
	"alpha2": "BM",
	"alpha3": "BMU",
	"numeric": "060",
	"currencies": [{
		"alpha3": "BMD",
		"numeric": "60",
		"decimals": 2,
		"name": "Bermudian dollar (customarily known as Bermuda dollar)"
	}],
	"dialing_codes": ["+1 441"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Bermuda"
}, {
	"alpha2": "ML",
	"alpha3": "MLI",
	"numeric": "466",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+223"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Mali"
}, {
	"alpha2": "TC",
	"alpha3": "TCA",
	"numeric": "796",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 649"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Turks And Caicos Islands"
}, {
	"alpha2": "CI",
	"alpha3": "CIV",
	"numeric": "384",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+225"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "C??te d'Ivoire"
}, {
	"alpha2": "IO",
	"alpha3": "IOT",
	"numeric": "086",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+246"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "British Indian Ocean Territory"
}, {
	"alpha2": "BB",
	"alpha3": "BRB",
	"numeric": "052",
	"currencies": [{
		"alpha3": "BBD",
		"numeric": "52",
		"decimals": 2,
		"name": "Barbados dollar"
	}],
	"dialing_codes": ["+1 246"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Barbados"
}, {
	"alpha2": "RS",
	"alpha3": "SRB",
	"numeric": "688",
	"currencies": [{
		"alpha3": "RSD",
		"numeric": "941",
		"decimals": 2,
		"name": "Serbian dinar"
	}],
	"dialing_codes": ["+381"],
	"languages": [{
		"alpha2": "sr",
		"alpha3": "srp",
		"name": "Serbian"
	}],
	"name": "Serbia"
}, {
	"alpha2": "AI",
	"alpha3": "AIA",
	"numeric": "660",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 264"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Anguilla"
}, {
	"alpha2": "BI",
	"alpha3": "BDI",
	"numeric": "108",
	"currencies": [{
		"alpha3": "BIF",
		"numeric": "108",
		"decimals": 0,
		"name": "Burundian franc"
	}],
	"dialing_codes": ["+257"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Burundi"
}, {
	"alpha2": "CU",
	"alpha3": "CUB",
	"numeric": "192",
	"currencies": [{
		"alpha3": "CUP",
		"numeric": "192",
		"decimals": 2,
		"name": "Cuban peso"
	}, {
		"alpha3": "CUC",
		"numeric": "931",
		"decimals": 2,
		"name": "Cuban convertible peso"
	}],
	"dialing_codes": ["+53"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Cuba"
}, {
	"alpha2": "DM",
	"alpha3": "DMA",
	"numeric": "212",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 767"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Dominica"
}, {
	"alpha2": "HN",
	"alpha3": "HND",
	"numeric": "340",
	"currencies": [{
		"alpha3": "HNL",
		"numeric": "340",
		"decimals": 2,
		"name": "Honduran lempira"
	}],
	"dialing_codes": ["+504"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Honduras"
}, {
	"alpha2": "US",
	"alpha3": "USA",
	"numeric": "840",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "United States"
}, {
	"alpha2": "CO",
	"alpha3": "COL",
	"numeric": "170",
	"currencies": [{
		"alpha3": "COP",
		"numeric": "170",
		"decimals": 2,
		"name": "Colombian peso"
	}, {
		"alpha3": "COU",
		"numeric": "970",
		"decimals": 2,
		"name": "Unidad de Valor Real"
	}],
	"dialing_codes": ["+57"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Colombia"
}, {
	"alpha2": "VC",
	"alpha3": "VCT",
	"numeric": "670",
	"currencies": [{
		"alpha3": "XCD",
		"numeric": "951",
		"decimals": 2,
		"name": "East Caribbean dollar"
	}],
	"dialing_codes": ["+1 784"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Saint Vincent And The Grenadines"
}, {
	"alpha2": "MO",
	"alpha3": "MAC",
	"numeric": "446",
	"currencies": [{
		"alpha3": "MOP",
		"numeric": "446",
		"decimals": 2,
		"name": "Macanese pataca"
	}],
	"dialing_codes": ["+853"],
	"languages": [{
		"alpha2": "zh",
		"alpha3": "zho",
		"name": "Chinese"
	}, {
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Macao"
}, {
	"alpha2": "MP",
	"alpha3": "MNP",
	"numeric": "580",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 670"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Northern Mariana Islands"
}, {
	"alpha2": "PE",
	"alpha3": "PER",
	"numeric": "604",
	"currencies": [{
		"alpha3": "PEN",
		"numeric": "604",
		"decimals": 2,
		"name": "Peruvian nuevo sol"
	}],
	"dialing_codes": ["+51"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}, {
		"alpha2": "ay",
		"alpha3": "aym",
		"name": "Aymara"
	}, {
		"alpha2": "qu",
		"alpha3": "que",
		"name": "Quechua"
	}],
	"name": "Peru"
}, {
	"alpha2": "PH",
	"alpha3": "PHL",
	"numeric": "608",
	"currencies": [{
		"alpha3": "PHP",
		"numeric": "608",
		"decimals": 2,
		"name": "Philippine peso"
	}],
	"dialing_codes": ["+63"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Philippines"
}, {
	"alpha2": "RE",
	"alpha3": "REU",
	"numeric": "638",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+262"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Reunion"
}, {
	"alpha2": "SH",
	"alpha3": "SHN",
	"numeric": "654",
	"currencies": [{
		"alpha3": "SHP",
		"numeric": "654",
		"decimals": 2,
		"name": "Saint Helena pound"
	}],
	"dialing_codes": ["+290"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Saint Helena, Ascension And Tristan Da Cunha"
}, {
	"alpha2": "AL",
	"alpha3": "ALB",
	"numeric": "008",
	"currencies": [{
		"alpha3": "ALL",
		"numeric": "8",
		"decimals": 2,
		"name": "Albanian lek"
	}],
	"dialing_codes": ["+355"],
	"languages": [{
		"alpha2": "sq",
		"alpha3": "sqi",
		"name": "Albanian"
	}],
	"name": "Albania"
}, {
	"alpha2": "CA",
	"alpha3": "CAN",
	"numeric": "124",
	"currencies": [{
		"alpha3": "CAD",
		"numeric": "124",
		"decimals": 2,
		"name": "Canadian dollar"
	}],
	"dialing_codes": ["+1"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Canada"
}, {
	"alpha2": "PL",
	"alpha3": "POL",
	"numeric": "616",
	"currencies": [{
		"alpha3": "PLN",
		"numeric": "985",
		"decimals": 2,
		"name": "Polish z??oty"
	}],
	"dialing_codes": ["+48"],
	"languages": [{
		"alpha2": "pl",
		"alpha3": "pol",
		"name": "Polish"
	}],
	"name": "Poland"
}, {
	"alpha2": "BQ",
	"alpha3": "BES",
	"numeric": "535",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+599"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Bonaire, Saint Eustatius And Saba"
}, {
	"alpha2": "LT",
	"alpha3": "LTU",
	"numeric": "440",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+370"],
	"languages": [{
		"alpha2": "lt",
		"alpha3": "lit",
		"name": "Lithuanian"
	}],
	"name": "Lithuania"
}, {
	"alpha2": "PR",
	"alpha3": "PRI",
	"numeric": "630",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 787", "+1 939"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Puerto Rico"
}, {
	"alpha2": "TD",
	"alpha3": "TCD",
	"numeric": "148",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+235"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Chad"
}, {
	"alpha2": "VI",
	"alpha3": "VIR",
	"numeric": "850",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 340"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Virgin Islands (US)"
}, {
	"alpha2": "WF",
	"alpha3": "WLF",
	"numeric": "876",
	"currencies": [{
		"alpha3": "XPF",
		"numeric": "953",
		"decimals": 0,
		"name": "CFP franc"
	}],
	"dialing_codes": ["+681"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Wallis And Futuna"
}, {
	"alpha2": "BH",
	"alpha3": "BHR",
	"numeric": "048",
	"currencies": [{
		"alpha3": "BHD",
		"numeric": "48",
		"decimals": 3,
		"name": "Bahraini dinar"
	}],
	"dialing_codes": ["+973"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Bahrain"
}, {
	"alpha2": "BO",
	"alpha3": "BOL",
	"numeric": "068",
	"currencies": [{
		"alpha3": "BOB",
		"numeric": "68",
		"decimals": 2,
		"name": "Boliviano"
	}, {
		"alpha3": "BOV",
		"numeric": "984",
		"decimals": 2,
		"name": "Bolivian Mvdol (funds code)"
	}],
	"dialing_codes": ["+591"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}, {
		"alpha2": "ay",
		"alpha3": "aym",
		"name": "Aymara"
	}, {
		"alpha2": "qu",
		"alpha3": "que",
		"name": "Quechua"
	}],
	"name": "Bolivia, Plurinational State Of"
}, {
	"alpha2": "GQ",
	"alpha3": "GNQ",
	"numeric": "226",
	"currencies": [{
		"alpha3": "XAF",
		"numeric": "950",
		"decimals": 0,
		"name": "CFA franc BEAC"
	}],
	"dialing_codes": ["+240"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "pt",
		"alpha3": "por",
		"name": "Portuguese"
	}],
	"name": "Equatorial Guinea"
}, {
	"alpha2": "PN",
	"alpha3": "PCN",
	"numeric": "612",
	"currencies": [{
		"alpha3": "NZD",
		"numeric": "554",
		"decimals": 2,
		"name": "New Zealand dollar"
	}],
	"dialing_codes": ["+872"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Pitcairn"
}, {
	"alpha2": "CZ",
	"alpha3": "CZE",
	"numeric": "203",
	"currencies": [{
		"alpha3": "CZK",
		"numeric": "203",
		"decimals": 2,
		"name": "Czech koruna"
	}],
	"dialing_codes": ["+420"],
	"languages": [{
		"alpha2": "cs",
		"alpha3": "ces",
		"name": "Czech"
	}],
	"name": "Czech Republic"
}, {
	"alpha2": "DJ",
	"alpha3": "DJI",
	"numeric": "262",
	"currencies": [{
		"alpha3": "DJF",
		"numeric": "262",
		"decimals": 0,
		"name": "Djiboutian franc"
	}],
	"dialing_codes": ["+253"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Djibouti"
}, {
	"alpha2": "ET",
	"alpha3": "ETH",
	"numeric": "231",
	"currencies": [{
		"alpha3": "ETB",
		"numeric": "230",
		"decimals": 2,
		"name": "Ethiopian birr"
	}],
	"dialing_codes": ["+251"],
	"languages": [{
		"alpha2": "am",
		"alpha3": "amh",
		"name": "Amharic"
	}],
	"name": "Ethiopia"
}, {
	"alpha2": "SR",
	"alpha3": "SUR",
	"numeric": "740",
	"currencies": [{
		"alpha3": "SRD",
		"numeric": "968",
		"decimals": 2,
		"name": "Surinamese dollar"
	}],
	"dialing_codes": ["+597"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Suriname"
}, {
	"alpha2": "TR",
	"alpha3": "TUR",
	"numeric": "792",
	"currencies": [{
		"alpha3": "TRY",
		"numeric": "949",
		"decimals": 2,
		"name": "Turkish lira"
	}],
	"dialing_codes": ["+90"],
	"languages": [{
		"alpha2": "tr",
		"alpha3": "tur",
		"name": "Turkish"
	}],
	"name": "Turkey"
}, {
	"alpha2": "GU",
	"alpha3": "GUM",
	"numeric": "316",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 671"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Guam"
}, {
	"alpha2": "KR",
	"alpha3": "KOR",
	"numeric": "410",
	"currencies": [{
		"alpha3": "KRW",
		"numeric": "410",
		"decimals": 0,
		"name": "South Korean won"
	}],
	"dialing_codes": ["+82"],
	"languages": [{
		"alpha2": "ko",
		"alpha3": "kor",
		"name": "Korean"
	}],
	"name": "Korea, Republic Of"
}, {
	"alpha2": "LI",
	"alpha3": "LIE",
	"numeric": "438",
	"currencies": [{
		"alpha3": "CHF",
		"numeric": "756",
		"decimals": 2,
		"name": "Swiss franc"
	}],
	"dialing_codes": ["+423"],
	"languages": [{
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}],
	"name": "Liechtenstein"
}, {
	"alpha2": "PK",
	"alpha3": "PAK",
	"numeric": "586",
	"currencies": [{
		"alpha3": "PKR",
		"numeric": "586",
		"decimals": 2,
		"name": "Pakistani rupee"
	}],
	"dialing_codes": ["+92"],
	"languages": [{
		"alpha2": "ur",
		"alpha3": "urd",
		"name": "Urdu"
	}, {
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Pakistan"
}, {
	"alpha2": "CL",
	"alpha3": "CHL",
	"numeric": "152",
	"currencies": [{
		"alpha3": "CLP",
		"numeric": "152",
		"decimals": 0,
		"name": "Chilean peso"
	}, {
		"alpha3": "CLF",
		"numeric": "990",
		"decimals": 0,
		"name": "Unidad de Fomento (funds code)"
	}],
	"dialing_codes": ["+56"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Chile"
}, {
	"alpha2": "HU",
	"alpha3": "HUN",
	"numeric": "348",
	"currencies": [{
		"alpha3": "HUF",
		"numeric": "348",
		"decimals": 2,
		"name": "Hungarian forint"
	}],
	"dialing_codes": ["+36"],
	"languages": [{
		"alpha2": "hu",
		"alpha3": "hun",
		"name": "Hungarian"
	}],
	"name": "Hungary"
}, {
	"alpha2": "IM",
	"alpha3": "IMN",
	"numeric": "833",
	"currencies": [{
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}],
	"dialing_codes": ["+44"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "gv",
		"alpha3": "glv",
		"name": "Manx"
	}],
	"name": "Isle Of Man"
}, {
	"alpha2": "LV",
	"alpha3": "LVA",
	"numeric": "428",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+371"],
	"languages": [{
		"alpha2": "lv",
		"alpha3": "lav",
		"name": "Latvian"
	}],
	"name": "Latvia"
}, {
	"alpha2": "FJ",
	"alpha3": "FJI",
	"numeric": "242",
	"currencies": [{
		"alpha3": "FJD",
		"numeric": "242",
		"decimals": 2,
		"name": "Fiji dollar"
	}],
	"dialing_codes": ["+679"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "fj",
		"alpha3": "fij",
		"name": "Fijian"
	}],
	"name": "Fiji"
}, {
	"alpha2": "NO",
	"alpha3": "NOR",
	"numeric": "578",
	"currencies": [{
		"alpha3": "NOK",
		"numeric": "578",
		"decimals": 2,
		"name": "Norwegian krone"
	}],
	"dialing_codes": ["+47"],
	"languages": [{
		"alpha2": "no",
		"alpha3": "nor",
		"name": "Norwegian"
	}],
	"name": "Norway"
}, {
	"alpha2": "SA",
	"alpha3": "SAU",
	"numeric": "682",
	"currencies": [{
		"alpha3": "SAR",
		"numeric": "682",
		"decimals": 2,
		"name": "Saudi riyal"
	}],
	"dialing_codes": ["+966"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Saudi Arabia"
}, {
	"alpha2": "NR",
	"alpha3": "NRU",
	"numeric": "520",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+674"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "na",
		"alpha3": "nau",
		"name": "Nauru"
	}],
	"name": "Nauru"
}, {
	"alpha2": "AW",
	"alpha3": "ABW",
	"numeric": "533",
	"currencies": [{
		"alpha3": "AWG",
		"numeric": "533",
		"decimals": 2,
		"name": "Aruban florin"
	}],
	"dialing_codes": ["+297"],
	"languages": [{
		"alpha2": "nl",
		"alpha3": "nld",
		"name": "Dutch"
	}],
	"name": "Aruba"
}, {
	"alpha2": "CC",
	"alpha3": "CCK",
	"numeric": "166",
	"currencies": [{
		"alpha3": "AUD",
		"numeric": "36",
		"decimals": 2,
		"name": "Australian dollar"
	}],
	"dialing_codes": ["+61"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Cocos (Keeling) Islands"
}, {
	"alpha2": "FK",
	"alpha3": "FLK",
	"numeric": "238",
	"currencies": [{
		"alpha3": "FKP",
		"numeric": "238",
		"decimals": 2,
		"name": "Falkland Islands pound"
	}],
	"dialing_codes": ["+500"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Falkland Islands"
}, {
	"alpha2": "IQ",
	"alpha3": "IRQ",
	"numeric": "368",
	"currencies": [{
		"alpha3": "IQD",
		"numeric": "368",
		"decimals": 3,
		"name": "Iraqi dinar"
	}],
	"dialing_codes": ["+964"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "ku",
		"alpha3": "kur",
		"name": "Kurdish"
	}],
	"name": "Iraq"
}, {
	"alpha2": "MG",
	"alpha3": "MDG",
	"numeric": "450",
	"currencies": [{
		"alpha3": "MGA",
		"numeric": "969",
		"decimals": 0,
		"name": "Malagasy ariary"
	}],
	"dialing_codes": ["+261"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}, {
		"alpha2": "mg",
		"alpha3": "mlg",
		"name": "Malagasy"
	}],
	"name": "Madagascar"
}, {
	"alpha2": "MV",
	"alpha3": "MDV",
	"numeric": "462",
	"currencies": [{
		"alpha3": "MVR",
		"numeric": "462",
		"decimals": 2,
		"name": "Maldivian rufiyaa"
	}],
	"dialing_codes": ["+960"],
	"languages": [{
		"alpha2": "dv",
		"alpha3": "div",
		"name": "Dhivehi"
	}],
	"name": "Maldives"
}, {
	"alpha2": "NC",
	"alpha3": "NCL",
	"numeric": "540",
	"currencies": [{
		"alpha3": "XPF",
		"numeric": "953",
		"decimals": 0,
		"name": "CFP franc"
	}],
	"dialing_codes": ["+687"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "New Caledonia"
}, {
	"alpha2": "PG",
	"alpha3": "PNG",
	"numeric": "598",
	"currencies": [{
		"alpha3": "PGK",
		"numeric": "598",
		"decimals": 2,
		"name": "Papua New Guinean kina"
	}],
	"dialing_codes": ["+675"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "Papua New Guinea"
}, {
	"alpha2": "SS",
	"alpha3": "SSD",
	"numeric": "728",
	"currencies": [{
		"alpha3": "SSP",
		"numeric": "728",
		"decimals": 2,
		"name": "South Sudanese pound"
	}],
	"dialing_codes": ["+211"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}],
	"name": "South Sudan"
}, {
	"alpha2": "AR",
	"alpha3": "ARG",
	"numeric": "032",
	"currencies": [{
		"alpha3": "ARS",
		"numeric": "32",
		"decimals": 2,
		"name": "Argentine peso"
	}],
	"dialing_codes": ["+54"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Argentina"
}, {
	"alpha2": "KH",
	"alpha3": "KHM",
	"numeric": "116",
	"currencies": [{
		"alpha3": "KHR",
		"numeric": "116",
		"decimals": 2,
		"name": "Cambodian riel"
	}],
	"dialing_codes": ["+855"],
	"languages": [{
		"alpha2": "km",
		"alpha3": "khm",
		"name": "Central Khmer"
	}],
	"name": "Cambodia"
}, {
	"alpha2": "VA",
	"alpha3": "VAT",
	"numeric": "336",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+379", "+39"],
	"languages": [{
		"alpha2": "it",
		"alpha3": "ita",
		"name": "Italian"
	}],
	"name": "Vatican City State"
}, {
	"alpha2": "GR",
	"alpha3": "GRC",
	"numeric": "300",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+30"],
	"languages": [{
		"alpha2": "el",
		"alpha3": "ell",
		"name": "Greek, Modern (1453-)"
	}],
	"name": "Greece"
}, {
	"alpha2": "KM",
	"alpha3": "COM",
	"numeric": "174",
	"currencies": [{
		"alpha3": "KMF",
		"numeric": "174",
		"decimals": 0,
		"name": "Comoro franc"
	}],
	"dialing_codes": ["+269"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}, {
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Comoros"
}, {
	"alpha2": "AT",
	"alpha3": "AUT",
	"numeric": "040",
	"currencies": [{
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+43"],
	"languages": [{
		"alpha2": "de",
		"alpha3": "deu",
		"name": "German"
	}],
	"name": "Austria"
}, {
	"alpha2": "SV",
	"alpha3": "SLV",
	"numeric": "222",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+503"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "El Salvador"
}, {
	"alpha2": "UG",
	"alpha3": "UGA",
	"numeric": "800",
	"currencies": [{
		"alpha3": "UGX",
		"numeric": "800",
		"decimals": 2,
		"name": "Ugandan shilling"
	}],
	"dialing_codes": ["+256"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "sw",
		"alpha3": "swa",
		"name": "Swahili"
	}],
	"name": "Uganda"
}, {
	"alpha2": "YE",
	"alpha3": "YEM",
	"numeric": "887",
	"currencies": [{
		"alpha3": "YER",
		"numeric": "886",
		"decimals": 2,
		"name": "Yemeni rial"
	}],
	"dialing_codes": ["+967"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Yemen"
}, {
	"alpha2": "ZW",
	"alpha3": "ZWE",
	"numeric": "716",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}, {
		"alpha3": "ZAR",
		"numeric": "710",
		"decimals": 2,
		"name": "South African rand"
	}, {
		"alpha3": "BWP",
		"numeric": "72",
		"decimals": 2,
		"name": "Botswana pula"
	}, {
		"alpha3": "GBP",
		"numeric": "826",
		"decimals": 2,
		"name": "Pound sterling"
	}, {
		"alpha3": "EUR",
		"numeric": "978",
		"decimals": 2,
		"name": "Euro"
	}],
	"dialing_codes": ["+263"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "sn",
		"alpha3": "sna",
		"name": "Shona"
	}, {
		"alpha2": "nd",
		"alpha3": "nde",
		"name": "Ndebele, North"
	}],
	"name": "Zimbabwe"
}, {
	"alpha2": "DO",
	"alpha3": "DOM",
	"numeric": "214",
	"currencies": [{
		"alpha3": "DOP",
		"numeric": "214",
		"decimals": 2,
		"name": "Dominican peso"
	}],
	"dialing_codes": ["+1 809", "+1 829", "+1 849"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Dominican Republic"
}, {
	"alpha2": "PA",
	"alpha3": "PAN",
	"numeric": "591",
	"currencies": [{
		"alpha3": "PAB",
		"numeric": "590",
		"decimals": 2,
		"name": "Panamanian balboa"
	}, {
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+507"],
	"languages": [{
		"alpha2": "es",
		"alpha3": "spa",
		"name": "Castilian"
	}],
	"name": "Panama"
}, {
	"alpha2": "AS",
	"alpha3": "ASM",
	"numeric": "016",
	"currencies": [{
		"alpha3": "USD",
		"numeric": "840",
		"decimals": 2,
		"name": "United States dollar"
	}],
	"dialing_codes": ["+1 684"],
	"languages": [{
		"alpha2": "en",
		"alpha3": "eng",
		"name": "English"
	}, {
		"alpha2": "sm",
		"alpha3": "smo",
		"name": "Samoan"
	}],
	"name": "American Samoa"
}, {
	"alpha2": "JO",
	"alpha3": "JOR",
	"numeric": "400",
	"currencies": [{
		"alpha3": "JOD",
		"numeric": "400",
		"decimals": 3,
		"name": "Jordanian dinar"
	}],
	"dialing_codes": ["+962"],
	"languages": [{
		"alpha2": "ar",
		"alpha3": "ara",
		"name": "Arabic"
	}],
	"name": "Jordan"
}, {
	"alpha2": "TG",
	"alpha3": "TGO",
	"numeric": "768",
	"currencies": [{
		"alpha3": "XOF",
		"numeric": "952",
		"decimals": 0,
		"name": "CFA franc BCEAO"
	}],
	"dialing_codes": ["+228"],
	"languages": [{
		"alpha2": "fr",
		"alpha3": "fra",
		"name": "French"
	}],
	"name": "Togo"
}, {
	"alpha2": "UZ",
	"alpha3": "UZB",
	"numeric": "860",
	"currencies": [{
		"alpha3": "UZS",
		"numeric": "860",
		"decimals": 2,
		"name": "Uzbekistan som"
	}],
	"dialing_codes": ["+998"],
	"languages": [{
		"alpha2": "uz",
		"alpha3": "uzb",
		"name": "Uzbek"
	}, {
		"alpha2": "ru",
		"alpha3": "rus",
		"name": "Russian"
	}],
	"name": "Uzbekistan"
}]`)

func dataDataJsonBytes() ([]byte, error) {
	return _dataDataJson, nil
}

func dataDataJson() (*asset, error) {
	bytes, err := dataDataJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/data.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"data/data.json": dataDataJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"data.json": &bintree{dataDataJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
