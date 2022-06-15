package db

import (
	"encoding/binary"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Task struct {
	Key   int
	Value string
}

var taskBucket = []byte("Tasks")

var db *bolt.DB

func InitDB(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: time.Second * 1})
	if err != nil {
		fmt.Println("there was an error opening the database", err)
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}

func ListTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		//fmt.Println("there was an error list tasks", err)
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	if err != nil {
		fmt.Println("del ", err)
	}
	return nil
}

func itob(k int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(k))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
