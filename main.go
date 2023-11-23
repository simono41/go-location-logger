package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Location represents the JSON payload structure
type Location struct {
	Type string  `json:"_type"`
	Tst  int64   `json:"tst"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
	Tid  string  `json:"tid"`
	Batt int     `json:"batt"`
	Vac  int     `json:"vac"`
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	log.Info("Server started. Listening on :8080")

	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	log.Info("Handling request")

	if r.Method != http.MethodPost {
		log.Warn("Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse query parameters
	device := r.URL.Query().Get("d")
	user := r.URL.Query().Get("u")

	log.Infof("Received request with device=%s and user=%s", device, user)

	var loc Location
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loc)
	if err != nil {
		log.Error("Invalid JSON payload: ", err)
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	log.Infof("Received JSON payload: %+v", loc)

	if loc.Type == "location" {
		db, err := sql.Open("mysql", getDBConnectionString())
		if err != nil {
			log.Error("Database connection error: ", err)
			http.Error(w, fmt.Sprintf("Database connection error: %v", err), http.StatusInternalServerError)
			return
		}
		defer db.Close()

		dt := time.Unix(loc.Tst, 0).Format("2006-01-02 15:04:05")

		_, err = db.Exec("INSERT INTO locations (dt, tid, lat, lon, batt, vac, device, user) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			dt, loc.Tid, loc.Lat, loc.Lon, loc.Batt, loc.Vac, device, user)
		if err != nil {
			log.Error("Database insertion error: ", err)
			http.Error(w, fmt.Sprintf("Database insertion error: %v", err), http.StatusInternalServerError)
			return
		}

		log.Info("Location data inserted into the database")
	}

	response := make(map[string]interface{})
	// Optionally add objects to return to the app (e.g., friends or cards)
	json.NewEncoder(w).Encode(response)

	log.Info("Request handled successfully")
}

func getDBConnectionString() string {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
}
