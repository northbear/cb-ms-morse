# cb-ms-morse

## Server Applcation - Morse Responder

The Server Application translates IP address for client to Morse code representation and sends it as response.
By default, it binds to port 5000 for serving requests. If another port number is reqired it should be provided by environment variable APP_PORT
with port number as value. Example:
```APP_PORT=4000 ./morse-responder```

The application is supposed to be run inside docker container. The dockerfile for building image is `Dockerfile.morse-responder`.
To build proper image use command:
```docker build -t morse-responder -f Dockerfile.morse-responder .```

To run container with given the application port number use shell command:
```docker run morse-responder --name morse-responder-test -p 5000:4000 --env APP_PORT=4000```

## Client Application - Morse Receiver

The Client Application sends requests to the server application and print the response to output. Server application ip address and port should be provided 
with environment variable `SERVICE\_ADDRESS`. 
```SERVICE_ADDRESS=myhost.mydomain.com:5000 ./morse-receiver```

If the environment variable is not defined the client application tries to connect at address 127.0.0.1:5000. 
The client application is supposed to be run inside docker container. The dockerfile for building image is `Dockerfile.morse-receiver`.
To build proper image use command:

```docker build -t morse-receiver -f Dockerfile.morse-receiver .```

To run container with given the application port number use shell command:

```docker run morse-receiver --name morse-receiver-test --env SERVICE_ADDRESS=myhost.mydomain.com:4000```


