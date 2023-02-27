package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PokharaCounty interface {
    feud.County
    BBesisahar贝西沙尔() 	feud.Barony
    BButwal布德沃尔() 	feud.Barony
    BChame磋梅() 	feud.Barony
    BMarpha莽尔法() 	feud.Barony
    BMuktinath木底那他() 	feud.Barony
    BPokhara布佉罗() 	feud.Barony
    BPrithbinarayan毕哩体微那罗延() 	feud.Barony
}

type 布佉罗PokharaCounty struct {
	feud.BaseCounty
	贝西沙尔Besisahar 	feud.Barony
	布德沃尔Butwal 	feud.Barony
	磋梅Chame 	feud.Barony
	莽尔法Marpha 	feud.Barony
	木底那他Muktinath 	feud.Barony
	布佉罗Pokhara 	feud.Barony
	毕哩体微那罗延Prithbinarayan 	feud.Barony
}

func (c *布佉罗PokharaCounty) BBesisahar贝西沙尔() feud.Barony {
	return c.贝西沙尔Besisahar
}
    
func (c *布佉罗PokharaCounty) BButwal布德沃尔() feud.Barony {
	return c.布德沃尔Butwal
}
    
func (c *布佉罗PokharaCounty) BChame磋梅() feud.Barony {
	return c.磋梅Chame
}
    
func (c *布佉罗PokharaCounty) BMarpha莽尔法() feud.Barony {
	return c.莽尔法Marpha
}
    
func (c *布佉罗PokharaCounty) BMuktinath木底那他() feud.Barony {
	return c.木底那他Muktinath
}
    
func (c *布佉罗PokharaCounty) BPokhara布佉罗() feud.Barony {
	return c.布佉罗Pokhara
}
    
func (c *布佉罗PokharaCounty) BPrithbinarayan毕哩体微那罗延() feud.Barony {
	return c.毕哩体微那罗延Prithbinarayan
}
    
var CPokhara布佉罗 PokharaCounty = &布佉罗PokharaCounty{}

func init() {
	f := CPokhara布佉罗.(*布佉罗PokharaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1475",
		Title:     "pokhara",
		TitleName: "布佉罗",
		TitleCode: "c_pokhara",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝西沙尔Besisahar = BBesisahar贝西沙尔
	f.贝西沙尔Besisahar.SetParent(f)
	
	f.布德沃尔Butwal = BButwal布德沃尔
	f.布德沃尔Butwal.SetParent(f)
	
	f.磋梅Chame = BChame磋梅
	f.磋梅Chame.SetParent(f)
	
	f.莽尔法Marpha = BMarpha莽尔法
	f.莽尔法Marpha.SetParent(f)
	
	f.木底那他Muktinath = BMuktinath木底那他
	f.木底那他Muktinath.SetParent(f)
	
	f.布佉罗Pokhara = BPokhara布佉罗
	f.布佉罗Pokhara.SetParent(f)
	
	f.毕哩体微那罗延Prithbinarayan = BPrithbinarayan毕哩体微那罗延
	f.毕哩体微那罗延Prithbinarayan.SetParent(f)
	
}
