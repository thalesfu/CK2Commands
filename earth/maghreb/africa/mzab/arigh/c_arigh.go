package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArighCounty interface {
	feud.County
	BAjlu埃吉卢() feud.Barony
	BArigh阿里格() feud.Barony
	BBouchaoui布沙维() feud.Barony
	BMgadid姆加迪() feud.Barony
	BRabta拉卜塔() feud.Barony
	BTimoktene蒂莫克滕() feud.Barony
	BZemours泽穆尔() feud.Barony
}

type 阿里格ArighCounty struct {
	feud.BaseCounty
	埃吉卢Ajlu       feud.Barony
	阿里格Arigh      feud.Barony
	布沙维Bouchaoui  feud.Barony
	姆加迪Mgadid     feud.Barony
	拉卜塔Rabta      feud.Barony
	蒂莫克滕Timoktene feud.Barony
	泽穆尔Zemours    feud.Barony
}

func (c *阿里格ArighCounty) BAjlu埃吉卢() feud.Barony {
	return c.埃吉卢Ajlu
}

func (c *阿里格ArighCounty) BArigh阿里格() feud.Barony {
	return c.阿里格Arigh
}

func (c *阿里格ArighCounty) BBouchaoui布沙维() feud.Barony {
	return c.布沙维Bouchaoui
}

func (c *阿里格ArighCounty) BMgadid姆加迪() feud.Barony {
	return c.姆加迪Mgadid
}

func (c *阿里格ArighCounty) BRabta拉卜塔() feud.Barony {
	return c.拉卜塔Rabta
}

func (c *阿里格ArighCounty) BTimoktene蒂莫克滕() feud.Barony {
	return c.蒂莫克滕Timoktene
}

func (c *阿里格ArighCounty) BZemours泽穆尔() feud.Barony {
	return c.泽穆尔Zemours
}

var CArigh阿里格 ArighCounty = &阿里格ArighCounty{}

func init() {
	f := CArigh阿里格.(*阿里格ArighCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1727",
		Title:     "arigh",
		TitleName: "阿里格",
		TitleCode: "c_arigh",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃吉卢Ajlu = BAjlu埃吉卢
	f.埃吉卢Ajlu.SetParent(f)

	f.阿里格Arigh = BArigh阿里格
	f.阿里格Arigh.SetParent(f)

	f.布沙维Bouchaoui = BBouchaoui布沙维
	f.布沙维Bouchaoui.SetParent(f)

	f.姆加迪Mgadid = BMgadid姆加迪
	f.姆加迪Mgadid.SetParent(f)

	f.拉卜塔Rabta = BRabta拉卜塔
	f.拉卜塔Rabta.SetParent(f)

	f.蒂莫克滕Timoktene = BTimoktene蒂莫克滕
	f.蒂莫克滕Timoktene.SetParent(f)

	f.泽穆尔Zemours = BZemours泽穆尔
	f.泽穆尔Zemours.SetParent(f)

}
