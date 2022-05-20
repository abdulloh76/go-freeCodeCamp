package lead

import (
	"fibercrm/database"

	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(fiberCtx *fiber.Ctx) {
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	fiberCtx.JSON(leads)
}

func GetLead(fiberCtx *fiber.Ctx) {
	id := fiberCtx.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	fiberCtx.JSON(lead)
}

func NewLead(fiberCtx *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := fiberCtx.BodyParser(lead); err != nil {
		fiberCtx.Status(503).Send(err)
		return
	}
	db.Create(lead)
	fiberCtx.JSON(lead)
}

func DeleteLead(fiberCtx *fiber.Ctx) {
	id := fiberCtx.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		fiberCtx.Status(500).Send("no lead found with id")
		return
	}

	db.Delete(&lead)
	fiberCtx.Send("lead successfully deleted")

}
