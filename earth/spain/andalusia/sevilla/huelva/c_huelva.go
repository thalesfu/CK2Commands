package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HuelvaCounty interface {
	feud.County
	BAljaraque阿尔哈拉克() feud.Barony
	BAyamonte阿亚蒙特() feud.Barony
	BGibraleon希夫拉莱昂() feud.Barony
	BHuelva韦尔瓦() feud.Barony
	BLepe莱佩() feud.Barony
	BMoguer莫格尔() feud.Barony
	BPortichuelo波蒂舒埃洛() feud.Barony
}

type c_huelvaHuelvaCounty struct {
	feud.BaseCounty
	阿尔哈拉克Aljaraque   feud.Barony
	阿亚蒙特Ayamonte     feud.Barony
	希夫拉莱昂Gibraleon   feud.Barony
	韦尔瓦Huelva        feud.Barony
	莱佩Lepe           feud.Barony
	莫格尔Moguer        feud.Barony
	波蒂舒埃洛Portichuelo feud.Barony
}

func (c *c_huelvaHuelvaCounty) BAljaraque阿尔哈拉克() feud.Barony {
	return c.阿尔哈拉克Aljaraque
}

func (c *c_huelvaHuelvaCounty) BAyamonte阿亚蒙特() feud.Barony {
	return c.阿亚蒙特Ayamonte
}

func (c *c_huelvaHuelvaCounty) BGibraleon希夫拉莱昂() feud.Barony {
	return c.希夫拉莱昂Gibraleon
}

func (c *c_huelvaHuelvaCounty) BHuelva韦尔瓦() feud.Barony {
	return c.韦尔瓦Huelva
}

func (c *c_huelvaHuelvaCounty) BLepe莱佩() feud.Barony {
	return c.莱佩Lepe
}

func (c *c_huelvaHuelvaCounty) BMoguer莫格尔() feud.Barony {
	return c.莫格尔Moguer
}

func (c *c_huelvaHuelvaCounty) BPortichuelo波蒂舒埃洛() feud.Barony {
	return c.波蒂舒埃洛Portichuelo
}

var CHuelvac_huelva HuelvaCounty = &c_huelvaHuelvaCounty{}

func init() {
	f := CHuelvac_huelva.(*c_huelvaHuelvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2014",
		Title:     "huelva",
		TitleName: "c_huelva",
		TitleCode: "c_huelva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔哈拉克Aljaraque = BAljaraque阿尔哈拉克
	f.阿尔哈拉克Aljaraque.SetParent(f)

	f.阿亚蒙特Ayamonte = BAyamonte阿亚蒙特
	f.阿亚蒙特Ayamonte.SetParent(f)

	f.希夫拉莱昂Gibraleon = BGibraleon希夫拉莱昂
	f.希夫拉莱昂Gibraleon.SetParent(f)

	f.韦尔瓦Huelva = BHuelva韦尔瓦
	f.韦尔瓦Huelva.SetParent(f)

	f.莱佩Lepe = BLepe莱佩
	f.莱佩Lepe.SetParent(f)

	f.莫格尔Moguer = BMoguer莫格尔
	f.莫格尔Moguer.SetParent(f)

	f.波蒂舒埃洛Portichuelo = BPortichuelo波蒂舒埃洛
	f.波蒂舒埃洛Portichuelo.SetParent(f)

}
