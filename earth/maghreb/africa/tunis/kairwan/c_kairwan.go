package kairwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KairwanCounty interface {
    feud.County
    BChebika舍比凯() 	feud.Barony
    BHaffouz哈富兹() 	feud.Barony
    BJalula杰卢拉() 	feud.Barony
    BKairouan凯鲁万() 	feud.Barony
    BMeknassy梅克纳萨伊() 	feud.Barony
    BRaqqada雷加达() 	feud.Barony
    BSidibouzid西迪布济德() 	feud.Barony
}

type 凯鲁万KairwanCounty struct {
	feud.BaseCounty
	舍比凯Chebika 	feud.Barony
	哈富兹Haffouz 	feud.Barony
	杰卢拉Jalula 	feud.Barony
	凯鲁万Kairouan 	feud.Barony
	梅克纳萨伊Meknassy 	feud.Barony
	雷加达Raqqada 	feud.Barony
	西迪布济德Sidibouzid 	feud.Barony
}

func (c *凯鲁万KairwanCounty) BChebika舍比凯() feud.Barony {
	return c.舍比凯Chebika
}
    
func (c *凯鲁万KairwanCounty) BHaffouz哈富兹() feud.Barony {
	return c.哈富兹Haffouz
}
    
func (c *凯鲁万KairwanCounty) BJalula杰卢拉() feud.Barony {
	return c.杰卢拉Jalula
}
    
func (c *凯鲁万KairwanCounty) BKairouan凯鲁万() feud.Barony {
	return c.凯鲁万Kairouan
}
    
func (c *凯鲁万KairwanCounty) BMeknassy梅克纳萨伊() feud.Barony {
	return c.梅克纳萨伊Meknassy
}
    
func (c *凯鲁万KairwanCounty) BRaqqada雷加达() feud.Barony {
	return c.雷加达Raqqada
}
    
func (c *凯鲁万KairwanCounty) BSidibouzid西迪布济德() feud.Barony {
	return c.西迪布济德Sidibouzid
}
    
var CKairwan凯鲁万 KairwanCounty = &凯鲁万KairwanCounty{}

func init() {
	f := CKairwan凯鲁万.(*凯鲁万KairwanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "815",
		Title:     "kairwan",
		TitleName: "凯鲁万",
		TitleCode: "c_kairwan",
		Baronies:  map[string]feud.Barony{},
	}

	f.舍比凯Chebika = BChebika舍比凯
	f.舍比凯Chebika.SetParent(f)
	
	f.哈富兹Haffouz = BHaffouz哈富兹
	f.哈富兹Haffouz.SetParent(f)
	
	f.杰卢拉Jalula = BJalula杰卢拉
	f.杰卢拉Jalula.SetParent(f)
	
	f.凯鲁万Kairouan = BKairouan凯鲁万
	f.凯鲁万Kairouan.SetParent(f)
	
	f.梅克纳萨伊Meknassy = BMeknassy梅克纳萨伊
	f.梅克纳萨伊Meknassy.SetParent(f)
	
	f.雷加达Raqqada = BRaqqada雷加达
	f.雷加达Raqqada.SetParent(f)
	
	f.西迪布济德Sidibouzid = BSidibouzid西迪布济德
	f.西迪布济德Sidibouzid.SetParent(f)
	
}
