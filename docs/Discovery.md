# Discovery

## On First run

Found out that we had a running vault instance, the root token is set in the docker-compose environment

### Vault
```
       | You may need to set the following environment variable:
vault_1       |
vault_1       |     $ export VAULT_ADDR='http://0.0.0.0:8200'
vault_1       |
vault_1       | The unseal key and root token are displayed below in case you want to
vault_1       | seal/unseal the Vault or re-authenticate.
vault_1       |
vault_1       | Unseal Key: ZlFDUiRRdmcF4sWK4UHsYIzx1tL70p7YTWV/5FYC4No=
vault_1       | Root Token: 8fb95528-57c6-422e-9722-d2147bcba8ed

```

### API

We are running a *fake* version of the Accounts API

Weirdly it tries to connect to a specific IP on up.

```
 | panic: dial tcp 172.18.0.2:5432: connect: connection refused
accountapi_1  |
```

It applies a db migration which is not mounted on a volume.

```
accountapi_1  | @timestamp="2020-05-21T10:48:33.763Z" level=info message="applied 1 database migrations!\n" service_name=interview_accountapi stack=
```

The API is running the [GIN framework](https://github.com/gin-gonic/gin)

```
accountapi_1  | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
   accountapi_1  |  - using env:	export GIN_MODE=release
   accountapi_1  |  - using code:	gin.SetMode(gin.ReleaseMode)
```

The API exposes multiple endpoints as requested to use in the client.

```
  | [GIN-debug] GET    /v1/health                --> github.com/form3tech/interview-accountapi/internal/app/interview-accountapi/api.HandleGetHealth (3 handlers)
accountapi_1  | [GIN-debug] GET    /v1/organisation/accounts/:id --> github.com/form3tech/interview-accountapi/internal/app/interview-accountapi/api.WithUserContext.func1 (4 handlers)
accountapi_1  | [GIN-debug] DELETE /v1/organisation/accounts/:id --> github.com/form3tech/interview-accountapi/internal/app/interview-accountapi/api.WithUserContext.func1 (4 handlers)
accountapi_1  | [GIN-debug] GET    /v1/organisation/accounts --> github.com/form3tech/interview-accountapi/internal/app/interview-accountapi/api.WithUserContext.func1 (4 handlers)
accountapi_1  | [GIN-debug] POST   /v1/organisation/accounts --> github.com/form3tech/interview-accountapi/internal/app/interview-accountapi/api.WithUserContext.func1 (4 handlers)
```

Also, includes a health check.

### Database

Postgres mounts a volume to run init scripts which includes `scripts/db/10-init.sql
` for the database and user creation.

## Poking the API

### List
Looking into the [documentation](https://api-docs.form3.tech/api.html#organisation-accounts) we find different ways to access the API which with little change allow us to use the fake API.

```
curl -X GET https://api.form3.tech//v1/organisation/accounts/ \
-H "Accept: application/vnd.api+json"
```

Documentation seems to have a trailing slash in all endpoints and that does not work with fake API returning a 301.

Correcting the request we end up with the empty Account List resource.

```
[GIN-debug] redirecting request 301: /v1/organisation/accounts/ --> /v1/organisation/accounts

<a href="/v1/organisation/accounts">Moved Permanently</a>.
< HTTP/1.1 301 Moved Permanently
< Content-Type: text/html; charset=utf-8
< Location: /v1/organisation/accounts

< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 21 May 2020 11:40:02 GMT
< Content-Length: 13
<
* Connection #0 to host localhost left intact
{"data":null}*
``` 

### Create

Create is also acessible based on documentation

```
curl -v -X POST -H "Content-Type: application/vnd.api+json" \
http://localhost:8080/v1/organisation/accounts \
-d \
'{
   "data": {
     "type": "accounts",
     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
     "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
     "attributes": {
       "country": "GB",
       "base_currency": "GBP",
       "bank_id": "400300",
       "bank_id_code": "GBDSC",
       "bic": "NWBKGB22"
     }
   }
 }'
```

List after Create

```
curl -v -X GET http://localhost:8080/v1/organisation/accounts \
-H "Accept: application/vnd.api+json"
Note: Unnecessary use of -X or --request, GET is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> GET /v1/organisation/accounts HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: application/vnd.api+json
>
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=utf-8
< Date: Thu, 21 May 2020 12:26:20 GMT
< Content-Length: 534
<
* Connection #0 to host localhost left intact
{"data":[{"attributes":{"alternative_bank_account_names":null,"bank_id":"400300","bank_id_code":"GBDSC","base_currency":"GBP","bic":"NWBKGB22","country":"GB"},"created_on":"2020-05-21T12:25:19.740Z","id":"ad27e265-9605-4b4b-a0e5-3003ea9cc4dc","modified_on":"2020-05-21T12:25:19.740Z","organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c","type":"accounts","version":0}],"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}* Closing connection 0

```

Attempt same Create

```
curl -v -X POST -H "Content-Type: application/vnd.api+json" \
http://localhost:8080/v1/organisation/accounts \
-d \
'{
   "data": {
     "type": "accounts",
     "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",                        <....
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /v1/organisation/accounts HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> Content-Type: application/vnd.api+json
> Content-Length: 330
>
* upload completely sent off: 330 out of 330 bytes
< HTTP/1.1 409 Conflict
< Content-Type: application/json; charset=utf-8
< Date: Thu, 21 May 2020 12:27:13 GMT
< Content-Length: 83
<
* Connection #0 to host localhost left intact
{"error_message":"Account cannot be created as it violates a duplicate constraint"}* Closing connection 0
```

Attempting the same Create twice returns the 409 Conflict for an existing resource as well as error message in body.

This gave me the error message structure.

### Database schema for API

There is no postgres migration to create a schema for the API so I looked into the container and found the following file being set up

```
 cat api/migrations/001_account.sql 
-- +migrate Up
CREATE TABLE IF NOT EXISTS "Account"
(
  id              UUID PRIMARY KEY NOT NULL,
  organisation_id UUID             NOT NULL,
  version         INT              NOT NULL,
  is_deleted      BOOLEAN          NOT NULL,
  is_locked       BOOLEAN          NOT NULL,
  created_on      TIMESTAMP,
  modified_on     TIMESTAMP,
  record          JSONB,
  pagination_id   SERIAL
);

CREATE UNIQUE INDEX ON "Account" (id);
CREATE UNIQUE INDEX Account_paginationid ON "Account" (pagination_id);

-- +migrate Down
DROP TABLE IF EXISTS "Account";
```

Logging into the database we can find a `gorp_migrations` table which belongs to the [sql-migrate](https://godoc.org/github.com/rubenv/sql-migrate
) package.

After running the Create request we can now see an example of an Account in the Database
```
check db entry 
ad27e265-9605-4b4b-a0e5-3003ea9cc4dc,
eb0bd6f5-c3f5-44b2-b677-acd23cdde73c,
0,
false,
false,
2020-05-21 12:25:19.740549,2020-05-21 12:25:19.740549,
"{""bic"": ""NWBKGB22"", ""bank_id"": ""400300"", ""country"": ""GB"", ""bank_id_code"": ""GBDSC"", ""base_currency"": ""GBP"", ""alternative_bank_account_names"": null}",
1
```