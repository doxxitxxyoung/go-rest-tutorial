package main


/*
go get go.mongodb.org/mongo-driver/mongo

*/
import (
    "fmt"
    "log"
//    "time"
    "net/http"
    "context"
    "github.com/gorilla/mux"

    //  mongo driver
//    "github.com/mongodb/mongo-go-driver/bson/primitive"
//    "github.com/mongodb/mongo-go-driver/mongo"
//    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    //  import models
//    "github.com/doxxitxxyoung/go-rest-tutorial/models"

)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage Endpoint hit")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

var client *mongo.Client

//const CollectionName = "Samples"


func main() {


    const AtlasUrl string = "mongodb+srv://doxxitxxyoung:Tele63741@cluster0-9usdz.gcp.mongodb.net"
    const DBName string = "glit_db_json"
    const CollectionName string = "Drugs"


    //  Atlas DB client
    fmt.Println("Starting the application...")

//    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //  skip errors ATM
//    client, _ = mongo.Connect(ctx, AtlasUrl)
//    client, _ = mongo.Connect(ctx,"mongodb+srv://doxxitxxyoung:Tele63741@cluster0-9usdz.gcp.mongodb.net")
//    client, _ = mongo.Connect(ctx,"mongodb://localhost:27017")

    //  configure connection url first
    /*
    client, err := mongo.NewClient(options.Client().ApplyURI(AtlasURL))
    if err != nil {
        log.Fatal(err)
    }
    */


//    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second) //  skip errors ATM

    //  https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial
    
    //  set Client options
    clientOptions := options.Client().ApplyURI(AtlasUrl)
    
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    //  Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }


    router := mux.NewRouter()
    http.ListenAndServe(":8080", router)

    drugsCollection := client.Database(DBName).Collection("Drugs")
    samplesCollection := client.Database(DBName).Collection("Samples")

    fmt.Println("Successfully connected to Atlas MongoDB")

    //handleRequests()
}
