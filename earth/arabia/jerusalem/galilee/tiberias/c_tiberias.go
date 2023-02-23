package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TiberiasCounty interface {
	feud.County
	BAshtera艾特拉() feud.Barony
	BBethsan贝特谢安() feud.Barony
	BCaymont凯蒙特() feud.Barony
	BHattin哈丁() feud.Barony
	BLafeve拉费夫() feud.Barony
	BMnttabor塔博尔() feud.Barony
	BNazareth拿撒勒() feud.Barony
	BTiberias太巴列() feud.Barony
}

type 太巴列TiberiasCounty struct {
	feud.BaseCounty
	艾特拉Ashtera  feud.Barony
	贝特谢安Bethsan feud.Barony
	凯蒙特Caymont  feud.Barony
	哈丁Hattin    feud.Barony
	拉费夫Lafeve   feud.Barony
	塔博尔Mnttabor feud.Barony
	拿撒勒Nazareth feud.Barony
	太巴列Tiberias feud.Barony
}

func (c *太巴列TiberiasCounty) BAshtera艾特拉() feud.Barony {
	return c.艾特拉Ashtera
}

func (c *太巴列TiberiasCounty) BBethsan贝特谢安() feud.Barony {
	return c.贝特谢安Bethsan
}

func (c *太巴列TiberiasCounty) BCaymont凯蒙特() feud.Barony {
	return c.凯蒙特Caymont
}

func (c *太巴列TiberiasCounty) BHattin哈丁() feud.Barony {
	return c.哈丁Hattin
}

func (c *太巴列TiberiasCounty) BLafeve拉费夫() feud.Barony {
	return c.拉费夫Lafeve
}

func (c *太巴列TiberiasCounty) BMnttabor塔博尔() feud.Barony {
	return c.塔博尔Mnttabor
}

func (c *太巴列TiberiasCounty) BNazareth拿撒勒() feud.Barony {
	return c.拿撒勒Nazareth
}

func (c *太巴列TiberiasCounty) BTiberias太巴列() feud.Barony {
	return c.太巴列Tiberias
}

var CTiberias太巴列 TiberiasCounty = &太巴列TiberiasCounty{}

func init() {
	f := CTiberias太巴列.(*太巴列TiberiasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "773",
		Title:     "tiberias",
		TitleName: "太巴列",
		TitleCode: "c_tiberias",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾特拉Ashtera = BAshtera艾特拉
	f.艾特拉Ashtera.SetParent(f)

	f.贝特谢安Bethsan = BBethsan贝特谢安
	f.贝特谢安Bethsan.SetParent(f)

	f.凯蒙特Caymont = BCaymont凯蒙特
	f.凯蒙特Caymont.SetParent(f)

	f.哈丁Hattin = BHattin哈丁
	f.哈丁Hattin.SetParent(f)

	f.拉费夫Lafeve = BLafeve拉费夫
	f.拉费夫Lafeve.SetParent(f)

	f.塔博尔Mnttabor = BMnttabor塔博尔
	f.塔博尔Mnttabor.SetParent(f)

	f.拿撒勒Nazareth = BNazareth拿撒勒
	f.拿撒勒Nazareth.SetParent(f)

	f.太巴列Tiberias = BTiberias太巴列
	f.太巴列Tiberias.SetParent(f)

}
