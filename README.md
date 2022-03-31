# client-sever-multithreading-downloader
## Server
The server was build using nodejs.
The server will run on port 3000.
> npm install
Then
> node server.js

## Client
Client built using go lang.
After running the server, the following command download the file using 20 threads.
> go run main.go -url http://localhost:3000/Theartofmulticore.pdf -t 20 -checksum 15a8243a75507e9aaf6d532e8131244c75aa29c3f17488680c6fd7c2ab9e30f1.
### command line arguments
-