// Hexsatisfaction microservice
//
// The purpose of this application to provide users basic functionality
//
//    Schemes: http
//    Host: localhost:8000
//    Version: 0.0.1
//
//    Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import "github.com/JesusG2000/hexsatisfaction/internal/app"

const configPath = "config/main"

func main() {
	app.Run(configPath)
}
