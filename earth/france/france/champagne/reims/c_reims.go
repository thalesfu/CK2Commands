package reims

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ReimsCounty interface {
	feud.County
	BAttigny阿蒂尼() feud.Barony
	BChalons沙隆() feud.Barony
	BChatillon沙蒂永() feud.Barony
	BDampierre当皮耶尔() feud.Barony
	BEpernay埃佩尔奈() feud.Barony
	BReims兰斯() feud.Barony
	BRethel勒泰勒() feud.Barony
	BRoucy鲁西() feud.Barony
}

type 兰斯ReimsCounty struct {
	feud.BaseCounty
	阿蒂尼Attigny    feud.Barony
	沙隆Chalons     feud.Barony
	沙蒂永Chatillon  feud.Barony
	当皮耶尔Dampierre feud.Barony
	埃佩尔奈Epernay   feud.Barony
	兰斯Reims       feud.Barony
	勒泰勒Rethel     feud.Barony
	鲁西Roucy       feud.Barony
}

func (c *兰斯ReimsCounty) BAttigny阿蒂尼() feud.Barony {
	return c.阿蒂尼Attigny
}

func (c *兰斯ReimsCounty) BChalons沙隆() feud.Barony {
	return c.沙隆Chalons
}

func (c *兰斯ReimsCounty) BChatillon沙蒂永() feud.Barony {
	return c.沙蒂永Chatillon
}

func (c *兰斯ReimsCounty) BDampierre当皮耶尔() feud.Barony {
	return c.当皮耶尔Dampierre
}

func (c *兰斯ReimsCounty) BEpernay埃佩尔奈() feud.Barony {
	return c.埃佩尔奈Epernay
}

func (c *兰斯ReimsCounty) BReims兰斯() feud.Barony {
	return c.兰斯Reims
}

func (c *兰斯ReimsCounty) BRethel勒泰勒() feud.Barony {
	return c.勒泰勒Rethel
}

func (c *兰斯ReimsCounty) BRoucy鲁西() feud.Barony {
	return c.鲁西Roucy
}

var CReims兰斯 ReimsCounty = &兰斯ReimsCounty{}

func init() {
	f := CReims兰斯.(*兰斯ReimsCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "114",
		Title:     "reims",
		TitleName: "兰斯",
		TitleCode: "c_reims",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿蒂尼Attigny = BAttigny阿蒂尼
	f.阿蒂尼Attigny.SetParent(f)

	f.沙隆Chalons = BChalons沙隆
	f.沙隆Chalons.SetParent(f)

	f.沙蒂永Chatillon = BChatillon沙蒂永
	f.沙蒂永Chatillon.SetParent(f)

	f.当皮耶尔Dampierre = BDampierre当皮耶尔
	f.当皮耶尔Dampierre.SetParent(f)

	f.埃佩尔奈Epernay = BEpernay埃佩尔奈
	f.埃佩尔奈Epernay.SetParent(f)

	f.兰斯Reims = BReims兰斯
	f.兰斯Reims.SetParent(f)

	f.勒泰勒Rethel = BRethel勒泰勒
	f.勒泰勒Rethel.SetParent(f)

	f.鲁西Roucy = BRoucy鲁西
	f.鲁西Roucy.SetParent(f)

}
