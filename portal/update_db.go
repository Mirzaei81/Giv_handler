package portal

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Update_Resault struct {
	Success  bool `json:"success"`
	Total    int  `json:"total"`
	Count    int  `json:"count"`
	Products []struct {
		ID         int      `json:"id"`
		Title      string   `json:"title"`
		Caption    any      `json:"caption"`
		Image      any      `json:"image"`
		Slug       string   `json:"slug"`
		URL        string   `json:"url"`
		Rate       any      `json:"rate"`
		RateCount  any      `json:"rate_count"`
		Price      any      `json:"price"`
		Stock      int      `json:"stock"`
		Stats      int      `json:"stats"`
		Comments   any      `json:"comments"`
		Position   int      `json:"position"`
		Status     []string `json:"status"`
		Expiration any      `json:"expiration"`
		Published  struct {
			Year      string `json:"year"`
			Month     string `json:"month"`
			MonthName string `json:"month_name"`
			Day       string `json:"day"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Universal string `json:"universal"`
			Timestamp int    `json:"timestamp"`
			Subtract  string `json:"subtract"`
			Past      bool   `json:"past"`
		} `json:"published"`
		Created struct {
			Year      string `json:"year"`
			Month     string `json:"month"`
			MonthName string `json:"month_name"`
			Day       string `json:"day"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Universal string `json:"universal"`
			Timestamp int    `json:"timestamp"`
			Subtract  string `json:"subtract"`
			Past      bool   `json:"past"`
		} `json:"created"`
		Creator struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Nickname any    `json:"nickname"`
			Avatar   any    `json:"avatar"`
		} `json:"creator"`
	} `json:"products"`
}
type get_product_Struct struct {
	Success bool `json:"success"`
	Product struct {
		ID                int      `json:"id"`
		Version           string   `json:"version"`
		Title             string   `json:"title"`
		Caption           any      `json:"caption"`
		Description       any      `json:"description"`
		Image             any      `json:"image"`
		Slug              string   `json:"slug"`
		URL               string   `json:"url"`
		Rate              any      `json:"rate"`
		RateCount         any      `json:"rate_count"`
		Password          any      `json:"password"`
		Layout            any      `json:"layout"`
		CommentingEnabled bool     `json:"commenting_enabled"`
		MetaTitle         any      `json:"meta_title"`
		MetaDescription   any      `json:"meta_description"`
		MetaKeywords      any      `json:"meta_keywords"`
		MetaRobots        any      `json:"meta_robots"`
		CanonicalURL      any      `json:"canonical_url"`
		Redirect          any      `json:"redirect"`
		Stats             int      `json:"stats"`
		Comments          any      `json:"comments"`
		Position          int      `json:"position"`
		Status            []string `json:"status"`
		Contents          any      `json:"contents"`
		Fields            any      `json:"fields"`
		Images            any      `json:"images"`
		Category          any      `json:"category"`
		Categories        any      `json:"categories"`
		Filters           any      `json:"filters"`
		Attributes        any      `json:"attributes"`
		Variants          []struct {
			ID           int      `json:"id"`
			ProductID    int      `json:"product_id"`
			Title        string   `json:"title"`
			Price        int      `json:"price"`
			ComparePrice any      `json:"compare_price"`
			Tax          any      `json:"tax"`
			Shipping     any      `json:"shipping"`
			Weight       any      `json:"weight"`
			Length       any      `json:"length"`
			Width        any      `json:"width"`
			Height       any      `json:"height"`
			Stock        int      `json:"stock"`
			Minimum      any      `json:"minimum"`
			Maximum      any      `json:"maximum"`
			Sku          *string  `json:"sku"`
			Image        any      `json:"image"`
			Type         string   `json:"type"`
			Status       []string `json:"status"`
			Files        any      `json:"files"`
		} `json:"variants"`
		Relates    any `json:"relates"`
		Expiration any `json:"expiration"`
		Published  struct {
			Year      string `json:"year"`
			Month     string `json:"month"`
			MonthName string `json:"month_name"`
			Day       string `json:"day"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Universal string `json:"universal"`
			Timestamp int    `json:"timestamp"`
			Subtract  string `json:"subtract"`
			Past      bool   `json:"past"`
		} `json:"published"`
		Created struct {
			Year      string `json:"year"`
			Month     string `json:"month"`
			MonthName string `json:"month_name"`
			Day       string `json:"day"`
			Date      string `json:"date"`
			Time      string `json:"time"`
			Universal string `json:"universal"`
			Timestamp int    `json:"timestamp"`
			Subtract  string `json:"subtract"`
			Past      bool   `json:"past"`
		} `json:"created"`
		Creator struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Name     string `json:"name"`
			Nickname any    `json:"nickname"`
			Avatar   any    `json:"avatar"`
		} `json:"creator"`
	} `json:"product"`
}

type GP_PP struct {
	product_id int
}

func Update_db(token string) {
	url := "https://batkap.com/site/api/v1/manage/store/products"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var Update_Resault Update_Resault
	decoder.Decode(&Update_Resault)
	for _, p := range Update_Resault.Products {
		go get_product(p.ID, token)
	}
}
func get_product(id int, token string) get_product_Struct {
	url := fmt.Sprintf("https://batkap.com/site/api/v1/manage/store/products/%d", id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var Product get_product_Struct
	decoder.Decode(&Product)
	key := Product.Product.Variants[0].Sku
	item := GP_PP{product_id: Product.Product.ID}
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err = enc.Encode(item)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if key != nil {
		DB.Write(*key, buf.Bytes())
	}
	return Product
}
