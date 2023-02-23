package jylland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JyllandCounty interface {
	feud.County
	BBurlun柏伦() feud.Barony
	BHjorring约灵() feud.Barony
	BHorns霍恩斯() feud.Barony
	BHvet韦特() feud.Barony
	BJarlslef耶斯莱乌() feud.Barony
	BSkagen斯卡恩() feud.Barony
	BThystadth齐斯泰兹() feud.Barony
}

type 斯卡恩JyllandCounty struct {
	feud.BaseCounty
	柏伦Burlun      feud.Barony
	约灵Hjorring    feud.Barony
	霍恩斯Horns      feud.Barony
	韦特Hvet        feud.Barony
	耶斯莱乌Jarlslef  feud.Barony
	斯卡恩Skagen     feud.Barony
	齐斯泰兹Thystadth feud.Barony
}

func (c *斯卡恩JyllandCounty) BBurlun柏伦() feud.Barony {
	return c.柏伦Burlun
}

func (c *斯卡恩JyllandCounty) BHjorring约灵() feud.Barony {
	return c.约灵Hjorring
}

func (c *斯卡恩JyllandCounty) BHorns霍恩斯() feud.Barony {
	return c.霍恩斯Horns
}

func (c *斯卡恩JyllandCounty) BHvet韦特() feud.Barony {
	return c.韦特Hvet
}

func (c *斯卡恩JyllandCounty) BJarlslef耶斯莱乌() feud.Barony {
	return c.耶斯莱乌Jarlslef
}

func (c *斯卡恩JyllandCounty) BSkagen斯卡恩() feud.Barony {
	return c.斯卡恩Skagen
}

func (c *斯卡恩JyllandCounty) BThystadth齐斯泰兹() feud.Barony {
	return c.齐斯泰兹Thystadth
}

var CJylland斯卡恩 JyllandCounty = &斯卡恩JyllandCounty{}

func init() {
	f := CJylland斯卡恩.(*斯卡恩JyllandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "267",
		Title:     "jylland",
		TitleName: "斯卡恩",
		TitleCode: "c_jylland",
		Baronies:  map[string]feud.Barony{},
	}

	f.柏伦Burlun = BBurlun柏伦
	f.柏伦Burlun.SetParent(f)

	f.约灵Hjorring = BHjorring约灵
	f.约灵Hjorring.SetParent(f)

	f.霍恩斯Horns = BHorns霍恩斯
	f.霍恩斯Horns.SetParent(f)

	f.韦特Hvet = BHvet韦特
	f.韦特Hvet.SetParent(f)

	f.耶斯莱乌Jarlslef = BJarlslef耶斯莱乌
	f.耶斯莱乌Jarlslef.SetParent(f)

	f.斯卡恩Skagen = BSkagen斯卡恩
	f.斯卡恩Skagen.SetParent(f)

	f.齐斯泰兹Thystadth = BThystadth齐斯泰兹
	f.齐斯泰兹Thystadth.SetParent(f)

}
