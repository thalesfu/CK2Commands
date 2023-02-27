package tsambagarav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsambagaravCounty interface {
    feud.County
    BAchit_tsambagarav阿奇特() 	feud.Barony
    BKhovd_tsambagarav科布多() 	feud.Barony
    BKokorya科科里亚() 	feud.Barony
    BKyzyl_khaya克孜勒哈亚() 	feud.Barony
    BTashanta塔尚塔() 	feud.Barony
    BTsambagarav查木巴嘎拉布() 	feud.Barony
    BUureg乌雷格() 	feud.Barony
}

type 查木巴嘎拉布TsambagaravCounty struct {
	feud.BaseCounty
	阿奇特Achit_tsambagarav 	feud.Barony
	科布多Khovd_tsambagarav 	feud.Barony
	科科里亚Kokorya 	feud.Barony
	克孜勒哈亚Kyzyl_khaya 	feud.Barony
	塔尚塔Tashanta 	feud.Barony
	查木巴嘎拉布Tsambagarav 	feud.Barony
	乌雷格Uureg 	feud.Barony
}

func (c *查木巴嘎拉布TsambagaravCounty) BAchit_tsambagarav阿奇特() feud.Barony {
	return c.阿奇特Achit_tsambagarav
}
    
func (c *查木巴嘎拉布TsambagaravCounty) BKhovd_tsambagarav科布多() feud.Barony {
	return c.科布多Khovd_tsambagarav
}
    
func (c *查木巴嘎拉布TsambagaravCounty) BKokorya科科里亚() feud.Barony {
	return c.科科里亚Kokorya
}
    
func (c *查木巴嘎拉布TsambagaravCounty) BKyzyl_khaya克孜勒哈亚() feud.Barony {
	return c.克孜勒哈亚Kyzyl_khaya
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

	f.阿奇特Achit_tsambagarav = BAchit_tsambagarav阿奇特
	f.阿奇特Achit_tsambagarav.SetParent(f)
	
	f.科布多Khovd_tsambagarav = BKhovd_tsambagarav科布多
	f.科布多Khovd_tsambagarav.SetParent(f)
	
	f.科科里亚Kokorya = BKokorya科科里亚
	f.科科里亚Kokorya.SetParent(f)
	
	f.克孜勒哈亚Kyzyl_khaya = BKyzyl_khaya克孜勒哈亚
	f.克孜勒哈亚Kyzyl_khaya.SetParent(f)
	
	f.塔尚塔Tashanta = BTashanta塔尚塔
	f.塔尚塔Tashanta.SetParent(f)
	
	f.查木巴嘎拉布Tsambagarav = BTsambagarav查木巴嘎拉布
	f.查木巴嘎拉布Tsambagarav.SetParent(f)
	
	f.乌雷格Uureg = BUureg乌雷格
	f.乌雷格Uureg.SetParent(f)
	
}
