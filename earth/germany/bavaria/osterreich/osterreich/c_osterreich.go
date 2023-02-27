package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OsterreichCounty interface {
    feud.County
    BBaden_wien巴登() 	feud.Barony
    BHeiligenkreuz海利根克罗伊茨() 	feud.Barony
    BKlosterneuburg克洛斯特新堡() 	feud.Barony
    BKorneuburg科尔新堡() 	feud.Barony
    BModling默德灵() 	feud.Barony
    BWagram瓦格拉姆() 	feud.Barony
    BWien维也纳() 	feud.Barony
}

type 维恩OsterreichCounty struct {
	feud.BaseCounty
	巴登Baden_wien 	feud.Barony
	海利根克罗伊茨Heiligenkreuz 	feud.Barony
	克洛斯特新堡Klosterneuburg 	feud.Barony
	科尔新堡Korneuburg 	feud.Barony
	默德灵Modling 	feud.Barony
	瓦格拉姆Wagram 	feud.Barony
	维也纳Wien 	feud.Barony
}

func (c *维恩OsterreichCounty) BBaden_wien巴登() feud.Barony {
	return c.巴登Baden_wien
}
    
func (c *维恩OsterreichCounty) BHeiligenkreuz海利根克罗伊茨() feud.Barony {
	return c.海利根克罗伊茨Heiligenkreuz
}
    
func (c *维恩OsterreichCounty) BKlosterneuburg克洛斯特新堡() feud.Barony {
	return c.克洛斯特新堡Klosterneuburg
}
    
func (c *维恩OsterreichCounty) BKorneuburg科尔新堡() feud.Barony {
	return c.科尔新堡Korneuburg
}
    
func (c *维恩OsterreichCounty) BModling默德灵() feud.Barony {
	return c.默德灵Modling
}
    
func (c *维恩OsterreichCounty) BWagram瓦格拉姆() feud.Barony {
	return c.瓦格拉姆Wagram
}
    
func (c *维恩OsterreichCounty) BWien维也纳() feud.Barony {
	return c.维也纳Wien
}
    
var COsterreich维恩 OsterreichCounty = &维恩OsterreichCounty{}

func init() {
	f := COsterreich维恩.(*维恩OsterreichCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "449",
		Title:     "osterreich",
		TitleName: "维恩",
		TitleCode: "c_osterreich",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴登Baden_wien = BBaden_wien巴登
	f.巴登Baden_wien.SetParent(f)
	
	f.海利根克罗伊茨Heiligenkreuz = BHeiligenkreuz海利根克罗伊茨
	f.海利根克罗伊茨Heiligenkreuz.SetParent(f)
	
	f.克洛斯特新堡Klosterneuburg = BKlosterneuburg克洛斯特新堡
	f.克洛斯特新堡Klosterneuburg.SetParent(f)
	
	f.科尔新堡Korneuburg = BKorneuburg科尔新堡
	f.科尔新堡Korneuburg.SetParent(f)
	
	f.默德灵Modling = BModling默德灵
	f.默德灵Modling.SetParent(f)
	
	f.瓦格拉姆Wagram = BWagram瓦格拉姆
	f.瓦格拉姆Wagram.SetParent(f)
	
	f.维也纳Wien = BWien维也纳
	f.维也纳Wien.SetParent(f)
	
}
