package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/iamrafaelsoares/go-admin/database"
	"github.com/iamrafaelsoares/go-admin/routes"
)

func main() {

	database.Connect()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)
	app.Listen(":3000")

}

// func populateDb(db *gorm.DB) {
// 	// fmt.Println("start populate")
// 	// faker := faker.New()

// 	// for i := 0; i <= 5; i++ {
// 	// 	person := faker.Person()
// 	// 	order := models.Order{
// 	// 		FirstName: person.FirstName(),
// 	// 		LastName:  person.LastName(),
// 	// 		Email:     person.Contact().Email,
// 	// 		UpdatedAt: "01/12/2020",
// 	// 		CreatedAt: "01/12/2020",
// 	// 	}

// 	// 	db.Create(&order)
// 	// 	fmt.Println(order)
// 	// }

// 	// for i := 0; i <= 36; i++ {
// 	// 	fk := faker.New()
// 	// 	orderItem := models.OrderItem{
// 	// 		OrderId:      uint(i),
// 	// 		ProductTitle: fk.Food().Vegetable(),
// 	// 		Price:        fk.Generator.Float32(),
// 	// 		Quantity:     1,
// 	// 	}

// 	// 	db.Create(&orderItem)
// 	}
