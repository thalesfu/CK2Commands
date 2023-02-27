package santiago

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SantiagoCounty interface {
    feud.County
    BPadron帕德龙() 	feud.Barony
    BPontevedra蓬特韦德拉() 	feud.Barony
    BSantiago圣地亚哥() 	feud.Barony
    BTuy图伊() 	feud.Barony
    BVerin贝林() 	feud.Barony
    BVigo维戈() 	feud.Barony
    BVilagarcia比拉加尔西亚() 	feud.Barony
}

type 圣地亚哥SantiagoCounty struct {
	feud.BaseCounty
	帕德龙Padron 	feud.Barony
	蓬特韦德拉Pontevedra 	feud.Barony
	圣地亚哥Santiago 	feud.Barony
	图伊Tuy 	feud.Barony
	贝林Verin 	feud.Barony
	维戈Vigo 	feud.Barony
	比拉加尔西亚Vilagarcia 	feud.Barony
}

func (c *圣地亚哥SantiagoCounty) BPadron帕德龙() feud.Barony {
	return c.帕德龙Padron
}
    
func (c *圣地亚哥SantiagoCounty) BPontevedra蓬特韦德拉() feud.Barony {
	return c.蓬特韦德拉Pontevedra
}
    
func (c *圣地亚哥SantiagoCounty) BSantiago圣地亚哥() feud.Barony {
	return c.圣地亚哥Santiago
}
    
func (c *圣地亚哥SantiagoCounty) BTuy图伊() feud.Barony {
	return c.图伊Tuy
}
    
func (c *圣地亚哥SantiagoCounty) BVerin贝林() feud.Barony {
	return c.贝林Verin
}
    
func (c *圣地亚哥SantiagoCounty) BVigo维戈() feud.Barony {
	return c.维戈Vigo
}
    
func (c *圣地亚哥SantiagoCounty) BVilagarcia比拉加尔西亚() feud.Barony {
	return c.比拉加尔西亚Vilagarcia
}
    
var CSantiago圣地亚哥 SantiagoCounty = &圣地亚哥SantiagoCounty{}

func init() {
	f := CSantiago圣地亚哥.(*圣地亚哥SantiagoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "157",
		Title:     "santiago",
		TitleName: "圣地亚哥",
		TitleCode: "c_santiago",
		Baronies:  map[string]feud.Barony{},
	}

	f.帕德龙Padron = BPadron帕德龙
	f.帕德龙Padron.SetParent(f)
	
	f.蓬特韦德拉Pontevedra = BPontevedra蓬特韦德拉
	f.蓬特韦德拉Pontevedra.SetParent(f)
	
	f.圣地亚哥Santiago = BSantiago圣地亚哥
	f.圣地亚哥Santiago.SetParent(f)
	
	f.图伊Tuy = BTuy图伊
	f.图伊Tuy.SetParent(f)
	
	f.贝林Verin = BVerin贝林
	f.贝林Verin.SetParent(f)
	
	f.维戈Vigo = BVigo维戈
	f.维戈Vigo.SetParent(f)
	
	f.比拉加尔西亚Vilagarcia = BVilagarcia比拉加尔西亚
	f.比拉加尔西亚Vilagarcia.SetParent(f)
	
}
