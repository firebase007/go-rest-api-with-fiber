
package main 

import (
"github.com/gofiber/fiber" // import the fiber package

"log"

"github.com/firebase007/go-rest-api-with-fiber/database"

"github.com/firebase007/go-rest-api-with-fiber/router"


_ "github.com/lib/pq"

)


func main() { // entry point to our program

	// Connect to database
if err := database.Connect(); err != nil {
	log.Fatal(err)
}

app := fiber.New()  // call the New() method - used to instantiate a new Fiber App

router.SetupRoutes(app)

app.Listen(3000) // listen/Serve the new Fiber app on port 3000

}