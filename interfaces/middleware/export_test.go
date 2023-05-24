package middleware

var Validate = (*AuthMiddleware).validate
var IsPublicPath = (*AuthMiddleware).isPublicPath
