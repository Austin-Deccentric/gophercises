package urlshort

import (
	"log"
	"fmt"
	"github.com/boltdb/bolt"
	"encoding/json"
	"net/http"
)

var(
    result []byte
)


// UpdateDB takes a Slice of Json data and a connection to an Open BoltDb database 
// creates a bucket and updates the bucket with the JSON data.
//
// NB: JSON should have this format:
// {
// 	"path": <URL Path>,
// 	"url": "<The URL to be accessed>"
// }
//Example:
// {
// 	"path": "/bolt",
// 	"url": "https://github.com/boltdb/bolt"
// }
func UpdateDB(jsn []byte,db *bolt.DB ) {
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}

		var pathUrls []pathURL
		if err = json.Unmarshal(jsn, &pathUrls); err != nil{
			return fmt.Errorf("Could not Unmarshal: %s", err)
		}
		for _,data := range pathUrls{
			//fmt.Println("Path:",data.Path,"Url:",data.URL)
			err = b.Put([]byte(data.Path), []byte(data.URL))

			if err != nil{
				return fmt.Errorf("Could not insert into db: %v", err)
			}
		}
		return nil
		
	})
		
		if err != nil{
			log.Fatal(err)
		}
}

// getdata is responsible for retrieving keys from the database.
// The key is the Url path.
func getdata(path string,db  *bolt.DB) ([]byte, error){
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(path))
		//fmt.Printf("The answer is: %s\n", v)
		result = v
		return nil
	})

	if err != nil{
		return nil, err
	}

	//fmt.Printf("Returned from db: %s\n",result)
	return result, nil
}

// DBHandler takes an open Boltdb connection and a  fallback http handler
// that is to called if the Url path is not found.
//
// Queries the database for the string and redirects to the URL if found.
// Otherwise is redirects to serves the Fallback
func DBHandler(fallback http.Handler,db *bolt.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Check for URL in the database
		result,err := getdata(path,db); if err !=nil{
			log.Fatal(err)
		}
		//fmt.Println("The URL is:", result)
		if result == nil {
			log.Println("Url not found. back to home")
			fallback.ServeHTTP(w,r)
			return
		}
		log.Println("Redirecting")
		http.Redirect(w,r,string(result), http.StatusFound)
	}
}
