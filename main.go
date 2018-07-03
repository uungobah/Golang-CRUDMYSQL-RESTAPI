package main

import(
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

)

func  main()  {
	router := mux.NewRouter()

	router.HandleFunc("/employee", GetAllEmployee).Methods("GET")
	router.HandleFunc("/employee/{id}", GetEmployee).Methods("GET")
	/*router.HandleFunc("/employee/{id}", CreateEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", DeleteEmployee).Methods("DELETE")*/
	log.Fatal(http.ListenAndServe(":8005", router))
}
type Employee struct{
	Id int
	Name string
	City string
}

type Info struct{
	Status string
	Error string
}

type Result struct {
	Employee [] Employee
	Info Info
}


func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Nggakpake29"
	dbName := "my_database"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	var employee []Employee
	var result Result

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	info := Info{}

	info.Status = "Sukses"
	info.Error = "500"

	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
		/*result = append(result, Result{Employee{Id:id,Name:name,City:city},Info{Status:"Sukses",error:"No Error"}})*/
		employee = append(employee, emp)
	}
		result.Employee = employee
		result.Info = info
	defer db.Close()
	json.NewEncoder(w).Encode(result)
}

// Display a single data
func GetEmployee(w http.ResponseWriter, r *http.Request) {
	/*var employee []Employee*/
	db := dbConn()
	params := mux.Vars(r)
	nId := params["id"]
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city string
		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.City = city
	}
	defer db.Close()

	json.NewEncoder(w).Encode(emp)
}

/*
// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}*/
