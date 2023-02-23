package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OverijsselCounty interface {
	feud.County
	BCoevorden库福尔登() feud.Barony
	BDeventer代芬特尔() feud.Barony
	BDrenthe德伦特() feud.Barony
	BEnschede恩斯赫德() feud.Barony
	BKampen坎彭() feud.Barony
	BVollenhove福伦霍弗() feud.Barony
	BZwolle兹沃勒() feud.Barony
}

type 上艾瑟尔OverijsselCounty struct {
	feud.BaseCounty
	库福尔登Coevorden  feud.Barony
	代芬特尔Deventer   feud.Barony
	德伦特Drenthe     feud.Barony
	恩斯赫德Enschede   feud.Barony
	坎彭Kampen       feud.Barony
	福伦霍弗Vollenhove feud.Barony
	兹沃勒Zwolle      feud.Barony
}

func (c *上艾瑟尔OverijsselCounty) BCoevorden库福尔登() feud.Barony {
	return c.库福尔登Coevorden
}

func (c *上艾瑟尔OverijsselCounty) BDeventer代芬特尔() feud.Barony {
	return c.代芬特尔Deventer
}

func (c *上艾瑟尔OverijsselCounty) BDrenthe德伦特() feud.Barony {
	return c.德伦特Drenthe
}

func (c *上艾瑟尔OverijsselCounty) BEnschede恩斯赫德() feud.Barony {
	return c.恩斯赫德Enschede
}

func (c *上艾瑟尔OverijsselCounty) BKampen坎彭() feud.Barony {
	return c.坎彭Kampen
}

func (c *上艾瑟尔OverijsselCounty) BVollenhove福伦霍弗() feud.Barony {
	return c.福伦霍弗Vollenhove
}

func (c *上艾瑟尔OverijsselCounty) BZwolle兹沃勒() feud.Barony {
	return c.兹沃勒Zwolle
}

var COverijssel上艾瑟尔 OverijsselCounty = &上艾瑟尔OverijsselCounty{}

func init() {
	f := COverijssel上艾瑟尔.(*上艾瑟尔OverijsselCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1989",
		Title:     "overijssel",
		TitleName: "上艾瑟尔",
		TitleCode: "c_overijssel",
		Baronies:  map[string]feud.Barony{},
	}

	f.库福尔登Coevorden = BCoevorden库福尔登
	f.库福尔登Coevorden.SetParent(f)

	f.代芬特尔Deventer = BDeventer代芬特尔
	f.代芬特尔Deventer.SetParent(f)

	f.德伦特Drenthe = BDrenthe德伦特
	f.德伦特Drenthe.SetParent(f)

	f.恩斯赫德Enschede = BEnschede恩斯赫德
	f.恩斯赫德Enschede.SetParent(f)

	f.坎彭Kampen = BKampen坎彭
	f.坎彭Kampen.SetParent(f)

	f.福伦霍弗Vollenhove = BVollenhove福伦霍弗
	f.福伦霍弗Vollenhove.SetParent(f)

	f.兹沃勒Zwolle = BZwolle兹沃勒
	f.兹沃勒Zwolle.SetParent(f)

}
