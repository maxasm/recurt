package main

import (
    "github.com/joho/godotenv"
    "encoding/json"
    "fmt"
    "os"
    "context"
    "net/http"
    "io"
    "errors"
    "strings"

    "github.com/labstack/echo/v4"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
)

const SAMPLE_TEXT = ``

// database connection
var database_client *mongo.Client
const DATABASE_URL = "mongodb://127.0.0.1:27017/?connect=direct" 

type APIResponse struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Text string `json:"text"`
}

// helper function to send JSON response to the client
func sendResponse(c echo.Context, ar APIResponse) error {
    return c.JSON(ar.Code, ar) 
}

// helper function to connect to Mongo DB
func connect_to_database() {
    // create the client options
    client_opts := options.Client()
    client_opts.ApplyURI(DATABASE_URL)

    client, err_connect := mongo.Connect(context.TODO(), client_opts) 
    if err_connect != nil {
        fmt.Printf("Error connecting to database: %s\n", err_connect) 
        os.Exit(1)
    }
    
    fmt.Printf("\nconnected to database successfully ...\n")
    database_client = client
}

type TextEntry struct {
    Text string `bson:"text" json:"text"`
    Tokens int `bson:"tokens" json:"tokens"`
    Cost float64 `bson:"cost" json:"cost"`
    IP string `bson:"ip" json:"ip"`
}

func createRewriteEntry(ip string, text string) (string, error) {
    
    if database_client == nil {
        return "", errors.New("No connection to database")
    } 
    
    // insert the new text to the database
    collection := database_client.Database("text").Collection("data")     

    te := TextEntry{
        Text: text,
        IP: ip, 
    }

    insert_res, err_insert_text := collection.InsertOne(context.TODO(), te)
    
    if err_insert_text != nil {
        fmt.Printf("Error inserting text into database\n") 
        return "", err_insert_text 
    }
    
    mongo_insert_id := insert_res.InsertedID
    insert_id, ok := mongo_insert_id.(primitive.ObjectID) 

    if !ok {
        fmt.Printf("error converting types\n") 
        return "", errors.New("Error converting types")
    }

    return insert_id.Hex(), nil 
}

func handleRewriteRequest(c echo.Context) error {
    // get IP address of the user sending the request  
    req_ip := c.RealIP()
    
    // read the incomming text
    data, err_read_data := io.ReadAll(c.Request().Body)
    if err_read_data != nil {
        return sendResponse(c, APIResponse { Code: http.StatusBadRequest, Message: "Invalid request: no content in body", Text: ""}) 
    }
    
    // convert data to string 
    data_as_str := string(data)
    
    // check if data is empty
    if len(strings.Trim(data_as_str, " \t\r\n")) == 0 {
        return sendResponse(c, APIResponse{Code: http.StatusBadRequest, Message: "Invalid request: len(trim(content)) = 0", Text: ""}) 
    }

    req_id, err_create_entry := createRewriteEntry(req_ip, data_as_str)
    
    if err_create_entry != nil {
        return sendResponse(c, APIResponse{Code: http.StatusInternalServerError, Message: "Internal server error: failed to uplad text for rewriting", Text: ""}) 
    }
    
    return sendResponse(c, APIResponse{Code: http.StatusOK, Message: "Text uploaded successfully", Text: req_id})
}

func start_server() {
    e := echo.New() 
    
    // handle POST request submitting the text to be rewritten
    e.POST("/rewrite", handleRewriteRequest) 
    
    if err_start_server := e.Start(":8081"); err_start_server != nil {
        fmt.Printf("Error starting server: %s\n", err_start_server)
        os.Exit(1)
    }
}

// the main function
func main() {
	// load environment varibles and access tokens
	godotenv.Load("access_tokens")
    
    // connect to database
    connect_to_database()

    // start the server
    start_server()

	// run the main app
	rp := run(SAMPLE_TEXT)
    
    rp_json, err_rp_json := json.MarshalIndent(rp, "", "\t") 
    if err_rp_json != nil { 
        fmt.Printf("Error: %s\n", err_rp_json)    
        os.Exit(1)
    }
    
    fmt.Printf("\n ---- resp as JSON ----\n\n%s\n", rp_json)
}
