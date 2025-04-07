package apiV1

import "github.com/ChocolateAceCream/telescope/backend/service"

var authService = new(service.AuthService)
var sseService = new(service.SSEService)
var awsService = new(service.AwsService)
var localeService = new(service.LocaleService)
var sketchService = new(service.SketchService)
var projectService = new(service.ProjectService)
