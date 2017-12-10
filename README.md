# Document Validator [![Build Status](https://travis-ci.org/aymone/doc-validator.svg?branch=master)](https://travis-ci.org/aymone/doc-validator)

## Getting Started

Install project (First run)
```bash
$ make install
```

Up containers and run project
```bash
$ make up
```

Stop containers
```bash
$ make stop
```

Watch logs from docker containers
```bash
$ make logs
```

Update dependencies on containers
```bash
$ make deps
```

Running unit tests
```bash
$ make test
```

Watching client unit tests
```bash
$ make client-watch
```

Running client unit tests
```bash
$ make client-test
```

Running backend unit tests
```bash
$ make test
```

Running backend and client tests
```bash
$ make full-test
```

## Directory Layout
```
src/                    --> All of the source files for the backend application
www/                    --> All of the source files for the client application
    index.html          --> HTML main file
    app/
        app.js/         --> Angular main file
        app.css/        --> Style file
        services/       --> Angular service files
        components/     --> Angular directive files
```

## References

### Backend
The backend was written in golang, theres a docker container ready to run it.
Gin go-wrapper are used to auto-build go application inside container when you change the code (auto reload).

For routing layer was used [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter) and logs by [logrus].
(https://github.com/sirupsen/logrus)

An instance of mongoDB are created by docker compose and connected through http://mongodb:27017.

There's too an nginx container to serve web on http://localhost:8080.

### Client

The SPA runs on nginx container mentioned above.

We have two kinds of dependencies in this project: tools and Angular framework code. The tools help
us manage and test the application.

The client was written with angularJs from [angular-seed](https://github.com/angular/angular-seed) project and codestyle was adapted to [johnPapa styleguide](https://github.com/johnpapa/angular-styleguide/blob/master/a1/README.md). Node's npm installs only tools like bower, karma and jasmine for tests.
Bower has the responsibility about angular application dependencies.

We have preconfigured `npm` to automatically run `bower` so we can simply do:
```
npm --prefix ./www i
```

But dont worry with it, just use the make instructions to install and build all project.

### Continuous Integration
[Travis CI](https://travis-ci.org/aymone/doc-validator) are configured to automated test builds.

## Prerequisites

[Docker](https://www.docker.com/)

[Docker compose](https://docs.docker.com/compose/install/)

[NodeJs](https://nodejs.org/en/)
