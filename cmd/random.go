package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "get beer randomly",
	Long:  "get beer randomly",
	Run: func(cmd *cobra.Command, args []string) {
		res := random()

		for _, v := range *res {
			fmt.Println(v.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func random() *Beers {
	endpoint := "https://api.punkapi.com/v2/beers/random"
	req, _ := http.NewRequest("GET", endpoint, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	jsonBytes := ([]byte)(byteArray)
	data := new(Beers)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	return data
}

type Beers []Beer

type Beer struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Tagline          string  `json:"tagline"`
	FirstBrewed      string  `json:"first_brewed"`
	Description      string  `json:"description"`
	ImageURL         string  `json:"image_url"`
	Abv              float64 `json:"abv"`
	Ibu              float64 `json:"ibu"`
	TargetFg         float64 `json:"target_fg"`
	TargetOg         float64 `json:"target_og"`
	Ebc              float64 `json:"ebc"`
	Srm              float64 `json:"srm"`
	Ph               float64 `json:"ph"`
	AttenuationLevel float64 `json:"attenuation_level"`
	Volume           struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"volume"`
	BoilVolume struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"boil_volume"`
	Method struct {
		MashTemp []struct {
			Temp struct {
				Value int    `json:"value"`
				Unit  string `json:"unit"`
			} `json:"temp"`
			Duration int `json:"duration"`
		} `json:"mash_temp"`
		Fermentation struct {
			Temp struct {
				Value int    `json:"value"`
				Unit  string `json:"unit"`
			} `json:"temp"`
		} `json:"fermentation"`
		Twist string `json:"twist"`
	} `json:"method"`
	Ingredients struct {
		Malt []struct {
			Name   string `json:"name"`
			Amount struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"amount"`
		} `json:"malt"`
		Hops []struct {
			Name   string `json:"name"`
			Amount struct {
				Value float64 `json:"value"`
				Unit  string  `json:"unit"`
			} `json:"amount"`
			Add       string `json:"add"`
			Attribute string `json:"attribute"`
		} `json:"hops"`
		Yeast string `json:"yeast"`
	} `json:"ingredients"`
	FoodPairing   []string `json:"food_pairing"`
	BrewersTips   string   `json:"brewers_tips"`
	ContributedBy string   `json:"contributed_by"`
}
