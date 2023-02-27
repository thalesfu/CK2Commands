package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BarcelonaCounty interface {
    feud.County
    BBarcelona巴塞罗那() 	feud.Barony
    BBerga贝尔加() 	feud.Barony
    BIgualada伊瓜拉达() 	feud.Barony
    BManresa曼雷萨() 	feud.Barony
    BOsona乌孙纳() 	feud.Barony
    BProvencana普罗文萨纳() 	feud.Barony
    BVallparadis巴利帕拉迪斯() 	feud.Barony
    BVic比克() 	feud.Barony
}

type 巴塞罗那BarcelonaCounty struct {
	feud.BaseCounty
	巴塞罗那Barcelona 	feud.Barony
	贝尔加Berga 	feud.Barony
	伊瓜拉达Igualada 	feud.Barony
	曼雷萨Manresa 	feud.Barony
	乌孙纳Osona 	feud.Barony
	普罗文萨纳Provencana 	feud.Barony
	巴利帕拉迪斯Vallparadis 	feud.Barony
	比克Vic 	feud.Barony
}

func (c *巴塞罗那BarcelonaCounty) BBarcelona巴塞罗那() feud.Barony {
	return c.巴塞罗那Barcelona
}
    
func (c *巴塞罗那BarcelonaCounty) BBerga贝尔加() feud.Barony {
	return c.贝尔加Berga
}
    
func (c *巴塞罗那BarcelonaCounty) BIgualada伊瓜拉达() feud.Barony {
	return c.伊瓜拉达Igualada
}
    
func (c *巴塞罗那BarcelonaCounty) BManresa曼雷萨() feud.Barony {
	return c.曼雷萨Manresa
}
    
func (c *巴塞罗那BarcelonaCounty) BOsona乌孙纳() feud.Barony {
	return c.乌孙纳Osona
}
    
func (c *巴塞罗那BarcelonaCounty) BProvencana普罗文萨纳() feud.Barony {
	return c.普罗文萨纳Provencana
}
    
func (c *巴塞罗那BarcelonaCounty) BVallparadis巴利帕拉迪斯() feud.Barony {
	return c.巴利帕拉迪斯Vallparadis
}
    
func (c *巴塞罗那BarcelonaCounty) BVic比克() feud.Barony {
	return c.比克Vic
}
    
var CBarcelona巴塞罗那 BarcelonaCounty = &巴塞罗那BarcelonaCounty{}

func init() {
	f := CBarcelona巴塞罗那.(*巴塞罗那BarcelonaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "204",
		Title:     "barcelona",
		TitleName: "巴塞罗那",
		TitleCode: "c_barcelona",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴塞罗那Barcelona = BBarcelona巴塞罗那
	f.巴塞罗那Barcelona.SetParent(f)
	
	f.贝尔加Berga = BBerga贝尔加
	f.贝尔加Berga.SetParent(f)
	
	f.伊瓜拉达Igualada = BIgualada伊瓜拉达
	f.伊瓜拉达Igualada.SetParent(f)
	
	f.曼雷萨Manresa = BManresa曼雷萨
	f.曼雷萨Manresa.SetParent(f)
	
	f.乌孙纳Osona = BOsona乌孙纳
	f.乌孙纳Osona.SetParent(f)
	
	f.普罗文萨纳Provencana = BProvencana普罗文萨纳
	f.普罗文萨纳Provencana.SetParent(f)
	
	f.巴利帕拉迪斯Vallparadis = BVallparadis巴利帕拉迪斯
	f.巴利帕拉迪斯Vallparadis.SetParent(f)
	
	f.比克Vic = BVic比克
	f.比克Vic.SetParent(f)
	
}
