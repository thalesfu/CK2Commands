package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarbalaCounty interface {
	feud.County
	BAinaltamur艾因塔姆尔() feud.Barony
	BHamzah哈姆扎赫() feud.Barony
	BHindiya辛蒂亚() feud.Barony
	BKarbala卡尔巴拉() feud.Barony
	BOfak奥法可() feud.Barony
	BQisair奇瑟() feud.Barony
	BShamiyah沙弥亚() feud.Barony
	BUkhaidir乌海迪尔() feud.Barony
}

type 卡尔巴拉KarbalaCounty struct {
	feud.BaseCounty
	艾因塔姆尔Ainaltamur feud.Barony
	哈姆扎赫Hamzah      feud.Barony
	辛蒂亚Hindiya      feud.Barony
	卡尔巴拉Karbala     feud.Barony
	奥法可Ofak         feud.Barony
	奇瑟Qisair        feud.Barony
	沙弥亚Shamiyah     feud.Barony
	乌海迪尔Ukhaidir    feud.Barony
}

func (c *卡尔巴拉KarbalaCounty) BAinaltamur艾因塔姆尔() feud.Barony {
	return c.艾因塔姆尔Ainaltamur
}

func (c *卡尔巴拉KarbalaCounty) BHamzah哈姆扎赫() feud.Barony {
	return c.哈姆扎赫Hamzah
}

func (c *卡尔巴拉KarbalaCounty) BHindiya辛蒂亚() feud.Barony {
	return c.辛蒂亚Hindiya
}

func (c *卡尔巴拉KarbalaCounty) BKarbala卡尔巴拉() feud.Barony {
	return c.卡尔巴拉Karbala
}

func (c *卡尔巴拉KarbalaCounty) BOfak奥法可() feud.Barony {
	return c.奥法可Ofak
}

func (c *卡尔巴拉KarbalaCounty) BQisair奇瑟() feud.Barony {
	return c.奇瑟Qisair
}

func (c *卡尔巴拉KarbalaCounty) BShamiyah沙弥亚() feud.Barony {
	return c.沙弥亚Shamiyah
}

func (c *卡尔巴拉KarbalaCounty) BUkhaidir乌海迪尔() feud.Barony {
	return c.乌海迪尔Ukhaidir
}

var CKarbala卡尔巴拉 KarbalaCounty = &卡尔巴拉KarbalaCounty{}

func init() {
	f := CKarbala卡尔巴拉.(*卡尔巴拉KarbalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "694",
		Title:     "karbala",
		TitleName: "卡尔巴拉",
		TitleCode: "c_karbala",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因塔姆尔Ainaltamur = BAinaltamur艾因塔姆尔
	f.艾因塔姆尔Ainaltamur.SetParent(f)

	f.哈姆扎赫Hamzah = BHamzah哈姆扎赫
	f.哈姆扎赫Hamzah.SetParent(f)

	f.辛蒂亚Hindiya = BHindiya辛蒂亚
	f.辛蒂亚Hindiya.SetParent(f)

	f.卡尔巴拉Karbala = BKarbala卡尔巴拉
	f.卡尔巴拉Karbala.SetParent(f)

	f.奥法可Ofak = BOfak奥法可
	f.奥法可Ofak.SetParent(f)

	f.奇瑟Qisair = BQisair奇瑟
	f.奇瑟Qisair.SetParent(f)

	f.沙弥亚Shamiyah = BShamiyah沙弥亚
	f.沙弥亚Shamiyah.SetParent(f)

	f.乌海迪尔Ukhaidir = BUkhaidir乌海迪尔
	f.乌海迪尔Ukhaidir.SetParent(f)

}
