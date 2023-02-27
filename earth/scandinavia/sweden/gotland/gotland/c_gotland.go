package gotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GotlandCounty interface {
    feud.County
    BAlva阿尔瓦() 	feud.Barony
    BGeatish_roma罗玛() 	feud.Barony
    BHejde黑德() 	feud.Barony
    BHemse海姆瑟() 	feud.Barony
    BOthem乌特海姆() 	feud.Barony
    BSlite斯利特() 	feud.Barony
    BVisborg维斯堡() 	feud.Barony
    BVisby维斯比() 	feud.Barony
}

type 哥特兰GotlandCounty struct {
	feud.BaseCounty
	阿尔瓦Alva 	feud.Barony
	罗玛Geatish_roma 	feud.Barony
	黑德Hejde 	feud.Barony
	海姆瑟Hemse 	feud.Barony
	乌特海姆Othem 	feud.Barony
	斯利特Slite 	feud.Barony
	维斯堡Visborg 	feud.Barony
	维斯比Visby 	feud.Barony
}

func (c *哥特兰GotlandCounty) BAlva阿尔瓦() feud.Barony {
	return c.阿尔瓦Alva
}
    
func (c *哥特兰GotlandCounty) BGeatish_roma罗玛() feud.Barony {
	return c.罗玛Geatish_roma
}
    
func (c *哥特兰GotlandCounty) BHejde黑德() feud.Barony {
	return c.黑德Hejde
}
    
func (c *哥特兰GotlandCounty) BHemse海姆瑟() feud.Barony {
	return c.海姆瑟Hemse
}
    
func (c *哥特兰GotlandCounty) BOthem乌特海姆() feud.Barony {
	return c.乌特海姆Othem
}
    
func (c *哥特兰GotlandCounty) BSlite斯利特() feud.Barony {
	return c.斯利特Slite
}
    
func (c *哥特兰GotlandCounty) BVisborg维斯堡() feud.Barony {
	return c.维斯堡Visborg
}
    
func (c *哥特兰GotlandCounty) BVisby维斯比() feud.Barony {
	return c.维斯比Visby
}
    
var CGotland哥特兰 GotlandCounty = &哥特兰GotlandCounty{}

func init() {
	f := CGotland哥特兰.(*哥特兰GotlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "301",
		Title:     "gotland",
		TitleName: "哥特兰",
		TitleCode: "c_gotland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔瓦Alva = BAlva阿尔瓦
	f.阿尔瓦Alva.SetParent(f)
	
	f.罗玛Geatish_roma = BGeatish_roma罗玛
	f.罗玛Geatish_roma.SetParent(f)
	
	f.黑德Hejde = BHejde黑德
	f.黑德Hejde.SetParent(f)
	
	f.海姆瑟Hemse = BHemse海姆瑟
	f.海姆瑟Hemse.SetParent(f)
	
	f.乌特海姆Othem = BOthem乌特海姆
	f.乌特海姆Othem.SetParent(f)
	
	f.斯利特Slite = BSlite斯利特
	f.斯利特Slite.SetParent(f)
	
	f.维斯堡Visborg = BVisborg维斯堡
	f.维斯堡Visborg.SetParent(f)
	
	f.维斯比Visby = BVisby维斯比
	f.维斯比Visby.SetParent(f)
	
}
