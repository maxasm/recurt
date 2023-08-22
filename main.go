package main

import (
    "github.com/joho/godotenv"
    "encoding/json"
    "fmt"
    "os"
)

const SAMPLE_TEXT = `The landscape of American politics, originating from the visionary reflections of the Founding Fathers, has experienced metamorphoses throughout the duration of centuries, culminating in a multifaceted tapestry of ideologies, interests, and institutions. As the vanguard democratic republic across the globe, the United States has observed notable modifications in its political sphere, encompassing struggles pertaining to slavery and civil rights, the ideological confrontations of the Cold War, and the divided discussions of the 21st century. These shifts not only reflect the nation's evolving priorities and values but also assume a pivotal role in shaping global political paradigms. As we delve into the intricate realm of U.S. politics, it becomes imperative to apprehend the historical, cultural, and societal components that have exerted influence on its trajectory.`

// the main function
func main() {
	// load environment varibles and access tokens
	godotenv.Load("access_tokens")
	
	// run the main app
	rp := run(SAMPLE_TEXT)
    
    rp_json, err_rp_json := json.MarshalIndent(rp, "", "\t") 
    if err_rp_json != nil { 
        fmt.Printf("Error: %s\n", err_rp_json)    
        os.Exit(1)
    }
    
    fmt.Printf("\n ---- resp as JSON ----\n\n%s\n", rp_json)
}
