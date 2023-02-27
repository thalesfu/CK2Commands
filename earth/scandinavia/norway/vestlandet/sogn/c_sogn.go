package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SognCounty interface {
    feud.County
    BAurland艾于兰() 	feud.Barony
    BFortun福通() 	feud.Barony
    BKaupanger凯于庞厄尔() 	feud.Barony
    BKvamsoy克瓦姆岛() 	feud.Barony
    BLeikanger莱康厄尔() 	feud.Barony
    BSogndal松达尔() 	feud.Barony
    BVik维克() 	feud.Barony
}

type 松恩SognCounty struct {
	feud.BaseCounty
	艾于兰Aurland 	feud.Barony
	福通Fortun 	feud.Barony
	凯于庞厄尔Kaupanger 	feud.Barony
	克瓦姆岛Kvamsoy 	feud.Barony
	莱康厄尔Leikanger 	feud.Barony
	松达尔Sogndal 	feud.Barony
	维克Vik 	feud.Barony
}

func (c *松恩SognCounty) BAurland艾于兰() feud.Barony {
	return c.艾于兰Aurland
}
    
func (c *松恩SognCounty) BFortun福通() feud.Barony {
	return c.福通Fortun
}
    
func (c *松恩SognCounty) BKaupanger凯于庞厄尔() feud.Barony {
	return c.凯于庞厄尔Kaupanger
}
    
func (c *松恩SognCounty) BKvamsoy克瓦姆岛() feud.Barony {
	return c.克瓦姆岛Kvamsoy
}
    
func (c *松恩SognCounty) BLeikanger莱康厄尔() feud.Barony {
	return c.莱康厄尔Leikanger
}
    
func (c *松恩SognCounty) BSogndal松达尔() feud.Barony {
	return c.松达尔Sogndal
}
    
func (c *松恩SognCounty) BVik维克() feud.Barony {
	return c.维克Vik
}
    
var CSogn松恩 SognCounty = &松恩SognCounty{}

func init() {
	f := CSogn松恩.(*松恩SognCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1625",
		Title:     "sogn",
		TitleName: "松恩",
		TitleCode: "c_sogn",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾于兰Aurland = BAurland艾于兰
	f.艾于兰Aurland.SetParent(f)
	
	f.福通Fortun = BFortun福通
	f.福通Fortun.SetParent(f)
	
	f.凯于庞厄尔Kaupanger = BKaupanger凯于庞厄尔
	f.凯于庞厄尔Kaupanger.SetParent(f)
	
	f.克瓦姆岛Kvamsoy = BKvamsoy克瓦姆岛
	f.克瓦姆岛Kvamsoy.SetParent(f)
	
	f.莱康厄尔Leikanger = BLeikanger莱康厄尔
	f.莱康厄尔Leikanger.SetParent(f)
	
	f.松达尔Sogndal = BSogndal松达尔
	f.松达尔Sogndal.SetParent(f)
	
	f.维克Vik = BVik维克
	f.维克Vik.SetParent(f)
	
}
