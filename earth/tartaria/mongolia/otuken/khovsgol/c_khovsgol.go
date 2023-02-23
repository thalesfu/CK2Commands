package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhovsgolCounty interface {
	feud.County
	BArshan阿尔尚() feud.Barony
	BBotogol博托戈尔() feud.Barony
	BKhovsgol库苏古尔() feud.Barony
	BMondy蒙德() feud.Barony
	BOnot奥诺特() feud.Barony
	BOrlik奥尔利克() feud.Barony
	BSorok索罗克() feud.Barony
}

type 库苏古尔KhovsgolCounty struct {
	feud.BaseCounty
	阿尔尚Arshan    feud.Barony
	博托戈尔Botogol  feud.Barony
	库苏古尔Khovsgol feud.Barony
	蒙德Mondy      feud.Barony
	奥诺特Onot      feud.Barony
	奥尔利克Orlik    feud.Barony
	索罗克Sorok     feud.Barony
}

func (c *库苏古尔KhovsgolCounty) BArshan阿尔尚() feud.Barony {
	return c.阿尔尚Arshan
}

func (c *库苏古尔KhovsgolCounty) BBotogol博托戈尔() feud.Barony {
	return c.博托戈尔Botogol
}

func (c *库苏古尔KhovsgolCounty) BKhovsgol库苏古尔() feud.Barony {
	return c.库苏古尔Khovsgol
}

func (c *库苏古尔KhovsgolCounty) BMondy蒙德() feud.Barony {
	return c.蒙德Mondy
}

func (c *库苏古尔KhovsgolCounty) BOnot奥诺特() feud.Barony {
	return c.奥诺特Onot
}

func (c *库苏古尔KhovsgolCounty) BOrlik奥尔利克() feud.Barony {
	return c.奥尔利克Orlik
}

func (c *库苏古尔KhovsgolCounty) BSorok索罗克() feud.Barony {
	return c.索罗克Sorok
}

var CKhovsgol库苏古尔 KhovsgolCounty = &库苏古尔KhovsgolCounty{}

func init() {
	f := CKhovsgol库苏古尔.(*库苏古尔KhovsgolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1903",
		Title:     "khovsgol",
		TitleName: "库苏古尔",
		TitleCode: "c_khovsgol",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔尚Arshan = BArshan阿尔尚
	f.阿尔尚Arshan.SetParent(f)

	f.博托戈尔Botogol = BBotogol博托戈尔
	f.博托戈尔Botogol.SetParent(f)

	f.库苏古尔Khovsgol = BKhovsgol库苏古尔
	f.库苏古尔Khovsgol.SetParent(f)

	f.蒙德Mondy = BMondy蒙德
	f.蒙德Mondy.SetParent(f)

	f.奥诺特Onot = BOnot奥诺特
	f.奥诺特Onot.SetParent(f)

	f.奥尔利克Orlik = BOrlik奥尔利克
	f.奥尔利克Orlik.SetParent(f)

	f.索罗克Sorok = BSorok索罗克
	f.索罗克Sorok.SetParent(f)

}
