# berke.in

Git-repository for the [berke.in](https://berke.in) website.

## Building the application
Build the docker image with the following command:
```
docker build -t gosite .
```

## Running the application
Run the image with the following command:
```
docker run -it --rm -p 80:80 gosite
```
The application now runs on localhost. A request sent to `localhost/health` should return `204 - No Content`.
