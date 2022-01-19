package config

import(
	"fmt"
	"github.com/boltdb/bolt"
)

var(
	db *bolt.DB
)

func Connect(){
	// Opens up DB -> returns DB object
	d, err := bolt.Open("test.db", 0600, nil)
	if err!=nil{
		panic(err)
	}
	db = d
	fmt.Println("DB Connected !!")

	// Setup DB if not buckets not exist
	err = setupDB(db)
	if err!=nil{
		panic(err)
	}
}

func GetDB() *bolt.DB {
	return db
}

func setupDB(db *bolt.DB) (error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("urls"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	
	if err != nil {
		return fmt.Errorf("could not set up url buckets, %v", err)
	}
	fmt.Println("DB Setup Done!")

	// Remove here
	err =  db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("urls"))
		err := b.Put([]byte("/path"), []byte( "https://example.com/demo/link"))
		return err
	})
	if err != nil {
		return err
	}

	fmt.Printf("ADDED: %s - %s\n", "/path", "https://example.com/demo/link")
	// till here

	return nil
}