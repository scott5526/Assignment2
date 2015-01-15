/*
File: timeserver.go
Author: Robinson Thompson

Description: Runs a simple timeserver to pull up a URL page displaying the current time.  		     Support was verified for Windows 7 OS.  Support has not been tested for other OS

Copyright:  All code was written originally by Robinson Thompson with assistance from various
	    free online resourses.
*/
package main

import (
"flag"
"fmt"
"net/http"
"os"
"strconv"
"time"
)

/*
Greeting message
*/
func greetingHandler(w http.ResponseWriter, req *http.Request) {
}

/*
Login handler.  Displays a html generated login form for the user to provide a name.  Creates a cookie for the user name and redirects them to the home page if a valid user name was provided.  If no valid user name was provided, outputs an error message
*/
func loginHandler(w http.ResponseWriter, req *http.Request) {
}


/*
Logout handler.  Clears user cookie, displays goodbye message for 10 seconds, then redirects user to login form
*/
func logoutHandler(w http.ResponseWriter, req *http.Request) {
}


/*
Handler for time requests.  Outputs the current time in the format:

Hour:Minute:Second PM/AM
*/
func timeHandler(w http.ResponseWriter, r *http.Request) {
    currTime := time.Now().Format("03:04:05 PM")

    fmt.Fprintf(w, "<html>" +
    "<head>" +
    "<style>p" +
    "{font-size: xx-large} p2 {color: red}" +
    "</style>" +
    "</head>" +
    "<body>" +
    "<p>The time is now <p2>" +
    currTime +
    "</p2>.</p>" +
    "</body>" +
    "</html>")
}

/*
Handler for invalid requests.  Outputs a 404 error message and a cheeky message
*/
func badHandler(w http.ResponseWriter, req *http.Request) {
    http.NotFound(w, req)
    w.Write([]byte("These are not the URLs you're looking for."))
}

/*
Main
*/
func main() {
    //Version output & port selection
    version := flag.Bool("V", false, "Version 2.0") //Create a bool flag for version  
    						    //and default to no false

    portNO := flag.Int("port", 8080, "")	    //Create a int flag for port selection
					            //and default to port 8080
    flag.Parse()

    if *version == true {		//If version outputting selected, output version and 
        fmt.Println("Version 1.0")	//terminate program with 0 error code
        os.Exit(0)
    }

    // URL handling
    http.HandleFunc("/", greetingHandler)
    http.HandleFunc("/index.html", greetingHandler)
    http.HandleFunc("/login?name=", loginHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/time", timeHandler)
    
    //Check localhose:(specified port #) for incomming connections
    error := http.ListenAndServe("host:" + strconv.Itoa(*portNO), nil)

    if error != nil {				// If the specified port is already in use, 
	fmt.Println("Port already in use")	// output a error message and exit with a 
	os.Exit(1)				// non-zero error code
    }
}
