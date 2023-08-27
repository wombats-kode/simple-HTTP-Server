package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Set CLI variables to hold config data
	port := flag.String("p", "8080", "port to serve files on")
	directory := flag.String("dir", "static", "the directory of static files to host")
	url := flag.String("url", "/", "string used to obfuscate default url resource")

	// Add 'secure' option to permit support for self-signed SSL protection
	secure := flag.Bool("secure", false, "enable/disable SSL protection")
	cert := flag.String("cert", "certs/server.pem", "location of SSL public cert")
	key := flag.String("key", "certs/server.key", "location of the SSL private key")

	flag.Usage = func() {
		fmt.Println("Serve is a simple HTTP/HTTPS file server used to serve files from a local folder ('static' by default).")
		fmt.Println("Use the 'secure' flag to enable SSL - requires your own OpenSSL cert/key files.")
		fmt.Println("To obfuscate the root url, specific a string to prevent accidental downloading by casual users.")
		fmt.Println()
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.Parse()

	// Validate that CLI flags make sense
	// Port numbers have to be positive integers between 0-65535
	myPort, err := strconv.Atoi(*port)
	if err != nil {
		log.Fatal("port has to be an integer value.")
	} else if myPort < 1 || myPort > 65535 {
		log.Fatal("port has to be a positive integer less that 65535")
	}

	// Directory has to exist and be accessible
	if _, err := os.Stat(*directory); os.IsNotExist(err) {
		log.Fatal("folder '", *directory, "' does not exist")
	}

	// URI location has to start and end with '/' to be correctly parsed.
	if *url != "/" {
		*url = strings.TrimPrefix(*url, "/")
		*url = strings.TrimSuffix(*url, "/")
		isAlphaNumeric := regexp.MustCompile(`^[a-zA-Z0-9]*$`).MatchString(*url)
		if isAlphaNumeric {
			// Add leading and trailing '/' for correct parsing of file location
			*url = fmt.Sprintf("/%s/", *url)
		} else {
			log.Fatal("url is limited to alpha-numeric characters only")
		}
	}

	http.Handle(*url, http.StripPrefix(*url, http.FileServer(http.Dir(*directory))))

	if *secure {
		log.Printf("Starting HTTPS server for folder '%s' on port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServeTLS(":"+*port, *cert, *key, nil))

	} else {
		log.Printf("Starting HTTP server for folder '%s' on port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))

	}
}
