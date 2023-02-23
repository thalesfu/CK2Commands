package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OtukenCounty interface {
	feud.County
	BKutulik库图利克() feud.Barony
	BKyren基伦() feud.Barony
	BOtuken于都斤() feud.Barony
	BSayan萨彦() feud.Barony
	BTenger腾格尔() feud.Barony
	BTulun图伦() feud.Barony
}

type 于都斤OtukenCounty struct {
	feud.BaseCounty
	库图利克Kutulik feud.Barony
	基伦Kyren     feud.Barony
	于都斤Otuken   feud.Barony
	萨彦Sayan     feud.Barony
	腾格尔Tenger   feud.Barony
	图伦Tulun     feud.Barony
}

func (c *于都斤OtukenCounty) BKutulik库图利克() feud.Barony {
	return c.库图利克Kutulik
}

func (c *于都斤OtukenCounty) BKyren基伦() feud.Barony {
	return c.基伦Kyren
}

func (c *于都斤OtukenCounty) BOtuken于都斤() feud.Barony {
	return c.于都斤Otuken
}

func (c *于都斤OtukenCounty) BSayan萨彦() feud.Barony {
	return c.萨彦Sayan
}

func (c *于都斤OtukenCounty) BTenger腾格尔() feud.Barony {
	return c.腾格尔Tenger
}

func (c *于都斤OtukenCounty) BTulun图伦() feud.Barony {
	return c.图伦Tulun
}

var COtuken于都斤 OtukenCounty = &于都斤OtukenCounty{}

func init() {
	f := COtuken于都斤.(*于都斤OtukenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1461",
		Title:     "otuken",
		TitleName: "于都斤",
		TitleCode: "c_otuken",
		Baronies:  map[string]feud.Barony{},
	}

	f.库图利克Kutulik = BKutulik库图利克
	f.库图利克Kutulik.SetParent(f)

	f.基伦Kyren = BKyren基伦
	f.基伦Kyren.SetParent(f)

	f.于都斤Otuken = BOtuken于都斤
	f.于都斤Otuken.SetParent(f)

	f.萨彦Sayan = BSayan萨彦
	f.萨彦Sayan.SetParent(f)

	f.腾格尔Tenger = BTenger腾格尔
	f.腾格尔Tenger.SetParent(f)

	f.图伦Tulun = BTulun图伦
	f.图伦Tulun.SetParent(f)

}
