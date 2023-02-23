package trapani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrapaniCounty interface {
	feud.County
	BAlcarno阿尔卡莫() feud.Barony
	BCastelvertrano卡斯泰尔韦特拉诺() feud.Barony
	BErice埃里切() feud.Barony
	BMarsala马尔萨拉() feud.Barony
	BMazara马扎拉() feud.Barony
	BSanguiseppelato圣朱塞佩亚托() feud.Barony
	BSantaninfa圣宁法() feud.Barony
	BTrapani特拉帕尼() feud.Barony
}

type 特拉帕尼TrapaniCounty struct {
	feud.BaseCounty
	阿尔卡莫Alcarno            feud.Barony
	卡斯泰尔韦特拉诺Castelvertrano feud.Barony
	埃里切Erice               feud.Barony
	马尔萨拉Marsala            feud.Barony
	马扎拉Mazara              feud.Barony
	圣朱塞佩亚托Sanguiseppelato  feud.Barony
	圣宁法Santaninfa          feud.Barony
	特拉帕尼Trapani            feud.Barony
}

func (c *特拉帕尼TrapaniCounty) BAlcarno阿尔卡莫() feud.Barony {
	return c.阿尔卡莫Alcarno
}

func (c *特拉帕尼TrapaniCounty) BCastelvertrano卡斯泰尔韦特拉诺() feud.Barony {
	return c.卡斯泰尔韦特拉诺Castelvertrano
}

func (c *特拉帕尼TrapaniCounty) BErice埃里切() feud.Barony {
	return c.埃里切Erice
}

func (c *特拉帕尼TrapaniCounty) BMarsala马尔萨拉() feud.Barony {
	return c.马尔萨拉Marsala
}

func (c *特拉帕尼TrapaniCounty) BMazara马扎拉() feud.Barony {
	return c.马扎拉Mazara
}

func (c *特拉帕尼TrapaniCounty) BSanguiseppelato圣朱塞佩亚托() feud.Barony {
	return c.圣朱塞佩亚托Sanguiseppelato
}

func (c *特拉帕尼TrapaniCounty) BSantaninfa圣宁法() feud.Barony {
	return c.圣宁法Santaninfa
}

func (c *特拉帕尼TrapaniCounty) BTrapani特拉帕尼() feud.Barony {
	return c.特拉帕尼Trapani
}

var CTrapani特拉帕尼 TrapaniCounty = &特拉帕尼TrapaniCounty{}

func init() {
	f := CTrapani特拉帕尼.(*特拉帕尼TrapaniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "341",
		Title:     "trapani",
		TitleName: "特拉帕尼",
		TitleCode: "c_trapani",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔卡莫Alcarno = BAlcarno阿尔卡莫
	f.阿尔卡莫Alcarno.SetParent(f)

	f.卡斯泰尔韦特拉诺Castelvertrano = BCastelvertrano卡斯泰尔韦特拉诺
	f.卡斯泰尔韦特拉诺Castelvertrano.SetParent(f)

	f.埃里切Erice = BErice埃里切
	f.埃里切Erice.SetParent(f)

	f.马尔萨拉Marsala = BMarsala马尔萨拉
	f.马尔萨拉Marsala.SetParent(f)

	f.马扎拉Mazara = BMazara马扎拉
	f.马扎拉Mazara.SetParent(f)

	f.圣朱塞佩亚托Sanguiseppelato = BSanguiseppelato圣朱塞佩亚托
	f.圣朱塞佩亚托Sanguiseppelato.SetParent(f)

	f.圣宁法Santaninfa = BSantaninfa圣宁法
	f.圣宁法Santaninfa.SetParent(f)

	f.特拉帕尼Trapani = BTrapani特拉帕尼
	f.特拉帕尼Trapani.SetParent(f)

}
