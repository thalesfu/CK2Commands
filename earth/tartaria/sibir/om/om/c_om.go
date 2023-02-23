package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OmCounty interface {
	feud.County
	BBorodinka博罗金卡() feud.Barony
	BButovka布托夫卡() feud.Barony
	BOm鄂木() feud.Barony
	BOmskiy鄂木斯基() feud.Barony
	BRostovka罗斯托夫卡() feud.Barony
	BStepnoy斯捷普诺伊() feud.Barony
	BZotino佐季诺() feud.Barony
}

type 鄂木OmCounty struct {
	feud.BaseCounty
	博罗金卡Borodinka feud.Barony
	布托夫卡Butovka   feud.Barony
	鄂木Om          feud.Barony
	鄂木斯基Omskiy    feud.Barony
	罗斯托夫卡Rostovka feud.Barony
	斯捷普诺伊Stepnoy  feud.Barony
	佐季诺Zotino     feud.Barony
}

func (c *鄂木OmCounty) BBorodinka博罗金卡() feud.Barony {
	return c.博罗金卡Borodinka
}

func (c *鄂木OmCounty) BButovka布托夫卡() feud.Barony {
	return c.布托夫卡Butovka
}

func (c *鄂木OmCounty) BOm鄂木() feud.Barony {
	return c.鄂木Om
}

func (c *鄂木OmCounty) BOmskiy鄂木斯基() feud.Barony {
	return c.鄂木斯基Omskiy
}

func (c *鄂木OmCounty) BRostovka罗斯托夫卡() feud.Barony {
	return c.罗斯托夫卡Rostovka
}

func (c *鄂木OmCounty) BStepnoy斯捷普诺伊() feud.Barony {
	return c.斯捷普诺伊Stepnoy
}

func (c *鄂木OmCounty) BZotino佐季诺() feud.Barony {
	return c.佐季诺Zotino
}

var COm鄂木 OmCounty = &鄂木OmCounty{}

func init() {
	f := COm鄂木.(*鄂木OmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1861",
		Title:     "om",
		TitleName: "鄂木",
		TitleCode: "c_om",
		Baronies:  map[string]feud.Barony{},
	}

	f.博罗金卡Borodinka = BBorodinka博罗金卡
	f.博罗金卡Borodinka.SetParent(f)

	f.布托夫卡Butovka = BButovka布托夫卡
	f.布托夫卡Butovka.SetParent(f)

	f.鄂木Om = BOm鄂木
	f.鄂木Om.SetParent(f)

	f.鄂木斯基Omskiy = BOmskiy鄂木斯基
	f.鄂木斯基Omskiy.SetParent(f)

	f.罗斯托夫卡Rostovka = BRostovka罗斯托夫卡
	f.罗斯托夫卡Rostovka.SetParent(f)

	f.斯捷普诺伊Stepnoy = BStepnoy斯捷普诺伊
	f.斯捷普诺伊Stepnoy.SetParent(f)

	f.佐季诺Zotino = BZotino佐季诺
	f.佐季诺Zotino.SetParent(f)

}
