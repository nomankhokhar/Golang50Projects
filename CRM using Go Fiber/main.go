package main

import (
	"CrmGoFiber/db"
	"CrmGoFiber/lead"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)


func initDatabase(){
	var err error
	db.DBConn , err = gorm.Open("sqlite3", "leads.db")

	if err != nil {
		panic("failed to connect database")
	}
	db.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database is Migrated")
}

func main() {
	app := fiber.New()
	
	initDatabase()
	
	setupRoutes(app)
	app.Listen(3000)
	
	defer db.DBConn.Close()
}


func setupRoutes(app *fiber.App){
	app.Get("api/v1/leads" , lead.GetLeads)
	app.Get("api/v1/lead/:id" , lead.GetLead)
	app.Post("api/v1/lead" , lead.NewLead)
	app.Delete("api/v1/lead/:id" , lead.DeleteLead)
}