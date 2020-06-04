# Bootstrap

After the Discovery stage I moved on to the bootstrapping of the project.

Looking into sources of inspiration for how to correctly manage a Golang project at [awesome-go](https://github.com/avelino/awesome-go
), specially at the [Project Layout](https://github.com/golang-standards/project-layout
)

After this I moved to the *Walking Skeleton* part where I prepare the bare bones project to build, run tests and use a CI.

As the project is on Github I'm also attempting to use Github Actions as the CI service.

See [.github](.github) folder where I include a: 

### - Golang build and test stage 
 * build project
 * test project (no external dependencies)
 
### - docker-compose stage
 * Run Postgres and Vault as original
 * AccountAPI test service allows AccountAPI to wait for Postgres
 * A test service will run all tests as in the previous stage - plus end2end tests that require the AccountAPI service (as request in exercise)

Using [Wait-for-it.sh](https://github.com/vishnubob/wait-for-it/) to active wait for dependencies.

Existing the test service with error will send the error back as exit code.

### [golangci](https://github.com/golangci) stage

## Also

I've separately started a [Sonarqube](https://www.sonarqube.org/) service to analyse code quality in the project.

At the end of this stage I can `make docker` to run all tests.
As well as a Health Check request to the AccountAPI
