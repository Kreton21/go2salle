package bolted

import (
	"fmt"

	"github.com/boltdb/bolt"
)

func Cbck(db *bolt.DB, bucket string) error {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return nil
}
func Wdb(bucket, key, value []byte) {
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	Cbck(db, "bucket")

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucket)
		err := b.Put(key, value)
		return err
	})
	db.Close()
}

func Rdb(bucket, key string) []byte {
	var result []byte
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		result = v

		//	fmt.Printf(string(v))

		return nil

	})
	return result
}
