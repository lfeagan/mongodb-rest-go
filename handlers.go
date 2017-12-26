package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//"strconv"

	"github.com/gorilla/mux"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func getClient() *mgo.Session {
	url := "localhost:27017"
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	return session
}

func DbIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// get Mongo client for this session based on cookie or create new client
	client := getClient()
	// fetch database names
	dbNames, err := client.DatabaseNames()
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(dbNames); err != nil {
		panic(err)
	}
}

func CollectionIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	dbName := vars["dbName"]
	// get Mongo client for this session based on cookie or create new client
	client := getClient()
	// fetch collection names
	collectionNames,err := client.DB(dbName).CollectionNames()
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(collectionNames); err != nil {
		panic(err)
	}
}

func QueryCollection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	// get Mongo client for this session based on cookie or create new client
	client := getClient()
	// query collection
	vars := mux.Vars(r)
	dbName := vars["dbName"]
	collectionName := vars["collectionName"]
	collection := client.DB(dbName).C(collectionName)
	query,err := extractQuery(r)
	if err != nil {
		panic(err)
	}
	fields,err := extractFields(r)
	if err != nil {
		panic(err)
	}
	sort,err := extractSort(r)
	if err != nil {
		panic(err)
	}

	var result []bson.M
	if fields == nil {
		if sort == nil {
			err = collection.Find(query).Iter().All(&result)
		} else {
			err = collection.Find(query).Sort(sort...).Iter().All(&result)
		}
	} else {
		if sort == nil {
			err = collection.Find(query).Select(fields).Iter().All(&result)
		} else {
			err = collection.Find(query).Select(fields).Sort(sort...).Iter().All(&result)
		}
	}
	if err != nil {
		panic(err)
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func extractQuery(r *http.Request) (bson.M,error) {
	rQuery := r.URL.Query()
	sQuery := rQuery.Get("query")
	if sQuery == "" {
		return nil,nil
	}
	var bQuery bson.M
	var err error
	err = json.NewDecoder(strings.NewReader(sQuery)).Decode(&bQuery)
	if err != nil {
		return nil, err
	} else {
		return bQuery, nil
	}
}

func extractFields(r *http.Request) (bson.M,error) {
	rQuery := r.URL.Query()
	sFields := rQuery.Get("fields")
	if sFields == "" {
		return nil,nil
	}
	var bQuery bson.M
	var err error
	err = json.NewDecoder(strings.NewReader(sFields)).Decode(&bQuery)
	if err != nil {
		return nil, err
	} else {
		return bQuery, nil
	}
}

func extractSort(r *http.Request) ([]string,error) {
	rQuery := r.URL.Query()
	sSort := rQuery.Get("sort")
	if sSort == "" {
		return nil,nil
	}
	var fields []string
	var err error
	err = json.NewDecoder(strings.NewReader(sSort)).Decode(&fields)
	if err != nil {
		return nil, err
	} else {
		return fields, nil
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func InsertCollection(w http.ResponseWriter, r *http.Request) {
	var document bson.M
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &document); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// get Mongo client for this session based on cookie or create new client
	client := getClient()
	// insert to collection
	vars := mux.Vars(r)
	dbName := vars["dbName"]
	collectionName := vars["collectionName"]
	collection := client.DB(dbName).C(collectionName)

	err = collection.Insert(document)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(document); err != nil {
		panic(err)
	}
}

func extractDocument(r *http.Request) (bson.M,error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var document bson.M
	//var err error
	//err = json.NewDecoder(r.Body).Decode(&document)
	err = json.Unmarshal(body, &document)
	//defer r.Body.Close()
	if err != nil {
		return nil, err
	} else {
		return document, nil
	}
}

