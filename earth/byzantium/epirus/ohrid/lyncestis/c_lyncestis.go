package lyncestis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LyncestisCounty interface {
    feud.County
    BBitola比托拉() 	feud.Barony
    BHeraclea_lyncestis赫拉克利亚() 	feud.Barony
    BLyncestis吕恩盖斯提斯() 	feud.Barony
    BMagarevo马加雷沃() 	feud.Barony
    BNeoljani内奥利亚尼() 	feud.Barony
    BPelagonia佩拉戈尼亚() 	feud.Barony
    BPrilap普里莱普() 	feud.Barony
}

type 比托拉LyncestisCounty struct {
	feud.BaseCounty
	比托拉Bitola 	feud.Barony
	赫拉克利亚Heraclea_lyncestis 	feud.Barony
	吕恩盖斯提斯Lyncestis 	feud.Barony
	马加雷沃Magarevo 	feud.Barony
	内奥利亚尼Neoljani 	feud.Barony
	佩拉戈尼亚Pelagonia 	feud.Barony
	普里莱普Prilap 	feud.Barony
}

func (c *比托拉LyncestisCounty) BBitola比托拉() feud.Barony {
	return c.比托拉Bitola
}
    
func (c *比托拉LyncestisCounty) BHeraclea_lyncestis赫拉克利亚() feud.Barony {
	return c.赫拉克利亚Heraclea_lyncestis
}
    
func (c *比托拉LyncestisCounty) BLyncestis吕恩盖斯提斯() feud.Barony {
	return c.吕恩盖斯提斯Lyncestis
}
    
func (c *比托拉LyncestisCounty) BMagarevo马加雷沃() feud.Barony {
	return c.马加雷沃Magarevo
}
    
func (c *比托拉LyncestisCounty) BNeoljani内奥利亚尼() feud.Barony {
	return c.内奥利亚尼Neoljani
}
    
func (c *比托拉LyncestisCounty) BPelagonia佩拉戈尼亚() feud.Barony {
	return c.佩拉戈尼亚Pelagonia
}
    
func (c *比托拉LyncestisCounty) BPrilap普里莱普() feud.Barony {
	return c.普里莱普Prilap
}
    
var CLyncestis比托拉 LyncestisCounty = &比托拉LyncestisCounty{}

func init() {
	f := CLyncestis比托拉.(*比托拉LyncestisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1887",
		Title:     "lyncestis",
		TitleName: "比托拉",
		TitleCode: "c_lyncestis",
		Baronies:  map[string]feud.Barony{},
	}

	f.比托拉Bitola = BBitola比托拉
	f.比托拉Bitola.SetParent(f)
	
	f.赫拉克利亚Heraclea_lyncestis = BHeraclea_lyncestis赫拉克利亚
	f.赫拉克利亚Heraclea_lyncestis.SetParent(f)
	
	f.吕恩盖斯提斯Lyncestis = BLyncestis吕恩盖斯提斯
	f.吕恩盖斯提斯Lyncestis.SetParent(f)
	
	f.马加雷沃Magarevo = BMagarevo马加雷沃
	f.马加雷沃Magarevo.SetParent(f)
	
	f.内奥利亚尼Neoljani = BNeoljani内奥利亚尼
	f.内奥利亚尼Neoljani.SetParent(f)
	
	f.佩拉戈尼亚Pelagonia = BPelagonia佩拉戈尼亚
	f.佩拉戈尼亚Pelagonia.SetParent(f)
	
	f.普里莱普Prilap = BPrilap普里莱普
	f.普里莱普Prilap.SetParent(f)
	
}
