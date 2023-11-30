package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khairulharu/averincrud/domain"
	"github.com/khairulharu/averincrud/dto"
	"github.com/khairulharu/averincrud/internal/util"
)

type patientApi struct {
	patientService domain.PatientService
}

func NewPatient(patientService domain.PatientService, app *fiber.App) {
	h := patientApi{
		patientService: patientService,
	}

	app.Get("/patients", h.GetPatients)
	app.Get("/patient/find/", h.FindPatient)
	app.Post("/patient/new/", h.NewPatient)
	app.Post("/patient/update/", h.PatientUpdating)
	app.Delete("/patient/delete/", h.DeletingPatient)
}

func (p patientApi) GetPatients(ctx *fiber.Ctx) error {
	res := p.patientService.GetAll(ctx.Context())
	return ctx.Status(util.GetHttpStatus(res.Code)).JSON(res)
}

func (p patientApi) FindPatient(ctx *fiber.Ctx) error {
	queryName := ctx.Query("name")
	res := p.patientService.FindPatient(ctx.Context(), dto.ReqPatient{
		Name: queryName,
	})
	return ctx.Status(util.GetHttpStatus(res.Code)).JSON(res)
}

func (p patientApi) NewPatient(ctx *fiber.Ctx) error {
	var reqPatient dto.ReqPatient
	if err := ctx.BodyParser(&reqPatient); err != nil {
		return ctx.SendStatus(400)
	}
	res := p.patientService.StorePatient(ctx.Context(), reqPatient)
	return ctx.Status(util.GetHttpStatus(res.Code)).JSON(res)
}

func (p patientApi) PatientUpdating(ctx *fiber.Ctx) error {
	queryName := ctx.Query("name")
	var reqPatient dto.ReqPatient
	if err := ctx.BodyParser(&reqPatient); err != nil {
		return ctx.SendStatus(400)
	}
	res := p.patientService.UpdatePatient(ctx.Context(), queryName, reqPatient)
	return ctx.Status(util.GetHttpStatus(res.Code)).JSON(res)
}

func (p patientApi) DeletingPatient(ctx *fiber.Ctx) error {
	queryName := ctx.Query("name")
	res := p.patientService.DeletePatient(ctx.Context(), dto.ReqPatient{
		Name: queryName,
	})
	return ctx.Status(util.GetHttpStatus(res.Code)).JSON(res)
}
