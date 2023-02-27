package sticht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type StichtCounty interface {
    feud.County
    BBuren比伦() 	feud.Barony
    BDorestad多雷斯塔德() 	feud.Barony
    BGorinchem霍林赫姆() 	feud.Barony
    BIjsselstein艾瑟尔斯泰因() 	feud.Barony
    BOudewater奥德瓦特() 	feud.Barony
    BUtrecht乌德勒支() 	feud.Barony
    BWoerden武尔登() 	feud.Barony
    BZeist泽斯特() 	feud.Barony
}

type 施蒂希特StichtCounty struct {
	feud.BaseCounty
	比伦Buren 	feud.Barony
	多雷斯塔德Dorestad 	feud.Barony
	霍林赫姆Gorinchem 	feud.Barony
	艾瑟尔斯泰因Ijsselstein 	feud.Barony
	奥德瓦特Oudewater 	feud.Barony
	乌德勒支Utrecht 	feud.Barony
	武尔登Woerden 	feud.Barony
	泽斯特Zeist 	feud.Barony
}

func (c *施蒂希特StichtCounty) BBuren比伦() feud.Barony {
	return c.比伦Buren
}
    
func (c *施蒂希特StichtCounty) BDorestad多雷斯塔德() feud.Barony {
	return c.多雷斯塔德Dorestad
}
    
func (c *施蒂希特StichtCounty) BGorinchem霍林赫姆() feud.Barony {
	return c.霍林赫姆Gorinchem
}
    
func (c *施蒂希特StichtCounty) BIjsselstein艾瑟尔斯泰因() feud.Barony {
	return c.艾瑟尔斯泰因Ijsselstein
}
    
func (c *施蒂希特StichtCounty) BOudewater奥德瓦特() feud.Barony {
	return c.奥德瓦特Oudewater
}
    
func (c *施蒂希特StichtCounty) BUtrecht乌德勒支() feud.Barony {
	return c.乌德勒支Utrecht
}
    
func (c *施蒂希特StichtCounty) BWoerden武尔登() feud.Barony {
	return c.武尔登Woerden
}
    
func (c *施蒂希特StichtCounty) BZeist泽斯特() feud.Barony {
	return c.泽斯特Zeist
}
    
var CSticht施蒂希特 StichtCounty = &施蒂希特StichtCounty{}

func init() {
	f := CSticht施蒂希特.(*施蒂希特StichtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "82",
		Title:     "sticht",
		TitleName: "施蒂希特",
		TitleCode: "c_sticht",
		Baronies:  map[string]feud.Barony{},
	}

	f.比伦Buren = BBuren比伦
	f.比伦Buren.SetParent(f)
	
	f.多雷斯塔德Dorestad = BDorestad多雷斯塔德
	f.多雷斯塔德Dorestad.SetParent(f)
	
	f.霍林赫姆Gorinchem = BGorinchem霍林赫姆
	f.霍林赫姆Gorinchem.SetParent(f)
	
	f.艾瑟尔斯泰因Ijsselstein = BIjsselstein艾瑟尔斯泰因
	f.艾瑟尔斯泰因Ijsselstein.SetParent(f)
	
	f.奥德瓦特Oudewater = BOudewater奥德瓦特
	f.奥德瓦特Oudewater.SetParent(f)
	
	f.乌德勒支Utrecht = BUtrecht乌德勒支
	f.乌德勒支Utrecht.SetParent(f)
	
	f.武尔登Woerden = BWoerden武尔登
	f.武尔登Woerden.SetParent(f)
	
	f.泽斯特Zeist = BZeist泽斯特
	f.泽斯特Zeist.SetParent(f)
	
}
