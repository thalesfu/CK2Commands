package amman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmmanCounty interface {
	feud.County
	BAmman安曼() feud.Barony
	BDhiban济班() feud.Barony
	BFuheis富海斯() feud.Barony
	BJerash杰拉什() feud.Barony
	BMahis马希斯() feud.Barony
	BSalt萨勒特() feud.Barony
}

type 安曼AmmanCounty struct {
	feud.BaseCounty
	安曼Amman   feud.Barony
	济班Dhiban  feud.Barony
	富海斯Fuheis feud.Barony
	杰拉什Jerash feud.Barony
	马希斯Mahis  feud.Barony
	萨勒特Salt   feud.Barony
}

func (c *安曼AmmanCounty) BAmman安曼() feud.Barony {
	return c.安曼Amman
}

func (c *安曼AmmanCounty) BDhiban济班() feud.Barony {
	return c.济班Dhiban
}

func (c *安曼AmmanCounty) BFuheis富海斯() feud.Barony {
	return c.富海斯Fuheis
}

func (c *安曼AmmanCounty) BJerash杰拉什() feud.Barony {
	return c.杰拉什Jerash
}

func (c *安曼AmmanCounty) BMahis马希斯() feud.Barony {
	return c.马希斯Mahis
}

func (c *安曼AmmanCounty) BSalt萨勒特() feud.Barony {
	return c.萨勒特Salt
}

var CAmman安曼 AmmanCounty = &安曼AmmanCounty{}

func init() {
	f := CAmman安曼.(*安曼AmmanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "724",
		Title:     "amman",
		TitleName: "安曼",
		TitleCode: "c_amman",
		Baronies:  map[string]feud.Barony{},
	}

	f.安曼Amman = BAmman安曼
	f.安曼Amman.SetParent(f)

	f.济班Dhiban = BDhiban济班
	f.济班Dhiban.SetParent(f)

	f.富海斯Fuheis = BFuheis富海斯
	f.富海斯Fuheis.SetParent(f)

	f.杰拉什Jerash = BJerash杰拉什
	f.杰拉什Jerash.SetParent(f)

	f.马希斯Mahis = BMahis马希斯
	f.马希斯Mahis.SetParent(f)

	f.萨勒特Salt = BSalt萨勒特
	f.萨勒特Salt.SetParent(f)

}
