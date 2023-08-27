# simple-HTTP-server

A simple HTTP/HTTPS fileserver written in GO, designed as a quick and simple way to share files.

By default, the app will share Any files/folder from a relative child folder named 'static' using an insecure HTTP connection hosted on port 8080.

### CLI Directives:
- use '**-p**' to change the default HTTP/HTTPS port from '8080'.  Note: root permissions will be required to open any common ports on Linux systems.
- use '**-dir**' to change the location of files to share.
- use '**-url**' to change the default '/' resource location to obfuscate the sharing of files to casual browsers.  Any alpha-numeric string is valid.
- use '**-secure** to enable SSL protection on the server.  This will require the user to supply a valid self-signed SSL key pair.  By default the app will look for appropriate files in relative 'certs' folder containing 'server.pem' and 'server.key' files.  Use the appropraite CLI commands to point to different locations for these files.

### Example1 - default:
``serve`` - will start an HTTP server on 0.0.0.0 and shares all file/folder in the folder 'static' on  port 8080.  Accessible locally as http://<local_IP>:8080

### Example2 - Secure Server
``serve`` -p=443 -url=testing -dir=/temp -secure -cert=cert.pem -key=mykey.key - will start an HTTPS service using SSL key/cert stored in the local directory, sharing the contents of the root directory '/temp' at the url "https://<local_IP>/testing/"




