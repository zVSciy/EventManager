# Payment service

## Setup

### Certificates

Generating self signed certificates has never been this easy!

The following command will generate the necessary certificates

```sh
docker compose -f compose-generate-certs.yaml run --rm cert-gen
```

If you wish to run the development stage on a server, also generate a client certificate and import it into your browser

If not, you can ignore the following command

```sh
docker compose -f compose-generate-certs.yaml run --rm client-cert-gen # optional
```

### Reverse proxy config

Open `nginx.conf`

Modify `server_name` to your needs

> [!CAUTION]
> Remove the `ssl_client_certificate`, `ssl_verify_client` and `ssl_verify_depth` directives if you don't want to use a client certificate


### Environment variables

> [!TIP]
> Everything will work with the fallback values, so a .env file is not required for functionality

If you want to change something, rename `example.env` to `.env` and then you can modify the values.


## Running the containers

This project utilizes multi-stage builds.

The payment api image has 4 stages:
- build
- `development`
- tests
- `production`

Choose `development` for hot reload and development purposes

Choose `production` to have only the compiled binary in the image

> [!WARNING]
> Do `not` select `tests` or `build` stage in the compose file

> [!NOTE]
> The tests will automatically run when building production and cancel the build if not all tests pass.

#### If everything is ready to go, you can start the containers 

```sh
docker compose up -d --build
```

## How to use

You can try out the api under `/swagger`


## Testing

### Automatic

Tests are run automatically through GitHub Actions and when building the production build.

The GitHub workflow can be found in the repository root under `.github/workflows/go.yml`


### Manual

To manually run tests, you can use:

```sh
docker compose -f compose-manual-tests.yaml run --rm api-payment-tests
```
