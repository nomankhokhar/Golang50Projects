package lead

import (
	"CrmGoFiber/db"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name  		string
	Company 	string
	Email 		string
	Phone 		int
}
func GetLeads(c *fiber.Ctx){
	
}

func GetLead(c *fiber.Ctx){
 	id := c.Params("id")
	db := db.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx){}

func DeleteLead(c *fiber.Ctx){}