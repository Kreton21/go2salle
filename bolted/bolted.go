package bolted

import (
	"fmt"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func InitDay(salles int) error {
	day := time.Now()
	date := day.Format("2006-01-02")

	fmt.Println(date)
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(date))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		b := tx.Bucket([]byte(date))
		for i := 0; i < salles; i++ {
			_, err := b.CreateBucket([]byte(strconv.Itoa(i)))
			if err != nil {
				return fmt.Errorf("create demi-bucket: %s", err)
			}
		}

		return nil
	})

	return nil

}
func Reserv(datenow, date, time, salle, id string) error {
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	keyval := date + time + salle + id

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(datenow))
		c := b.Bucket([]byte(salle))
		var keye int
		if err := b.ForEach(func(k, v []byte) error {
			if err != nil {
				fmt.Println(err)
				k = []byte("1")
			}
			keye, err = strconv.Atoi(string(k))
			if err != nil {
				// handle error
				return err
			}
			return nil
		}); err != nil {
			return err
		}
		key := []byte(strconv.Itoa(keye))
		err := c.Put(key, []byte(keyval))
		return err
	})
	fmt.Println(keyval)
	return nil
}
func Test(date, salle string) {
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(date))
		c := b.Bucket([]byte(salle))
		err = c.ForEach(func(k, v []byte) error {
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("La cle", string(k), "et la valeur", string(v))
			return err
		})
		return err
	})
}

func Wdb(bucket *bolt.Bucket, key, value []byte) {
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		err := bucket.Put(key, value)
		return err
	})
	db.Close()
}

func Rdb(bucket *bolt.Bucket, key string) []byte {
	var result []byte
	db, err := bolt.Open("reserv.db", 0600, nil)
	if err != nil {
		fmt.Print(err)
	}

	db.View(func(tx *bolt.Tx) error {
		v := bucket.Get([]byte(key))
		result = v

		//	fmt.Printf(string(v))

		return nil

	})
	return result
}
