# wallet-api


#### Full list of tools and libraries used in this project
* [Gin](https://github.com/labstack/echo) - Web framework
* [viper](https://github.com/spf13/viper) - Go configuration library
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
* [logrus](https://github.com/sirupsen/logrus) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [jwt-go](https://github.com/dgrijalva/jwt-go) - JSON Web Tokens (JWT)
* [uuid](https://github.com/google/uuid) - UUID
* [bluemonday](https://github.com/microcosm-cc/bluemonday) - HTML sanitizer
* [swag](https://github.com/swaggo/swag) - Swagger
* [gomock](https://github.com/golang/mock) - Mocking framework
* [Docker](https://www.docker.com/) - Containerization platform


#### Running in Docker
    make develop // run all containers in development mode
    make prod // Run all containers in production mode

#### Generating mocks
    make mock // Generate mocks for all services

#### Running unit tests
    make test // Run all unit tests


