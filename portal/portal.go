package portal

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	givsoft "giv/givsoft"
	"io"
	"log"
	"net/http"
	net_url "net/url"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"time"

	"github.com/peterbourgon/diskv/v3"
	Jalaali "github.com/yaa110/go-persian-calendar"
)

var DB *diskv.Diskv

// var color_names = []string{"مشکی", "قهوه‌ای", "عسلی", "قرمز", "بنفش", "سفید"}
// var color_ids = []int{153239968, 153239969, 169088539, 153239973, 153239966, 153239964}
//
//	var category_to_portalID = map[string][]int{
//		"001": {157759245, 157760171},
//		"002": {169074151},
//		"003": {157760171},
//		"004": {169075112, 169075118},
//		"005": {169074735},
//		"006": {169121062, 169121073},
//		"007": {157759245},
//		"008": {169075118},
//		"009": {153239952},
//		"010": {169074599, 169075052},
//		"011": {169074599, 169075052},
//		"012": {153239952},
//		"013": {169074727, 159213279},
//		"014": {169074922},
//		"015": {153239952},
//		"016": {153239952},
//		"017": {157759245, 157760171},
//		"018": {153239952},
//		"019": {158916151},
//		"020": {153239952},
//		"021": {153239952},
//		"022": {153239952},
//		"023": {169075351},
//		"024": {153239952},
//		"025": {169075351},
//		"026": {153239952},
//		"027": {169074919},
//		"028": {153239952},
//		"029": {153239952},
//		"030": {169075351},
//		"031": {169075354},
//		"032": {157760108},
//		"033": {169075118},
//		"034": {169075027},
//		"035": {169075354},
//		"036": {169075027},
//		"037": {169075351},
//		"038": {169075103},
//		"039": {169075112, 169075118},
//		"040": {169075038},
//		"041": {169075354},
//		"042": {157760278},
//		"043": {157758929},
//		"044": {157760171},
//		"045": {169074151},
//		"046": {169075351},
//		"047": {153239952},
//		"048": {153239952},
//		"049": {159213279},
//		"050": {153239952},
//		"051": {153239952},
//	}
var base_url string = "https://batkap.com/site/"

