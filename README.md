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

## This documentation
Tables are describing the directives. *Occurence* columns means how many times this must appear in a block of comments

## General information

### Directive descriptions

| Directive | occurence | Description|
|----------------------|-----|----------------|
| ``@APITitle {title string}``| 1 | API title |
| ``@APIVersion {version string}``| 0/1 | API version |
| ``@APIDescription {description string}`` | 0/1 | API description |
| ``@APIConsumes {content-type string}`` | 0/n | What consumes the API |
| ``@APIProduces {content-type string}`` | 0/n | What produces the API |
| ``@Contact {email string}`` | 0/1 | API contact |
| ``@TermsOfServiceUrl {url string}`` | 0/1 | URL or the terms of service |
| ``@License {licence string}`` | 0/1 | Type of licence |
| ``@LicenseUrl {url string}`` | 0/1 | URL of the licence |

### Example
Here is an example of the general information of your API.
```
// @APIVersion 1.0.0
// @APITitle Landru descriptor
// @APIDescription This is an example of the swagger descriptor
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

## Sub-routes
Sub-routes are declared like following (comments must be in the same block).

### Directive descriptions

| Directive | occurence | Description|
|----------------------|-----|----------------|
| ``@SubApi {title string} [{resource string}]``| 1 | Declare a sub-route |
| ``@SubApi {description string}``| 0/1 | Sub-route description |
| ``@Router {route string}``| 0/1 | Sub-route value; if not defined route will be the resource |
| ``@Resource {subRoute string}`` | 0/1 | Parent sub-route |

### Example

In this exemple, `/user` is an internal reference (@Resource)

```
// @SubApi Users [/users]
// @SubApi Allows you access to different features of the users , login , get status etc [/users]
```

## Routes

### Directive descriptions

| Directive | occurence | Description|
|----------------------|-----|----------------|
| ``@Title {title string}``| 1 | Route title |
| ``@Description {description string}`` | 1 | Route description |
| ``@Deprecated`` | 0/1 | Makes the route deprecated |
| ``@Accept {content-type string}`` | 0/n | What consumes the route |
| ``@Produces {content-type string}`` | 0/n | What produces the route |
| ``@Success {code int} {type string} {description string}`` | 0/n | Success return code |
| ``@Failure {code int} {type string} {description string}`` | 0/n | Failure return code |
| ``@Resource {subRoute string}`` | 1 | Sub-route |
| ``@Router {uri string} [{method string}]`` | 1 | Method and uri |

### Example

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
