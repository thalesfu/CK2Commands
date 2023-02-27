package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KostiCounty interface {
    feud.County
    BAbbeit阿贝特() 	feud.Barony
    BDalang迪灵() 	feud.Barony
    BKaduqli卡杜格利() 	feud.Barony
    BKosti科斯提() 	feud.Barony
    BLagawa拉嘎瓦() 	feud.Barony
    BTandalti坦达勒提() 	feud.Barony
    BUmm_ruwaba乌姆鲁瓦巴() 	feud.Barony
}

type 科斯提KostiCounty struct {
	feud.BaseCounty
	阿贝特Abbeit 	feud.Barony
	迪灵Dalang 	feud.Barony
	卡杜格利Kaduqli 	feud.Barony
	科斯提Kosti 	feud.Barony
	拉嘎瓦Lagawa 	feud.Barony
	坦达勒提Tandalti 	feud.Barony
	乌姆鲁瓦巴Umm_ruwaba 	feud.Barony
}

func (c *科斯提KostiCounty) BAbbeit阿贝特() feud.Barony {
	return c.阿贝特Abbeit
}
    
func (c *科斯提KostiCounty) BDalang迪灵() feud.Barony {
	return c.迪灵Dalang
}
    
func (c *科斯提KostiCounty) BKaduqli卡杜格利() feud.Barony {
	return c.卡杜格利Kaduqli
}
    
func (c *科斯提KostiCounty) BKosti科斯提() feud.Barony {
	return c.科斯提Kosti
}
    
func (c *科斯提KostiCounty) BLagawa拉嘎瓦() feud.Barony {
	return c.拉嘎瓦Lagawa
}
    
func (c *科斯提KostiCounty) BTandalti坦达勒提() feud.Barony {
	return c.坦达勒提Tandalti
}
    
func (c *科斯提KostiCounty) BUmm_ruwaba乌姆鲁瓦巴() feud.Barony {
	return c.乌姆鲁瓦巴Umm_ruwaba
}
    
var CKosti科斯提 KostiCounty = &科斯提KostiCounty{}

func init() {
	f := CKosti科斯提.(*科斯提KostiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1335",
		Title:     "kosti",
		TitleName: "科斯提",
		TitleCode: "c_kosti",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿贝特Abbeit = BAbbeit阿贝特
	f.阿贝特Abbeit.SetParent(f)
	
	f.迪灵Dalang = BDalang迪灵
	f.迪灵Dalang.SetParent(f)
	
	f.卡杜格利Kaduqli = BKaduqli卡杜格利
	f.卡杜格利Kaduqli.SetParent(f)
	
	f.科斯提Kosti = BKosti科斯提
	f.科斯提Kosti.SetParent(f)
	
	f.拉嘎瓦Lagawa = BLagawa拉嘎瓦
	f.拉嘎瓦Lagawa.SetParent(f)
	
	f.坦达勒提Tandalti = BTandalti坦达勒提
	f.坦达勒提Tandalti.SetParent(f)
	
	f.乌姆鲁瓦巴Umm_ruwaba = BUmm_ruwaba乌姆鲁瓦巴
	f.乌姆鲁瓦巴Umm_ruwaba.SetParent(f)
	
}
