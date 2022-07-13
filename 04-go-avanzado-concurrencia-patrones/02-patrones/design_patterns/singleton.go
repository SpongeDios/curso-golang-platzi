package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct {

}

func (d Database) CreateSingletonConnection() {
	fmt.Println("Creating singleton for database")
	time.Sleep(2 * time.Second)
	fmt.Println("La creacion a terminado")
}

var db *Database
var lock sync.Mutex

func getDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creando una conexion a la bd")
		db = &Database{}
		db.CreateSingletonConnection()
	}else{
		fmt.Println("Ya existe la conexion a la DB")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance()
		}()
	}
	wg.Wait()
}