package author

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/300brand/coverage/author"
	"github.com/300brand/coverage/downloader"
	"ioutil"
	"log"
)

var xpath string
var jsonFileName string
var urls []string

type XPath struct {
	Host  string   `json:Host`
	Paths []string `json:Paths`
}

type XPathCollection struct {
	XPaths []XPath
}

func init() {
	flag.StringVar(&xpath, "xpath", "", "Represents XPath to use to search for the author. Attempts to use XPaths in the .json file if not defined.")
	flag.StringVar(&jsonFileName, "f", "xpaths.json", "Represents the file name of the .json of Hosts and XPaths to find the author. Defaults to \"xpaths.json\"")
	urls = flag.Args()
}

func main() {
	var jsonFileData XPathCollection
	var data = &jsonFileData.XPaths
	//reads the file into a []byte
	jsonByteFileData, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		log.Fatal(err)
	}
	//decode the json into structs
	err = json.Unmarshal(jsonByteFileData, &data)
	if err != nil {
		log.Fatal(err)
	}

	//check if xpath is defined and grab first one associated with host in json if not
	if xpath == nil || xpath == "" {

	} else { //else add this one to json as first and use

	}

	var urlBodys map[string][]byte
	//loops through non-flag arguments and downloads; tries each xpath for host until author successfully extracted
	for _, url := range urls {
		urlBodys[url], err = downloader.Fetch(url).Body
		if err != nil {
			log.Printf("Error downloading: URL %s - %s", url, err)
		}
	}
}
