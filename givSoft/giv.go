package givsoft

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var base_url = "http://91.92.214.97:8201"
var M map[string]string // map for updating env variables
type Itemdetail struct {
	ItemDetailID int    `json:"ItemDetailID"`
	OrderID      int    `json:"OrderID"`
	ItemID       int    `json:"ItemID"`
	ItemBarcode  string `json:"ItemBarcode"`
	Quantity     int    `json:"Quantity"`
	Fee          int    `json:"Fee"`
	DateCreated  string `json:"DateCreated"`
	DateChanged  string `json:"DateChanged"`
}
type order_resault struct {
	Code         int    `json:"Code"`
	Message      string `json:"Message"`
	PageIndex    any    `json:"PageIndex"`
	PageSize     any    `json:"PageSize"`
	ResultSize   any    `json:"ResultSize"`
	TotalCount   any    `json:"TotalCount"`
	LastDatetime any    `json:"LastDatetime"`
	Value        struct {
		OrderID            int     `json:"OrderID"`
		SourceID           int     `json:"SourceID"`
		Type               string  `json:"Type"`
		No                 string  `json:"No"`
		Date               string  `json:"Date"`
		EffectiveDate      string  `json:"EffectiveDate"`
		PersonID           int     `json:"PersonID"`
		CouponCode         string  `json:"CouponCode"`
		Description        string  `json:"Description"`
		TotalQuantity      float64 `json:"TotalQuantity"`
		TotalPrice         float64 `json:"TotalPrice"`
		TotalDiscount      float64 `json:"TotalDiscount"`
		PackingCost        float64 `json:"PackingCost"`
		TransferCost       float64 `json:"TransferCost"`
		PostRefCode        any     `json:"PostRefCode"`
		ReceiverName       string  `json:"ReceiverName"`
		ReceiverProvinceID int     `json:"ReceiverProvinceID"`
		ReceiverCity       string  `json:"ReceiverCity"`
		ReceiverAddress    string  `json:"ReceiverAddress"`
		ReceiverTel        string  `json:"ReceiverTel"`
		ReceiverMobile     string  `json:"ReceiverMobile"`
		ReceiverPostalCode string  `json:"ReceiverPostalCode"`
		PaymentBank        string  `json:"PaymentBank"`
		PaymentType        string  `json:"PaymentType"`
		PaymentStatus      string  `json:"PaymentStatus"`
		PaymentBankRefCode string  `json:"PaymentBankRefCode"`
		DateCreated        string  `json:"DateCreated"`
		DateChanged        string  `json:"DateChanged"`
		CreditUsed         float64 `json:"CreditUsed"`
	} `json:"Value"`
}
type Order_detail struct {
	OrderID            int    `json:"OrderID"`
	SourceID           int    `json:"SourceID"`
	Type               string `json:"Type"`
	No                 int    `json:"No"`
	Date               string `json:"Date"`
	EffectiveDate      string `json:"EffectiveDate"`
	PersonID           string `json:"PersonID"`
	CouponCode         string `json:"CouponCode"`
	Description        string `json:"Description"`
	TotalQuantity      int    `json:"TotalQuantity"`
	TotalPrice         int    `json:"TotalPrice"`
	TotalDiscount      int    `json:"TotalDiscount"`
	PackingCost        int    `json:"PackingCost"`
	TransferCost       int    `json:"TransferCost"`
	PostRefCode        string `json:"PostRefCode"`
	ReceiverName       string `json:"ReceiverName"`
	ReceiverProvinceID int    `json:"ReceiverProvinceID"`
	ReceiverCity       string `json:"ReceiverCity"`
	ReceiverAddress    string `json:"ReceiverAddress"`
	ReceiverTel        string `json:"ReceiverTel"`
	ReceiverMobile     string `json:"ReceiverMobile"`
	ReceiverPostalCode string `json:"ReceiverPostalCode"`
	PaymentBank        string `json:"PaymentBank"`
	PaymentType        string `json:"PaymentType"`
	PaymentStatus      string `json:"PaymentStatus"`
	PaymentBankRefCode string `json:"PaymentBankRefCode"`
	DateCreated        string `json:"DateCreated"`
	DateChanged        string `json:"DateChanged"`
	ItemDetail         []Itemdetail
}

