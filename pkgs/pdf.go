package pkgs

import (
	"log"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/repository"
)

func Genprint() string {
	customFont := "微软雅黑"
	customFontFile := "E:/fonts/fonts/msyh.ttf"
	customFonts, err := repository.New().
		AddUTF8Font(customFont, fontstyle.Normal, customFontFile).
		AddUTF8Font(customFont, fontstyle.Italic, customFontFile).
		AddUTF8Font(customFont, fontstyle.Bold, customFontFile).
		AddUTF8Font(customFont, fontstyle.BoldItalic, customFontFile).
		Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	builder := config.NewBuilder().
		WithDimensions(241, 93). // 设置页面宽高
		WithTopMargin(2).
		WithBottomMargin(2).
		WithLeftMargin(10).
		WithRightMargin(10).
		WithCustomFonts(customFonts)

	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).
		Build()

	m := maroto.New(cfg)

	// 2. 定义样式
	cellStyle := &props.Cell{
		BorderType:  border.Full,
		BorderColor: &props.Color{Red: 0, Green: 0, Blue: 0},
	}
	labelStyle := props.Text{Size: 12, Align: align.Center}
	contentStyle := props.Text{Size: 12, Align: align.Center}

	// 3. 标题与单号
	m.AddRows(
		row.New(6).Add(col.New(12).Add(text.New("辽宁盛源生物质燃料有限公司", props.Text{Align: align.Center, Size: 18}))),
		row.New(8).Add(col.New(12).Add(text.New("过 磅 单", props.Text{Align: align.Center, Size: 18, Style: fontstyle.Bold}))),
		row.New(6).Add(
			col.New(4).Add(text.New("供方：", props.Text{Size: 12})),
			col.New(5).Add(text.New("2026/1/28 15:38:25", props.Text{Align: align.Center, Size: 12})),
			col.New(3).Add(text.New("NO: 4955431651", props.Text{Align: align.Right, Size: 12})),
		),
	)

	// 4. 表格主体
	// 第一行：燃料 + 车牌
	m.AddRows(row.New(10).Add(
		col.New(1).Add(text.New("燃 料", labelStyle)).WithStyle(cellStyle),
		col.New(5).Add(text.New("玉米秸秆", contentStyle)).WithStyle(cellStyle),
		col.New(3).Add(text.New("车牌：", labelStyle)).WithStyle(cellStyle),
		col.New(3).Add(text.New("辽A·88888", contentStyle)).WithStyle(cellStyle),
	))

	// 第二行：毛重 + 过磅时间 + 霉腐 + 杂质
	m.AddRows(row.New(10).Add(
		col.New(2).Add(text.New("毛重 (吨)", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("35.50", contentStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("过磅时间", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("15:30", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("霉腐(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("1.0", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("杂质(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("0.5", contentStyle)).WithStyle(cellStyle),
	))

	// 第三行：皮重 + 过磅时间 + 水分 + 包长
	m.AddRows(row.New(10).Add(
		col.New(2).Add(text.New("皮重 (吨)", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("15.20", contentStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("过磅时间", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("15:50", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("水分(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("12.0", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("包长(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("52", contentStyle)).WithStyle(cellStyle),
	))

	// 第四行：净重 + 准重 + 其他 + 包重
	m.AddRows(row.New(10).Add(
		col.New(2).Add(text.New("净重 (吨)", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("20.30", contentStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("准重 (吨)", labelStyle)).WithStyle(cellStyle),
		col.New(2).Add(text.New("20.10", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("其他(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("45", contentStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("包重(%)", labelStyle)).WithStyle(cellStyle),
		col.New(1).Add(text.New("55", contentStyle)).WithStyle(cellStyle),
	))

	// 5. 底部签名
	m.AddRows(
		row.New(5),
		row.New(10).Add(
			col.New(4).Add(text.New("过磅员：", props.Text{Size: 12})),
			col.New(4).Add(text.New("业务员：", props.Text{Size: 12})),
			col.New(4).Add(text.New("供货人：", props.Text{Size: 12})),
		),
	)

	// 保存结果
	doc, _ := m.Generate()
	// _ = doc.Save("weight_ticket_v2.pdf")
	return doc.GetBase64()
}
func Brcode() {

	customFont := "微软雅黑"
	customFontFile := "E:/fonts/fonts/msyh.ttf"
	customFonts, err := repository.New().
		AddUTF8Font(customFont, fontstyle.Normal, customFontFile).
		AddUTF8Font(customFont, fontstyle.Italic, customFontFile).
		AddUTF8Font(customFont, fontstyle.Bold, customFontFile).
		AddUTF8Font(customFont, fontstyle.BoldItalic, customFontFile).
		Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	builder := config.NewBuilder().
		WithDimensions(80, 20). // 设置页面宽高
		WithTopMargin(2).
		WithBottomMargin(2).
		WithLeftMargin(5).
		WithRightMargin(5).
		WithCustomFonts(customFonts)

	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).
		Build()

	m := maroto.New(cfg)
	m.AddRows(
		row.New(11).Add(
			code.NewBarCol(12, "123456789", props.Barcode{
				Type:    barcode.Code128,
				Percent: 100,
				Center:  true,
			}),
		),
	)

	m.AddRows(
		row.New(5).Add(
			col.New(12).Add(
				text.New("123456789", props.Text{
					Size:  12,
					Align: align.Center,
					Top:   1,
				}),
			),
		),
	)
	doc, _ := m.Generate()
	_ = doc.Save("brcode.pdf")
}
