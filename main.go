package main

import (
	"GoCrudChallange/initializers"
	"GoCrudChallange/routes"
)

func init() {
 initializers.LoadEnvVariables();
}

func main() {

routes.PersonRoutes();

}