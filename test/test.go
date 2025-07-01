package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Product_resault struct {
	Success bool `json:"success"`
}
type csvPath struct {
	Success     bool   `json:"success"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

var base_url = "https://batkap.com"

//		func main() {
//			godotenv.Load()
//			log.SetFlags(log.LstdFlags | log.Lshortfile)
//			flatTransform := func(s string) []string { return []string{} }
//			db := diskv.New(diskv.Options{
//				BasePath:     "portal_DB",
//				Transform:    flatTransform,
//				CacheSizeMax: 1024 * 1024,
//			})
//			db.Write("LASTGIVODER", []byte("2024-12-02 10:52:32"))
//	}
func main() {
	f, _ := os.OpenFile("test.txt", os.O_WRONLY, 0755)
	f.Seek(0, 0)
	info, _ := f.Stat()
	buffer := make([]byte, info.Size())
	f.Write(buffer)
	f.Close()
}

//	func main() {
//		godotenv.Load()
//		token := portal.Make_session()
//		log.SetFlags(log.LstdFlags | log.Lshortfile)
//		flatTransform := func(s string) []string { return []string{} }
//		db := diskv.New(diskv.Options{
//			BasePath:     "portal_DB",
//			Transform:    flatTransform,
//			CacheSizeMax: 1024 * 1024,
//		})
//		portal.DB = db
//		givsoft.DB = db
//		update.DB = db
//		file ,_ := os.Open("./bk.csv")
//		defer file.Close()
//		reader:= csv.NewReader(file)
//		readCsvbk(reader,token)
//
// }
//
//	func getCsv() {
//		url := "https://batkap.com/site/api/v1/manage/store/products/variants/export"
//		method := "GET"
//		req, err := http.NewRequest(method, url, nil)
//		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token)k)
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(-1)
//		}
//		client := &http.Client{}
//		res, err := client.Do(req)
//		defer res.Body.Close()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(-1)
//		}
//		decoder := json.NewDecoder(res.Body)
//		csvPath := new(csvPath)
//		decoder.Decode(csvPath)
//		readCsv(csvPath.Path, token)
//	}
func readCsv(uri string, token string) {
	url := base_url + uri
	method := "GET"
	req, err := http.NewRequest(method, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	fmt.Println("getting new Csv File", url)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	reader := csv.NewReader(bytes.NewBuffer(body))
	_, err = reader.Read()
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				log.Println(err)
				break
			}
		}
		if line[8] != "" {
			itemId, _ := strconv.ParseInt(line[0], 10, 64)
			fmt.Printf("Updating variant : %s with Sku Of %s with id of %d", line[2], line[8], itemId)
		}
	}
}
func readCsvbk(reader *csv.Reader, token string) {
	_, err := reader.Read()
	if err != nil {
		fmt.Print(err)
		os.Exit(-1)
	}
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				log.Println(err)
				break
			}
		}
		if line[2] != "" {
			itemId, _ := strconv.ParseInt(line[0], 10, 64)
			fmt.Printf("Updating variant : %s with Sku Of %s with id of %d\n", line[1], line[2], itemId)
			//givsoft.QuantityOnhand_byitem(token, line[2], int(itemId), true)
		}
	}
}

