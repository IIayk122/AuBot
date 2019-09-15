package au

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type SearchResult struct {
	Count int `json:"Count"`
	List  []struct {
		AnswersCount int     `json:"AnswersCount"`
		BidsCount    int     `json:"BidsCount"`
		BlitzPrice   float64 `json:"BlitzPrice"`
		Categories   []struct {
			ID       int    `json:"Id"`
			Name     string `json:"Name"`
			ParentID int    `json:"ParentId"`
		} `json:"Categories"`
		CurrentPrice float64 `json:"CurrentPrice"`
		EndDate      int     `json:"EndDate"`
		ID           int     `json:"Id"`
		Images       []struct {
			Big     string `json:"Big"`
			Medium  string `json:"Medium"`
			Small   string `json:"Small"`
			Origin  string `json:"Origin"`
			ImageID string `json:"ImageId"`
		} `json:"Images"`
		IsFavorite    bool   `json:"IsFavorite"`
		IsHighlighted bool   `json:"IsHighlighted"`
		IsKingSize    bool   `json:"IsKingSize"`
		Name          string `json:"Name"`
		Owner         struct {
			Avatars struct {
				Small    string `json:"Small"`
				Big      string `json:"Big"`
				Original string `json:"Original"`
			} `json:"Avatars"`
			Guilds struct {
				STANDARD struct {
					DateClose time.Time `json:"DateClose"`
					DateEnd   time.Time `json:"DateEnd"`
					DateStart time.Time `json:"DateStart"`
					ID        string    `json:"Id"`
					IsActive  bool      `json:"IsActive"`
					Name      string    `json:"Name"`
					Priority  int       `json:"Priority"`
					UserID    int       `json:"UserId"`
				} `json:"STANDARD"`
			} `json:"Guilds"`
			IsBanned        bool    `json:"IsBanned"`
			IsPhoneVerified bool    `json:"IsPhoneVerified"`
			IsPro           bool    `json:"IsPro"`
			Login           string  `json:"Login"`
			Rating          float64 `json:"Rating"`
		} `json:"Owner"`
		RegionID   int     `json:"RegionId"`
		StartDate  int     `json:"StartDate"`
		StartPrice float64 `json:"StartPrice"`
		Status     int     `json:"Status"`
		Type       int     `json:"Type"`
		ViewsCount int     `json:"ViewsCount"`
	} `json:"List"`
}

func Search(searchString string) SearchResult {

	res, err := http.Get("https://apidemo.au.ru/v1/auction?page=1&pageSize=200&regionId=464&query=" + searchString + "&sort=ItemId&dir=DESC&isOtherCities=false")
	if err != nil {
		log.Fatal(err)
	}

	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var app = SearchResult{}
	err1 := json.Unmarshal(robots, &app)
	if err1 != nil {
		log.Fatal(err1)
	}
	return app
	/*i := 0
	fmt.Println(app.Count)
	for _, row := range app.List {
		i++
		fmt.Println(row.ID, row.Name, i)
	}
	*/
}
