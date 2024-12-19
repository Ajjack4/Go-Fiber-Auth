package main

import "github.com/gofiber/fiber/v2"

type SignupRequest struct {
	Name string
	Email    string
	Password string
}

func main() {
	app := fiber.New()

    app.Post("/signup", func(c *fiber.Ctx) error{
		req:=new(SignupRequest)
		if err:=c.BodyParser(req); err!=nil{
			return err
		}
		if req.Name=="" || req.Email==""|| req.Password==""{
			return fiber.NewError(fiber.ErrBadRequest.Code,"Invalid signup request")
		}
		return nil
	})
	app.Post("/login", func(c *fiber.Ctx) error{
		
		return nil
	})
	app.Get("/public", func(c *fiber.Ctx) error{
		return c.JSON(fiber.Map{"success": true,"path": "public"})
	})
	app.Post("/private", func(c *fiber.Ctx) error{
		return c.JSON(fiber.Map{"success": true,"path": "private"})
	})

	if err:=app.Listen(":3000");err!=nil{
		panic(err)
	}
}