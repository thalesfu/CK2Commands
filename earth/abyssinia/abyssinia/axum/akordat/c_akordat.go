package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AkordatCounty interface {
	feud.County
	BAkordat阿科达特() feud.Barony
	BAntalla安塔拉() feud.Barony
	BBarentu巴伦图() feud.Barony
	BDghe德赫() feud.Barony
	BMogolo莫格洛() feud.Barony
	BShambuko沙姆布克() feud.Barony
	BTeseney特瑟内() feud.Barony
}

type 阿科达特AkordatCounty struct {
	feud.BaseCounty
	阿科达特Akordat  feud.Barony
	安塔拉Antalla   feud.Barony
	巴伦图Barentu   feud.Barony
	德赫Dghe       feud.Barony
	莫格洛Mogolo    feud.Barony
	沙姆布克Shambuko feud.Barony
	特瑟内Teseney   feud.Barony
}

func (c *阿科达特AkordatCounty) BAkordat阿科达特() feud.Barony {
	return c.阿科达特Akordat
}

func (c *阿科达特AkordatCounty) BAntalla安塔拉() feud.Barony {
	return c.安塔拉Antalla
}

func (c *阿科达特AkordatCounty) BBarentu巴伦图() feud.Barony {
	return c.巴伦图Barentu
}

func (c *阿科达特AkordatCounty) BDghe德赫() feud.Barony {
	return c.德赫Dghe
}

func (c *阿科达特AkordatCounty) BMogolo莫格洛() feud.Barony {
	return c.莫格洛Mogolo
}

func (c *阿科达特AkordatCounty) BShambuko沙姆布克() feud.Barony {
	return c.沙姆布克Shambuko
}

func (c *阿科达特AkordatCounty) BTeseney特瑟内() feud.Barony {
	return c.特瑟内Teseney
}

var CAkordat阿科达特 AkordatCounty = &阿科达特AkordatCounty{}

func init() {
	f := CAkordat阿科达特.(*阿科达特AkordatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "876",
		Title:     "akordat",
		TitleName: "阿科达特",
		TitleCode: "c_akordat",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿科达特Akordat = BAkordat阿科达特
	f.阿科达特Akordat.SetParent(f)

	f.安塔拉Antalla = BAntalla安塔拉
	f.安塔拉Antalla.SetParent(f)

	f.巴伦图Barentu = BBarentu巴伦图
	f.巴伦图Barentu.SetParent(f)

	f.德赫Dghe = BDghe德赫
	f.德赫Dghe.SetParent(f)

	f.莫格洛Mogolo = BMogolo莫格洛
	f.莫格洛Mogolo.SetParent(f)

	f.沙姆布克Shambuko = BShambuko沙姆布克
	f.沙姆布克Shambuko.SetParent(f)

	f.特瑟内Teseney = BTeseney特瑟内
	f.特瑟内Teseney.SetParent(f)

}
