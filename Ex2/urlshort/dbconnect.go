package urlshort

import (
	"log"
	"fmt"
	"github.com/boltdb/bolt"
	"encoding/json"
)
var(
	db  *bolt.DB
    result []byte

)

// OpenDb opens a connection to a BoltDb database.
func OpenDb() *bolt.DB{
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	conn, err := bolt.Open("pathurls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	fmt.Println("Connected to database")
	return db
	
}

// UpdateDB loads a json data into the in memory database
func UpdateDB(jsn []byte) {
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
			fmt.Println("Path:",data.Path,"Url:",data.URL)
			err = b.Put([]byte(data.Path), []byte(data.URL))
			//log.Println(err)
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

func getdata(path string) ([]byte, error){
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(path))
		fmt.Printf("The answer is: %s\n", v)
		//copy(result,v)
		result = v
		return nil
	})
	if err != nil{
		return nil, err
	}
	fmt.Printf("Returned from db: %s\n",result)
	return result, nil
}
