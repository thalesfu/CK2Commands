package gemer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GemerCounty interface {
	feud.County
	BBalassagyarmat包洛绍焦尔毛特() feud.Barony
	BDobsina多布希瑙() feud.Barony
	BGomor格默尔() feud.Barony
	BJolsva约尔什沃() feud.Barony
	BLosonc洛雄茨() feud.Barony
	BNagyroce大勒采() feud.Barony
	BNyustya纽什焦() feud.Barony
	BRozsnyo罗日纽() feud.Barony
}

type 格梅尔GemerCounty struct {
	feud.BaseCounty
	包洛绍焦尔毛特Balassagyarmat feud.Barony
	多布希瑙Dobsina           feud.Barony
	格默尔Gomor              feud.Barony
	约尔什沃Jolsva            feud.Barony
	洛雄茨Losonc             feud.Barony
	大勒采Nagyroce           feud.Barony
	纽什焦Nyustya            feud.Barony
	罗日纽Rozsnyo            feud.Barony
}

func (c *格梅尔GemerCounty) BBalassagyarmat包洛绍焦尔毛特() feud.Barony {
	return c.包洛绍焦尔毛特Balassagyarmat
}

func (c *格梅尔GemerCounty) BDobsina多布希瑙() feud.Barony {
	return c.多布希瑙Dobsina
}

func (c *格梅尔GemerCounty) BGomor格默尔() feud.Barony {
	return c.格默尔Gomor
}

func (c *格梅尔GemerCounty) BJolsva约尔什沃() feud.Barony {
	return c.约尔什沃Jolsva
}

func (c *格梅尔GemerCounty) BLosonc洛雄茨() feud.Barony {
	return c.洛雄茨Losonc
}

func (c *格梅尔GemerCounty) BNagyroce大勒采() feud.Barony {
	return c.大勒采Nagyroce
}

func (c *格梅尔GemerCounty) BNyustya纽什焦() feud.Barony {
	return c.纽什焦Nyustya
}

func (c *格梅尔GemerCounty) BRozsnyo罗日纽() feud.Barony {
	return c.罗日纽Rozsnyo
}

var CGemer格梅尔 GemerCounty = &格梅尔GemerCounty{}

func init() {
	f := CGemer格梅尔.(*格梅尔GemerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "524",
		Title:     "gemer",
		TitleName: "格梅尔",
		TitleCode: "c_gemer",
		Baronies:  map[string]feud.Barony{},
	}

	f.包洛绍焦尔毛特Balassagyarmat = BBalassagyarmat包洛绍焦尔毛特
	f.包洛绍焦尔毛特Balassagyarmat.SetParent(f)

	f.多布希瑙Dobsina = BDobsina多布希瑙
	f.多布希瑙Dobsina.SetParent(f)

	f.格默尔Gomor = BGomor格默尔
	f.格默尔Gomor.SetParent(f)

	f.约尔什沃Jolsva = BJolsva约尔什沃
	f.约尔什沃Jolsva.SetParent(f)

	f.洛雄茨Losonc = BLosonc洛雄茨
	f.洛雄茨Losonc.SetParent(f)

	f.大勒采Nagyroce = BNagyroce大勒采
	f.大勒采Nagyroce.SetParent(f)

	f.纽什焦Nyustya = BNyustya纽什焦
	f.纽什焦Nyustya.SetParent(f)

	f.罗日纽Rozsnyo = BRozsnyo罗日纽
	f.罗日纽Rozsnyo.SetParent(f)

}
