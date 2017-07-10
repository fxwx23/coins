package coin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// Rate now supports jpy based only in this projects.
type Rate struct {
	JpyBased Currency `json:"jpy"`
}

// Currency covers all encrypted currency which supported by cc.
type Currency struct {
	// US $
	Usd string `json:"usd"`
	// Encrypted currency
	Btc  string `json:"btc"`
	Eth  string `json:"eth"`
	Etc  string `json:"etc"`
	Dao  string `json:"dao"`
	Lsk  string `json:"lsk"`
	Fct  string `json:"fct"`
	Xmr  string `json:"xmr"`
	Rep  string `json:"rep"`
	Xrp  string `json:"xrp"`
	Zec  string `json:"zec"`
	Xem  string `json:"xem"`
	Ltc  string `json:"ltc"`
	Dash string `json:"dash"`
}

const ccAPIURL = "https://coincheck.com/api/rate/all"

func All(c echo.Context) error {
	res, err := getAllRate()
	if err != nil {
		return c.String(http.StatusNoContent, "no content.")
	}

	mes := `
	*Current encrypted curreny status*

	USD:
	` + res.JpyBased.Usd +
		`
	BTC:
	` + res.JpyBased.Btc

	json := map[string]string{
		"response_type": "in_channel",
		"text":          mes,
	}

	return c.JSON(http.StatusOK, json)
}

func getAllRate() (*Rate, error) {
	client := new(http.Client)
	req, err := http.NewRequest("GET", ccAPIURL, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	rate := new(Rate)
	e := json.Unmarshal(b, rate)
	if e != nil {
		return nil, e
	}

	return rate, nil
}
