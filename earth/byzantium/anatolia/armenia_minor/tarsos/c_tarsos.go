package tarsos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarsosCounty interface {
    feud.County
    BBardzerben巴德泽本() 	feud.Barony
    BCastabala卡斯塔巴拉() 	feud.Barony
    BKorikos科里科斯() 	feud.Barony
    BLamas拉马斯() 	feud.Barony
    BLampron拉姆拉恩() 	feud.Barony
    BPendosis朋多斯() 	feud.Barony
    BTarsos塔索斯() 	feud.Barony
    BZephyrium泽胡里尤() 	feud.Barony
}

type 塔索斯TarsosCounty struct {
	feud.BaseCounty
	巴德泽本Bardzerben 	feud.Barony
	卡斯塔巴拉Castabala 	feud.Barony
	科里科斯Korikos 	feud.Barony
	拉马斯Lamas 	feud.Barony
	拉姆拉恩Lampron 	feud.Barony
	朋多斯Pendosis 	feud.Barony
	塔索斯Tarsos 	feud.Barony
	泽胡里尤Zephyrium 	feud.Barony
}

func (c *塔索斯TarsosCounty) BBardzerben巴德泽本() feud.Barony {
	return c.巴德泽本Bardzerben
}
    
func (c *塔索斯TarsosCounty) BCastabala卡斯塔巴拉() feud.Barony {
	return c.卡斯塔巴拉Castabala
}
    
func (c *塔索斯TarsosCounty) BKorikos科里科斯() feud.Barony {
	return c.科里科斯Korikos
}
    
func (c *塔索斯TarsosCounty) BLamas拉马斯() feud.Barony {
	return c.拉马斯Lamas
}
    
func (c *塔索斯TarsosCounty) BLampron拉姆拉恩() feud.Barony {
	return c.拉姆拉恩Lampron
}
    
func (c *塔索斯TarsosCounty) BPendosis朋多斯() feud.Barony {
	return c.朋多斯Pendosis
}
    
func (c *塔索斯TarsosCounty) BTarsos塔索斯() feud.Barony {
	return c.塔索斯Tarsos
}
    
func (c *塔索斯TarsosCounty) BZephyrium泽胡里尤() feud.Barony {
	return c.泽胡里尤Zephyrium
}
    
var CTarsos塔索斯 TarsosCounty = &塔索斯TarsosCounty{}

func init() {
	f := CTarsos塔索斯.(*塔索斯TarsosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "761",
		Title:     "tarsos",
		TitleName: "塔索斯",
		TitleCode: "c_tarsos",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴德泽本Bardzerben = BBardzerben巴德泽本
	f.巴德泽本Bardzerben.SetParent(f)
	
	f.卡斯塔巴拉Castabala = BCastabala卡斯塔巴拉
	f.卡斯塔巴拉Castabala.SetParent(f)
	
	f.科里科斯Korikos = BKorikos科里科斯
	f.科里科斯Korikos.SetParent(f)
	
	f.拉马斯Lamas = BLamas拉马斯
	f.拉马斯Lamas.SetParent(f)
	
	f.拉姆拉恩Lampron = BLampron拉姆拉恩
	f.拉姆拉恩Lampron.SetParent(f)
	
	f.朋多斯Pendosis = BPendosis朋多斯
	f.朋多斯Pendosis.SetParent(f)
	
	f.塔索斯Tarsos = BTarsos塔索斯
	f.塔索斯Tarsos.SetParent(f)
	
	f.泽胡里尤Zephyrium = BZephyrium泽胡里尤
	f.泽胡里尤Zephyrium.SetParent(f)
	
}
