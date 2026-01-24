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
To also remove the images and shared volume, add the flags `--rmi all` and `-v`.  

The application now runs on `localhost`. Any request sent to `localhost/health` should return `204 - No Content`. Localhost itself contains the portfolio.
