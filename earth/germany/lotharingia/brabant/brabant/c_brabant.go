package brabant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrabantCounty interface {
	feud.County
	BAarschot阿尔斯霍特() feud.Barony
	BAntwerpen安特卫普() feud.Barony
	BBrussel布鲁塞尔() feud.Barony
	BGrimbergen赫林贝亨() feud.Barony
	BHerstal赫尔斯塔尔() feud.Barony
	BLeuven勒芬() feud.Barony
	BLier利尔() feud.Barony
	BMechelen梅赫伦() feud.Barony
}

type 布拉班特BrabantCounty struct {
	feud.BaseCounty
	阿尔斯霍特Aarschot  feud.Barony
	安特卫普Antwerpen  feud.Barony
	布鲁塞尔Brussel    feud.Barony
	赫林贝亨Grimbergen feud.Barony
	赫尔斯塔尔Herstal   feud.Barony
	勒芬Leuven       feud.Barony
	利尔Lier         feud.Barony
	梅赫伦Mechelen    feud.Barony
}

func (c *布拉班特BrabantCounty) BAarschot阿尔斯霍特() feud.Barony {
	return c.阿尔斯霍特Aarschot
}

func (c *布拉班特BrabantCounty) BAntwerpen安特卫普() feud.Barony {
	return c.安特卫普Antwerpen
}

func (c *布拉班特BrabantCounty) BBrussel布鲁塞尔() feud.Barony {
	return c.布鲁塞尔Brussel
}

func (c *布拉班特BrabantCounty) BGrimbergen赫林贝亨() feud.Barony {
	return c.赫林贝亨Grimbergen
}

func (c *布拉班特BrabantCounty) BHerstal赫尔斯塔尔() feud.Barony {
	return c.赫尔斯塔尔Herstal
}

func (c *布拉班特BrabantCounty) BLeuven勒芬() feud.Barony {
	return c.勒芬Leuven
}

func (c *布拉班特BrabantCounty) BLier利尔() feud.Barony {
	return c.利尔Lier
}

func (c *布拉班特BrabantCounty) BMechelen梅赫伦() feud.Barony {
	return c.梅赫伦Mechelen
}

var CBrabant布拉班特 BrabantCounty = &布拉班特BrabantCounty{}

func init() {
	f := CBrabant布拉班特.(*布拉班特BrabantCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "117",
		Title:     "brabant",
		TitleName: "布拉班特",
		TitleCode: "c_brabant",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔斯霍特Aarschot = BAarschot阿尔斯霍特
	f.阿尔斯霍特Aarschot.SetParent(f)

	f.安特卫普Antwerpen = BAntwerpen安特卫普
	f.安特卫普Antwerpen.SetParent(f)

	f.布鲁塞尔Brussel = BBrussel布鲁塞尔
	f.布鲁塞尔Brussel.SetParent(f)

	f.赫林贝亨Grimbergen = BGrimbergen赫林贝亨
	f.赫林贝亨Grimbergen.SetParent(f)

	f.赫尔斯塔尔Herstal = BHerstal赫尔斯塔尔
	f.赫尔斯塔尔Herstal.SetParent(f)

	f.勒芬Leuven = BLeuven勒芬
	f.勒芬Leuven.SetParent(f)

	f.利尔Lier = BLier利尔
	f.利尔Lier.SetParent(f)

	f.梅赫伦Mechelen = BMechelen梅赫伦
	f.梅赫伦Mechelen.SetParent(f)

}
