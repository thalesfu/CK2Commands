package hebron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HebronCounty interface {
    feud.County
    BAlsamoa艾萨摩亚() 	feud.Barony
    BBethgibelin贝特吉伯兰() 	feud.Barony
    BBethlehem伯利恒() 	feud.Barony
    BDeimachar德马赤() 	feud.Barony
    BHebron希伯伦() 	feud.Barony
    BLatrun拉特伦() 	feud.Barony
    BSaintcharlton圣查尔顿() 	feud.Barony
    BSyrbelmont贝尔蒙特() 	feud.Barony
}

type 希伯伦HebronCounty struct {
	feud.BaseCounty
	艾萨摩亚Alsamoa 	feud.Barony
	贝特吉伯兰Bethgibelin 	feud.Barony
	伯利恒Bethlehem 	feud.Barony
	德马赤Deimachar 	feud.Barony
	希伯伦Hebron 	feud.Barony
	拉特伦Latrun 	feud.Barony
	圣查尔顿Saintcharlton 	feud.Barony
	贝尔蒙特Syrbelmont 	feud.Barony
}

func (c *希伯伦HebronCounty) BAlsamoa艾萨摩亚() feud.Barony {
	return c.艾萨摩亚Alsamoa
}
    
func (c *希伯伦HebronCounty) BBethgibelin贝特吉伯兰() feud.Barony {
	return c.贝特吉伯兰Bethgibelin
}
    
func (c *希伯伦HebronCounty) BBethlehem伯利恒() feud.Barony {
	return c.伯利恒Bethlehem
}
    
func (c *希伯伦HebronCounty) BDeimachar德马赤() feud.Barony {
	return c.德马赤Deimachar
}
    
func (c *希伯伦HebronCounty) BHebron希伯伦() feud.Barony {
	return c.希伯伦Hebron
}
    
func (c *希伯伦HebronCounty) BLatrun拉特伦() feud.Barony {
	return c.拉特伦Latrun
}
    
func (c *希伯伦HebronCounty) BSaintcharlton圣查尔顿() feud.Barony {
	return c.圣查尔顿Saintcharlton
}
    
func (c *希伯伦HebronCounty) BSyrbelmont贝尔蒙特() feud.Barony {
	return c.贝尔蒙特Syrbelmont
}
    
var CHebron希伯伦 HebronCounty = &希伯伦HebronCounty{}

func init() {
	f := CHebron希伯伦.(*希伯伦HebronCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "776",
		Title:     "hebron",
		TitleName: "希伯伦",
		TitleCode: "c_hebron",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾萨摩亚Alsamoa = BAlsamoa艾萨摩亚
	f.艾萨摩亚Alsamoa.SetParent(f)
	
	f.贝特吉伯兰Bethgibelin = BBethgibelin贝特吉伯兰
	f.贝特吉伯兰Bethgibelin.SetParent(f)
	
	f.伯利恒Bethlehem = BBethlehem伯利恒
	f.伯利恒Bethlehem.SetParent(f)
	
	f.德马赤Deimachar = BDeimachar德马赤
	f.德马赤Deimachar.SetParent(f)
	
	f.希伯伦Hebron = BHebron希伯伦
	f.希伯伦Hebron.SetParent(f)
	
	f.拉特伦Latrun = BLatrun拉特伦
	f.拉特伦Latrun.SetParent(f)
	
	f.圣查尔顿Saintcharlton = BSaintcharlton圣查尔顿
	f.圣查尔顿Saintcharlton.SetParent(f)
	
	f.贝尔蒙特Syrbelmont = BSyrbelmont贝尔蒙特
	f.贝尔蒙特Syrbelmont.SetParent(f)
	
}
