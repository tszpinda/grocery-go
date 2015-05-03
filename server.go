package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rcrowley/go-tigertonic"

	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	mux *tigertonic.TrieServeMux
)

func main() {
	mux := tigertonic.NewTrieServeMux()

	db := openDB()
	db.LogMode(true)

	fruitDao := FruitDao{DB: db}
	fruitResource := FruitResource{fruitDao}
	dbHealthCheck := DbHealthCheck{DB: db}
	//health check
	bindGet(mux, "/healthCheck/database", dbHealthCheck.Check)

	bindGet(mux, "/fruits", fruitResource.findFruits)
	bindGet(mux, "/fruits/search", fruitResource.searchFruits)
	//POST
	bindPost(mux, "/fruits/{id}/price", fruitResource.updateFruitPrice)
	//PUT add override
	bindPut(mux, "/fruits/addOrOverride", fruitResource.addOrOverride)

	startHttpServer(mux)
}

func startHttpServer(mux *tigertonic.TrieServeMux) {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	server := tigertonic.NewServer(":"+port, tigertonic.ApacheLogged(mux))

	err := server.ListenAndServe()
	if nil != err {
		log.Println(err)
	}
}

func handler(path string, method interface{}) http.Handler {
	handler := tigertonic.Timed(
		tigertonic.Marshaled(method),
		path,
		nil,
	)

	return handler
}

func bindPost(mux *tigertonic.TrieServeMux, path string, method interface{}) {
	handler := handler(path, method)
	mux.Handle("POST", path, handler)
}

func bindGet(mux *tigertonic.TrieServeMux, path string, method interface{}) {
	handler := handler(path, method)
	mux.Handle("GET", path, handler)
}
func bindPut(mux *tigertonic.TrieServeMux, path string, method interface{}) {
	handler := handler(path, method)
	mux.Handle("PUT", path, handler)
}

func openDB() *gorm.DB {
	dbSql, _ := openSqlDB()
	db, err := gorm.Open("sqlite3", dbSql)
	if err != nil {
		log.Fatalln(err)
	}
	db.DB()
	db.AutoMigrate(&Fruit{})

	return &db
}

func openSqlDB() (*sql.DB, *string) {
	openPath := ":memory:"

	outDir, err := ioutil.TempDir(os.TempDir(), "grocery")
	if err != nil {
		log.Fatalln(err)
	}
	openPath = filepath.Join(outDir, "grocery.db")
	db, err := sql.Open("sqlite3", openPath)

	if err != nil {
		log.Fatalln(err)
	}
	return db, &openPath
}
