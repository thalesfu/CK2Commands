package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PamirCounty interface {
	feud.County
	BDarmadar达麦大() feud.Barony
	BKhorugh霍罗格() feud.Barony
	BTraguladar特拉古拉达尔() feud.Barony
	BVanj万吉() feud.Barony
}

type 播密PamirCounty struct {
	feud.BaseCounty
	达麦大Darmadar      feud.Barony
	霍罗格Khorugh       feud.Barony
	特拉古拉达尔Traguladar feud.Barony
	万吉Vanj           feud.Barony
}

func (c *播密PamirCounty) BDarmadar达麦大() feud.Barony {
	return c.达麦大Darmadar
}

func (c *播密PamirCounty) BKhorugh霍罗格() feud.Barony {
	return c.霍罗格Khorugh
}

func (c *播密PamirCounty) BTraguladar特拉古拉达尔() feud.Barony {
	return c.特拉古拉达尔Traguladar
}

func (c *播密PamirCounty) BVanj万吉() feud.Barony {
	return c.万吉Vanj
}

var CPamir播密 PamirCounty = &播密PamirCounty{}

func init() {
	f := CPamir播密.(*播密PamirCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1511",
		Title:     "pamir",
		TitleName: "播密",
		TitleCode: "c_pamir",
		Baronies:  map[string]feud.Barony{},
	}

	f.达麦大Darmadar = BDarmadar达麦大
	f.达麦大Darmadar.SetParent(f)

	f.霍罗格Khorugh = BKhorugh霍罗格
	f.霍罗格Khorugh.SetParent(f)

	f.特拉古拉达尔Traguladar = BTraguladar特拉古拉达尔
	f.特拉古拉达尔Traguladar.SetParent(f)

	f.万吉Vanj = BVanj万吉
	f.万吉Vanj.SetParent(f)

}
