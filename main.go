package main

import (
	"giv/givsoft"
	"giv/portal"
	SyncPortal "giv/update"
	"log"
	"os"
	"sync"
	"time"

	godotenv "github.com/joho/godotenv"
	"github.com/peterbourgon/diskv/v3"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var env_vars = [7]string{
	"WEB_TOKEN",
	"ITEM_DETAIL_ID",
	"PORTAL_PASS",
	"PORTAL_USER",
	"LAST_PORTAL_PURCHASE",
	"LAST_GIV_PURCHASE",
	"SUCCESS",
}

// func main() {
// 	godotenv.Load()
// 	itemPrice := givsoft.GetItemDetail(83550001019)
// 	fmt.Println(int(*itemPrice))
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while Loading .env file  %s \n", err)

		os.Exit(1)
	}
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./main.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     10,
		Compress:   true,
	})
	flatTransform := func(s string) []string { return []string{} }
	db := diskv.New(diskv.Options{
		BasePath:     "portal_DB",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	portal.DB = db
	givsoft.DB = db
	SyncPortal.DB = db
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err != nil {
		log.Printf("Error:error while  loading .env %s\n", err)
	}

	token := portal.Make_session()
	wg := new(sync.WaitGroup)
	portal.GetVariants(token)
	for range time.Tick(time.Minute * 5) {
		log.Printf("Syncing Begineing ...\n")
		wg.Add(2)
		go portal.Get_orders(token, wg)    //Syncing giv Items via portal orders
		go givsoft.GetNewOrders(token, wg) //updating portal Product with  giv quantity on hand
		wg.Wait()
	}
	for range time.Tick(time.Minute * 35) {
		log.Println("Sync Interal Database From new Portal Enteries")
		portal.GetVariants(token)
	}
}
