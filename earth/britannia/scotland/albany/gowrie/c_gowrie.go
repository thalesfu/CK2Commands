package gowrie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GowrieCounty interface {
    feud.County
    BAbernethy阿伯内西() 	feud.Barony
    BClunie克卢尼() 	feud.Barony
    BDundee邓迪() 	feud.Barony
    BDunkeld邓凯尔德() 	feud.Barony
    BErrol埃勒尔() 	feud.Barony
    BForteviot福蒂维尔特() 	feud.Barony
    BPerth珀斯() 	feud.Barony
    BScone斯昆() 	feud.Barony
}

type 高里GowrieCounty struct {
	feud.BaseCounty
	阿伯内西Abernethy 	feud.Barony
	克卢尼Clunie 	feud.Barony
	邓迪Dundee 	feud.Barony
	邓凯尔德Dunkeld 	feud.Barony
	埃勒尔Errol 	feud.Barony
	福蒂维尔特Forteviot 	feud.Barony
	珀斯Perth 	feud.Barony
	斯昆Scone 	feud.Barony
}

func (c *高里GowrieCounty) BAbernethy阿伯内西() feud.Barony {
	return c.阿伯内西Abernethy
}
    
func (c *高里GowrieCounty) BClunie克卢尼() feud.Barony {
	return c.克卢尼Clunie
}
    
func (c *高里GowrieCounty) BDundee邓迪() feud.Barony {
	return c.邓迪Dundee
}
    
func (c *高里GowrieCounty) BDunkeld邓凯尔德() feud.Barony {
	return c.邓凯尔德Dunkeld
}
    
func (c *高里GowrieCounty) BErrol埃勒尔() feud.Barony {
	return c.埃勒尔Errol
}
    
func (c *高里GowrieCounty) BForteviot福蒂维尔特() feud.Barony {
	return c.福蒂维尔特Forteviot
}
    
func (c *高里GowrieCounty) BPerth珀斯() feud.Barony {
	return c.珀斯Perth
}
    
func (c *高里GowrieCounty) BScone斯昆() feud.Barony {
	return c.斯昆Scone
}
    
var CGowrie高里 GowrieCounty = &高里GowrieCounty{}

func init() {
	f := CGowrie高里.(*高里GowrieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "43",
		Title:     "gowrie",
		TitleName: "高里",
		TitleCode: "c_gowrie",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯内西Abernethy = BAbernethy阿伯内西
	f.阿伯内西Abernethy.SetParent(f)
	
	f.克卢尼Clunie = BClunie克卢尼
	f.克卢尼Clunie.SetParent(f)
	
	f.邓迪Dundee = BDundee邓迪
	f.邓迪Dundee.SetParent(f)
	
	f.邓凯尔德Dunkeld = BDunkeld邓凯尔德
	f.邓凯尔德Dunkeld.SetParent(f)
	
	f.埃勒尔Errol = BErrol埃勒尔
	f.埃勒尔Errol.SetParent(f)
	
	f.福蒂维尔特Forteviot = BForteviot福蒂维尔特
	f.福蒂维尔特Forteviot.SetParent(f)
	
	f.珀斯Perth = BPerth珀斯
	f.珀斯Perth.SetParent(f)
	
	f.斯昆Scone = BScone斯昆
	f.斯昆Scone.SetParent(f)
	
}
