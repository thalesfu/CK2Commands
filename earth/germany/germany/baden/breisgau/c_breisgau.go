package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BreisgauCounty interface {
	feud.County
	BBreisach布赖萨赫() feud.Barony
	BFreiburg弗赖堡() feud.Barony
	BLahr拉尔() feud.Barony
	BLohrrach勒拉赫() feud.Barony
	BOffenburg奥芬堡() feud.Barony
	BRottweil罗特韦尔() feud.Barony
	BStblasien圣布拉辛() feud.Barony
	BZahringen扎林根() feud.Barony
}

type 布赖施高BreisgauCounty struct {
	feud.BaseCounty
	布赖萨赫Breisach  feud.Barony
	弗赖堡Freiburg   feud.Barony
	拉尔Lahr        feud.Barony
	勒拉赫Lohrrach   feud.Barony
	奥芬堡Offenburg  feud.Barony
	罗特韦尔Rottweil  feud.Barony
	圣布拉辛Stblasien feud.Barony
	扎林根Zahringen  feud.Barony
}

func (c *布赖施高BreisgauCounty) BBreisach布赖萨赫() feud.Barony {
	return c.布赖萨赫Breisach
}

func (c *布赖施高BreisgauCounty) BFreiburg弗赖堡() feud.Barony {
	return c.弗赖堡Freiburg
}

func (c *布赖施高BreisgauCounty) BLahr拉尔() feud.Barony {
	return c.拉尔Lahr
}

func (c *布赖施高BreisgauCounty) BLohrrach勒拉赫() feud.Barony {
	return c.勒拉赫Lohrrach
}

func (c *布赖施高BreisgauCounty) BOffenburg奥芬堡() feud.Barony {
	return c.奥芬堡Offenburg
}

func (c *布赖施高BreisgauCounty) BRottweil罗特韦尔() feud.Barony {
	return c.罗特韦尔Rottweil
}

func (c *布赖施高BreisgauCounty) BStblasien圣布拉辛() feud.Barony {
	return c.圣布拉辛Stblasien
}

func (c *布赖施高BreisgauCounty) BZahringen扎林根() feud.Barony {
	return c.扎林根Zahringen
}

var CBreisgau布赖施高 BreisgauCounty = &布赖施高BreisgauCounty{}

func init() {
	f := CBreisgau布赖施高.(*布赖施高BreisgauCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "250",
		Title:     "breisgau",
		TitleName: "布赖施高",
		TitleCode: "c_breisgau",
		Baronies:  map[string]feud.Barony{},
	}

	f.布赖萨赫Breisach = BBreisach布赖萨赫
	f.布赖萨赫Breisach.SetParent(f)

	f.弗赖堡Freiburg = BFreiburg弗赖堡
	f.弗赖堡Freiburg.SetParent(f)

	f.拉尔Lahr = BLahr拉尔
	f.拉尔Lahr.SetParent(f)

	f.勒拉赫Lohrrach = BLohrrach勒拉赫
	f.勒拉赫Lohrrach.SetParent(f)

	f.奥芬堡Offenburg = BOffenburg奥芬堡
	f.奥芬堡Offenburg.SetParent(f)

	f.罗特韦尔Rottweil = BRottweil罗特韦尔
	f.罗特韦尔Rottweil.SetParent(f)

	f.圣布拉辛Stblasien = BStblasien圣布拉辛
	f.圣布拉辛Stblasien.SetParent(f)

	f.扎林根Zahringen = BZahringen扎林根
	f.扎林根Zahringen.SetParent(f)

}
