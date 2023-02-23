package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NassauCounty interface {
	feud.County
	BFalkenstein法尔肯施泰因() feud.Barony
	BFritzlar弗里茨拉尔() feud.Barony
	BFulda富尔达() feud.Barony
	BHaerulfisfeld海鲁尔菲斯费尔德() feud.Barony
	BHersfeld赫斯费尔德() feud.Barony
	BIsenburg伊森堡() feud.Barony
	BNassau拿骚() feud.Barony
}

type 弗里茨拉尔NassauCounty struct {
	feud.BaseCounty
	法尔肯施泰因Falkenstein     feud.Barony
	弗里茨拉尔Fritzlar         feud.Barony
	富尔达Fulda              feud.Barony
	海鲁尔菲斯费尔德Haerulfisfeld feud.Barony
	赫斯费尔德Hersfeld         feud.Barony
	伊森堡Isenburg           feud.Barony
	拿骚Nassau              feud.Barony
}

func (c *弗里茨拉尔NassauCounty) BFalkenstein法尔肯施泰因() feud.Barony {
	return c.法尔肯施泰因Falkenstein
}

func (c *弗里茨拉尔NassauCounty) BFritzlar弗里茨拉尔() feud.Barony {
	return c.弗里茨拉尔Fritzlar
}

func (c *弗里茨拉尔NassauCounty) BFulda富尔达() feud.Barony {
	return c.富尔达Fulda
}

func (c *弗里茨拉尔NassauCounty) BHaerulfisfeld海鲁尔菲斯费尔德() feud.Barony {
	return c.海鲁尔菲斯费尔德Haerulfisfeld
}

func (c *弗里茨拉尔NassauCounty) BHersfeld赫斯费尔德() feud.Barony {
	return c.赫斯费尔德Hersfeld
}

func (c *弗里茨拉尔NassauCounty) BIsenburg伊森堡() feud.Barony {
	return c.伊森堡Isenburg
}

func (c *弗里茨拉尔NassauCounty) BNassau拿骚() feud.Barony {
	return c.拿骚Nassau
}

var CNassau弗里茨拉尔 NassauCounty = &弗里茨拉尔NassauCounty{}

func init() {
	f := CNassau弗里茨拉尔.(*弗里茨拉尔NassauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "121",
		Title:     "nassau",
		TitleName: "弗里茨拉尔",
		TitleCode: "c_nassau",
		Baronies:  map[string]feud.Barony{},
	}

	f.法尔肯施泰因Falkenstein = BFalkenstein法尔肯施泰因
	f.法尔肯施泰因Falkenstein.SetParent(f)

	f.弗里茨拉尔Fritzlar = BFritzlar弗里茨拉尔
	f.弗里茨拉尔Fritzlar.SetParent(f)

	f.富尔达Fulda = BFulda富尔达
	f.富尔达Fulda.SetParent(f)

	f.海鲁尔菲斯费尔德Haerulfisfeld = BHaerulfisfeld海鲁尔菲斯费尔德
	f.海鲁尔菲斯费尔德Haerulfisfeld.SetParent(f)

	f.赫斯费尔德Hersfeld = BHersfeld赫斯费尔德
	f.赫斯费尔德Hersfeld.SetParent(f)

	f.伊森堡Isenburg = BIsenburg伊森堡
	f.伊森堡Isenburg.SetParent(f)

	f.拿骚Nassau = BNassau拿骚
	f.拿骚Nassau.SetParent(f)

}