type QuantityOnhand struct {
	Code         int    `json:"Code"`
	Message      string `json:"Message"`
	PageIndex    int    `json:"PageIndex"`
	PageSize     int    `json:"PageSize"`
	ResultSize   int    `json:"ResultSize"`
	TotalCount   int    `json:"TotalCount"`
	LastDatetime string `json:"LastDatetime"`
	Value        []struct {
		ItemID             int64   `json:"ItemID"`
		Item               any     `json:"Item"`
		ItemQuantityOnHand float64 `json:"ItemQuantityOnHand"`
		LastDate           string  `json:"LastDate"`
		IsActive           bool    `json:"IsActive"`
	} `json:"Value"`
}
type MakeAnOrder struct {
	Code         int    `json:"Code"`
	Message      string `json:"Message"`
	PageIndex    any    `json:"PageIndex"`
	PageSize     any    `json:"PageSize"`
	ResultSize   any    `json:"ResultSize"`
	TotalCount   any    `json:"TotalCount"`
	LastDatetime any    `json:"LastDatetime"`
	Value        struct {
		OrderID            int     `json:"OrderID"`
		SourceID           int     `json:"SourceID"`
		Type               string  `json:"Type"`
		No                 string  `json:"No"`
		Date               string  `json:"Date"`
		EffectiveDate      string  `json:"EffectiveDate"`
		PersonID           int     `json:"PersonID"`
		CouponCode         string  `json:"CouponCode"`
		Description        string  `json:"Description"`
		TotalQuantity      float64 `json:"TotalQuantity"`
		TotalPrice         float64 `json:"TotalPrice"`
		TotalDiscount      float64 `json:"TotalDiscount"`
		PackingCost        float64 `json:"PackingCost"`
		TransferCost       float64 `json:"TransferCost"`
		PostRefCode        string  `json:"PostRefCode"`
		ReceiverName       string  `json:"ReceiverName"`
		ReceiverProvinceID int     `json:"ReceiverProvinceID"`
		ReceiverCity       string  `json:"ReceiverCity"`
		ReceiverAddress    string  `json:"ReceiverAddress"`
		ReceiverTel        string  `json:"ReceiverTel"`
		ReceiverMobile     string  `json:"ReceiverMobile"`
		ReceiverPostalCode string  `json:"ReceiverPostalCode"`
		PaymentBank        string  `json:"PaymentBank"`
		PaymentType        string  `json:"PaymentType"`
		PaymentStatus      string  `json:"PaymentStatus"`
		PaymentBankRefCode string  `json:"PaymentBankRefCode"`
		DateCreated        string  `json:"DateCreated"`
		DateChanged        string  `json:"DateChanged"`
		CreditUsed         any     `json:"CreditUsed"`
	} `json:"Value"`
}
type CreateCustomer struct {
	Code         int    `json:"Code"`
	Message      string `json:"Message"`
	PageIndex    any    `json:"PageIndex"`
	PageSize     any    `json:"PageSize"`
	ResultSize   any    `json:"ResultSize"`
	TotalCount   any    `json:"TotalCount"`
	LastDatetime any    `json:"LastDatetime"`
	Value        struct {
		PersonID              int    `json:"PersonID"`
		FirstName             string `json:"FirstName"`
		LastName              string `json:"LastName"`
		Address               string `json:"Address"`
		Tel                   string `json:"Tel"`
		Mobile                string `json:"Mobile"`
		IsActive              bool   `json:"IsActive"`
		Email                 any    `json:"Email"`
		ProvinceID            any    `json:"ProvinceId"`
		City                  string `json:"City"`
		PostalCode            string `json:"PostalCode"`
		SexCode               any    `json:"SexCode"`
		BirthDate             any    `json:"BirthDate"`
		WeddingDate           any    `json:"WeddingDate"`
		HousbandBirthDate     any    `json:"HousbandBirthDate"`
		ImportantDate         any    `json:"ImportantDate"`
		ImportantDateDesc     any    `json:"ImportantDateDesc"`
		SpecialDiscountRate   any    `json:"SpecialDiscountRate"`
		GradeCode             any    `json:"GradeCode"`
		ClassCode             any    `json:"ClassCode"`
		SpecialDiscountAmount any    `json:"SpecialDiscountAmount"`
		SpecialDiscountType   any    `json:"SpecialDiscountType"`
		VIPCode               any    `json:"VIPCode"`
		VIPCardIssueDate      any    `json:"VIPCardIssueDate"`
		DateCreated           string `json:"DateCreated"`
		Occupation            any    `json:"Occupation"`
		HousbandOccupation    any    `json:"HousbandOccupation"`
		Nationality           any    `json:"Nationality"`
		NationalIDNo          any    `json:"NationalIDNo"`
		LastDate              string `json:"LastDate"`
		DateChanged           any    `json:"DateChanged"`
		Description           any    `json:"Description"`
	} `json:"Value"`
}

