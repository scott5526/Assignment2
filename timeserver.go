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

var currUser string
var portNO *int
var cookieMap = make(map[string]http.Cookie)

/*
Greeting message
*/
func greetingHandler(w http.ResponseWriter, r *http.Request) {
    //Check for existing cookie here <------------------------------
    //Redirect to login form if necessary here <--------------------

    fmt.Fprintf(w, "Greetings, ")
}

/*
Login handler.  Displays a html generated login form for the user to provide a name.  Creates a cookie for the user name and redirects them to the home page if a valid user name was provided.  If no valid user name was provided, outputs an error message
*/
func loginHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html>" +
    "<body>" +
    "<form method=post action=login>" +
  	"What is your name, Earthling?" +
    "<input type=text name=name size=50>" +
    "<input type=submit name=submit>" +
    "</form>" +
    "</p>" +
    "</body>" +
    "</html>")
    r.ParseForm()
    name := r.PostFormValue("name")
    submit := r.PostFormValue("submit") 
    if submit == "Submit" {
    	if name == "" {
    		fmt.Fprintf(w, "C'mon, I need a name.")
    	} else {
		//Add new cookie here <---------------------------

		fmt.Fprintf(w, "<html>" +
		"<head>" +
        	"<META http-equiv=refresh content=0;URL=http://localhost:" + strconv.Itoa(*portNO) + "/index.html>" +
    		"</head>")
    	}
    }
}

/*
Logout handler.  Clears user cookie, displays goodbye message for 10 seconds, then redirects user to login form
*/
func logoutHandler(w http.ResponseWriter, req *http.Request) {
    //clear cookie here <--------------------------------
    fmt.Fprintf(w, "<html>" +
    "<head>" +
    "<META http-equiv=refresh content=10;URL=http://localhost:" + strconv.Itoa(*portNO) + "/index.html>"+
    "<body>" +
    "<p>Good-bye.</p>" +
    "</body>" +
    "</html>")
}


/*
Handler for time requests.  Outputs the current time in the format:

Hour:Minute:Second PM/AM
*/
func timeHandler(w http.ResponseWriter, r *http.Request) {
    currTime := time.Now().Format("03:04:05 PM")
    utcTime := time.Now().UTC()
    utcTime = time.Date(
        time.Now().UTC().Year(),
        time.Now().UTC().Month(),
        time.Now().UTC().Day(),
        time.Now().UTC().Hour(),
        time.Now().UTC().Minute(),
        time.Now().UTC().Second(),
        time.Now().UTC().Nanosecond(),
        time.UTC,
    )
    utcTime.UTC()
    //utcTime.Format("03:04:05 07")

    fmt.Fprintf(w, "<html>" +
    "<head>" +
    "<style>p" +
    "{font-size: xx-large} p2 {color: red}" +
    "</style>" +
    "</head>" +
    "<body>" +
    "<p>The time is now <p2>" +
    currTime +
    "</p2><p3>  (" +
    utcTime.Format("03:04:05") + 
    " UTC), " +
    //name goes here <-------------------------------
    "</p3></p>" +
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
    version := flag.Bool("V", false, "Version 2.3") //Create a bool flag for version  
    						    //and default to no false

    portNO = flag.Int("port", 8080, "")	    //Create a int flag for port selection
					            //and default to port 8080
    flag.Parse()

    if *version == true {		//If version outputting selected, output version and 
        fmt.Println("Version 1.0")	//terminate program with 0 error code
        os.Exit(0)
    }

    // URL handling
    //http.HandleFunc("/", greetingHandler)
    http.HandleFunc("/index.html", greetingHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/time", timeHandler)
    
    //Check host:(specified port #) for incomming connections
    error := http.ListenAndServe("localhost:" + strconv.Itoa(*portNO), nil)

    if error != nil {				// If the specified port is already in use, 
	fmt.Println("Port already in use")	// output a error message and exit with a 
	os.Exit(1)				// non-zero error code
    }
}
