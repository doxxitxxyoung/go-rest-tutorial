package repository

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

const AtlasUrl string = "mongodb+srv://doxxitxxyoung:Tele63741@cluster0-9usdz.gcp.mongodb.net"
const DBName string = "glit_db_json"


//func ConnectDB() *mongo.Collection {
func ConnectDB() *mongo.Database {
    clientOptions := options.Client().ApplyURI(AtlasUrl)

    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Atlas Connected.")


    database := client.Database(DBName)

    return database

}

type ErrorResponse struct {
    StatusCode int `json:"status"`
    ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {
    log.Fatal(err.Error())
    var response = ErrorResponse{
        StatusCode: http.StatusInternalServerError,
        ErrorMessage: err.Error(),
    }

    message, _ := json.Marshal(response)

    w.WriteHeader(response.StatusCode)
    w.Write(message)
}
