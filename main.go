package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"tagprint/pkgs"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

var ctx = context.Background()

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		// gencode("sY2026-01-26")

		return c.SendString("http启动")
	})
	//waf禁用ip
	app.Get("/brcode", func(c *fiber.Ctx) error {
		ip := c.Query("code")
		// address := c.Query("addr")
		if ip != "" {
			gencode(ip)
		}
		return c.SendString("success")
	})
	// go subc()
	pkgs.Brcode()
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
		base := gencode(msg.Payload)
		// base := pkgs.Genprint()
		err = rdb.Publish(ctx, "golang", base).Err()
		pkgs.Err(err)
	}
}
func gencode(content string) string {
	// --- 参数设置 ---
	barcodeW, barcodeH := 285, 30
	padding := 10   // 四周留白高度
	textSpace := 35 // 底部文字区域高度
	charWidth := 7  // basicfont.Face7x13 的平均字符宽度
	// ----------------

	// 1. 生成条形码
	bc, _ := code128.Encode(content)
	scaledBC, _ := barcode.Scale(bc, barcodeW, barcodeH)

	// 2. 计算画布总尺寸
	// 总宽 = 条形码宽 + 左右边距
	// 总高 = 条形码高 + 文字空间 + 上下边距
	canvasW := barcodeW + (padding * 2)
	canvasH := barcodeH + textSpace + (padding * 2)
	canvas := image.NewRGBA(image.Rect(0, 0, canvasW, canvasH))

	// 背景填白
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// 3. 绘制条形码 (向下移动 padding 的距离，向右移动 padding 的距离)
	barcodeRect := image.Rect(padding, padding, padding+barcodeW, padding+barcodeH)
	draw.Draw(canvas, barcodeRect, scaledBC, image.Point{}, draw.Over)

	// 4. 绘制居中文字
	// 计算文字起始 X 坐标: (总宽 - 文字总长度) / 2
	textWidth := len(content) * charWidth
	startX := (canvasW - textWidth) / 2

	// Y 坐标: 条形码底部 + padding + 文字上偏移
	startY := padding + barcodeH + 25

	point := fixed.Point26_6{
		X: fixed.I(startX),
		Y: fixed.I(startY),
	}

	d := &font.Drawer{
		Dst:  canvas,
		Src:  image.NewUniform(color.Black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(content)

	// 5. 保存
	// file, _ := os.Create("barcode.png")
	// defer file.Close()
	// png.Encode(file, canvas)
	var buf bytes.Buffer

	// 2. 将图片以 PNG 格式写入缓冲区
	// 你也可以用 jpeg.Encode，但条形码建议用 PNG 保持清晰度
	err := png.Encode(&buf, canvas)
	if err != nil {
		pkgs.Err(err)
		// return ""
	}

	// 3. 将字节流转换为 Base64 字符串
	// StdEncoding 是标准编码，如果你要在网页显示，可以直接拼接前缀
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return imgBase64Str
}
