package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MorayCounty interface {
	feud.County
	BCawdor考德() feud.Barony
	BElgin埃尔金() feud.Barony
	BForres福里斯() feud.Barony
	BInverness因弗内斯() feud.Barony
	BKinloss金洛斯() feud.Barony
	BLochindorb洛欣多布() feud.Barony
	BNairn奈恩() feud.Barony
	BUrquhart厄克特() feud.Barony
}

type 马里MorayCounty struct {
	feud.BaseCounty
	考德Cawdor       feud.Barony
	埃尔金Elgin       feud.Barony
	福里斯Forres      feud.Barony
	因弗内斯Inverness  feud.Barony
	金洛斯Kinloss     feud.Barony
	洛欣多布Lochindorb feud.Barony
	奈恩Nairn        feud.Barony
	厄克特Urquhart    feud.Barony
}

func (c *马里MorayCounty) BCawdor考德() feud.Barony {
	return c.考德Cawdor
}

func (c *马里MorayCounty) BElgin埃尔金() feud.Barony {
	return c.埃尔金Elgin
}

func (c *马里MorayCounty) BForres福里斯() feud.Barony {
	return c.福里斯Forres
}

func (c *马里MorayCounty) BInverness因弗内斯() feud.Barony {
	return c.因弗内斯Inverness
}

func (c *马里MorayCounty) BKinloss金洛斯() feud.Barony {
	return c.金洛斯Kinloss
}

func (c *马里MorayCounty) BLochindorb洛欣多布() feud.Barony {
	return c.洛欣多布Lochindorb
}

func (c *马里MorayCounty) BNairn奈恩() feud.Barony {
	return c.奈恩Nairn
}

func (c *马里MorayCounty) BUrquhart厄克特() feud.Barony {
	return c.厄克特Urquhart
}

var CMoray马里 MorayCounty = &马里MorayCounty{}

func init() {
	f := CMoray马里.(*马里MorayCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "40",
		Title:     "moray",
		TitleName: "马里",
		TitleCode: "c_moray",
		Baronies:  map[string]feud.Barony{},
	}

	f.考德Cawdor = BCawdor考德
	f.考德Cawdor.SetParent(f)

	f.埃尔金Elgin = BElgin埃尔金
	f.埃尔金Elgin.SetParent(f)

	f.福里斯Forres = BForres福里斯
	f.福里斯Forres.SetParent(f)

	f.因弗内斯Inverness = BInverness因弗内斯
	f.因弗内斯Inverness.SetParent(f)

	f.金洛斯Kinloss = BKinloss金洛斯
	f.金洛斯Kinloss.SetParent(f)

	f.洛欣多布Lochindorb = BLochindorb洛欣多布
	f.洛欣多布Lochindorb.SetParent(f)

	f.奈恩Nairn = BNairn奈恩
	f.奈恩Nairn.SetParent(f)

	f.厄克特Urquhart = BUrquhart厄克特
	f.厄克特Urquhart.SetParent(f)

}
