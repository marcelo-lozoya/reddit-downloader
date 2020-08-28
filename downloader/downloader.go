package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/google/uuid"
)

var wg sync.WaitGroup

const (
	ENDPOINT = "https://www.reddit.com/r/"
)

type Image struct {
	Source struct {
		URL string `json:"url"`
	} `json:"source"`
}

type Post struct {
	Data struct {
		Title   string `json:"title"`
		Preview struct {
			Images []Image `json:"images"`
		} `json:"preview"`
	} `json:"data"`
}

type Data struct {
	Kind string `json:"kind"`
	Data struct {
		Dist     int    `json:"dist"`
		Children []Post `json:"children"`
	} `json:"data"`
}

// Our interface for the DirectoryHelper
type IDownloader interface {
	MakeRequestForReddit(subreddit string, limit int)
}

// Our concrete implementation of the DirectoryHelper
type CDownloader struct {
	DirectoryHelper IDirectoryHelper
}

// Our provider for the DirectoryHelper
func InitDownloader(dh IDirectoryHelper) IDownloader {
	return CDownloader{DirectoryHelper: dh}
}

// MakeRequestForReddit takes in subreddit name and number of images to download
func (d CDownloader) MakeRequestForReddit(subreddit string, limit int) {
	d.DirectoryHelper.CreateRequiredFolders(subreddit)

	uriString := buildEndpoint(subreddit, limit)

	client := &http.Client{}
	fmt.Printf("Grabbing From: %s\n", uriString)
	req, err := http.NewRequest("GET", uriString, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var need Data

	err = json.Unmarshal(data, &need)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range need.Data.Children {
		if v.Data.Preview.Images != nil && len(v.Data.Preview.Images) > 0 {
			wg.Add(1)
			f := getImageLink(v.Data.Preview.Images[0].Source.URL)
			fmt.Println("Grabbing image from: " + f)
			go downloadImage(f, subreddit)
		}
	}

	wg.Wait()
}

/***************************
     Private Functions
***************************/

func buildEndpoint(subr string, limit int) string {
	return fmt.Sprintf("%s%s%s%s%d", ENDPOINT, subr, ".json?", "limit=", limit)
}

func getImageLink(url string) string {
	return strings.Replace(url, "amp;", "", 1)
}

func downloadImage(url string, sub string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	homeDir, _ := os.UserHomeDir()

	fileName := fmt.Sprintf("%s%s%s%s%s", homeDir, "/Documents/reddit-downloader/", sub, "/", uuid.New())

	file, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	_, e := io.Copy(file, resp.Body)

	if e != nil {
		log.Fatal("ok")
	}
	fmt.Println("image written to : ", fileName)
	wg.Done()

	return true, nil
}
