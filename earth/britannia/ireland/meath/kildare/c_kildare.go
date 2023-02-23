package kildare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KildareCounty interface {
	feud.County
	BAthlone阿斯隆() feud.Barony
	BClonard克洛纳德() feud.Barony
	BDurrow达罗() feud.Barony
	BKildare基尔代尔() feud.Barony
	BKnockaulin康斯喀林() feud.Barony
	BMaynooth梅努斯() feud.Barony
	BRathangan拉桑根() feud.Barony
}

type 基尔代尔KildareCounty struct {
	feud.BaseCounty
	阿斯隆Athlone     feud.Barony
	克洛纳德Clonard    feud.Barony
	达罗Durrow       feud.Barony
	基尔代尔Kildare    feud.Barony
	康斯喀林Knockaulin feud.Barony
	梅努斯Maynooth    feud.Barony
	拉桑根Rathangan   feud.Barony
}

func (c *基尔代尔KildareCounty) BAthlone阿斯隆() feud.Barony {
	return c.阿斯隆Athlone
}

func (c *基尔代尔KildareCounty) BClonard克洛纳德() feud.Barony {
	return c.克洛纳德Clonard
}

func (c *基尔代尔KildareCounty) BDurrow达罗() feud.Barony {
	return c.达罗Durrow
}

func (c *基尔代尔KildareCounty) BKildare基尔代尔() feud.Barony {
	return c.基尔代尔Kildare
}

func (c *基尔代尔KildareCounty) BKnockaulin康斯喀林() feud.Barony {
	return c.康斯喀林Knockaulin
}

func (c *基尔代尔KildareCounty) BMaynooth梅努斯() feud.Barony {
	return c.梅努斯Maynooth
}

func (c *基尔代尔KildareCounty) BRathangan拉桑根() feud.Barony {
	return c.拉桑根Rathangan
}

var CKildare基尔代尔 KildareCounty = &基尔代尔KildareCounty{}

func init() {
	f := CKildare基尔代尔.(*基尔代尔KildareCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "10",
		Title:     "kildare",
		TitleName: "基尔代尔",
		TitleCode: "c_kildare",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯隆Athlone = BAthlone阿斯隆
	f.阿斯隆Athlone.SetParent(f)

	f.克洛纳德Clonard = BClonard克洛纳德
	f.克洛纳德Clonard.SetParent(f)

	f.达罗Durrow = BDurrow达罗
	f.达罗Durrow.SetParent(f)

	f.基尔代尔Kildare = BKildare基尔代尔
	f.基尔代尔Kildare.SetParent(f)

	f.康斯喀林Knockaulin = BKnockaulin康斯喀林
	f.康斯喀林Knockaulin.SetParent(f)

	f.梅努斯Maynooth = BMaynooth梅努斯
	f.梅努斯Maynooth.SetParent(f)

	f.拉桑根Rathangan = BRathangan拉桑根
	f.拉桑根Rathangan.SetParent(f)

}
