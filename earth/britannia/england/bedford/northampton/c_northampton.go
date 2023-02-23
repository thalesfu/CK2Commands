package northampton

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NorthamptonCounty interface {
	feud.County
	BCambridge剑桥() feud.Barony
	BCrowland克罗兰() feud.Barony
	BHuntingdon亨廷登() feud.Barony
	BKettering凯特林() feud.Barony
	BNorthampton北安普顿() feud.Barony
	BPeterborough彼得伯勒() feud.Barony
	BRamsey拉姆西() feud.Barony
	BRockingham罗金厄姆() feud.Barony
}

type 北安普顿NorthamptonCounty struct {
	feud.BaseCounty
	剑桥Cambridge      feud.Barony
	克罗兰Crowland      feud.Barony
	亨廷登Huntingdon    feud.Barony
	凯特林Kettering     feud.Barony
	北安普顿Northampton  feud.Barony
	彼得伯勒Peterborough feud.Barony
	拉姆西Ramsey        feud.Barony
	罗金厄姆Rockingham   feud.Barony
}

func (c *北安普顿NorthamptonCounty) BCambridge剑桥() feud.Barony {
	return c.剑桥Cambridge
}

func (c *北安普顿NorthamptonCounty) BCrowland克罗兰() feud.Barony {
	return c.克罗兰Crowland
}

func (c *北安普顿NorthamptonCounty) BHuntingdon亨廷登() feud.Barony {
	return c.亨廷登Huntingdon
}

func (c *北安普顿NorthamptonCounty) BKettering凯特林() feud.Barony {
	return c.凯特林Kettering
}

func (c *北安普顿NorthamptonCounty) BNorthampton北安普顿() feud.Barony {
	return c.北安普顿Northampton
}

func (c *北安普顿NorthamptonCounty) BPeterborough彼得伯勒() feud.Barony {
	return c.彼得伯勒Peterborough
}

func (c *北安普顿NorthamptonCounty) BRamsey拉姆西() feud.Barony {
	return c.拉姆西Ramsey
}

func (c *北安普顿NorthamptonCounty) BRockingham罗金厄姆() feud.Barony {
	return c.罗金厄姆Rockingham
}

var CNorthampton北安普顿 NorthamptonCounty = &北安普顿NorthamptonCounty{}

func init() {
	f := CNorthampton北安普顿.(*北安普顿NorthamptonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "68",
		Title:     "northampton",
		TitleName: "北安普顿",
		TitleCode: "c_northampton",
		Baronies:  map[string]feud.Barony{},
	}

	f.剑桥Cambridge = BCambridge剑桥
	f.剑桥Cambridge.SetParent(f)

	f.克罗兰Crowland = BCrowland克罗兰
	f.克罗兰Crowland.SetParent(f)

	f.亨廷登Huntingdon = BHuntingdon亨廷登
	f.亨廷登Huntingdon.SetParent(f)

	f.凯特林Kettering = BKettering凯特林
	f.凯特林Kettering.SetParent(f)

	f.北安普顿Northampton = BNorthampton北安普顿
	f.北安普顿Northampton.SetParent(f)

	f.彼得伯勒Peterborough = BPeterborough彼得伯勒
	f.彼得伯勒Peterborough.SetParent(f)

	f.拉姆西Ramsey = BRamsey拉姆西
	f.拉姆西Ramsey.SetParent(f)

	f.罗金厄姆Rockingham = BRockingham罗金厄姆
	f.罗金厄姆Rockingham.SetParent(f)

}
