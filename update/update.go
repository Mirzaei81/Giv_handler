package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"giv/types"
	"log"
	"net/http"
	"time"

	"github.com/peterbourgon/diskv/v3"
)

var DB *diskv.Diskv

func Update_Variants(token string, product_id int, stock *int, sku string, price float64, isTesting bool) {
	url := fmt.Sprintf("https://batkap.com/site/api/v1/manage/store/products/variants/%d", product_id)
	method := "PUT"
	variant := getVariant(token, product_id)
	if(variant==nil) {return}
	time.Sleep(time.Microsecond * 300)

	if stock != nil {
		variant.Stock = *stock
	}
	if (price != 0 && variant.ComparePrice == 0 && variant.Price == 0) || isTesting {
		variant.Price = int(price / 10)
		variant.ComparePrice = int(price / 10)
	}
	variant.Sku = sku
	product_byte, err := json.Marshal(variant)
	log.Printf("Updaing Variant %s\n", string(product_byte))
	fmt.Printf("Updaing Variant %s\n", string(product_byte))
	if err != nil {
		log.Println(string(product_byte))
	}
	payload := bytes.NewReader(product_byte)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
}

func getVariant(token string, variantid int) *types.Variant {
	url := fmt.Sprintf("https://batkap.com/site/api/v1/manage/store/products/variants/%d", variantid)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)
	variant := new(types.VariantResult)
	err = dec.Decode(variant)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &variant.Variant
}
