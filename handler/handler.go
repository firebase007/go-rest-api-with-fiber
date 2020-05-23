package handler

import (
	"log"
	"database/sql"

	"github.com/gofiber/fiber"

	"github.com/firebase007/go-rest-api-with-fiber/model"

	"github.com/firebase007/go-rest-api-with-fiber/database"
)


// GetAllProducts from db
func GetAllProducts(c *fiber.Ctx) {

	// query product table in the database
	rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error": err,
		  })
		return
	}

	defer rows.Close()

	result := model.Products{}

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
		// Exit if we get an error
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err,
			  })
			return
		}

		// Append Product to Products
		result.Products = append(result.Products, product)
	}

	// Return Products in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"product":  result,
		"message": "All product returned successfully",
	  }); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		  })
		return
	}
}


// GetSingleProduct from db
func GetSingleProduct(c *fiber.Ctx) {

	id := c.Params("id")
	product := model.Product{}

	// query product database
	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		  })
		return
	}

	defer row.Close()

	// iterate through the values of the row
	for row.Next() {
	switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category ); err {
		case sql.ErrNoRows:
			  log.Println("No rows were returned!")
			  c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			  })
		case nil:
  			log.Println(product.Name, product.Description, product.Category, product.Amount)
		default:
			//   panic(err)
			  c.Status(500).JSON(&fiber.Map{
				"success": false,
				"message": err,
			  })
	}

}
	
	// return product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched product",
		"product": product,
	  }); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message":  err,
		  })
		return
	}

	
}

// CreateProduct handler
func CreateProduct(c *fiber.Ctx) {

	// Instantiate new Product struct
	p := new(model.Product)

	//  Parse body into product struct
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		  })
		return
	}

	// Insert Product into database
	res, err := database.DB.Query("INSERT INTO products (name, description, category, amount) VALUES ($1, $2, $3, $4)" , p.Name, p.Description, p.Category, p.Amount )
	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		  })
		return
	}
	// Print result
	log.Println(res)
	
	// Return Product in JSON format
	if err := c.JSON(&fiber.Map{
		"success": true,
		"message": "Product successfully created",
		"product": p,
	  }); err != nil {
		c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message":  "Error creating product",
		  })
		return
	}
}


// DeleteProduct from db 
func DeleteProduct(c *fiber.Ctx) {

		id := c.Params("id")

		// query product table in database
		res, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
		if err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err,
			  })
			return
		}

		// Print result
		log.Println(res)

		// return product in JSON format
		if err := c.JSON(&fiber.Map{
			"success": true,
			"message": "product deleted successfully",
		  }); err != nil {
			c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error": err,
			  })
			return
		}
}