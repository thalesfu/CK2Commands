package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LaodikeiaCounty interface {
	feud.County
	BChonai肖奈() feud.Barony
	BFlaviupolis弗拉优波利斯() feud.Barony
	BGordes格尔德斯() feud.Barony
	BHieropolis黑洛波利斯() feud.Barony
	BKona科纳() feud.Barony
	BLaodikeia老底嘉() feud.Barony
	BRhoas洛阿斯() feud.Barony
}

type 老底嘉LaodikeiaCounty struct {
	feud.BaseCounty
	肖奈Chonai          feud.Barony
	弗拉优波利斯Flaviupolis feud.Barony
	格尔德斯Gordes        feud.Barony
	黑洛波利斯Hieropolis   feud.Barony
	科纳Kona            feud.Barony
	老底嘉Laodikeia      feud.Barony
	洛阿斯Rhoas          feud.Barony
}

func (c *老底嘉LaodikeiaCounty) BChonai肖奈() feud.Barony {
	return c.肖奈Chonai
}

func (c *老底嘉LaodikeiaCounty) BFlaviupolis弗拉优波利斯() feud.Barony {
	return c.弗拉优波利斯Flaviupolis
}

func (c *老底嘉LaodikeiaCounty) BGordes格尔德斯() feud.Barony {
	return c.格尔德斯Gordes
}

func (c *老底嘉LaodikeiaCounty) BHieropolis黑洛波利斯() feud.Barony {
	return c.黑洛波利斯Hieropolis
}

func (c *老底嘉LaodikeiaCounty) BKona科纳() feud.Barony {
	return c.科纳Kona
}

func (c *老底嘉LaodikeiaCounty) BLaodikeia老底嘉() feud.Barony {
	return c.老底嘉Laodikeia
}

func (c *老底嘉LaodikeiaCounty) BRhoas洛阿斯() feud.Barony {
	return c.洛阿斯Rhoas
}

var CLaodikeia老底嘉 LaodikeiaCounty = &老底嘉LaodikeiaCounty{}

func init() {
	f := CLaodikeia老底嘉.(*老底嘉LaodikeiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "748",
		Title:     "laodikeia",
		TitleName: "老底嘉",
		TitleCode: "c_laodikeia",
		Baronies:  map[string]feud.Barony{},
	}

	f.肖奈Chonai = BChonai肖奈
	f.肖奈Chonai.SetParent(f)

	f.弗拉优波利斯Flaviupolis = BFlaviupolis弗拉优波利斯
	f.弗拉优波利斯Flaviupolis.SetParent(f)

	f.格尔德斯Gordes = BGordes格尔德斯
	f.格尔德斯Gordes.SetParent(f)

	f.黑洛波利斯Hieropolis = BHieropolis黑洛波利斯
	f.黑洛波利斯Hieropolis.SetParent(f)

	f.科纳Kona = BKona科纳
	f.科纳Kona.SetParent(f)

	f.老底嘉Laodikeia = BLaodikeia老底嘉
	f.老底嘉Laodikeia.SetParent(f)

	f.洛阿斯Rhoas = BRhoas洛阿斯
	f.洛阿斯Rhoas.SetParent(f)

}
