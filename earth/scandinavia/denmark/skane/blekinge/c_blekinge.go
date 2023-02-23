package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BlekingeCounty interface {
	feud.County
	BAvaskar阿瓦谢尔() feud.Barony
	BElleholm埃勒霍尔姆() feud.Barony
	BLister利斯特() feud.Barony
	BLycka吕科() feud.Barony
	BLyckeby吕克比() feud.Barony
	BRonneby龙讷比() feud.Barony
	BSolvesborg瑟尔沃斯堡() feud.Barony
}

type 布莱金厄BlekingeCounty struct {
	feud.BaseCounty
	阿瓦谢尔Avaskar     feud.Barony
	埃勒霍尔姆Elleholm   feud.Barony
	利斯特Lister       feud.Barony
	吕科Lycka         feud.Barony
	吕克比Lyckeby      feud.Barony
	龙讷比Ronneby      feud.Barony
	瑟尔沃斯堡Solvesborg feud.Barony
}

func (c *布莱金厄BlekingeCounty) BAvaskar阿瓦谢尔() feud.Barony {
	return c.阿瓦谢尔Avaskar
}

func (c *布莱金厄BlekingeCounty) BElleholm埃勒霍尔姆() feud.Barony {
	return c.埃勒霍尔姆Elleholm
}

func (c *布莱金厄BlekingeCounty) BLister利斯特() feud.Barony {
	return c.利斯特Lister
}

func (c *布莱金厄BlekingeCounty) BLycka吕科() feud.Barony {
	return c.吕科Lycka
}

func (c *布莱金厄BlekingeCounty) BLyckeby吕克比() feud.Barony {
	return c.吕克比Lyckeby
}

func (c *布莱金厄BlekingeCounty) BRonneby龙讷比() feud.Barony {
	return c.龙讷比Ronneby
}

func (c *布莱金厄BlekingeCounty) BSolvesborg瑟尔沃斯堡() feud.Barony {
	return c.瑟尔沃斯堡Solvesborg
}

var CBlekinge布莱金厄 BlekingeCounty = &布莱金厄BlekingeCounty{}

func init() {
	f := CBlekinge布莱金厄.(*布莱金厄BlekingeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "6",
		Title:     "blekinge",
		TitleName: "布莱金厄",
		TitleCode: "c_blekinge",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿瓦谢尔Avaskar = BAvaskar阿瓦谢尔
	f.阿瓦谢尔Avaskar.SetParent(f)

	f.埃勒霍尔姆Elleholm = BElleholm埃勒霍尔姆
	f.埃勒霍尔姆Elleholm.SetParent(f)

	f.利斯特Lister = BLister利斯特
	f.利斯特Lister.SetParent(f)

	f.吕科Lycka = BLycka吕科
	f.吕科Lycka.SetParent(f)

	f.吕克比Lyckeby = BLyckeby吕克比
	f.吕克比Lyckeby.SetParent(f)

	f.龙讷比Ronneby = BRonneby龙讷比
	f.龙讷比Ronneby.SetParent(f)

	f.瑟尔沃斯堡Solvesborg = BSolvesborg瑟尔沃斯堡
	f.瑟尔沃斯堡Solvesborg.SetParent(f)

}
