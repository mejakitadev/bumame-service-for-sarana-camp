package config

import (
	"sarana-dafa-ai-service/controller"

	"github.com/gofiber/fiber/v2"
)

func BumameAuthRouter(app *fiber.App, cont controller.BumameAuthController) {
	app.Post("/auth/login", cont.Login)
	app.Get("/auth/read-token", cont.ReadToken)
}

func BumameB2BProductRouter(app *fiber.App, cont controller.BumameB2BProductController) {
	app.Put("/b2b-product/bulk-update", cont.BulkUpdate)

	// Move the specific route before the parameterized route
	app.Get("/b2b-product/name/:b2b_product_name", cont.FindByName)
	app.Get("/b2b-product", cont.FindAll)
	app.Get("/b2b-product/:b2b_product_id", cont.FindById)
	app.Post("/b2b-product", cont.Create)
	app.Put("/b2b-product/:b2b_product_id", cont.Update)
	app.Delete("/b2b-product/:b2b_product_id", cont.Delete)

	app.Post("/b2b-product/generate-slugs", cont.GenerateSlugs)
}

func BumameB2BPasienRouter(app *fiber.App, cont controller.BumameB2BPasienController) {
	app.Get("/b2b-pasien/name/:b2b_pasien_name", cont.FindByName)
	app.Get("/b2b-pasien", cont.FindAll)
	app.Get("/b2b-pasien/:b2b_pasien_id", cont.FindById)
	app.Post("/b2b-pasien", cont.Create) // Already using Create
	app.Put("/b2b-pasien/:b2b_pasien_id", cont.Update)
	app.Delete("/b2b-pasien/:b2b_pasien_id", cont.Delete)
}

func BumameB2BDokterRouter(app *fiber.App, cont controller.BumameB2BDokterController) {
	app.Get("/b2b-dokter/name/:b2b_dokter_name", cont.FindByName)
	app.Get("/b2b-dokter", cont.FindAll)
	app.Get("/b2b-dokter/:b2b_dokter_id", cont.FindById)
	app.Post("/b2b-dokter", cont.Create) // Already using Create
	app.Put("/b2b-dokter/:b2b_dokter_id", cont.Update)
	app.Delete("/b2b-dokter/:b2b_dokter_id", cont.Delete)
}
