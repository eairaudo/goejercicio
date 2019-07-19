package miapi

import (
"../../utils/apierrors"
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
)

type Site struct {
	ID							string `json:"id"`
	Name						string `json:"name"`
	CountryID					string `json:"country_id"`
	SaleFeesMode				string `json:"sale_fees_mode"`
	MercadoPagoVersion			int8 `json:"mercadopago_version"`
	DefaultCurrencyID			string `json:"default_currency_id"`
	InmediatePayment			string `json:"immediate_payment"`
	InmediatePaymentMethodID	[]string`json:"payment_method_ids"`
	Settings	struct{
		IdentificationTypes 				[]string`json:"identification_types"`
		TaxpayerTypes 						[]string`json:"taxpayer_types"`
		IdentificationTypesRules 			[] struct{
			IdentificationType			string `json:"identification_type"`
			Rules 			[] struct{
				EnabledTaxpayerTypes 	[]string`json:"enabled_taxpayer_types"`
				BeginsWith 				string`json:"begins_with"`
				Type 					string`json:"type"`
				MinLength 				int8`json:"min_length"`
				MaxLength 				int8`json:"max_length"`

			}`json:"rules"`
		}`json:"identification_types_rules"`
	}`json:"settings"`
	Currencies [] struct{
		ID				string `json:"id"`
		Symbol			string `json:"symbol"`
	}`json:"currencies"`
	Categories [] struct{
		ID				string `json:"id"`
		name			string `json:"name"`
	}`json:"categories"`

}

const urlSite  = "http://localhost:8081/site/"

func (site *Site) Get() *apierrors.ApiError  {

	url := fmt.Sprintf( urlSite + site.ID )

	res , err := http.Get(url)

	if err != nil{
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data , err := ioutil.ReadAll(res.Body)

	if err != nil{
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	if err = json.Unmarshal(data, &site); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}

	}

	return nil

}