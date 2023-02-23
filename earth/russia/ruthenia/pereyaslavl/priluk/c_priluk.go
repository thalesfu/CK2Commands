package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PrilukCounty interface {
	feud.County
	BBobrovytsia博布罗维齐亚() feud.Barony
	BIchina伊奇纳() feud.Barony
	BNizhyn涅任() feud.Barony
	BNosivka诺维斯卡() feud.Barony
	BPriluk普里卢基() feud.Barony
	BRomny罗姆内() feud.Barony
	BVarva瓦尔瓦() feud.Barony
}

type 普里卢基PrilukCounty struct {
	feud.BaseCounty
	博布罗维齐亚Bobrovytsia feud.Barony
	伊奇纳Ichina         feud.Barony
	涅任Nizhyn          feud.Barony
	诺维斯卡Nosivka       feud.Barony
	普里卢基Priluk        feud.Barony
	罗姆内Romny          feud.Barony
	瓦尔瓦Varva          feud.Barony
}

func (c *普里卢基PrilukCounty) BBobrovytsia博布罗维齐亚() feud.Barony {
	return c.博布罗维齐亚Bobrovytsia
}

func (c *普里卢基PrilukCounty) BIchina伊奇纳() feud.Barony {
	return c.伊奇纳Ichina
}

func (c *普里卢基PrilukCounty) BNizhyn涅任() feud.Barony {
	return c.涅任Nizhyn
}

func (c *普里卢基PrilukCounty) BNosivka诺维斯卡() feud.Barony {
	return c.诺维斯卡Nosivka
}

func (c *普里卢基PrilukCounty) BPriluk普里卢基() feud.Barony {
	return c.普里卢基Priluk
}

func (c *普里卢基PrilukCounty) BRomny罗姆内() feud.Barony {
	return c.罗姆内Romny
}

func (c *普里卢基PrilukCounty) BVarva瓦尔瓦() feud.Barony {
	return c.瓦尔瓦Varva
}

var CPriluk普里卢基 PrilukCounty = &普里卢基PrilukCounty{}

func init() {
	f := CPriluk普里卢基.(*普里卢基PrilukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1675",
		Title:     "priluk",
		TitleName: "普里卢基",
		TitleCode: "c_priluk",
		Baronies:  map[string]feud.Barony{},
	}

	f.博布罗维齐亚Bobrovytsia = BBobrovytsia博布罗维齐亚
	f.博布罗维齐亚Bobrovytsia.SetParent(f)

	f.伊奇纳Ichina = BIchina伊奇纳
	f.伊奇纳Ichina.SetParent(f)

	f.涅任Nizhyn = BNizhyn涅任
	f.涅任Nizhyn.SetParent(f)

	f.诺维斯卡Nosivka = BNosivka诺维斯卡
	f.诺维斯卡Nosivka.SetParent(f)

	f.普里卢基Priluk = BPriluk普里卢基
	f.普里卢基Priluk.SetParent(f)

	f.罗姆内Romny = BRomny罗姆内
	f.罗姆内Romny.SetParent(f)

	f.瓦尔瓦Varva = BVarva瓦尔瓦
	f.瓦尔瓦Varva.SetParent(f)

}
