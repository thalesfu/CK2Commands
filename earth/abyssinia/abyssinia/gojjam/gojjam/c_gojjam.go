package gojjam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GojjamCounty interface {
	feud.County
	BBaremma巴雷马() feud.Barony
	BEliyas伊利亚斯() feud.Barony
	BGojjam戈贾姆() feud.Barony
	BMankorar曼科拉尔() feud.Barony
	BYegeleka耶格勒卡() feud.Barony
	BYejube耶朱贝() feud.Barony
}

type 戈贾姆GojjamCounty struct {
	feud.BaseCounty
	巴雷马Baremma   feud.Barony
	伊利亚斯Eliyas   feud.Barony
	戈贾姆Gojjam    feud.Barony
	曼科拉尔Mankorar feud.Barony
	耶格勒卡Yegeleka feud.Barony
	耶朱贝Yejube    feud.Barony
}

func (c *戈贾姆GojjamCounty) BBaremma巴雷马() feud.Barony {
	return c.巴雷马Baremma
}

func (c *戈贾姆GojjamCounty) BEliyas伊利亚斯() feud.Barony {
	return c.伊利亚斯Eliyas
}

func (c *戈贾姆GojjamCounty) BGojjam戈贾姆() feud.Barony {
	return c.戈贾姆Gojjam
}

func (c *戈贾姆GojjamCounty) BMankorar曼科拉尔() feud.Barony {
	return c.曼科拉尔Mankorar
}

func (c *戈贾姆GojjamCounty) BYegeleka耶格勒卡() feud.Barony {
	return c.耶格勒卡Yegeleka
}

func (c *戈贾姆GojjamCounty) BYejube耶朱贝() feud.Barony {
	return c.耶朱贝Yejube
}

var CGojjam戈贾姆 GojjamCounty = &戈贾姆GojjamCounty{}

func init() {
	f := CGojjam戈贾姆.(*戈贾姆GojjamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1280",
		Title:     "gojjam",
		TitleName: "戈贾姆",
		TitleCode: "c_gojjam",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴雷马Baremma = BBaremma巴雷马
	f.巴雷马Baremma.SetParent(f)

	f.伊利亚斯Eliyas = BEliyas伊利亚斯
	f.伊利亚斯Eliyas.SetParent(f)

	f.戈贾姆Gojjam = BGojjam戈贾姆
	f.戈贾姆Gojjam.SetParent(f)

	f.曼科拉尔Mankorar = BMankorar曼科拉尔
	f.曼科拉尔Mankorar.SetParent(f)

	f.耶格勒卡Yegeleka = BYegeleka耶格勒卡
	f.耶格勒卡Yegeleka.SetParent(f)

	f.耶朱贝Yejube = BYejube耶朱贝
	f.耶朱贝Yejube.SetParent(f)

}
