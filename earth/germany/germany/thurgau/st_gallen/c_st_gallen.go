package st_gallen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type St_gallenCounty interface {
    feud.County
    BAltstatten阿尔特施泰滕() 	feud.Barony
    BAppenzell阿彭策尔() 	feud.Barony
    BFrauenfeld弗劳恩费尔德() 	feud.Barony
    BHerisau黑里绍() 	feud.Barony
    BLichtensteig利希滕施泰格() 	feud.Barony
    BRheineck赖内克() 	feud.Barony
    BStgallen圣加仑() 	feud.Barony
    BVaduz瓦杜兹() 	feud.Barony
}

type 图尔高St_gallenCounty struct {
	feud.BaseCounty
	阿尔特施泰滕Altstatten 	feud.Barony
	阿彭策尔Appenzell 	feud.Barony
	弗劳恩费尔德Frauenfeld 	feud.Barony
	黑里绍Herisau 	feud.Barony
	利希滕施泰格Lichtensteig 	feud.Barony
	赖内克Rheineck 	feud.Barony
	圣加仑Stgallen 	feud.Barony
	瓦杜兹Vaduz 	feud.Barony
}

func (c *图尔高St_gallenCounty) BAltstatten阿尔特施泰滕() feud.Barony {
	return c.阿尔特施泰滕Altstatten
}
    
func (c *图尔高St_gallenCounty) BAppenzell阿彭策尔() feud.Barony {
	return c.阿彭策尔Appenzell
}
    
func (c *图尔高St_gallenCounty) BFrauenfeld弗劳恩费尔德() feud.Barony {
	return c.弗劳恩费尔德Frauenfeld
}
    
func (c *图尔高St_gallenCounty) BHerisau黑里绍() feud.Barony {
	return c.黑里绍Herisau
}
    
func (c *图尔高St_gallenCounty) BLichtensteig利希滕施泰格() feud.Barony {
	return c.利希滕施泰格Lichtensteig
}
    
func (c *图尔高St_gallenCounty) BRheineck赖内克() feud.Barony {
	return c.赖内克Rheineck
}
    
func (c *图尔高St_gallenCounty) BStgallen圣加仑() feud.Barony {
	return c.圣加仑Stgallen
}
    
func (c *图尔高St_gallenCounty) BVaduz瓦杜兹() feud.Barony {
	return c.瓦杜兹Vaduz
}
    
var CSt_gallen图尔高 St_gallenCounty = &图尔高St_gallenCounty{}

func init() {
	f := CSt_gallen图尔高.(*图尔高St_gallenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "248",
		Title:     "st_gallen",
		TitleName: "图尔高",
		TitleCode: "c_st_gallen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔特施泰滕Altstatten = BAltstatten阿尔特施泰滕
	f.阿尔特施泰滕Altstatten.SetParent(f)
	
	f.阿彭策尔Appenzell = BAppenzell阿彭策尔
	f.阿彭策尔Appenzell.SetParent(f)
	
	f.弗劳恩费尔德Frauenfeld = BFrauenfeld弗劳恩费尔德
	f.弗劳恩费尔德Frauenfeld.SetParent(f)
	
	f.黑里绍Herisau = BHerisau黑里绍
	f.黑里绍Herisau.SetParent(f)
	
	f.利希滕施泰格Lichtensteig = BLichtensteig利希滕施泰格
	f.利希滕施泰格Lichtensteig.SetParent(f)
	
	f.赖内克Rheineck = BRheineck赖内克
	f.赖内克Rheineck.SetParent(f)
	
	f.圣加仑Stgallen = BStgallen圣加仑
	f.圣加仑Stgallen.SetParent(f)
	
	f.瓦杜兹Vaduz = BVaduz瓦杜兹
	f.瓦杜兹Vaduz.SetParent(f)
	
}
