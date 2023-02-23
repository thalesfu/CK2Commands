package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PalmyraCounty interface {
	feud.County
	BAlhuwaysis胡韦西斯() feud.Barony
	BAlqasr盖斯尔() feud.Barony
	BArak阿拉克() feud.Barony
	BArraml赖姆莱() feud.Barony
	BJihar吉哈尔() feud.Barony
	BKhirbat希尔拜特() feud.Barony
	BPalmyra巴尔米拉() feud.Barony
	BTiace提采() feud.Barony
}

type 巴尔米拉PalmyraCounty struct {
	feud.BaseCounty
	胡韦西斯Alhuwaysis feud.Barony
	盖斯尔Alqasr      feud.Barony
	阿拉克Arak        feud.Barony
	赖姆莱Arraml      feud.Barony
	吉哈尔Jihar       feud.Barony
	希尔拜特Khirbat    feud.Barony
	巴尔米拉Palmyra    feud.Barony
	提采Tiace        feud.Barony
}

func (c *巴尔米拉PalmyraCounty) BAlhuwaysis胡韦西斯() feud.Barony {
	return c.胡韦西斯Alhuwaysis
}

func (c *巴尔米拉PalmyraCounty) BAlqasr盖斯尔() feud.Barony {
	return c.盖斯尔Alqasr
}

func (c *巴尔米拉PalmyraCounty) BArak阿拉克() feud.Barony {
	return c.阿拉克Arak
}

func (c *巴尔米拉PalmyraCounty) BArraml赖姆莱() feud.Barony {
	return c.赖姆莱Arraml
}

func (c *巴尔米拉PalmyraCounty) BJihar吉哈尔() feud.Barony {
	return c.吉哈尔Jihar
}

func (c *巴尔米拉PalmyraCounty) BKhirbat希尔拜特() feud.Barony {
	return c.希尔拜特Khirbat
}

func (c *巴尔米拉PalmyraCounty) BPalmyra巴尔米拉() feud.Barony {
	return c.巴尔米拉Palmyra
}

func (c *巴尔米拉PalmyraCounty) BTiace提采() feud.Barony {
	return c.提采Tiace
}

var CPalmyra巴尔米拉 PalmyraCounty = &巴尔米拉PalmyraCounty{}

func init() {
	f := CPalmyra巴尔米拉.(*巴尔米拉PalmyraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "730",
		Title:     "palmyra",
		TitleName: "巴尔米拉",
		TitleCode: "c_palmyra",
		Baronies:  map[string]feud.Barony{},
	}

	f.胡韦西斯Alhuwaysis = BAlhuwaysis胡韦西斯
	f.胡韦西斯Alhuwaysis.SetParent(f)

	f.盖斯尔Alqasr = BAlqasr盖斯尔
	f.盖斯尔Alqasr.SetParent(f)

	f.阿拉克Arak = BArak阿拉克
	f.阿拉克Arak.SetParent(f)

	f.赖姆莱Arraml = BArraml赖姆莱
	f.赖姆莱Arraml.SetParent(f)

	f.吉哈尔Jihar = BJihar吉哈尔
	f.吉哈尔Jihar.SetParent(f)

	f.希尔拜特Khirbat = BKhirbat希尔拜特
	f.希尔拜特Khirbat.SetParent(f)

	f.巴尔米拉Palmyra = BPalmyra巴尔米拉
	f.巴尔米拉Palmyra.SetParent(f)

	f.提采Tiace = BTiace提采
	f.提采Tiace.SetParent(f)

}
