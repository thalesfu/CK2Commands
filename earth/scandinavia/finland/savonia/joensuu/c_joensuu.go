package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JoensuuCounty interface {
	feud.County
	BEno埃诺() feud.Barony
	BIisalmi伊萨尔米() feud.Barony
	BIlomantsi伊洛曼齐() feud.Barony
	BJoensuu约恩苏() feud.Barony
	BKontiolahti孔蒂奥拉赫蒂() feud.Barony
	BLiperi利佩里() feud.Barony
	BPyhaselka皮海湖() feud.Barony
}

type 约恩苏JoensuuCounty struct {
	feud.BaseCounty
	埃诺Eno             feud.Barony
	伊萨尔米Iisalmi       feud.Barony
	伊洛曼齐Ilomantsi     feud.Barony
	约恩苏Joensuu        feud.Barony
	孔蒂奥拉赫蒂Kontiolahti feud.Barony
	利佩里Liperi         feud.Barony
	皮海湖Pyhaselka      feud.Barony
}

func (c *约恩苏JoensuuCounty) BEno埃诺() feud.Barony {
	return c.埃诺Eno
}

func (c *约恩苏JoensuuCounty) BIisalmi伊萨尔米() feud.Barony {
	return c.伊萨尔米Iisalmi
}

func (c *约恩苏JoensuuCounty) BIlomantsi伊洛曼齐() feud.Barony {
	return c.伊洛曼齐Ilomantsi
}

func (c *约恩苏JoensuuCounty) BJoensuu约恩苏() feud.Barony {
	return c.约恩苏Joensuu
}

func (c *约恩苏JoensuuCounty) BKontiolahti孔蒂奥拉赫蒂() feud.Barony {
	return c.孔蒂奥拉赫蒂Kontiolahti
}

func (c *约恩苏JoensuuCounty) BLiperi利佩里() feud.Barony {
	return c.利佩里Liperi
}

func (c *约恩苏JoensuuCounty) BPyhaselka皮海湖() feud.Barony {
	return c.皮海湖Pyhaselka
}

var CJoensuu约恩苏 JoensuuCounty = &约恩苏JoensuuCounty{}

func init() {
	f := CJoensuu约恩苏.(*约恩苏JoensuuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1599",
		Title:     "joensuu",
		TitleName: "约恩苏",
		TitleCode: "c_joensuu",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃诺Eno = BEno埃诺
	f.埃诺Eno.SetParent(f)

	f.伊萨尔米Iisalmi = BIisalmi伊萨尔米
	f.伊萨尔米Iisalmi.SetParent(f)

	f.伊洛曼齐Ilomantsi = BIlomantsi伊洛曼齐
	f.伊洛曼齐Ilomantsi.SetParent(f)

	f.约恩苏Joensuu = BJoensuu约恩苏
	f.约恩苏Joensuu.SetParent(f)

	f.孔蒂奥拉赫蒂Kontiolahti = BKontiolahti孔蒂奥拉赫蒂
	f.孔蒂奥拉赫蒂Kontiolahti.SetParent(f)

	f.利佩里Liperi = BLiperi利佩里
	f.利佩里Liperi.SetParent(f)

	f.皮海湖Pyhaselka = BPyhaselka皮海湖
	f.皮海湖Pyhaselka.SetParent(f)

}