func Create_customer(PersonId string, FirstName string, City string, Address string, Mobile string, PostalCode string) string {
	url := base_url + "/api/customer"
	method := "POST"
	date := time.Now().Format("2020-01-02 15:04:05")
	payload_string := fmt.Sprintf(`{
    "PersonID":%s,
	"FirstName":"%s",
	"LastName":"",
    "City": "%s",
    "Address":"%s",
	"Mobile":"%s",
    "DateCreated": "%s",
	"IsActive":true,
    "PostalCode": "%s"
}`, PersonId, FirstName, City, Address, Mobile, date, PostalCode)
	payload := strings.NewReader(payload_string)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("WEB_TOKEN", os.Getenv("WEB_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var newCustomer CreateCustomer
	decoder.Decode(&newCustomer)
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, uint32(newCustomer.Value.PersonID))
	if err != nil {
		log.Println(err)
	}
	return string(newCustomer.Value.PersonID)
}
func Make_Order(order_datail Order_detail) {
	url := base_url + "/api/order"
	method := "POST"
	payload_string, _ := json.Marshal(order_datail)
	payload := strings.NewReader(string(payload_string))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("WEB_TOKEN", os.Getenv("WEB_TOKEN"))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var order_resault order_resault
	decoder.Decode(&order_resault)
	lastItemId, err := strconv.Atoi(os.Getenv("ITEM_DETAIL_ID"))
	if err != nil {
		fmt.Print(err)
	}
	for idx, item := range order_datail.ItemDetail {
		ItemDetail := Itemdetail{
			ItemDetailID: lastItemId + idx,
			OrderID:      order_resault.Value.OrderID,
			ItemID:       item.ItemID,
			ItemBarcode:  item.ItemBarcode,
			Quantity:     item.Quantity,
			Fee:          item.Fee,
			DateCreated:  item.DateCreated,
			DateChanged:  item.DateChanged,
		}
		go Submit_order(ItemDetail)
	}
	os.Setenv("ITEM_DETAIL_ID", string(lastItemId+1))
}

func Submit_order(item_desc Itemdetail) {
	url := base_url + "/api/orderrow"
	method := "POST"
	payload_string := fmt.Sprintf(`{
	"OrderID":%d,
	"RowID": %d,
	"ItemID": %d,
	"ItemBarcode": %s,
	"Quantity": %d,
	"Fee": %d,
	"RowDiscount": 0,
	"TotalDiscount": 0,
	"VatValue": 0,
	"DateCreated": %s,
	"DateChanged": %s
	}`, item_desc.OrderID, item_desc.OrderID, item_desc.ItemID, item_desc.ItemBarcode, item_desc.Quantity, item_desc.Fee, item_desc.DateCreated, item_desc.DateChanged)
	payload := strings.NewReader(payload_string)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(body))
}
