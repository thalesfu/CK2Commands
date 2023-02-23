package tripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TripuriCounty interface {
	feud.County
	BGarha迦罗诃() feud.Barony
	BMandla门陀罗() feud.Barony
	BRasidpur罗悉补罗() feud.Barony
	BSamdong僧东() feud.Barony
	BTripuri帝利城() feud.Barony
	BTungaturti东伽咄帝() feud.Barony
	BVelan吠烂() feud.Barony
}

type 帝利城TripuriCounty struct {
	feud.BaseCounty
	迦罗诃Garha       feud.Barony
	门陀罗Mandla      feud.Barony
	罗悉补罗Rasidpur   feud.Barony
	僧东Samdong      feud.Barony
	帝利城Tripuri     feud.Barony
	东伽咄帝Tungaturti feud.Barony
	吠烂Velan        feud.Barony
}

func (c *帝利城TripuriCounty) BGarha迦罗诃() feud.Barony {
	return c.迦罗诃Garha
}

func (c *帝利城TripuriCounty) BMandla门陀罗() feud.Barony {
	return c.门陀罗Mandla
}

func (c *帝利城TripuriCounty) BRasidpur罗悉补罗() feud.Barony {
	return c.罗悉补罗Rasidpur
}

func (c *帝利城TripuriCounty) BSamdong僧东() feud.Barony {
	return c.僧东Samdong
}

func (c *帝利城TripuriCounty) BTripuri帝利城() feud.Barony {
	return c.帝利城Tripuri
}

func (c *帝利城TripuriCounty) BTungaturti东伽咄帝() feud.Barony {
	return c.东伽咄帝Tungaturti
}

func (c *帝利城TripuriCounty) BVelan吠烂() feud.Barony {
	return c.吠烂Velan
}

var CTripuri帝利城 TripuriCounty = &帝利城TripuriCounty{}

func init() {
	f := CTripuri帝利城.(*帝利城TripuriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1271",
		Title:     "tripuri",
		TitleName: "帝利城",
		TitleCode: "c_tripuri",
		Baronies:  map[string]feud.Barony{},
	}

	f.迦罗诃Garha = BGarha迦罗诃
	f.迦罗诃Garha.SetParent(f)

	f.门陀罗Mandla = BMandla门陀罗
	f.门陀罗Mandla.SetParent(f)

	f.罗悉补罗Rasidpur = BRasidpur罗悉补罗
	f.罗悉补罗Rasidpur.SetParent(f)

	f.僧东Samdong = BSamdong僧东
	f.僧东Samdong.SetParent(f)

	f.帝利城Tripuri = BTripuri帝利城
	f.帝利城Tripuri.SetParent(f)

	f.东伽咄帝Tungaturti = BTungaturti东伽咄帝
	f.东伽咄帝Tungaturti.SetParent(f)

	f.吠烂Velan = BVelan吠烂
	f.吠烂Velan.SetParent(f)

}