type Order_result struct {
	Success bool `json:"success"`
	Order   struct {
		ID             int      `json:"id"`
		Description    any      `json:"description"`
		Status         []string `json:"status"`
		Quantity       int      `json:"quantity"`
		Weight         int      `json:"weight"`
		Shipping       int      `json:"shipping"`
		Subtotal       int      `json:"subtotal"`
		Discount       int      `json:"discount"`
		Tax            int      `json:"tax"`
		Price          int      `json:"price"`
		RemainingPrice int      `json:"remaining_price"`
		IP             string   `json:"ip"`
		Contact        struct {
			Name    string `json:"name"`
			Mobile  string `json:"mobile"`
			Phone   any    `json:"phone"`
			Email   any    `json:"email"`
			Country struct {
				ID        int     `json:"id"`
				Name      string  `json:"name"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"country"`
			State struct {
				ID        int     `json:"id"`
				Name      string  `json:"name"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"state"`
			City struct {
				ID        int     `json:"id"`
				Name      string  `json:"name"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"city"`
			Zipcode   string  `json:"zipcode"`
			Address   string  `json:"address"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"contact"`
		Items []struct {
			Variant  any     `json:"variant"`
			Product  any     `json:"product"`
			Title    string  `json:"title"`
			Price    int     `json:"price"`
			Quantity int     `json:"quantity"`
			Weight   any     `json:"weight"`
			Shipping any     `json:"shipping"`
			Discount any     `json:"discount"`
			Sku      *string `json:"sku"`
			Tax      int     `json:"tax"`
		} `json:"items"`
		Coupons   any `json:"coupons"`
		Shipments any `json:"shipments"`
		Payments  []struct {
			ID          int      `json:"id"`
			Description any      `json:"description"`
			ReferenceID string   `json:"reference_id"`
			Type        string   `json:"type"`
			Status      []string `json:"status"`
			SubStatus   any      `json:"sub_status"`
			Amount      int      `json:"amount"`
			Created     struct {
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
			Gateway struct {
				ID    int    `json:"id"`
				Title string `json:"title"`
				Type  string `json:"type"`
				Owner string `json:"owner"`
			} `json:"gateway"`
		} `json:"payments"`
		User struct {
			ID           int    `json:"id"`
			Username     string `json:"username"`
			Name         any    `json:"name"`
			Nickname     any    `json:"nickname"`
			NationalCode any    `json:"national_code"`
			Avatar       any    `json:"avatar"`
		} `json:"user"`
		Label         any `json:"label"`
		ShippingClass struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"shipping_class"`
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
		DueDate struct {
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
		} `json:"due_date"`
		Delivery any `json:"delivery"`
		Updated  any `json:"updated"`
	} `json:"order"`
}
type Order struct {
	ID          int      `json:"id"`
	Description any      `json:"description"`
	Status      []string `json:"status"`
	Quantity    int      `json:"quantity"`
	Weight      int      `json:"weight"`
	Shipping    int      `json:"shipping"`
	Subtotal    int      `json:"subtotal"`
	Discount    int      `json:"discount"`
	Tax         int      `json:"tax"`
	Price       int      `json:"price"`
	Payments    int      `json:"payments"`
	IP          string   `json:"ip"`
	User        struct {
		ID           int    `json:"id"`
		Username     string `json:"username"`
		Name         any    `json:"name"`
		Nickname     any    `json:"nickname"`
		NationalCode any    `json:"national_code"`
		Avatar       any    `json:"avatar"`
	} `json:"user"`
	Contact struct {
		Name    string `json:"name"`
		Mobile  string `json:"mobile"`
		Phone   any    `json:"phone"`
		Email   any    `json:"email"`
		Country struct {
			ID        int     `json:"id"`
			Name      string  `json:"name"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"country"`
		State struct {
			ID        int     `json:"id"`
			Name      string  `json:"name"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"state"`
		City struct {
			ID        int     `json:"id"`
			Name      string  `json:"name"`
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"city"`
		Zipcode   string  `json:"zipcode"`
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"contact"`
	Label         any `json:"label"`
	ShippingClass struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"shipping_class"`
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
	DueDate struct {
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
	} `json:"due_date"`
	Delivery any `json:"delivery"`
	Updated  any `json:"updated"`
}
type Orders struct {
	Success bool     `json:"success"`
	Total   int      `json:"total"`
	Count   int      `json:"count"`
	Orders  []*Order `json:"orders"`
}
type session_resault struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

func Make_session() string {
	url := base_url + "api/v1/user/create-session"
	method := "POST"
	user := os.Getenv("PORTAL_USER")
	password := os.Getenv("PORTAL_PASS")
	payload_string := fmt.Sprintf(`{
		"username": "%s",
		"password": "%s"
	}`, user, password)
	payload := strings.NewReader(payload_string)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	req.Close = true
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", "Application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	dec := json.NewDecoder(res.Body)
	var Token_res session_resault
	err = dec.Decode(&Token_res)
	if err != nil {
		log.Fatal(err)
	}
	return Token_res.Token
}
func Get_orders(token string) {
	todayJ := Jalaali.Now().Format("yyy/MM/dd")
	url := base_url + fmt.Sprintf("api/v1/manage/store/orders?page=1&size=20&status=fulfilled&start=%s", todayJ)
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "Application/json")
	req.Header.Set("Accept-Encoding", "identity")
	req.Close = true
	if err != nil {
		log.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("somthing is wrong with the Body %s\n", err)
	}
	decoder := json.NewDecoder(res.Body)
	var orders Orders
	err = decoder.Decode(&orders)
	if err != nil {
		log.Printf("There was an error while decoding orders %s the main Body : %s \n", err, body)
		return
	}
	lastPortalPurchase := os.Getenv("LAST_PORTAL_PURCHASE")
	for _, order := range orders.Orders {
		if string(order.ID) == lastPortalPurchase {
			go get_Order(token, order.ID)
		}
	}
	if len(orders.Orders) > 0 {
		os.Setenv("LAST_PORTAL_PURCHASE", string(orders.Orders[0].ID))
	}
}
func get_Order(token string, order_id int) {
	url := fmt.Sprintf("https://batkap.com/site/api/v1/manage/store/orders/%d", order_id)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	date := Jalaali.Now().Format("060102")
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("Content-Type", "Application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	log.Println(string(body))
	decder := json.NewDecoder(res.Body)
	var order_resault Order_result
	decder.Decode(&order_resault)
	if err != nil {
		log.Println(err)
		return
	}
	PersonId := string(order_resault.Order.User.ID)
	if DB.Has(PersonId) {
		PersonIdb, err := DB.Read(PersonId)
		if err != nil {
			log.Printf("Somthing Wen wrong oops %s\n", err)
		}
		PersonId = string(PersonIdb)
	} else {
		PersonId = givsoft.Create_customer(string(order_resault.Order.User.ID), order_resault.Order.Contact.Name, order_resault.Order.Contact.City.Name, order_resault.Order.Contact.Address, order_resault.Order.Contact.Mobile, order_resault.Order.Contact.Zipcode)
	}
	date_created, _ := time.Parse("02/01/2006 15:04:05", order_resault.Order.Payments[0].Created.Universal)
	date_created_formated := date_created.Format("2006/01/02 15:04:05")
	itemDetail := make([]givsoft.Itemdetail, 0)
	for _, item := range order_resault.Order.Items {
		if reflect.TypeOf(item.Sku) == reflect.TypeOf("string") {
			if item.Sku != nil {
				itemId, _ := strconv.Atoi(*item.Sku)
				itemDetail = append(itemDetail, givsoft.Itemdetail{
					ItemID:      itemId,
					ItemBarcode: *item.Sku,
					Quantity:    item.Quantity,
					Fee:         item.Price,
					DateCreated: date_created_formated,
					DateChanged: date_created_formated,
				})
			}
		}
	}
	order_detail := givsoft.Order_detail{
		OrderID:            -1,
		SourceID:           order_resault.Order.ID,
		Type:               "WEB",
		No:                 order_resault.Order.ID,
		Date:               date,
		EffectiveDate:      date,
		PersonID:           PersonId,
		CouponCode:         "",
		Description:        "سفارش خرید از پرتال",
		TotalQuantity:      order_resault.Order.Quantity,
		TotalPrice:         order_resault.Order.Price,
		TotalDiscount:      0,
		PackingCost:        0,
		TransferCost:       0,
		PostRefCode:        string(order_resault.Order.ID),
		ReceiverName:       order_resault.Order.Contact.Name,
		ReceiverCity:       order_resault.Order.Contact.City.Name,
		ReceiverAddress:    order_resault.Order.Contact.Address,
		ReceiverMobile:     order_resault.Order.Contact.Mobile,
		ReceiverPostalCode: order_resault.Order.Contact.Zipcode,
		PaymentBank:        order_resault.Order.Payments[0].Gateway.Title,
		PaymentType:        order_resault.Order.Payments[0].Gateway.Type,
		PaymentStatus:      order_resault.Order.Payments[0].Status[0],
		PaymentBankRefCode: order_resault.Order.Payments[0].ReferenceID,
		DateCreated:        date_created_formated,
		DateChanged:        date_created_formated,
	}
	go givsoft.Make_Order(order_detail)
}
func Update_Product(token string, product_id int, stock int) {
	url := fmt.Sprintf("https://batkap.com/site/api/v1/manage/store/products/%d", product_id)
	method := "PUT"
	product := get_product(product_id, token)
	product.Product.Variants[0].Stock = stock
	product_byte, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	payload := strings.NewReader(string(product_byte))
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
func Quantityonhand(token string) {
	url := fmt.Sprintf("http://91.92.214.97:8201/api/quantityonhand?lastDate=%s", net_url.QueryEscape(os.Getenv("LAST_GIV_PURCHASE")))
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	req.Close = true
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("WEB_TOKEN", os.Getenv("WEB_TOKEN"))
	client := &http.Client{}
	res, err := client.Do(req)
	fmt.Print(res)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println(res.StatusCode)
	defer res.Body.Close()
	var qoh givsoft.QuantityOnhand
	if err != nil {
		log.Printf("Error while Setting Env variable %s\n", err)
		os.Exit(1)
	}
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&qoh)
	err = os.Setenv("LAST_GIV_PURCHASE", qoh.LastDatetime)
	if err != nil {
		log.Print(err)
		debug.PrintStack()
		os.Exit(1)
	}
	fmt.Printf("the Last GIV Purchase updated %s ->last datetime to check : %s\n", os.Getenv("LAST_GIV_PURCHASE"), qoh.LastDatetime)
	for _, item := range qoh.Value {
		itemDescb, _ := DB.Read(string(item.ItemID)) //getting ItemId from Giv map to Portal Item Id",
		r := bytes.NewReader(itemDescb)
		dec := gob.NewDecoder(r)
		var itemDesc GP_PP
		err = dec.Decode(&itemDesc)
		if err != nil {
			log.Printf("error while Updating order from giv to portal %s\n", err)
			continue
		}
		Update_Product(token, itemDesc.product_id, int(item.ItemQuantityOnHand))
	}
}
