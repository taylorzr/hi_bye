package storage

import(
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/taylorzr/hibye/root"
)

var db             = connect()
var bucketName     = []byte("hibye")
var usersKey       = []byte("users")
var lastMessageKey = []byte("last_message")
var lastUsersKey = []byte("last_users")

// TODO: These Timestamp read/write functions are exactly the same except the
// key where the data is store. Make this abstract. Also how can we make more
// complex data abstract, e.g. the users data
func ReadUsersTimestamp() (time.Time, error) {
	timeData, err := read(lastUsersKey)

	if err != nil {
		return time.Time{}, err
	}

	t := time.Time{}

	err = t.UnmarshalText(timeData)

	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func WriteUsersTimestamp(t time.Time) error {
	timeText, _ := t.MarshalText()

	err := write(lastUsersKey, timeText)

	if err != nil {
		return err
	}

	return nil
}

func ReadMessageTimestamp() (time.Time, error) {
	timeData, err := read(lastMessageKey)

	if err != nil {
		return time.Time{}, err
	}

	t := time.Time{}

	err = t.UnmarshalText(timeData)

	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

func WriteMessageTimestamp(t time.Time) error {
	timeText, _ := t.MarshalText()

	err := write(lastMessageKey, timeText)

	if err != nil {
		return err
	}

	return nil
}

func ReadUsers() (root.Users, error) {
	usersData, err := read(usersKey)
	if err != nil {
		log.Fatal(err)
	}

	users := root.Users{}
	err = users.Decode(usersData)
	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func WriteUsers(users root.Users) error {
	usersData, err := users.Encode()

	if err != nil {
		log.Fatal(err)
	}

	err = write(usersKey, usersData)

	return err
}

func write(key []byte, data []byte) error {

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
        if bucket == nil {
            return fmt.Errorf("Bucket %q not found!", bucketName)
        }

		err := bucket.Put(key, data)
		if err != nil {
			return err
		}

		return nil
	})
}

func read(key []byte) ([]byte, error) {
	data := []byte{}

	err := db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket(bucketName)
        if bucket == nil {
            fmt.Errorf("Bucket %q not found!", bucketName)
        }

        data = bucket.Get(key)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}

func Exists() bool {
    err := db.View(func(tx *bolt.Tx) error {
        bucket := tx.Bucket(bucketName)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", bucketName)
		} else {
			return nil
		}
	})
	return err == nil
}

func Create() error {
    return db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})
}


func connect() *bolt.DB {
	db, err := bolt.Open("/Users/ztaylo43/bolt.db", 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}
