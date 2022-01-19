package models

import(
	"log"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/denilbhatt0814/practice/go-urls/pkg/config"
)

var(
	db *bolt.DB
)

type PathUrl struct{
	Path string	`json:"path"`
	URL  string	`json:"url"`
}

func init(){
	// initiating connection to DB
	config.Connect()
	db = config.GetDB()
}

func GetUrl(p string) (string,error) {
	/* GETS URL FROM KV PAIR DB */
	var url string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		v := b.Get([]byte(p))
		url = string(v)
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	fmt.Printf("GOT: %s - %s\n", p, url)
	return url, nil
}

func (pu *PathUrl)ShortenUrl() *PathUrl {
	err :=  db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte(pu.Path), []byte(pu.URL))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ADDED: %s - %s\n", pu.Path, pu.URL)
	return pu
}

// func main(){
// 	// initing DB
// 	db, err := setupDB()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()


// 	// Adding value to DB
// 	pu1 := pathUrl{"/path", "example.com/demo/link"}
// 	err = addUrl(db, pu1)
// 	if err != nil {
// 		log.Fatal(err)	
// 	}

// 	// Getting value from DB
// 	u, err := getUrl(db, "/path")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	fmt.Println(u)
// }

// func addUrl(db *bolt.DB, pu pathUrl) error {
// 	err :=  db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("urls"))
// 		err := b.Put([]byte(pu.Path), []byte(pu.URL))
// 		return err
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("ADDED: %s - %s\n", pu.Path, pu.URL)
// 	return nil
// }

// func getUrl(db *bolt.DB, p string) (string,error) {
// 	var url string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("urls"))
// 		v := b.Get([]byte(p))
// 		url = string(v)
// 		return nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	fmt.Printf("GOT: %s - %s\n", p, url)
// 	return url, nil
// }

// func setupDB() (*bolt.DB,error) {
// 	db, err := bolt.Open("test.db", 0600, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not open db, %v", err)
// 	}

// 	err = db.Update(func(tx *bolt.Tx) error {
// 		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}
// 		return nil
// 	})
	
// 	if err != nil {
// 		return nil, fmt.Errorf("could not set up url buckets, %v", err)
// 	}
// 	fmt.Println("DB Setup Done!")
// 	return db, nil
// }