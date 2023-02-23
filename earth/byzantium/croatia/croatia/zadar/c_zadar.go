package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZadarCounty interface {
	feud.County
	BBenkovac本科瓦茨() feud.Barony
	BKnin克宁() feud.Barony
	BNin尼恩() feud.Barony
	BNovigrad诺维格勒() feud.Barony
	BObrovac奥布罗瓦茨() feud.Barony
	BPag帕格() feud.Barony
	BSibenik希贝尼克() feud.Barony
	BZadar扎达尔() feud.Barony
}

type 扎达尔ZadarCounty struct {
	feud.BaseCounty
	本科瓦茨Benkovac feud.Barony
	克宁Knin       feud.Barony
	尼恩Nin        feud.Barony
	诺维格勒Novigrad feud.Barony
	奥布罗瓦茨Obrovac feud.Barony
	帕格Pag        feud.Barony
	希贝尼克Sibenik  feud.Barony
	扎达尔Zadar     feud.Barony
}

func (c *扎达尔ZadarCounty) BBenkovac本科瓦茨() feud.Barony {
	return c.本科瓦茨Benkovac
}

func (c *扎达尔ZadarCounty) BKnin克宁() feud.Barony {
	return c.克宁Knin
}

func (c *扎达尔ZadarCounty) BNin尼恩() feud.Barony {
	return c.尼恩Nin
}

func (c *扎达尔ZadarCounty) BNovigrad诺维格勒() feud.Barony {
	return c.诺维格勒Novigrad
}

func (c *扎达尔ZadarCounty) BObrovac奥布罗瓦茨() feud.Barony {
	return c.奥布罗瓦茨Obrovac
}

func (c *扎达尔ZadarCounty) BPag帕格() feud.Barony {
	return c.帕格Pag
}

func (c *扎达尔ZadarCounty) BSibenik希贝尼克() feud.Barony {
	return c.希贝尼克Sibenik
}

func (c *扎达尔ZadarCounty) BZadar扎达尔() feud.Barony {
	return c.扎达尔Zadar
}

var CZadar扎达尔 ZadarCounty = &扎达尔ZadarCounty{}

func init() {
	f := CZadar扎达尔.(*扎达尔ZadarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "465",
		Title:     "zadar",
		TitleName: "扎达尔",
		TitleCode: "c_zadar",
		Baronies:  map[string]feud.Barony{},
	}

	f.本科瓦茨Benkovac = BBenkovac本科瓦茨
	f.本科瓦茨Benkovac.SetParent(f)

	f.克宁Knin = BKnin克宁
	f.克宁Knin.SetParent(f)

	f.尼恩Nin = BNin尼恩
	f.尼恩Nin.SetParent(f)

	f.诺维格勒Novigrad = BNovigrad诺维格勒
	f.诺维格勒Novigrad.SetParent(f)

	f.奥布罗瓦茨Obrovac = BObrovac奥布罗瓦茨
	f.奥布罗瓦茨Obrovac.SetParent(f)

	f.帕格Pag = BPag帕格
	f.帕格Pag.SetParent(f)

	f.希贝尼克Sibenik = BSibenik希贝尼克
	f.希贝尼克Sibenik.SetParent(f)

	f.扎达尔Zadar = BZadar扎达尔
	f.扎达尔Zadar.SetParent(f)

}
