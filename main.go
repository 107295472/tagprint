package main

import (
	"context"
	"fmt"
	"tagprint/pkgs"

	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		// gencode("sY2026-01-26")
		// pkgs.Genprint()
		return c.SendString("http启动")
	})
	app.Get("/brcode", func(c *fiber.Ctx) error {
		// ip := c.Query("code")
		// // address := c.Query("addr")
		// if ip != "" {
		// 	gencode(ip)
		// }
		return c.SendString("success")
	})
	go subc()

	app.Listen(fmt.Sprintf(":%d", *pkgs.ServerConfig.HttpPort))
}
func subc() {
	rdb := pkgs.GetRedis()
	pubsub := rdb.Subscribe(ctx, "gocode")
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		// println(msg.Payload)
		base := pkgs.Gencode(msg.Payload)
		// base := pkgs.Brcode(msg.Payload)
		err = rdb.Publish(ctx, "golang", base).Err()
		pkgs.Err(err)
	}
}
