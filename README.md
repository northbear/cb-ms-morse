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

```docker run morse-responer --name morse-responder-test -p 5000:4000 --env APP_PORT=4000```

## Client Application - Morse Receiver


