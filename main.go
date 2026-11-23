package main

import (
        "encoding/json"
        "fmt"
        "net/http"
        "database/sql"
        _ "github.com/go-sql-driver/mysql"

)

type sensor struct {
        Light       float64 `json:"light"`
        Temperature float64 `json:"temperature"`
        Humidity    float64 `json:"humidity"`
}

var db *sql.DB
func test(rw http.ResponseWriter, req *http.Request) {
        decoder := json.NewDecoder(req.Body)
        var t sensor
        err := decoder.Decode(&t)
        if err != nil {
                panic(err)
        }
        _, err = db.Exec("INSERT INTO sensor(hum, temp, light) VALUES (?, ?, ?)", t.Humidity, t.Temperature, t.Light)
        if err != nil{
                fmt.Println(err)
        }
        fmt.Println(t.Light, t.Temperature, t.Humidity)
}

func main() {
var err error
        db, err = sql.Open("mysql", "root:1234@tcp(10.152.34.252:3307)/mydb")
        if err != nil{
                fmt.Println("Error connecting db")
                return;
        }
        defer db.Close()
        _, err = db.Exec("CREATE TABLE IF NOT EXISTS sensor (id INT AUTO_INCREMENT PRIMARY KEY, hum FLOAT, temp FLOAT, light FLOAT,timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP);")
        if err != nil {
                fmt.Println(err)
        }

        fmt.Println("Starting server on port 5000...")
        http.HandleFunc("/sensor/data", test)
        fmt.Print(http.ListenAndServe(":5000", nil))
}
