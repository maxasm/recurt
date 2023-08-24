package main

import (
    "github.com/labstack/echo/v4"
    "os"
    "fmt"
)

// file server for react single-page application
func serve_static_site(c echo.Context) error {
    // get the path of the file being requsted
    req_path := c.Request().RequestURI
    
    if req_path == "/" {
        req_path = "/index.html" 
    }

    base_dir := "./client/dist"
    full_path := base_dir + req_path
    
    // check if the file exists
    f, err := os.Open(full_path)
    defer f.Close()

    if err != nil && os.IsNotExist(err) {
        // the file does not exist, serve index.html to handle routing 
        full_path = base_dir + "/index.html" 
    }
    
    // serve the requsted file
    return c.File(full_path)
}
