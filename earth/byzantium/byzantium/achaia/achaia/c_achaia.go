package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AchaiaCounty interface {
	feud.County
	BAkova阿科瓦() feud.Barony
	BAndravida安兹拉维扎() feud.Barony
	BChalandritza哈兰兹里察() feud.Barony
	BGeraki耶拉基() feud.Barony
	BKalavryta卡拉夫律塔() feud.Barony
	BKarditza卡尔季察() feud.Barony
	BPatras帕特拉斯() feud.Barony
	BPyrgos皮尔戈斯() feud.Barony
}

type 亚该亚AchaiaCounty struct {
	feud.BaseCounty
	阿科瓦Akova          feud.Barony
	安兹拉维扎Andravida    feud.Barony
	哈兰兹里察Chalandritza feud.Barony
	耶拉基Geraki         feud.Barony
	卡拉夫律塔Kalavryta    feud.Barony
	卡尔季察Karditza      feud.Barony
	帕特拉斯Patras        feud.Barony
	皮尔戈斯Pyrgos        feud.Barony
}

func (c *亚该亚AchaiaCounty) BAkova阿科瓦() feud.Barony {
	return c.阿科瓦Akova
}

func (c *亚该亚AchaiaCounty) BAndravida安兹拉维扎() feud.Barony {
	return c.安兹拉维扎Andravida
}

func (c *亚该亚AchaiaCounty) BChalandritza哈兰兹里察() feud.Barony {
	return c.哈兰兹里察Chalandritza
}

func (c *亚该亚AchaiaCounty) BGeraki耶拉基() feud.Barony {
	return c.耶拉基Geraki
}

func (c *亚该亚AchaiaCounty) BKalavryta卡拉夫律塔() feud.Barony {
	return c.卡拉夫律塔Kalavryta
}

func (c *亚该亚AchaiaCounty) BKarditza卡尔季察() feud.Barony {
	return c.卡尔季察Karditza
}

func (c *亚该亚AchaiaCounty) BPatras帕特拉斯() feud.Barony {
	return c.帕特拉斯Patras
}

func (c *亚该亚AchaiaCounty) BPyrgos皮尔戈斯() feud.Barony {
	return c.皮尔戈斯Pyrgos
}

var CAchaia亚该亚 AchaiaCounty = &亚该亚AchaiaCounty{}

func init() {
	f := CAchaia亚该亚.(*亚该亚AchaiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "476",
		Title:     "achaia",
		TitleName: "亚该亚",
		TitleCode: "c_achaia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿科瓦Akova = BAkova阿科瓦
	f.阿科瓦Akova.SetParent(f)

	f.安兹拉维扎Andravida = BAndravida安兹拉维扎
	f.安兹拉维扎Andravida.SetParent(f)

	f.哈兰兹里察Chalandritza = BChalandritza哈兰兹里察
	f.哈兰兹里察Chalandritza.SetParent(f)

	f.耶拉基Geraki = BGeraki耶拉基
	f.耶拉基Geraki.SetParent(f)

	f.卡拉夫律塔Kalavryta = BKalavryta卡拉夫律塔
	f.卡拉夫律塔Kalavryta.SetParent(f)

	f.卡尔季察Karditza = BKarditza卡尔季察
	f.卡尔季察Karditza.SetParent(f)

	f.帕特拉斯Patras = BPatras帕特拉斯
	f.帕特拉斯Patras.SetParent(f)

	f.皮尔戈斯Pyrgos = BPyrgos皮尔戈斯
	f.皮尔戈斯Pyrgos.SetParent(f)

}
