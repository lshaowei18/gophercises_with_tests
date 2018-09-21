package db

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Key   int
	Value string
}

var taskBucket = []byte("tasks")
var db *bolt.DB

//Setup the db
func SetupDB(file string) error {
	var err error
	//Open the db with 1 second time out
	db, err = bolt.Open(file, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return fmt.Errorf("could not open db, %v", err)
	}
	//create bucket if it doesn't exist
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

//Create a task
func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		//Get the bucket named "taskBucket"
		b := tx.Bucket(taskBucket)
		//Get the next autoincrementing integer for the bucket
		id64, _ := b.NextSequence()
		//Convert the id into int and store it
		id = int(id64)
		//Change it into a 8byte
		key := itob(id)
		//Sets the value of the key in the bucket
		return b.Put(key, []byte(task))
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

//List all tasks
func ListAllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		//Get bucket
		b := tx.Bucket(taskBucket)
		//get cursor
		c := b.Cursor()
		//Using cursor to loop through all tasks sequentially
		for k, v := c.First(); k != nil; k, v = c.Next() {
			//create task
			task := Task{
				Key:   btoi(k),
				Value: string(v),
			}
			//Append to tasks
			tasks = append(tasks, task)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

//Function to delete a task
func DeleteTaskByID(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		//Get bucket
		b := tx.Bucket(taskBucket)
		//Get Key
		key := itob(id)
		//Delete the value and key
		return b.Delete(key)
	})
}

func GetTaskByID(id int) (Task, error) {
	var task Task
	task.Key = id
	err := db.View(func(tx *bolt.Tx) error {
		//get bucket
		b := tx.Bucket(taskBucket)
		//get Key
		key := itob(id)
		task.Value = string(b.Get(key))
		if task.Value == "" {
			return fmt.Errorf("Did not get the task back.")
		}
		return nil
	})
	return task, err
}

//Take a int and turn it into bytes
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

//Take bytes and convert into int
func btoi(v []byte) int {
	i := int(binary.BigEndian.Uint64(v))
	return i
}
