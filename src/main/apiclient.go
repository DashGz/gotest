package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Site struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	CountryID          string   `json:"country_id"`
	SaleFeesMode       string   `json:"sale_fees_mode"`
	MercadopagoVersion int      `json:"mercadopago_version"`
	DefaultCurrencyID  string   `json:"default_currency_id"`
	ImmediatePayment   string   `json:"immediate_payment"`
	PaymentMethodIds   []string `json:"payment_method_ids"`
	Settings struct {
		IdentificationTypes      []interface{} `json:"identification_types"`
		TaxpayerTypes            []interface{} `json:"taxpayer_types"`
		IdentificationTypesRules []interface{} `json:"identification_types_rules"`
	} `json:"settings"`
	Currencies []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Categories []struct {
	} `json:"categories"`
}

func main()  {
	var siteID string
	fmt.Println("Ingrese Site ID:")
	fmt.Scan(&siteID)

	response, err := http.Get(fmt.Sprintf("https://api.mercadolibre.com/sites/%s", siteID))
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var site Site
	err = json.Unmarshal(bytes, &site)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(site)
	fmt.Printf("%+v\n", site)


}