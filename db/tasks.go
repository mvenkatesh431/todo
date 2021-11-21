package db

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	bolt "go.etcd.io/bbolt"
)

var todoBucket = []byte("TodoList")
var db *bolt.DB

type Todo struct {
	Key   int
	Value string
}

// This is exported initialize function.
// We will save the DB file on user "Home" directory, We will use 'os.UserHomeDir()' function
func Initialize() error {

	// os.UserHomeDir, will give the homedirectory (/home/venkey/)
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(dirname + string(os.PathSeparator) + "todo.db")

	// Create the boltDB, It will be created if it doesn't exist.
	db, err = bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	defer db.Close()

	// Create the Bucket in DB, We are using "CreateBucketIfNotExists", So will create only if there is no bucket
	// db.update returns "nil", If everything is fine
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(todoBucket)
		if err != nil {
			return fmt.Errorf("failed to create bucket: %s", err)
		}
		fmt.Println("DB Initialzed successfully at", dbFile)
		return nil
	})
}

func CreateTodo(todo string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket(todoBucket)

		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(todo))
	})

	if err != nil {
		return -1, err
	}
	return id, nil

}

func ListTodos() ([]Todo, error) {
	var todos []Todo

	err := db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket(todoBucket)

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			todos = append(todos, Todo{
				Key:   btoi(k),
				Value: string(v),
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todos, nil

}

func DeleteTodo(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(todoBucket)
		return b.Delete(itob(key))
	})

	return err
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi returns a Uint64 from the binary
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
