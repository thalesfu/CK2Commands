package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HainautCounty interface {
    feud.County
    BAth阿特() 	feud.Barony
    BAvesnes阿韦讷() 	feud.Barony
    BCambrai康布雷() 	feud.Barony
    BCharleroi沙勒罗瓦() 	feud.Barony
    BChievres谢夫尔() 	feud.Barony
    BEnghien昂吉安() 	feud.Barony
    BMons蒙斯() 	feud.Barony
    BValenciennes瓦朗谢讷() 	feud.Barony
}

type 埃诺HainautCounty struct {
	feud.BaseCounty
	阿特Ath 	feud.Barony
	阿韦讷Avesnes 	feud.Barony
	康布雷Cambrai 	feud.Barony
	沙勒罗瓦Charleroi 	feud.Barony
	谢夫尔Chievres 	feud.Barony
	昂吉安Enghien 	feud.Barony
	蒙斯Mons 	feud.Barony
	瓦朗谢讷Valenciennes 	feud.Barony
}

func (c *埃诺HainautCounty) BAth阿特() feud.Barony {
	return c.阿特Ath
}
    
func (c *埃诺HainautCounty) BAvesnes阿韦讷() feud.Barony {
	return c.阿韦讷Avesnes
}
    
func (c *埃诺HainautCounty) BCambrai康布雷() feud.Barony {
	return c.康布雷Cambrai
}
    
func (c *埃诺HainautCounty) BCharleroi沙勒罗瓦() feud.Barony {
	return c.沙勒罗瓦Charleroi
}
    
func (c *埃诺HainautCounty) BChievres谢夫尔() feud.Barony {
	return c.谢夫尔Chievres
}
    
func (c *埃诺HainautCounty) BEnghien昂吉安() feud.Barony {
	return c.昂吉安Enghien
}
    
func (c *埃诺HainautCounty) BMons蒙斯() feud.Barony {
	return c.蒙斯Mons
}
    
func (c *埃诺HainautCounty) BValenciennes瓦朗谢讷() feud.Barony {
	return c.瓦朗谢讷Valenciennes
}
    
var CHainaut埃诺 HainautCounty = &埃诺HainautCounty{}

func init() {
	f := CHainaut埃诺.(*埃诺HainautCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "94",
		Title:     "hainaut",
		TitleName: "埃诺",
		TitleCode: "c_hainaut",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿特Ath = BAth阿特
	f.阿特Ath.SetParent(f)
	
	f.阿韦讷Avesnes = BAvesnes阿韦讷
	f.阿韦讷Avesnes.SetParent(f)
	
	f.康布雷Cambrai = BCambrai康布雷
	f.康布雷Cambrai.SetParent(f)
	
	f.沙勒罗瓦Charleroi = BCharleroi沙勒罗瓦
	f.沙勒罗瓦Charleroi.SetParent(f)
	
	f.谢夫尔Chievres = BChievres谢夫尔
	f.谢夫尔Chievres.SetParent(f)
	
	f.昂吉安Enghien = BEnghien昂吉安
	f.昂吉安Enghien.SetParent(f)
	
	f.蒙斯Mons = BMons蒙斯
	f.蒙斯Mons.SetParent(f)
	
	f.瓦朗谢讷Valenciennes = BValenciennes瓦朗谢讷
	f.瓦朗谢讷Valenciennes.SetParent(f)
	
}
