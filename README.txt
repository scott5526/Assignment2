timeserver.go README

Resources used
---------------------------------------------------------------------------------------------
http://golang.org/pkg/os/
http://golangtutorials.blogspot.com/2011/06/web-programming-with-go-first-web-hello.html
http://stackoverflow.com/questions/10105935/how-to-convert-a-int-value-to-string-in-go
http://golang.org/pkg/net/http/
http://golang.org/pkg/time/
http://grokbase.com/t/gg/golang-nuts/134kenh4xz/go-nuts-time-format-giving-unpredictable-results
---------------------------------------------------------------------------------------------

Running the timeserver.go file
---------------------------------------------------------------------------------------------
To run timeserver.go, open the Windows command prompt and move to the directory of timeserver.go.  To run the file, use "go run timeserver.go" with any applicable flags.

Applicable flags include:

-V ("go run timeserver.go -V)

Runs timeserver.go with the version flag enabled.  Will output the current version of the file and terminate the program with a zero error code.

-port # ("go run timeserver.go -port 9999)

Runs timeserver.go with a specified port (the default port # is 8080).
---------------------------------------------------------------------------------------------


Accessing the server from a web browser
---------------------------------------------------------------------------------------------
Enter the URL http://localhost:(port #)/time to access the timeserver.  Any URL beyond (port#)/ that doesn't match this specified URL will result in a 404 not found web page.
---------------------------------------------------------------------------------------------



Caveats
---------------------------------------------------------------------------------------------
When trying to run the server, if the specified port is already in use the program will terminate with a error message on a non-zero error code.

Any URL beyond http://localhost:(port #) that doesn't match the above specified URL will result in a 404 not found web page.
---------------------------------------------------------------------------------------------