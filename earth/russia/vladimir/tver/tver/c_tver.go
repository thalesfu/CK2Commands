package tver

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TverCounty interface {
	feud.County
	BBurashevo布拉舍沃() feud.Barony
	BKlin克林() feud.Barony
	BLotoshino洛托希诺() feud.Barony
	BPuchka普奇卡() feud.Barony
	BRodnya罗德尼亚() feud.Barony
	BTurayevo图拉耶沃() feud.Barony
	BTver特维尔() feud.Barony
}

type 特维尔TverCounty struct {
	feud.BaseCounty
	布拉舍沃Burashevo feud.Barony
	克林Klin        feud.Barony
	洛托希诺Lotoshino feud.Barony
	普奇卡Puchka     feud.Barony
	罗德尼亚Rodnya    feud.Barony
	图拉耶沃Turayevo  feud.Barony
	特维尔Tver       feud.Barony
}

func (c *特维尔TverCounty) BBurashevo布拉舍沃() feud.Barony {
	return c.布拉舍沃Burashevo
}

func (c *特维尔TverCounty) BKlin克林() feud.Barony {
	return c.克林Klin
}

func (c *特维尔TverCounty) BLotoshino洛托希诺() feud.Barony {
	return c.洛托希诺Lotoshino
}

func (c *特维尔TverCounty) BPuchka普奇卡() feud.Barony {
	return c.普奇卡Puchka
}

func (c *特维尔TverCounty) BRodnya罗德尼亚() feud.Barony {
	return c.罗德尼亚Rodnya
}

func (c *特维尔TverCounty) BTurayevo图拉耶沃() feud.Barony {
	return c.图拉耶沃Turayevo
}

func (c *特维尔TverCounty) BTver特维尔() feud.Barony {
	return c.特维尔Tver
}

var CTver特维尔 TverCounty = &特维尔TverCounty{}

func init() {
	f := CTver特维尔.(*特维尔TverCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "570",
		Title:     "tver",
		TitleName: "特维尔",
		TitleCode: "c_tver",
		Baronies:  map[string]feud.Barony{},
	}

	f.布拉舍沃Burashevo = BBurashevo布拉舍沃
	f.布拉舍沃Burashevo.SetParent(f)

	f.克林Klin = BKlin克林
	f.克林Klin.SetParent(f)

	f.洛托希诺Lotoshino = BLotoshino洛托希诺
	f.洛托希诺Lotoshino.SetParent(f)

	f.普奇卡Puchka = BPuchka普奇卡
	f.普奇卡Puchka.SetParent(f)

	f.罗德尼亚Rodnya = BRodnya罗德尼亚
	f.罗德尼亚Rodnya.SetParent(f)

	f.图拉耶沃Turayevo = BTurayevo图拉耶沃
	f.图拉耶沃Turayevo.SetParent(f)

	f.特维尔Tver = BTver特维尔
	f.特维尔Tver.SetParent(f)

}
