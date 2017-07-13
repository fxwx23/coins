package coin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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

	mes := `*Current cryptocurrency status*

*USD* : $` + res.JpyBased.Usd + `
*BTC* : ¥` + res.JpyBased.Btc + `
*ETH* : ¥` + res.JpyBased.Eth + `
*ETC* : ¥` + res.JpyBased.Etc + `
*DAO* : ¥` + res.JpyBased.Dao + `
*LSK* : ¥` + res.JpyBased.Lsk + `
*FCT* : ¥` + res.JpyBased.Fct + `
*XMR* : ¥` + res.JpyBased.Xmr + `
*REP* : ¥` + res.JpyBased.Rep + `
*XRP* : ¥` + res.JpyBased.Xrp + `
*ZEC* : ¥` + res.JpyBased.Zec + `
*XEM* : ¥` + res.JpyBased.Xem + `
*LTC* : ¥` + res.JpyBased.Ltc + `
*DASH* : ¥` + res.JpyBased.Dash

	if t := c.FormValue("text"); strings.Contains(t, "btc") {
		mes = "*BTC* : ¥" + res.JpyBased.Btc
	}

	if t := c.FormValue("text"); strings.Contains(t, "eth") {
		mes = "*ETH* : ¥" + res.JpyBased.Eth
	}

	if t := c.FormValue("text"); strings.Contains(t, "etc") {
		mes = "*ETC* : ¥" + res.JpyBased.Etc
	}

	if t := c.FormValue("text"); strings.Contains(t, "dao") {
		mes = "*DAO* : ¥" + res.JpyBased.Dao
	}

	if t := c.FormValue("text"); strings.Contains(t, "lsk") {
		mes = "*LSK* : ¥" + res.JpyBased.Lsk
	}

	if t := c.FormValue("text"); strings.Contains(t, "fct") {
		mes = "*FCT* : ¥" + res.JpyBased.Fct
	}

	if t := c.FormValue("text"); strings.Contains(t, "xmr") {
		mes = "*XMR* : ¥" + res.JpyBased.Xmr
	}

	if t := c.FormValue("text"); strings.Contains(t, "rep") {
		mes = "*REP* : ¥" + res.JpyBased.Rep
	}

	if t := c.FormValue("text"); strings.Contains(t, "xrp") {
		mes = "*XRP* : ¥" + res.JpyBased.Xrp
	}

	if t := c.FormValue("text"); strings.Contains(t, "zec") {
		mes = "*ZEC* : ¥" + res.JpyBased.Zec
	}

	if t := c.FormValue("text"); strings.Contains(t, "xem") {
		mes = "*XEM* : ¥" + res.JpyBased.Xem
	}

	if t := c.FormValue("text"); strings.Contains(t, "ltc") {
		mes = "*LTC* : ¥" + res.JpyBased.Ltc
	}

	if t := c.FormValue("text"); strings.Contains(t, "dash") {
		mes = "*DASH* : ¥" + res.JpyBased.Dash
	}

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
