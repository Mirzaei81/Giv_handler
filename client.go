package main

import (
	"giv/portal"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/peterbourgon/diskv/v3"
)

var env_vars = [6]string{
	"WEB_TOKEN",
	"ITEM_DETAIL_ID",
	"PORTAL_PASS",
	"PORTAL_USER",
	"LAST_PORTAL_PURCHASE",
	"LAST_GIV_PURCHASE",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while Loading .env file  %s \n", err)
		os.Exit(1)
	}
	f, err := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Couldn't Open Log.txt %s \n", err)
		os.Exit(1)
	}
	defer f.Close()
	w := io.MultiWriter(os.Stdout, f)
	log.SetOutput(w)
	flatTransform := func(s string) []string { return []string{} }
	M := make(map[string]string)
	portal.DB = diskv.New(diskv.Options{
		BasePath:     "GP_PP",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	portal.DB = diskv.New(diskv.Options{
		BasePath:     "PersonaId",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err != nil {
		log.Printf("Error:error while  loading .env %s\n", err)
	}
	token := portal.Make_session()
	//running Program every minute
	for range time.Tick(time.Minute) {
		log.Printf("Syncing Begineing ...\n")
		var wg sync.WaitGroup
		//Syncing giv Items via portal orders
		go portal.Get_orders(token)
		wg.Add(1)
		//updating portal Product with  giv quantity on hand
		go portal.Quantityonhand(token)
		wg.Add(1)
		wg.Wait()
		//updating env variables if failure on  happens
		for _, e := range env_vars {
			M[e] = os.Getenv(e)
		}
		log.Print(M)
		godotenv.Write(M, "./.env")
	}
	for range time.Tick(time.Hour * 24) {
		portal.Update_db(token)
	}
}
package main

import (
	"giv/portal"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/peterbourgon/diskv/v3"
)

var env_vars = [6]string{
	"WEB_TOKEN",
	"ITEM_DETAIL_ID",
	"PORTAL_PASS",
	"PORTAL_USER",
	"LAST_PORTAL_PURCHASE",
	"LAST_GIV_PURCHASE",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while Loading .env file  %s \n", err)
		os.Exit(1)
	}
	f, err := os.OpenFile("Log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("Couldn't Open Log.txt %s \n", err)
		os.Exit(1)
	}
	defer f.Close()
	w := io.MultiWriter(os.Stdout, f)
	log.SetOutput(w)
	flatTransform := func(s string) []string { return []string{} }
	M := make(map[string]string)
	portal.DB = diskv.New(diskv.Options{
		BasePath:     "GP_PP",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	portal.DB = diskv.New(diskv.Options{
		BasePath:     "PersonaId",
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err != nil {
		log.Printf("Error:error while  loading .env %s\n", err)
	}
	token := portal.Make_session()
	//running Program every minute
	for range time.Tick(time.Minute) {
		log.Printf("Syncing Begineing ...\n")
		var wg sync.WaitGroup
		//Syncing giv Items via portal orders
		go portal.Get_orders(token)
		wg.Add(1)
		//updating portal Product with  giv quantity on hand
		go portal.Quantityonhand(token)
		wg.Add(1)
		wg.Wait()
		//updating env variables if failure on  happens
		for _, e := range env_vars {
			M[e] = os.Getenv(e)
		}
		log.Print(M)
		godotenv.Write(M, "./.env")
	}
	for range time.Tick(time.Hour * 24) {
		portal.Update_db(token)
	}
}
