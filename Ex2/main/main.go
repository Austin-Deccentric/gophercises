package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"urlshort/urlshort"
	"github.com/boltdb/bolt"
)

func main() {
	// flgptr := flag.String("fpath","pathurls.yaml","Provide the absoulute path of a json/yaml file")
	fileptr := flag.String("file","pathurls.json","Provide the absoulute path of a json file to be used to update the database")
	flag.Parse()

	mux := defaultMux()
	

	// data,err:= ioutil.ReadFile(*flgptr); if err!= nil{
	// 	log.Fatal(err)
	// }

	dataDb,err:= ioutil.ReadFile(*fileptr); if err!= nil{
		log.Fatal(err)
	}

	db:= openDb()
	defer db.Close()

	urlshort.UpdateDB(dataDb,db)
	dbHandler := urlshort.DBHandler(mux,db)


	// yamlHandler, err := urlshort.YAMLHandler(data, dbHandler)
	// if err != nil {
	// 	panic(err)
	// }

	// jsonHandler, err := urlshort.JSONHandler(data, mapHandler); if err != nil{
	// 	panic(err)
	// }

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}



// OpenDb opens a connection to a BoltDb database.
func openDb() *bolt.DB{
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("pathurls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//db = conn
	fmt.Println("Connected to database")
	return db
	
}
