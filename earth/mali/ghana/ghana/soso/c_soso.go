package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SosoCounty interface {
	feud.County
	BDiado迪亚多() feud.Barony
	BKirina基里纳() feud.Barony
	BKoio库瓦奥() feud.Barony
	BSama萨马() feud.Barony
	BSebete塞贝泰() feud.Barony
	BSegou塞古() feud.Barony
	BSoso索索() feud.Barony
}

type 索索SosoCounty struct {
	feud.BaseCounty
	迪亚多Diado  feud.Barony
	基里纳Kirina feud.Barony
	库瓦奥Koio   feud.Barony
	萨马Sama    feud.Barony
	塞贝泰Sebete feud.Barony
	塞古Segou   feud.Barony
	索索Soso    feud.Barony
}

func (c *索索SosoCounty) BDiado迪亚多() feud.Barony {
	return c.迪亚多Diado
}

func (c *索索SosoCounty) BKirina基里纳() feud.Barony {
	return c.基里纳Kirina
}

func (c *索索SosoCounty) BKoio库瓦奥() feud.Barony {
	return c.库瓦奥Koio
}

func (c *索索SosoCounty) BSama萨马() feud.Barony {
	return c.萨马Sama
}

func (c *索索SosoCounty) BSebete塞贝泰() feud.Barony {
	return c.塞贝泰Sebete
}

func (c *索索SosoCounty) BSegou塞古() feud.Barony {
	return c.塞古Segou
}

func (c *索索SosoCounty) BSoso索索() feud.Barony {
	return c.索索Soso
}

var CSoso索索 SosoCounty = &索索SosoCounty{}

func init() {
	f := CSoso索索.(*索索SosoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1880",
		Title:     "soso",
		TitleName: "索索",
		TitleCode: "c_soso",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪亚多Diado = BDiado迪亚多
	f.迪亚多Diado.SetParent(f)

	f.基里纳Kirina = BKirina基里纳
	f.基里纳Kirina.SetParent(f)

	f.库瓦奥Koio = BKoio库瓦奥
	f.库瓦奥Koio.SetParent(f)

	f.萨马Sama = BSama萨马
	f.萨马Sama.SetParent(f)

	f.塞贝泰Sebete = BSebete塞贝泰
	f.塞贝泰Sebete.SetParent(f)

	f.塞古Segou = BSegou塞古
	f.塞古Segou.SetParent(f)

	f.索索Soso = BSoso索索
	f.索索Soso.SetParent(f)

}
