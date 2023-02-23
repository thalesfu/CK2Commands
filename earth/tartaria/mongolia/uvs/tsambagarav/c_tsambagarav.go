package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsambagaravCounty interface {
	feud.County
	BKokorya科科里亚() feud.Barony
	BTashanta塔尚塔() feud.Barony
	BTsambagarav查木巴嘎拉布() feud.Barony
	BUureg乌雷格() feud.Barony
}

type 查木巴嘎拉布TsambagaravCounty struct {
	feud.BaseCounty
	科科里亚Kokorya       feud.Barony
	塔尚塔Tashanta       feud.Barony
	查木巴嘎拉布Tsambagarav feud.Barony
	乌雷格Uureg          feud.Barony
}

func (c *查木巴嘎拉布TsambagaravCounty) BKokorya科科里亚() feud.Barony {
	return c.科科里亚Kokorya
}

func (c *查木巴嘎拉布TsambagaravCounty) BTashanta塔尚塔() feud.Barony {
	return c.塔尚塔Tashanta
}

func (c *查木巴嘎拉布TsambagaravCounty) BTsambagarav查木巴嘎拉布() feud.Barony {
	return c.查木巴嘎拉布Tsambagarav
}

func (c *查木巴嘎拉布TsambagaravCounty) BUureg乌雷格() feud.Barony {
	return c.乌雷格Uureg
}

var CTsambagarav查木巴嘎拉布 TsambagaravCounty = &查木巴嘎拉布TsambagaravCounty{}

func init() {
	f := CTsambagarav查木巴嘎拉布.(*查木巴嘎拉布TsambagaravCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1908",
		Title:     "tsambagarav",
		TitleName: "查木巴嘎拉布",
		TitleCode: "c_tsambagarav",
		Baronies:  map[string]feud.Barony{},
	}

	f.科科里亚Kokorya = BKokorya科科里亚
	f.科科里亚Kokorya.SetParent(f)

	f.塔尚塔Tashanta = BTashanta塔尚塔
	f.塔尚塔Tashanta.SetParent(f)

	f.查木巴嘎拉布Tsambagarav = BTsambagarav查木巴嘎拉布
	f.查木巴嘎拉布Tsambagarav.SetParent(f)

	f.乌雷格Uureg = BUureg乌雷格
	f.乌雷格Uureg.SetParent(f)

}
