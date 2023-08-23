package main

import (
    "github.com/joho/godotenv"
    "fmt"
    "os"
    "context"
    "net/http"
    "io"
    "errors"
    "strings"
    "encoding/json"

    "github.com/labstack/echo/v4"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    
    "golang.org/x/net/websocket" 
)

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

// update the database
func update_db(rp RewriteResponse, id primitive.ObjectID) error {
    
    collection := database_client.Database("text").Collection("data")
    _, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": bson.M{
        "tokens": rp.Tokens,
        "cost": rp.Cost, 
        "text": "",
    }})
     
    return err 
}

// handle incomming web socket connection 
func handleWebSocketRequest(c echo.Context) error {
    fmt.Printf("Received websocket connection request ...\n")
    // get the id of the request
    req_id := c.Param("id")

    // check if the ID exists in the database
    object_id, err_get_object_id := primitive.ObjectIDFromHex(req_id)    
    if err_get_object_id != nil {
        fmt.Printf("Error: %s\n", err_get_object_id)
        return c.NoContent(http.StatusBadRequest)
    }
    
    // find response
    var find_result TextEntry = TextEntry{}
    
    err_find := database_client.Database("text").Collection("data").FindOne(context.TODO(), bson.D{{"_id", object_id}}).Decode(&find_result)
    if err_find != nil {
        if err_find == mongo.ErrNoDocuments {
            fmt.Printf("Error: %s\n", err_find)
            return c.NoContent(http.StatusBadRequest)
        }  
        return c.NoContent(http.StatusBadRequest) 
    }

    websocket.Handler(func(ws *websocket.Conn){
        defer ws.Close() 
            
        // run the 'run' function which rewrites the text
        text := find_result.Text
        rp := run(text)
    
        err_update_db := update_db(rp, object_id) 
        if err_update_db != nil {
            fmt.Printf("error updating database\n")
        } else {
            fmt.Printf("updated database successfully\n") 
        }
 
        
        // convert rp to json
        json_b, err_json := json.Marshal(rp)
        if err_json != nil {
            fmt.Printf("Error converting to JSON string\n") 
        } else {     
            // send the json encoded response
            websocket.Message.Send(ws, string(json_b)) 
        }

    }).ServeHTTP(c.Response(), c.Request())
    
    return nil
}

func start_server() {
    e := echo.New() 

    e.Static("/", "./client")  

    // handle POST request submitting the text to be rewritten
    e.POST("/rewrite", handleRewriteRequest) 
    
    // get request to handle Websocket connections
    e.GET("/ws/:id", handleWebSocketRequest)    
    
    if err_start_server := e.Start(":8080"); err_start_server != nil {
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
}
