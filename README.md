# berke.in

Git-repository for the [berke.in](https://berke.in) website.

## Running the application
Build and run the docker containers with the following command:
```
docker compose up
```

## Stop the application
Stop the containers with the following command:
```
docker compose down
```
If you want to remove the images as well, add the flag `--rmi all`.  

The application now runs on localhost. Any request sent to `localhost/health` should return `204 - No Content`. Localhost itself contains my portfolio.
