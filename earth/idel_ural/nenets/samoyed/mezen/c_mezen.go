package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MezenCounty interface {
    feud.County
    BKimzha基姆扎() 	feud.Barony
    BKyzma克兹马() 	feud.Barony
    BLoptyuga洛普秋加() 	feud.Barony
    BMezen梅津() 	feud.Barony
    BPyoza皮奥扎() 	feud.Barony
    BPyssa佩萨() 	feud.Barony
    BSula苏拉() 	feud.Barony
}

type 梅津MezenCounty struct {
	feud.BaseCounty
	基姆扎Kimzha 	feud.Barony
	克兹马Kyzma 	feud.Barony
	洛普秋加Loptyuga 	feud.Barony
	梅津Mezen 	feud.Barony
	皮奥扎Pyoza 	feud.Barony
	佩萨Pyssa 	feud.Barony
	苏拉Sula 	feud.Barony
}

func (c *梅津MezenCounty) BKimzha基姆扎() feud.Barony {
	return c.基姆扎Kimzha
}
    
func (c *梅津MezenCounty) BKyzma克兹马() feud.Barony {
	return c.克兹马Kyzma
}
    
func (c *梅津MezenCounty) BLoptyuga洛普秋加() feud.Barony {
	return c.洛普秋加Loptyuga
}
    
func (c *梅津MezenCounty) BMezen梅津() feud.Barony {
	return c.梅津Mezen
}
    
func (c *梅津MezenCounty) BPyoza皮奥扎() feud.Barony {
	return c.皮奥扎Pyoza
}
    
func (c *梅津MezenCounty) BPyssa佩萨() feud.Barony {
	return c.佩萨Pyssa
}
    
func (c *梅津MezenCounty) BSula苏拉() feud.Barony {
	return c.苏拉Sula
}
    
var CMezen梅津 MezenCounty = &梅津MezenCounty{}

func init() {
	f := CMezen梅津.(*梅津MezenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1823",
		Title:     "mezen",
		TitleName: "梅津",
		TitleCode: "c_mezen",
		Baronies:  map[string]feud.Barony{},
	}

	f.基姆扎Kimzha = BKimzha基姆扎
	f.基姆扎Kimzha.SetParent(f)
	
	f.克兹马Kyzma = BKyzma克兹马
	f.克兹马Kyzma.SetParent(f)
	
	f.洛普秋加Loptyuga = BLoptyuga洛普秋加
	f.洛普秋加Loptyuga.SetParent(f)
	
	f.梅津Mezen = BMezen梅津
	f.梅津Mezen.SetParent(f)
	
	f.皮奥扎Pyoza = BPyoza皮奥扎
	f.皮奥扎Pyoza.SetParent(f)
	
	f.佩萨Pyssa = BPyssa佩萨
	f.佩萨Pyssa.SetParent(f)
	
	f.苏拉Sula = BSula苏拉
	f.苏拉Sula.SetParent(f)
	
}
