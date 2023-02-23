package dwin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DwinCounty interface {
	feud.County
	BAlaverdi阿拉韦尔迪() feud.Barony
	BDvin德温() feud.Barony
	BErebuni埃雷布尼() feud.Barony
	BEtchmiadzin埃奇米阿津() feud.Barony
	BKhorvirap霍尔维拉普() feud.Barony
	BMatsnaberd马兹纳贝德() feud.Barony
	BTumanyan图曼扬() feud.Barony
}

type 德温DwinCounty struct {
	feud.BaseCounty
	阿拉韦尔迪Alaverdi    feud.Barony
	德温Dvin           feud.Barony
	埃雷布尼Erebuni      feud.Barony
	埃奇米阿津Etchmiadzin feud.Barony
	霍尔维拉普Khorvirap   feud.Barony
	马兹纳贝德Matsnaberd  feud.Barony
	图曼扬Tumanyan      feud.Barony
}

func (c *德温DwinCounty) BAlaverdi阿拉韦尔迪() feud.Barony {
	return c.阿拉韦尔迪Alaverdi
}

func (c *德温DwinCounty) BDvin德温() feud.Barony {
	return c.德温Dvin
}

func (c *德温DwinCounty) BErebuni埃雷布尼() feud.Barony {
	return c.埃雷布尼Erebuni
}

func (c *德温DwinCounty) BEtchmiadzin埃奇米阿津() feud.Barony {
	return c.埃奇米阿津Etchmiadzin
}

func (c *德温DwinCounty) BKhorvirap霍尔维拉普() feud.Barony {
	return c.霍尔维拉普Khorvirap
}

func (c *德温DwinCounty) BMatsnaberd马兹纳贝德() feud.Barony {
	return c.马兹纳贝德Matsnaberd
}

func (c *德温DwinCounty) BTumanyan图曼扬() feud.Barony {
	return c.图曼扬Tumanyan
}

var CDwin德温 DwinCounty = &德温DwinCounty{}

func init() {
	f := CDwin德温.(*德温DwinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "672",
		Title:     "dwin",
		TitleName: "德温",
		TitleCode: "c_dwin",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉韦尔迪Alaverdi = BAlaverdi阿拉韦尔迪
	f.阿拉韦尔迪Alaverdi.SetParent(f)

	f.德温Dvin = BDvin德温
	f.德温Dvin.SetParent(f)

	f.埃雷布尼Erebuni = BErebuni埃雷布尼
	f.埃雷布尼Erebuni.SetParent(f)

	f.埃奇米阿津Etchmiadzin = BEtchmiadzin埃奇米阿津
	f.埃奇米阿津Etchmiadzin.SetParent(f)

	f.霍尔维拉普Khorvirap = BKhorvirap霍尔维拉普
	f.霍尔维拉普Khorvirap.SetParent(f)

	f.马兹纳贝德Matsnaberd = BMatsnaberd马兹纳贝德
	f.马兹纳贝德Matsnaberd.SetParent(f)

	f.图曼扬Tumanyan = BTumanyan图曼扬
	f.图曼扬Tumanyan.SetParent(f)

}
