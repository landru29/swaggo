# Swaggo

Swaggo is a swagger definition generator written in go.
Just  add comments in your code.

All swaggo comments begin with `@`

For more information:
```
./swaggo --help
```

## Compilation
```
go get
go build
```

## General information
Here is an example of the general information of your API. These comments must be in the `main` package
```
// @APIVersion 1.0.0
// @APITitle Landru parser
// @APIDescription This is an example of the swagger parser
// @Contact cmeichel@free.fr
// @TermsOfServiceUrl https://www.teamwork.com/termsofservice
// @License BSD
// @LicenseUrl http://opensource.org/licenses/BSD-2-Clause
// @APIProduces application/json
// @APIProduces application/xml
// @APIConsumes application/json
```
Other mendatory parameters are passed through the command line:
```
./swaggo --api-basepath /v1 --api-host localhost:8080 --api-scheme http --ouptut my-swagger.json
```

## Routes
### Sub-routes
Sub-routes are declared like following (comments must be in the same block). In this exemple, `/user` is an internal reference (@Resource)
```
// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users , login , get status etc [/users]
```

### Routes
Routes are declared like this (Note that there is the @Resource `/user`):
```
// @Title Get Users Information
// @Description Get Users Information
// @Accept application/json
// @Param userId path int true "User ID"
// @Success 200 {object} string "Success"
// @Failure 401 {object} string "Access denied"
// @Failure 404 {object} string "Not Found"
// @Resource /users
// @Router /:userId.json [get]
```
