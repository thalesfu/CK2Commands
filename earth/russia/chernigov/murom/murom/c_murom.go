package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MuromCounty interface {
	feud.County
	BArkhangel阿尔汉格尔() feud.Barony
	BKasimov卡西莫夫() feud.Barony
	BMelenki梅连基() feud.Barony
	BMurom穆罗姆() feud.Barony
	BUlanovka乌拉诺夫卡() feud.Barony
	BVyksa维克萨() feud.Barony
	BYelatma叶拉季马() feud.Barony
}

type 穆罗姆MuromCounty struct {
	feud.BaseCounty
	阿尔汉格尔Arkhangel feud.Barony
	卡西莫夫Kasimov    feud.Barony
	梅连基Melenki     feud.Barony
	穆罗姆Murom       feud.Barony
	乌拉诺夫卡Ulanovka  feud.Barony
	维克萨Vyksa       feud.Barony
	叶拉季马Yelatma    feud.Barony
}

func (c *穆罗姆MuromCounty) BArkhangel阿尔汉格尔() feud.Barony {
	return c.阿尔汉格尔Arkhangel
}

func (c *穆罗姆MuromCounty) BKasimov卡西莫夫() feud.Barony {
	return c.卡西莫夫Kasimov
}

func (c *穆罗姆MuromCounty) BMelenki梅连基() feud.Barony {
	return c.梅连基Melenki
}

func (c *穆罗姆MuromCounty) BMurom穆罗姆() feud.Barony {
	return c.穆罗姆Murom
}

func (c *穆罗姆MuromCounty) BUlanovka乌拉诺夫卡() feud.Barony {
	return c.乌拉诺夫卡Ulanovka
}

func (c *穆罗姆MuromCounty) BVyksa维克萨() feud.Barony {
	return c.维克萨Vyksa
}

func (c *穆罗姆MuromCounty) BYelatma叶拉季马() feud.Barony {
	return c.叶拉季马Yelatma
}

var CMurom穆罗姆 MuromCounty = &穆罗姆MuromCounty{}

func init() {
	f := CMurom穆罗姆.(*穆罗姆MuromCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "581",
		Title:     "murom",
		TitleName: "穆罗姆",
		TitleCode: "c_murom",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔汉格尔Arkhangel = BArkhangel阿尔汉格尔
	f.阿尔汉格尔Arkhangel.SetParent(f)

	f.卡西莫夫Kasimov = BKasimov卡西莫夫
	f.卡西莫夫Kasimov.SetParent(f)

	f.梅连基Melenki = BMelenki梅连基
	f.梅连基Melenki.SetParent(f)

	f.穆罗姆Murom = BMurom穆罗姆
	f.穆罗姆Murom.SetParent(f)

	f.乌拉诺夫卡Ulanovka = BUlanovka乌拉诺夫卡
	f.乌拉诺夫卡Ulanovka.SetParent(f)

	f.维克萨Vyksa = BVyksa维克萨
	f.维克萨Vyksa.SetParent(f)

	f.叶拉季马Yelatma = BYelatma叶拉季马
	f.叶拉季马Yelatma.SetParent(f)

}
