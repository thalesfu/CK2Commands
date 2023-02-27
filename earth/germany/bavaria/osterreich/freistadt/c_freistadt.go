package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FreistadtCounty interface {
    feud.County
    BAschach阿沙赫() 	feud.Barony
    BFreistadt弗赖施塔特() 	feud.Barony
    BGrein格赖恩() 	feud.Barony
    BPergkirchen佩尔格基兴() 	feud.Barony
    BSteyregg施泰尔埃格() 	feud.Barony
    BTaversheim塔费尔斯海姆() 	feud.Barony
    BZell采尔() 	feud.Barony
}

type 弗赖施塔特FreistadtCounty struct {
	feud.BaseCounty
	阿沙赫Aschach 	feud.Barony
	弗赖施塔特Freistadt 	feud.Barony
	格赖恩Grein 	feud.Barony
	佩尔格基兴Pergkirchen 	feud.Barony
	施泰尔埃格Steyregg 	feud.Barony
	塔费尔斯海姆Taversheim 	feud.Barony
	采尔Zell 	feud.Barony
}

func (c *弗赖施塔特FreistadtCounty) BAschach阿沙赫() feud.Barony {
	return c.阿沙赫Aschach
}
    
func (c *弗赖施塔特FreistadtCounty) BFreistadt弗赖施塔特() feud.Barony {
	return c.弗赖施塔特Freistadt
}
    
func (c *弗赖施塔特FreistadtCounty) BGrein格赖恩() feud.Barony {
	return c.格赖恩Grein
}
    
func (c *弗赖施塔特FreistadtCounty) BPergkirchen佩尔格基兴() feud.Barony {
	return c.佩尔格基兴Pergkirchen
}
    
func (c *弗赖施塔特FreistadtCounty) BSteyregg施泰尔埃格() feud.Barony {
	return c.施泰尔埃格Steyregg
}
    
func (c *弗赖施塔特FreistadtCounty) BTaversheim塔费尔斯海姆() feud.Barony {
	return c.塔费尔斯海姆Taversheim
}
    
func (c *弗赖施塔特FreistadtCounty) BZell采尔() feud.Barony {
	return c.采尔Zell
}
    
var CFreistadt弗赖施塔特 FreistadtCounty = &弗赖施塔特FreistadtCounty{}

func init() {
	f := CFreistadt弗赖施塔特.(*弗赖施塔特FreistadtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1698",
		Title:     "freistadt",
		TitleName: "弗赖施塔特",
		TitleCode: "c_freistadt",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿沙赫Aschach = BAschach阿沙赫
	f.阿沙赫Aschach.SetParent(f)
	
	f.弗赖施塔特Freistadt = BFreistadt弗赖施塔特
	f.弗赖施塔特Freistadt.SetParent(f)
	
	f.格赖恩Grein = BGrein格赖恩
	f.格赖恩Grein.SetParent(f)
	
	f.佩尔格基兴Pergkirchen = BPergkirchen佩尔格基兴
	f.佩尔格基兴Pergkirchen.SetParent(f)
	
	f.施泰尔埃格Steyregg = BSteyregg施泰尔埃格
	f.施泰尔埃格Steyregg.SetParent(f)
	
	f.塔费尔斯海姆Taversheim = BTaversheim塔费尔斯海姆
	f.塔费尔斯海姆Taversheim.SetParent(f)
	
	f.采尔Zell = BZell采尔
	f.采尔Zell.SetParent(f)
	
}
