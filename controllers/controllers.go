package controllers

import (
    "fmt"
    "context"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"

    "github.com/doxxitxxyoung/go-rest-tutorial/models"
    "github.com/doxxitxxyoung/go-rest-tutorial/repository"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"

//    "labix.org/v2/mgo/bson"

)

func GetDrugs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    //  create Drug array
    var drugs []models.Drug

    //  Connect to MongoDB Atlas with Repository
    drugCollection := repository.ConnectDB().Collection("drugs")
//    drugCollection := repository.ConnectDB().Database("glit_db_json").Collection("drugs")


    //  bson.M{} : empty filter.-> get all data
    cur, err := drugCollection.Find(context.TODO(), bson.M{})

    if err != nil {
        repository.GetError(err, w)
        return
    }

    defer cur.Close(context.TODO())

    for cur.Next(context.TODO()) {
        //  create a value into which a single doc can be decoded
        var drug models.Drug

        //  &book: memory address of book
        err := cur.Decode(&drug)

        if err != nil {
            log.Fatal(err)
        }

        drugs = append(drugs, drug)
    }

    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }

    json.NewEncoder(w).Encode(drugs)
}

func GetDrugById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var drug models.Drug
    var params = mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    drugCollection := repository.ConnectDB().Collection("drugs")


    filter := bson.M{"_id": id}
    err := drugCollection.FindOne(context.TODO(), filter).Decode(&drug)

    //  add scenario when there is no result matching the requesr param.
    //  Simply returning null, not stopping the whole server.

    fmt.Println(err)

    if err != nil {
        repository.GetError(err, w)
        return
    }

    json.NewEncoder(w).Encode(drug)
}

func GetDrugByDrugname(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var drug models.Drug
    var params = mux.Vars(r)
    drugname, _ := params["drugname"]

    drugCollection := repository.ConnectDB().Collection("drugs")


    filter := bson.M{"drugname": drugname}
    err := drugCollection.FindOne(context.TODO(), filter).Decode(&drug)

    //  add scenario when there is no result matching the requesr param.
    //  Simply returning null, not stopping the whole server.

    fmt.Println(err)

    if err != nil {
        repository.GetError(err, w)
        return
    }

    json.NewEncoder(w).Encode(drug)
}

func GetSamples(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var samples []models.Sample

    sampleCollection := repository.ConnectDB().Collection("samples")

    cur, err := sampleCollection.Find(context.TODO(), bson.M{})

    if err != nil {
        repository.GetError(err, w)
        return
    }

    defer cur.Close(context.TODO())

    for cur.Next(context.TODO()) {
        var sample models.Sample

        err := cur.Decode(&sample)

        if err != nil {
            log.Fatal(err)
        }

        samples = append(samples, sample)
    }

    if err := cur.Err(); err != nil {
        log.Fatal(err)
    }

    json.NewEncoder(w).Encode(samples)
}

func GetSampleById(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var sample models.Sample
    var params = mux.Vars(r)
    id, _ := primitive.ObjectIDFromHex(params["id"])

    sampleCollection := repository.ConnectDB().Collection("samples")


    filter := bson.M{"_id": id}
    err := sampleCollection.FindOne(context.TODO(), filter).Decode(&sample)

    fmt.Println(err)

    if err != nil {
        repository.GetError(err, w)
        return
    }

    json.NewEncoder(w).Encode(sample)
}

func GetSamplesByDrugname(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var samples []models.Sample
    var params = mux.Vars(r)
    drugname, _ := params["drugname"]

    sampleCollection := repository.ConnectDB().Collection("samples")

    filter := bson.M{"drugname": drugname}
    cur, err := sampleCollection.Find(context.TODO(), filter)

    fmt.Println(err)

    if err != nil {
        repository.GetError(err, w)
        return
    }

    defer cur.Close(context.TODO())

    for cur.Next(context.TODO()) {
        var sample models.Sample

        err := cur.Decode(&sample)

        if err != nil {
            fmt.Println(err)
        }

        samples = append(samples, sample)
    }

    json.NewEncoder(w).Encode(samples)
}
