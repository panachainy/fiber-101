package main

import (
	"fiber-101/mockresponse"
	"fiber-101/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: 1, Name: "abc", Completed: false},
	{ID: 2, Name: "def", Completed: true},
}

func main() {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's an fiber.*Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Send custom error page
			err = ctx.Status(code).SendFile(fmt.Sprintf("./%d.html", code))
			if err != nil {
				// In case the SendFile fails
				return ctx.Status(500).SendString("Internal Server Error")
			}

			// Return from handler
			return nil
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/version", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// Example Exception

	app.Get("/1", func(c *fiber.Ctx) error {
		// Pass error to Fiber
		return c.SendFile("file-does-not-exist")
	})

	app.Get("/2", func(c *fiber.Ctx) error {
		panic("This panic is catched by fiber")
	})

	app.Get("/3", func(c *fiber.Ctx) error {
		// 503 Service Unavailable
		return fiber.ErrServiceUnavailable
	})

	app.Get("/4", func(c *fiber.Ctx) error {
		// 503 On vacation!
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	})

	app.Get("/mock-response", func(c *fiber.Ctx) error {
		return c.SendString(mockresponse.GetResponse())
	})

	app.Get("/arrays", func(c *fiber.Ctx) error {
		langs := [4]string{
			"golang",
			"typescript",
		}
		fmt.Printf("1) langs: %v \n", langs)
		fmt.Printf("1) langs: %#v \n", langs)

		// Get value specific array
		n := langs[1]
		fmt.Printf("2) langs: %#v \n", n)

		// Set
		langs[0] = "golang level 2"
		fmt.Printf("3) langs: %#v \n", langs)

		return c.SendString("see your console")
	})

	app.Get("/slices", func(c *fiber.Ctx) error {
		// init
		langs := []string{"test", "ja"}
		fmt.Printf("1) slices langs: %#v \n", langs)
		fmt.Printf("1) slices langs length: %#v \n", len(langs))

		// append
		langs = append(langs, "C#")
		fmt.Printf("2) slices langs: %#v \n", langs)
		fmt.Printf("2) slices langs length: %#v \n", len(langs))

		// append
		langs = append(langs, "F#", "Python", "Js")
		fmt.Printf("3) slices langs: %#v \n", langs)
		fmt.Printf("3) slices langs length: %#v \n", len(langs))

		return c.SendString("see your console")
	})

	app.Get("/slicing", func(c *fiber.Ctx) error {
		fmt.Printf("1 slice ============= \n")

		langs := []string{"golang", "java", "js"}
		newLangs1 := langs[0:2] // get just 0 1
		fmt.Printf("1) slicing langs: %#v \n", newLangs1)
		fmt.Printf("1) slicing langs length: %#v \n", langs)

		newLangs2 := langs[1:3] // get just 1 2
		fmt.Printf("2) slicing langs: %#v \n", newLangs2)
		fmt.Printf("2) slicing langs old value: %#v \n", langs)

		newLangs3 := langs[1:len(langs)] // get just 1 2
		fmt.Printf("3) slicing langs: %#v \n", newLangs3)
		fmt.Printf("3) slicing langs old value: %#v \n", langs)

		a := langs[0:3]
		fmt.Printf("a) slicing langs: %#v \n", a)
		b := langs[:3]
		fmt.Printf("b) slicing langs: %#v \n", b)

		c2 := langs[0:]
		fmt.Printf("c) slicing langs: %#v \n", c2)

		d := langs[:]
		fmt.Printf("d) slicing langs: %#v \n", d)

		fmt.Printf("2 array to slice ============= \n")

		services.PrintSlice(langs)

		names := [2]string{"dam", "ping"}

		// services.PrintSlice(names) // Can't do like this
		services.PrintSlice(names[:])

		fmt.Printf("3 zero slice ============= \n")

		var counties []string
		fmt.Printf("counties : %#v \n", counties)

		if counties == nil {
			fmt.Printf("counties is nil \n")
		} else {
			fmt.Printf("counties is NOT nil, value:%#v\n", counties)
		}

		counties2 := []string{"thai"}
		fmt.Printf("counties : %#v \n", counties2)

		if counties == nil {
			fmt.Printf("counties2 is nil \n")
		} else {
			fmt.Printf("counties2 is NOT nil, value:%#v\n", counties2)
		}
		return c.SendString("see your console")
	})

	app.Get("/loop", func(c *fiber.Ctx) error {
		// for
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

		fmt.Println("===============================")

		// while
		i := 0
		for i < 10 {
			fmt.Println(i)
			i++
		}

		fmt.Println("===============================")

		// // Loop forever
		// for {
		// 	fmt.Println(i)
		// 	i++
		// }

		fmt.Println("===============================")

		langs := []string{"golang", "python", "typescript"}

		for i := 0; i < len(langs); i++ {
			fmt.Println(i, ":", langs[i])
		}

		fmt.Println("===============================")

		for index, value := range langs {
			fmt.Println(index, ":", value)
		}

		fmt.Println("===============================")

		for _, value := range langs {
			fmt.Println("only value:", value)
		}
		return c.SendString("see your console")
	})

	app.Get("/maps", func(c *fiber.Ctx) error {
		status := map[int]string{
			200: "OK",
			400: "Bad Request",
		}

		fmt.Printf("% #v\n", status)

		l := len(status)
		fmt.Printf("length: %d\n", l)

		fmt.Println()

		status[200] = "Okie"
		fmt.Printf("% #v\n", status)

		fmt.Println()

		status[285] = "wtf"
		l = len(status)
		fmt.Printf("length: %d\n", l)
		fmt.Printf("% #v\n", status)

		fmt.Println()

		delete(status, 285)
		l = len(status)
		fmt.Printf("length: %d\n", l)
		fmt.Printf("% #v\n", status)

		fmt.Println()

		// check if existing
		v, ok := status[222]
		if ok {
			fmt.Printf("% #v\n", v)
		} else {
			fmt.Println("not found")
		}

		// check if existing -> short for
		if v2, ok2 := status[222]; ok2 {
			fmt.Printf("% #v\n", v2)
		} else {
			fmt.Println("not found")
		}

		// check map nil

		var m map[string]string
		if m == nil {
			fmt.Printf("% #v\n", m)
		} else {
			fmt.Println("not found")
		}

		// Allocate memory before use.

		x := map[string]string{}

		if x == nil {
			fmt.Printf("is nil \n")
		} else {
			fmt.Printf("% #v\n", x)
		}

		return c.SendString("see your console")
	})

	app.Get("/pointer", func(c *fiber.Ctx) error {
		return c.SendString("see your console")
	})

	app.Listen(":5000")
}
